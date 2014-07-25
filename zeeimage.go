package main

import (
	"code.google.com/p/go-uuid/uuid"
	"crypto/md5"
	"fmt"
	"github.com/wmgaca/go-phash"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var PHashMap = make(map[uint64]*ZeeImage)

var MD5HashMap = make(map[[16]byte]*ZeeImage)

type ZeeImage struct {
	Path    string
	S3Paths map[string]string
	UUID    string
	Time    time.Time
	PHash   uint64
	MD5Hash [16]byte
}

func NewZeeImageFromRequest(r *http.Request, fieldName string) (*ZeeImage, error) {
	imageFile, imageFileHeader, err := r.FormFile(fieldName)
	if err != nil {
		return nil, err
	}

	tempFile, err := ioutil.TempFile(StaticDir, imageFileHeader.Filename+"-")
	defer tempFile.Close()
	if err != nil {
		return nil, err
	}

	log.Println(tempFile.Name())

	_, err = io.Copy(tempFile, imageFile)
	if err != nil {
		return nil, err
	}

	zeeImage := &ZeeImage{Path: tempFile.Name()}
	zeeImage.generateS3Paths()

	return zeeImage, nil
}

func (z *ZeeImage) generateS3Paths() {
	z.UUID = uuid.New()
	z.Time = time.Now()
	z.S3Paths = make(map[string]string)
}

func (z *ZeeImage) String() string {
	return fmt.Sprintf("%s (pHash: %d, MD5: %x)", z.Path, z.PHash, z.MD5Hash)
}

func (z *ZeeImage) computePHash() {
	result, err := phash.ImageHashDCT(z.Path)
	if err != nil {
		log.Println("Big Bad Error while computing pHash =>", err)
		return
	}
	z.PHash = result
}

func (z *ZeeImage) computeMD5Hash() {
	bytes, err := ioutil.ReadFile(z.Path)
	if err != nil {
		log.Println("Big Bad Error while computing MD5 hash =>", err)
		return
	}
	z.MD5Hash = md5.Sum(bytes)
}

func (z *ZeeImage) Compute(addToSet bool) {
	z.computePHash()
	z.computeMD5Hash()

	if addToSet {
		z.AddToSet()
	}
}

func (z *ZeeImage) AddToSet() {
	PHashMap[z.PHash] = z
	MD5HashMap[z.MD5Hash] = z
}
