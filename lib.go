package primitivedefaultvaluego

// defaultValue returns the default value for a given Go primitive type.
func DefaultValue(fieldType string) string {
	switch fieldType {
	case "string":
		return `""`
	case "int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64", "byte", "rune":
		return "0"
	case "float32", "float64":
		return "0.0"
	case "bool":
		return "false"
	case "complex64", "complex128":
		return "0 + 0i"
	default:
		// Reference types (slices, maps, pointers, channels) and function types default to nil.
		// For other types, such as structs or arrays, adjust as necessary.
		return "nil"
	}
}
