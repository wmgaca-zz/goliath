# Go
[![Build Status](https://travis-ci.org/wmgaca/goliath.svg?branch=master)](https://travis-ci.org/wmgaca/goliath)

Image processing POC written in Go:
* loads a set of images and keeps them in memory,
* provides an HTTP endpoint to check if given (posted) image exists
  in the set.

# Up and running

### Prerequisitories?
- Well, [The Go Programming Language](http://golang.org/doc/install)
  would be nice

### Set up your Go environment
```bash
mkdir ~/dev/go
export GOPATH=~/dev/go        # You want to add this to your .bashrc
export PATH=$PATH:$GOPATH/bin # This too
```

### Install
```bash
go get github.com/wmgaca/goliath
```

### Configure
Goliath reads the port and AWS credentials from the environment,
make sure you set them up before running the server:
```bash
export PORT=10666
export AWS_ACCESS_KEY_ID="YOUR-AWS-ACCESS-KEY-GOES-HERE"
export AWS_SECRET_ACCESS_KEY="YOUR-AWS-SECRET-ACCESS-KEY-GOES-HERE"
```

### Run
```bash
goliath
```
