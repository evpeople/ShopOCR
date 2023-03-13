package logic

import "github.com/evpeople/ShopOCR/service/ocr/imgReader"

func NewImageReader() imgReader.ImgReader {
	reader := imgReader.NewLocalReader("imgs")
	return &reader
}
