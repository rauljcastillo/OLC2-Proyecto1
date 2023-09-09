package prueba

const (
	STRING = iota
	FLOAT
	INT
	BOOL
	CHAR
	RTN
	BRK
	CONTINUE
	VOID
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

func Num(tipo string) int {
	switch tipo {
	case "break":
		return BRK
	case "continue":
		return CONTINUE
	case "return":
		return RTN
	case "Float":
		return FLOAT
	case "Int":
		return INT
	case "String":
		return STRING
	case "Bool":
		return BOOL
	}
	return 0
}
