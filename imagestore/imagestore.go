// Package takes care of all the necessary image manipulation
package imagestore

import (
	"io/ioutil"
	"path"
	"strings"
)

// Image type used to store
type Image struct {
	name string
	path string
}

// Image name getter
func (i *Image) Name() string {
	return i.name
}

// Image path getter
func (i *Image) Path() string {
	return i.path
}

// Create a new Image instance
func NewImage(fullPath string) *Image {
	fPath, fName := path.Split(fullPath)

	return &Image{
		name: fName,
		path: fPath,
	}
}

// MD5 hash -> *Image
var md5Map map[string]*Image

// pHash -> *Image
var pHashMap map[string]*Image

// OpenCV hash -> *Image
var openCVMap map[string]*Image

// Name -> *Image
var images map[string]*Image

// Init will load and process all of the images in given directory
func Init(imageSetPath string) {
	images = make(map[string]*Image)
	files, _ := ioutil.ReadDir(imageSetPath)
	for _, f := range files {
		if strings.HasPrefix(f.Name(), ".") {
			continue
		}

		image := NewImage(path.Join(imageSetPath, f.Name()))
		images[image.name] = image
	}
}

// Get the original image by its name
func Get(name string) (image *Image) {
	return images[name]
}

// Check if an image with this name already exists in the set
func Exists(name string) bool {
	_, ok := images[name]
	return ok
}
