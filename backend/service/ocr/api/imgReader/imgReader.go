package imgReader

type ImgReader interface {
	ReadImg() (string, error)
}
