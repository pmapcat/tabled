## `tabled` is a simple and fast tool to flatten/unflatten excel files

* TODO: rewrite readme
* TODO: debug unflatten
* TODO: make this thing blazing fast

### Rationale

* There are a lot of files that don't have line breaks
* Tey make __Emacs__ and __Cider__ go bad for some reason I don't know
* __grep__ and __ag__ do not work nicely with file without linebreaks 

This is a dead simple tool to insert line breaks after a text pattern. 
Here are some examples how it may be used

### Install 

`go install github.com/pmapcat/linebr@latest`


### Example 

```
echo '{:name "Автомобільні диски", :parent_id 4653880, :category_id 4624949} {:name "Плитка", :parent_id 4639424, :category_id 4625710} {:name "Посуд", :parent_id 2394287, :category_id 4626594}' | linebr "} " 

{:name "Автомобільні диски", :parent_id 4653880, :category_id 4624949} 
{:name "Плитка", :parent_id 4639424, :category_id 4625710} 
{:name "Посуд", :parent_id 2394287, :category_id 4626594}
```

#### Performance 

It will take __~800ms__ on my machine to go through __150mb__ oneline file.
This pefromance can be improved

* TODO Make it faster

```
time cat data/test_huge.txt | linebr "|" | tail 
4a53b4e7-e85a-4d60-b50c-768a1243dd55|
753b6ca9-4c69-4907-869e-48c1ed900a56|
bed998f0-0adf-4655-97b5-c81db40b846b|
1495d592-4131-4fad-94d9-37416762e777|
d660d940-0099-4cc4-915a-bb6a5084dfa0|
7f67e3ac-eadd-4034-8053-9f78631ed24c|
5bd738cc-de71-45cc-b570-660ac353f882|
6c6eab8f-6413-4857-ba95-1461621f5813|
d694b546-40d3-4d56-8b5a-90d2810bfa3c|
7e266fc9-88dc-445f-a76d-d4666abab2c9|
cat data/test_huge.txt  0,01s user 0,20s system 24% cpu 0,836 total
linebr "|"  0,71s user 0,25s system 114% cpu 0,836 total
tail  0,25s user 0,19s system 52% cpu 0,835 total
```


