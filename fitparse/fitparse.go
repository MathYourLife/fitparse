package fitparse

import(
	"fmt"
	"os"
)

type FitFile struct {
	fp *os.File
	FileHeader *FileHeader
}

func NewParser(fp *os.File) (*FitFile, error) {
	return &FitFile{
		fp: fp,
		FileHeader: &FileHeader{},
	}, nil
}

func (fit *FitFile) Parse() error {
	fmt.Println("starting parse")
	var err error
	err = fit.ReadFileHeader()
	if err != nil {
		return err
	}
	return nil
}
