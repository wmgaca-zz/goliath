# Goliath
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

### WAT? :scream:

#### Upload a new image
```bash
$ curl -v -include --form "image=@doge.jpg" goliath.go/{apiVersion:[0-9\.]+}/upload/
{
    "Status": true,
    "Path": "/a/relative/path/to/your/doge.jpg"
    "Time": "1089278"
}
```

### Check if a given image already exists
```bash
$ curl -v -include --form "image=@doge.jpg" goliath.go/{apiVersion:[0-9\.]+}/check/{methodName:[a-z0-9]+}/
{
    "Status": true
    "Exists": true,
    "Path":   "/a/relative/path/to/your/doge.jpg"
    "Time":   "1089278"
}
```

### Error response
```bash
$ curl goliath.go/no/idea/what/i/am/doing/
{
    "Status": false,
    "Error": "Your have no idea what you're doing."
}
```
