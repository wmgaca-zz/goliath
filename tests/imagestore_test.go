package main

import (
	"fmt"
	"github.com/wmgaca/goliath/imagestore"
	"path"
	"testing"
)

const TEST_SET_PATH = "test_data"

var existingFiles = []string{
	"cs8n2c08.png"}

var nonExistingFiles = []string{
	"foobar.png",
	"helloWorld.png"}

func TestExistsShouldReturnTrueForExistingFiles(t *testing.T) {
	imagestore.Init("test_data")

	for _, fileName := range existingFiles {
		if !imagestore.Exists(fileName) {
			t.FailNow()
		}
	}
}

func TestExistsShouldReturnFalseForNonExistingFiles(t *testing.T) {
	imagestore.Init("test_data")

	for _, fileName := range nonExistingFiles {
		if imagestore.Exists(fileName) {
			t.FailNow()
		}
	}
}

func TestInitShouldCreateImagesWithProperData(t *testing.T) {
	imagestore.Init(TEST_SET_PATH)

	for _, fileName := range existingFiles {
		image := imagestore.Get(fileName)

		if image.Name() != fileName {
			fmt.Printf("Got: %s, expected: %s", image.Name(), fileName)
			t.FailNow()
		}

		if path.Join(image.Path()) != path.Join(TEST_SET_PATH) {
			fmt.Printf("Got: %s, expected: %s\n", image.Path(), TEST_SET_PATH)
			t.FailNow()
		}
	}
}
