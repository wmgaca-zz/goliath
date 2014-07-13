[![Build Status](https://travis-ci.org/wmgaca/goliath.svg?branch=master)](https://travis-ci.org/wmgaca/goliath)

# Goliath
Image processing POC written in Go:
* loads a set of images and keeps them in memory,
* provides an HTTP endpoint to check if given (posted) image exists in the set.

# Up and running

### Prerequisitories?
- Well, [The Go Programming Language](http://golang.org/doc/install) would be nice

### Set up your Go environment
```bash
mkdir ~/dev/go
export GOPATH=~/dev/go        # You want to add this to your .bashrc
export PATH=$PATH:$GOPATH/bin # This too
```

### Install goliath
```bash
go get github.com/wmgaca/goliath
```

### Run goliath
```bash
PORT=5000 goliath
```
