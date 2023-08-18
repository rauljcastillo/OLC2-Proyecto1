package prueba

type Ambiente struct {
	anterior  *Ambiente
	variables map[string]*Variable
}

func NewAmbiente(ant *Ambiente) *Ambiente {
	return &Ambiente{anterior: ant, variables: make(map[string]*Variable)}
}

func (l *Ambiente) save(tipo int, id string, value any) {
	if _, ok := l.variables[id]; !ok {
		l.variables[id] = &Variable{tipo: tipo, id: id, val: value}
	}

}

func (l *Ambiente) getVar(id string) interface{} {
	for l != nil {
		if valor, ok := l.variables[id]; ok {
			return Valor{val: valor.val, tipo: valor.tipo}
		}
		l = l.anterior
	}

	return Valor{}
}

func (l *Ambiente) updateVar(id string, newVal any) {
	for l != nil {
		if mapita, ok := l.variables[id]; ok {
			mapita.val = newVal
			break
		}
		l = l.anterior
	}
}
