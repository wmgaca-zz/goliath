package main

import (
	"crypto/md5"
	"fmt"
	"github.com/wmgaca/go-phash"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

var PHashMap = make(map[uint64]*ZeeImage)

var MD5HashMap = make(map[[16]byte]*ZeeImage)

type ZeeImage struct {
	Path    string
	PHash   uint64
	MD5Hash [16]byte
}

func NewZeeImageFromRequest(r *http.Request, fieldName string) (*ZeeImage, error) {
	imageFile, imageFileHeader, err := r.FormFile("image")
	if err != nil {
		return nil, err
	}

	tempFile, err := ioutil.TempFile(StaticDir, imageFileHeader.Filename+"-")
	defer tempFile.Close()
	if err != nil {
		log.Println("ioutil.TempFile failed with", err)
		return nil, err
	}

	_, err = io.Copy(tempFile, imageFile)
	if err != nil {
		log.Println("io.Copy failed with", err)
		return nil, err
	}

	log.Println("Looks good, dude.")

	return &ZeeImage{Path: tempFile.Name()}, nil
}

func (z *ZeeImage) String() string {
	return fmt.Sprintf("%s (phash: %d, md5: %x)", z.Path, z.PHash, z.MD5Hash)
}

func (z *ZeeImage) computePHash() {
	result, err := phash.ImageHashDCT(z.Path)
	if err != nil {
		log.Println("Big Bad Error =>", err)
		return
	}
	z.PHash = result
}

func (z *ZeeImage) computeMD5Hash() {
	bytes, err := ioutil.ReadFile(z.Path)
	if err != nil {
		log.Println("Big Bad Error =>", err)
		return
	}

	z.MD5Hash = md5.Sum(bytes)
}

func (z *ZeeImage) Compute(addToSet bool) {
	log.Println("Compute =>", z.String())

	z.computePHash()
	z.computeMD5Hash()

	if addToSet {
		z.AddToSet()
	}

	log.Println("Finish =>", z.String())
}

func (z *ZeeImage) AddToSet() {
	PHashMap[z.PHash] = z
	MD5HashMap[z.MD5Hash] = z
}
