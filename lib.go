package primitivedefaultvaluego

import (
	"errors"
	"go/ast"
)

// defaultValue returns the default value for a given Go primitive type.
func DefaultValue(fieldType ast.Expr) (string, error) {
	switch ft := fieldType.(type) {
	case *ast.Ident:
		switch ft.Name {
		case "string":
			return `""`, nil
		case "int", "int8", "int16", "int32", "int64",
			"uint", "uint8", "uint16", "uint32", "uint64", "byte", "rune":
			return "0", nil
		case "float32", "float64":
			return "0.0", nil
		case "bool":
			return "false", nil
		case "complex64", "complex128":
			return "0 + 0i", nil
		default:
			// Reference types (slices, maps, pointers, channels) and function types default to nil.
			// For other types, such as structs or arrays, adjust as necessary.
			return "nil", nil
		}
	}
	return "", errors.New("unknown type")

}
