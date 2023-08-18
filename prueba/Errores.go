package prueba

type Error struct {
	mensaje string
}

func (l *Error) Agregar(message string) Error {
	return Error{mensaje: message}
}

func (l Error) GetMessage() string {
	return l.mensaje
}
