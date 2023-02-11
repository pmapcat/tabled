package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/xuri/excelize/v2"
	"os"
	"runtime"
	"strings"
)

const DELIM = " :: "

var PURPLE = func() string {
	if runtime.GOOS == "windows" {
		return ""
	}
	return "\033[35m"
}()

var RESET_COLOR = func() string {
	if runtime.GOOS == "windows" {
		return ""
	}
	return "\033[0m"
}()

const NEWLINE = "\n"

// data must be sorted to emit nicely working excel file
func consumeSortedFlatData(fpath string) {
	// ;; apparently, no error means file exists ::facepalm::
	_, err := os.Stat(fpath)
	if err == nil {
		panic("File " + fpath + " exists, overriding it !!!WILL CAUSE DATA LOSS!!!")
	}

	// build xlsx filling up
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	prevSheet, prevCol := "", ""
	colIndex, valIndex := 0, 0
	for scanner.Scan() {
		_s := strings.Split(scanner.Text(), DELIM)
		// poor's woman destructuring
		sheet, col, val := _s[0], _s[1], _s[2]
		if sheet != prevSheet {
			prevSheet = sheet
			colIndex, valIndex = 0, 0

			// Create a new sheet.
			_, err := f.NewSheet(sheet)
			if err != nil {
				panic(err)
			}
		}

		if col != prevCol {
			prevCol = col
			colIndex += 1
			valIndex = 0
		}

		_coord, err := excelize.CoordinatesToCellName(colIndex, valIndex)
		if err != nil {
			// wtf error here is supposed to mean?
			panic(err)
		}

		f.SetCellValue(sheet, _coord, val)
		valIndex += 1
	}
	if err := f.SaveAs(fpath); err != nil {
		panic(err)
	}
}

func emitFlatData(fpath string) {
	rowName := map[int]string{}

	f, err := excelize.OpenFile(fpath)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	fmtByLine := map[int]string{
		0: PURPLE,
	}

	for _, sheetName := range f.GetSheetMap() {
		rows, err := f.GetRows(sheetName)
		if err != nil {
			panic(err)
		}

		for rowIndex, rowValues := range rows {
			if rowIndex == 0 {
				// generate row names
				for columnIndex, columnValue := range rowValues {
					rowName[columnIndex] = columnValue
				}
			}
			if rowIndex != 0 {
				for columnIndex, columnValue := range rowValues {
					fmt.Print(fmtByLine[columnIndex] + sheetName + DELIM + rowName[columnIndex] + DELIM + columnValue + RESET_COLOR + NEWLINE)
				}
			}
		}
	}
}

func main() {
	infile := flag.String("in", "", "file to read")
	outfile := flag.String("out", "", "file to write result data to")
	flag.Parse()

	if *infile != "" {
		// "data/sample.xlsx"
		emitFlatData(*infile)
		return
	}

	if *outfile != "" {
		consumeSortedFlatData(*outfile)
	}

}
