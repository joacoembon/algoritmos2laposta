package comandos

type Comandos interface {
	//AgregarArchivo recibe un nombre de un archivo y guarda todo su contenido. Devuelve true si el archivo existe y false en caso contrario.
	AgregarArchivo(nombre_archivo string) bool

	//VerDoS devuelve un lista de strings con las ips de los ataques DoS.
	VerDoS() []string

	//VerVisitantes recibe una ip de origen y otra de final. Devuelve un array con las ips dentro de ese rango (incluyendo extremos).
	VerVisitantes(desdeIP, hastaIP string) []string

	//VerMasVisitados recibe un numero k entero. Devuelve un array ordenado de mayor a menor de la estructura urlUsos que tiene como miembros
	// URL (string) y Cant (int) con su url y la cantidad de veces visitado.
	VerMasVisitados(k int) []urlUsos
}
