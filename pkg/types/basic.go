package types

import stdtypes "go/types"

type Basic struct {
	Kind      stdtypes.BasicKind
	IsPointer bool
}

func (t *Basic) Name() string {
	return stdtypes.Typ[t.Kind].Name()
}

func (t *Basic) BitSize() int {
	switch t.Kind {
	case stdtypes.Int8, stdtypes.Uint8:
		return 8
	case stdtypes.Int16, stdtypes.Uint16:
		return 16
	case stdtypes.Int32, stdtypes.Float32, stdtypes.Uint32:
		return 32
	default: // for types.Int, types.Uint, types.Float64, types.Uint64, types.Int64 and other.
		return 64
	}
}

func (t *Basic) IsString() bool {
	return t.Kind == stdtypes.String
}

func (t *Basic) IsNumeric() bool {
	switch t.Kind {
	default:
		return false
	case stdtypes.Uint,
		stdtypes.Uint8,
		stdtypes.Uint16,
		stdtypes.Uint32,
		stdtypes.Uint64,
		stdtypes.Int,
		stdtypes.Int8,
		stdtypes.Int16,
		stdtypes.Int32,
		stdtypes.Int64,
		stdtypes.Float32,
		stdtypes.Float64:
		return true
	}
}

func (t *Basic) IsInteger() bool {
	return t.IsSigned() || t.IsUnsigned()
}

func (t *Basic) IsSigned() bool {
	switch t.Kind {
	case stdtypes.Int, stdtypes.Int8, stdtypes.Int16, stdtypes.Int32, stdtypes.Int64:
		return true
	}
	return false
}

func (t *Basic) IsInt() bool {
	return t.Kind == stdtypes.Int
}

func (t *Basic) IsInt8() bool {
	return t.Kind == stdtypes.Int8
}

func (t *Basic) IsInt16() bool {
	return t.Kind == stdtypes.Int16
}

func (t *Basic) IsInt32() bool {
	return t.Kind == stdtypes.Int32
}

func (t *Basic) IsInt64() bool {
	return t.Kind == stdtypes.Int64
}

func (t *Basic) IsUnsigned() bool {
	switch t.Kind {
	case stdtypes.Uint, stdtypes.Uint8, stdtypes.Uint16, stdtypes.Uint32, stdtypes.Uint64:
		return true
	}
	return false
}

func (t *Basic) IsUint() bool {
	return t.Kind == stdtypes.Uint
}

func (t *Basic) IsUint8() bool {
	return t.Kind == stdtypes.Uint8
}

func (t *Basic) IsUint16() bool {
	return t.Kind == stdtypes.Uint16
}

func (t *Basic) IsUint32() bool {
	return t.Kind == stdtypes.Uint32
}

func (t *Basic) IsUint64() bool {
	return t.Kind == stdtypes.Uint64
}

func (t *Basic) IsFloat() bool {
	switch t.Kind {
	case stdtypes.Float32, stdtypes.Float64, stdtypes.Complex128:
		return true
	}
	return false
}

func (t *Basic) IsFloat32() bool {
	return t.Kind == stdtypes.Float32
}

func (t *Basic) IsFloat64() bool {
	return t.Kind == stdtypes.Float64
}

func (t *Basic) IsBool() bool {
	return t.Kind == stdtypes.Bool
}

func (t *Basic) IsByte() bool {
	return t.Kind == stdtypes.Byte
}

var BasicTyp = []*Basic{
	stdtypes.Bool:       {stdtypes.Bool, false},
	stdtypes.Int:        {stdtypes.Int, false},
	stdtypes.Int8:       {stdtypes.Int8, false},
	stdtypes.Int16:      {stdtypes.Int16, false},
	stdtypes.Int32:      {stdtypes.Int32, false},
	stdtypes.Int64:      {stdtypes.Int64, false},
	stdtypes.Uint:       {stdtypes.Uint, false},
	stdtypes.Uint8:      {stdtypes.Uint8, false},
	stdtypes.Uint16:     {stdtypes.Uint16, false},
	stdtypes.Uint32:     {stdtypes.Uint32, false},
	stdtypes.Uint64:     {stdtypes.Uint64, false},
	stdtypes.Uintptr:    {stdtypes.Uintptr, false},
	stdtypes.Float32:    {stdtypes.Float32, false},
	stdtypes.Float64:    {stdtypes.Float64, false},
	stdtypes.Complex64:  {stdtypes.Complex64, false},
	stdtypes.Complex128: {stdtypes.Complex128, false},
	stdtypes.String:     {stdtypes.String, false},
}
