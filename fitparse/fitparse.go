package fitparse

import(
	"errors"
	"fmt"
	"os"
)

type FitFile struct {
	fp *os.File
	FileHeader *FileHeader
}

func NewFitFile(fp *os.File) (*FitFile, error) {
	ff := &FitFile{
		fp: fp,
		FileHeader: &FileHeader{},
	}

	err := ff.Parse()
	if err != nil {
		return nil, err
	}
	return ff, nil
}

func (fit *FitFile) Parse() error {
	fmt.Println("starting parse")
	var err error
	err = fit.ReadFileHeader()
	if err != nil {
		return err
	}
	rh, err := fit.ReadRecordHeader()
	fmt.Println(rh)

	if rh.MessageType {
		fmt.Println(fit.ReadDefinition())
	} else {
		return errors.New("Not yet implemented")
	}
	return nil
}
