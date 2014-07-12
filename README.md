# Goliath
Image processing POC written in GO:
* loads a set of images and keeps them in memory,
* provides an HTTP endpoint to check if given (posted) image exists in the set

## Prerequisites
* [go](http://golang.org) would be nice.

## Install
```bash
mkdir [my-awsome-dir]
cd [my-awsome-dir]
export GOPATH=`pwd`
go get github.com/wmgaca/goliath
```

## Run
```bash
bin/goliath [image-set-dir]
```

There you go, now check if it's running:
```bash
curl -i -F -filedata=~/foo/bar.png  http://localhost:8000/compare/
```
