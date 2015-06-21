package fitparse

import(
	"errors"
	"encoding/binary"
)

type FileHeader struct {
	Size uint8
	ProtocolVersion uint8
	Profile uint16
	DataSize uint32
	DataType [4]byte
	CRC uint16
}

func (fit *FitFile) ReadFileHeader() error {
  err := binary.Read(fit.fp, binary.LittleEndian, fit.FileHeader)
  if err != nil {
  	return err
  }
  if string(fit.FileHeader.DataType[:]) != ".FIT" {
  	return errors.New("Unrecognized FIT file format")
  }
  return nil
}
