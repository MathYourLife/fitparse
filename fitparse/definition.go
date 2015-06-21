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
	err = binary.Read(fit.fp, binary.LittleEndian, &def.GlobalMessageNum)
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

type Definition struct {
	Reserved byte
	Architecture byte
	GlobalMessageNum uint16
	FieldCount uint8
	Fields []DefinitionField
}

type DefinitionField struct {
	FieldDefinitionNumber uint8
	Size uint8
	BaseType uint8
}