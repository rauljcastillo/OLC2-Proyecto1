package prueba

var a = [4][4]int{
	{STRING, 4, 4, 4},
	{4, FLOAT, FLOAT, 4},
	{4, FLOAT, INT, 4},
	{4, 4, 4, BOOL},
}

/*
STRING = iota
FLOAT
INT
BOOL
*/
func TablaT(tipoA, tipoB int) int {
	//String String
	return a[tipoA][tipoB]
}
