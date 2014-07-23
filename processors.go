package main

import (
	"image"
)

type GoliathImage image.Image

func Copy(i image.Image) image.Image {
	i2 := i
	return i2
}

//
// func Process(i *image.Image) *image.Image {
//
// }
//
// func FixOrientation() {}
//
// func BoxFitResize() {}
//
// func Watermark() {}
//
// func Canvas() {}
//
// func DynamicResize() {}
//
// func BestFitResize() {}
