package fitparse

import(
	"encoding/binary"
)

func (fit *FitFile) ReadRecordHeader() (*RecordHeader, error) {
  var rh_byte uint8
  err := binary.Read(fit.fp, binary.LittleEndian, &rh_byte)
  if err != nil {
    return nil, err
  }
  return NewRecordHeader(rh_byte), nil
}

func NewRecordHeader(b uint8) *RecordHeader {
	return &RecordHeader{
		LocalMessageType: b&15,
		MessageType: bool(b>>6&1 == 1),
		NormalHeader: bool(b>>7&1 == 0),
	}
}

type RecordHeader struct {
	LocalMessageType uint8
	MessageType bool   // True = Definition Message, False = Data Message
	NormalHeader bool  // True if is a Normal Header (bit == 0)
}
