package fitparse

import(
	"fmt"
	"os"
)

type FitFile struct {
	fp *os.File
	FileHeader *FileHeader
	LocalMsgTypes map[uint8]*Definition
}

func NewFitFile(fp *os.File) (*FitFile, error) {
	ff := &FitFile{
		fp: fp,
		FileHeader: &FileHeader{},
		LocalMsgTypes: map[uint8]*Definition{},
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

	for i:=0; i<3; i++ {
		rh, err := fit.ReadRecordHeader()
		if err != nil {
			return err
		}
		fmt.Printf("%v\n",rh.Type)
		fmt.Printf("%v\n",rh.DataMsgType)
		fmt.Printf("%#v\n", rh)

		if rh.Type == RECORDHEADER_DEFINITION {
			def, err := fit.ReadDefinition()
			if err != nil {
				return err
			}
			fit.LocalMsgTypes[rh.LocalMessageId] = def
			fmt.Printf("%#v\n", def)
			fmt.Printf("Global Message Number: %d\n", def.GlobalMessageNum)
		} else {
			fit.ReadData(fit.LocalMsgTypes[rh.LocalMessageId])
		}
	}
	return nil
}
