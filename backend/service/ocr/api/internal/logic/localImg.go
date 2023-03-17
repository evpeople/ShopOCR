package logic

import "github.com/evpeople/ShopOCR/service/ocr/api/imgReader"

func NewImageReader() imgReader.ImgReader {
	reader := imgReader.NewLocalReader("imgs")
	return &reader
}
