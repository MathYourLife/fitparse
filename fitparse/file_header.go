package fitparse

import(
	"errors"
	"encoding/binary"
)

// Flexible & Interoperable Data Transfer (FIT) Protocol Rev 1.7
// Table 3-1. Byte Description of File Header
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
  // ASCII values for “.FIT”. A FIT binary file opened with a text editor will
	// contain a readable “.FIT” in the first line.
  if string(fit.FileHeader.DataType[:]) != ".FIT" {
  	return errors.New("Unrecognized FIT file format")
  }
  return nil
}
