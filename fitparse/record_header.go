package fitparse

import(
	"encoding/binary"
)

type RecordHeader struct {
	LocalMessageId uint8
	Type RecordHeaderType
	DataMsgType DataMsgType
}

func NewRecordHeader(b uint8) *RecordHeader {
	var msg_type RecordHeaderType
	if bool(b>>6&1 == 1) {
		msg_type = RECORDHEADER_DEFINITION
	} else {
		msg_type = RECORDHEADER_DATA
	}
	var data_msg_type DataMsgType
	if bool(b>>7&1 == 0) {
		data_msg_type = DATAMSGTYPE_NORMAL
	} else {
		data_msg_type = DATAMSGTYPE_COMPTS
	}
	return &RecordHeader{
		LocalMessageId: b&15,
		Type: msg_type,
		DataMsgType: data_msg_type,
	}
}

func (fit *FitFile) ReadRecordHeader() (*RecordHeader, error) {
  var rh_byte uint8
  err := binary.Read(fit.fp, binary.LittleEndian, &rh_byte)
  if err != nil {
    return nil, err
  }
  return NewRecordHeader(rh_byte), nil
}

/*
There are two kinds of data records:
Definition Messages – define the upcoming data messages. A definition message
will define a local message type and associate it to a specific FIT message,
and then designate the byte alignment and field contents of the upcoming data
message.
Data Messages – contain a local message type and populated data fields in the
format described by the preceding definition message.
*/
type RecordHeaderType int8
const(
	RECORDHEADER_DATA RecordHeaderType = 0
	RECORDHEADER_DEFINITION RecordHeaderType = 1
)
var RecordHeaderType_name = map[int8]string{
	0: "RECORDHEADER_DATA",
	1: "RECORDHEADER_DEFINITION",
}
var RecordHeaderType_value = map[string]int8{
	"RECORDHEADER_DATA": 0,
	"RECORDHEADER_DEFINITION": 1,
}
func (x RecordHeaderType) Enum() *RecordHeaderType {
	p := new(RecordHeaderType)
	*p = x
	return p
}
func (x RecordHeaderType) String() string {
	return RecordHeaderType_name[int8(x)]
}

/*
The definition message and its associated data messages will have matching
local message types. There are two types of data message:
- Normal Data Message
- Compressed Timestamp Data Message
*/
type DataMsgType int8
const(
	DATAMSGTYPE_NORMAL DataMsgType = 0
	DATAMSGTYPE_COMPTS DataMsgType = 1
)
var DataMsgType_name = map[int8]string{
	0: "DATAMSGTYPE_NORMAL",
	1: "DATAMSGTYPE_COMPTS",
}
var DataMsgType_value = map[string]int8{
	"DATAMSGTYPE_NORMAL": 0,
	"DATAMSGTYPE_COMPTS": 1,
}
func (x DataMsgType) Enum() *DataMsgType {
	p := new(DataMsgType)
	*p = x
	return p
}
func (x DataMsgType) String() string {
	return DataMsgType_name[int8(x)]
}
