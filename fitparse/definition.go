/*
Describes the architecture, format, and fields of upcoming data messages
*/

package fitparse

import(
	"encoding/binary"
)

// Flexible & Interoperable Data Transfer (FIT) Protocol Rev 1.7
// Table 4-3. Definition Message Contents
func (fit *FitFile) ReadDefinition() (*Definition, error) {
	var err error
  def := &Definition{}

	err = binary.Read(fit.fp, binary.LittleEndian, &def.Reserved)
	if err != nil {
		return nil, err
	}
	err = binary.Read(fit.fp, binary.LittleEndian, &def.Architecture)
	if err != nil {
		return nil, err
	}
	if def.Architecture == DEFENDIAN_LITTLE {
		err = binary.Read(fit.fp, binary.LittleEndian, &def.GlobalMessageNum)
	} else {
		err = binary.Read(fit.fp, binary.BigEndian, &def.GlobalMessageNum)
	}
	if err != nil {
		return nil, err
	}
	err = binary.Read(fit.fp, binary.LittleEndian, &def.FieldCount)
	if err != nil {
		return nil, err
	}

	// Flexible & Interoperable Data Transfer (FIT) Protocol Rev 1.7
	// Table 4-4. Field Definition Contents
	for i := uint8(0); i < def.FieldCount; i++ {
		var def_field DefinitionField
		err = binary.Read(fit.fp, binary.LittleEndian, &def_field)
		if err != nil {
			return nil, err
		}
		def.Fields = append(def.Fields, def_field)
	}

	return def, nil
}

type DefEndian int8
const(
	DEFENDIAN_LITTLE DefEndian = 0
	DEFENDIAN_BIG DefEndian = 1
)
var DefEndian_name = map[int8]string{
	0: "DEFENDIAN_LITTLE",
	1: "DEFENDIAN_BIG",
}
var DefEndian_value = map[string]int8{
	"DEFENDIAN_LITTLE": 0,
	"DEFENDIAN_BIG": 1,
}
func (x DefEndian) Enum() *DefEndian {
	p := new(DefEndian)
	*p = x
	return p
}
func (x DefEndian) String() string {
	return DefEndian_name[int8(x)]
}

type Definition struct {
	Reserved byte
	Architecture DefEndian
	GlobalMessageNum uint16
	FieldCount uint8
	Fields []DefinitionField
}

type DefinitionField struct {
	FieldDefinitionNumber uint8
	Size uint8
	BaseType uint8
}