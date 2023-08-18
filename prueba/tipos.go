package prueba

const (
	STRING = iota
	FLOAT
	INT
	BOOL
)

func Str(tipo int) string {
	switch tipo {
	case 0:
		return "String"
	case 1:
		return "Float"
	case 2:
		return "Int"
	case 3:
		return "Bool"
	}
	return ""
}
