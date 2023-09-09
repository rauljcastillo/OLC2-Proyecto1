package prueba

type Ambiente struct {
	anterior  *Ambiente
	variables map[string]*Variable
	functions map[string]Funcion
}

func NewAmbiente(ant *Ambiente) *Ambiente {
	return &Ambiente{anterior: ant, variables: make(map[string]*Variable), functions: make(map[string]Funcion)}
}

// Guarda variables
func (l *Ambiente) save(tipo int, id string, value any, nametemp string) bool {
	if _, ok := l.variables[id]; !ok {
		l.variables[id] = &Variable{tipo: tipo, id: id, val: value, nametem: nametemp}
		return true
	}
	return false
}

func (l *Ambiente) saveFunctions(tipe int, id string, params any, block any) bool {
	if _, ok := l.functions[id]; !ok {
		l.functions[id] = Funcion{tipo: tipe, ID: id, Params: params, Block: block}
		return true
	}
	return false
}

func (l *Ambiente) getFuction(id string) interface{} {
	for l != nil {
		if valor, ok := l.functions[id]; ok {
			return valor
		}
		l = l.anterior
	}

	return ""
}

func (l *Ambiente) getVar(id string) interface{} {
	for l != nil {
		if valor, ok := l.variables[id]; ok {
			return Valor{val: valor.val, tipo: valor.tipo, temporal: valor.nametem}
		}
		l = l.anterior
	}

	return 0
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
