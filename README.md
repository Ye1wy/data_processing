# About
In this project, I wanted to try to create 3 programs:
1. A program that reads a file from Json and Xml through the implementation of the interface, which will also convert the read file from json to xml and back, to visually compare with an existing xml or json file
2. A program that reads a file from json or xml, parses them, and then compares an existing json or xml file with the read one
3. A program that reads a file from txt and compares an existing txt file

# Install
For install xml and json reader
```
go build src/data_reader/data_reader.go
```
For install compare xml and json
```
go build src/data_compare/data_comapre.go
```
For install fs comapre
```
go build src/fs_compare/fs_compare.go
```

# Using

Data reader:\
Programm have flag -f (*your xml or json file*)

**example:**
```
data_reader -f xml/original_database.xml
```
Data compare:\
Programm have 2 flag:\
- --old - file that will be the basis
- --new - file that will be compared with the base file

**example:**
```
 data_compare --old xml/original_database.xml --new json/stolen_database.json
```
FS compare: work like data compare