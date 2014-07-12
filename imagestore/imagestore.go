package imagestore

import (
	// "fmt"
	"io/ioutil"
	// "os"
	"path"
	"strings"
	// "time"
)

type Image struct {
	name string
	path string
}

func NewImage(fullPath string) *Image {
	fPath, fName := path.Split(fullPath)
	return &Image{fName, fPath}
}

// MD5 hash -> *Image
var md5Map map[string]*Image

// pHash -> *Image
var pHashMap map[string]*Image

// OpenCV hash -> *Image
var openCVMap map[string]*Image

// Name -> *Image
var images map[string]*Image

func Init(imageSetPath string) {
	images = make(map[string]*Image)
	files, _ := ioutil.ReadDir(imageSetPath)
	for _, f := range files {
		if strings.HasPrefix(f.Name(), ".") {
			continue
		}

		image := NewImage(path.Join(imageSetPath, f.Name()))
		images[image.name] = image

		// fileName := f.Name()
		// filePath := path.Join(imageSetPath, fileName)

		// imageFile, _ := os.Open(image.path)
		// fmt.Println("=>", image.name, image.path)
	}
}

func Exists() (exists bool) {
	return true
}
