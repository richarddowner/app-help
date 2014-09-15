docs
========

Simple Golang CLI program to display application help for popular commands.

* docs.go programName
* Use some sort of external files for the docs
* Use a directory structure accessable by programName

### Install

```
go get github.com/richarddowner/docs
```

### Usage 

````
$ docs git
(1) init
(2) commit
(3) pull

>> select: 3
git pull origin
  
````

### Folder Structure

* docs/
* docs/git
* docs/git/init
* docs/git/commit
* docs/git/pull


