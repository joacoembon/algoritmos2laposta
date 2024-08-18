package diccionario

import (
	"fmt"
	"hash/fnv"
)

type estadoParClaveDato int

const (
	VACIO = estadoParClaveDato(iota)
	BORRADO
	OCUPADO
)

const (
	MARCADORBUSQUEDA = 1
	REDIMENSION      = 2
	CAPACIDADINICIAL = 7
	CAPACIDADMIN     = 20
	CAPACIDADMAX     = 70
	PORCENTAJE       = 100
)

type celdaHash[K comparable, V any] struct {
	clave  K
	dato   V
	estado estadoParClaveDato
}

type hashCerrado[K comparable, V any] struct {
	tabla    []celdaHash[K, V]
	cantidad int
	borrados int
}

type iterHashCerrado[K comparable, V any] struct {
	hash   *hashCerrado[K, V]
	actual int
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	tabla := make([]celdaHash[K, V], CAPACIDADINICIAL)
	return &hashCerrado[K, V]{tabla: tabla, cantidad: 0, borrados: 0}
}

func (hash *hashCerrado[K, V]) Guardar(clave K, dato V) {
	if (hash.cantidad + hash.borrados) > (CAPACIDADMAX * len(hash.tabla) / PORCENTAJE) {
		nueva_capacidad := len(hash.tabla) * REDIMENSION
		hash.rehash(nueva_capacidad)
	}
	posicion := hash.buscador(clave)

	if hash.tabla[posicion].clave == clave && hash.tabla[posicion].estado == OCUPADO {
		//Si se guardo un espacio vacio lo guarda con OCUPADO
		hash.tabla[posicion].dato = dato
	} else {
		hash.tabla[posicion].clave = clave
		hash.tabla[posicion].dato = dato
		hash.tabla[posicion].estado = OCUPADO
		hash.cantidad++
	}
}

func (hash *hashCerrado[K, V]) Pertenece(clave K) bool {
	posicion := hash.buscador(clave)
	return hash.cantidad > 0 && hash.tabla[posicion].clave == clave
}

func (hash *hashCerrado[K, V]) Obtener(clave K) V {
	posicion := hash.buscador(clave)

	if hash.tabla[posicion].clave != clave || hash.cantidad == 0 {
		panic("La clave no pertenece al diccionario")
	}
	return hash.tabla[posicion].dato
}

func (hash *hashCerrado[K, V]) Borrar(clave K) V {
	posicion := hash.buscador(clave)

	if hash.tabla[posicion].clave != clave || hash.cantidad == 0 {
		panic("La clave no pertenece al diccionario")
	}
	hash.tabla[posicion].estado = BORRADO
	hash.cantidad--
	hash.borrados++
	dato_borrado := hash.tabla[posicion].dato

	if (hash.cantidad + hash.borrados) < (CAPACIDADMIN * len(hash.tabla) / PORCENTAJE) {
		nueva_capacidad := len(hash.tabla) / REDIMENSION
		hash.rehash(nueva_capacidad)
	}
	return dato_borrado
}

func (hash *hashCerrado[K, V]) Cantidad() int {
	return hash.cantidad
}

func (hash *hashCerrado[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	if hash.cantidad == 0 {
		return
	}
	for posActual := 0; posActual < len(hash.tabla); posActual++ {
		if hash.tabla[posActual].estado != OCUPADO {
			continue
		}
		if !visitar(hash.tabla[posActual].clave, hash.tabla[posActual].dato) {
			return
		}
	}
}

func (hash *hashCerrado[K, V]) Iterador() IterDiccionario[K, V] {
	pos_actual := 0
	if hash.cantidad == 0 {
		pos_actual = len(hash.tabla)
	} else {
		for hash.tabla[pos_actual].estado != OCUPADO && pos_actual <= len(hash.tabla) {
			pos_actual++
		}
	}
	return &iterHashCerrado[K, V]{hash: hash, actual: pos_actual}

}

func (iter *iterHashCerrado[K, V]) HaySiguiente() bool {
	return iter.actual < len(iter.hash.tabla)
}

func (iter *iterHashCerrado[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return iter.hash.tabla[iter.actual].clave, iter.hash.tabla[iter.actual].dato
}

func (iter *iterHashCerrado[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	iter.actual++
	if iter.actual == len(iter.hash.tabla) {
		return
	}
	for iter.hash.tabla[iter.actual].estado != OCUPADO {
		iter.actual++
		if iter.actual == len(iter.hash.tabla) {
			return
		}
	}
}

func (hash *hashCerrado[K, V]) hash(clave K) int {
	hf := fnv.New32()
	hf.Write([]byte(fmt.Sprintf("%v", clave)))
	return int(hf.Sum32()) % len(hash.tabla)
}

func (hash *hashCerrado[K, V]) buscador(clave K) int {
	// Esta funcion busca la clave y devuelve la posicion de la clave o un espacio vacio.
	posicion := hash.hash(clave)

	for (hash.tabla[posicion].estado != VACIO && hash.tabla[posicion].clave != clave) || hash.tabla[posicion].estado == BORRADO {
		//Sigue buscando en el caso que encuentre uno con la misma clave pero BORRADO
		posicion = (posicion + MARCADORBUSQUEDA) % len(hash.tabla)
	}
	return posicion
}

func (hash *hashCerrado[K, V]) rehash(tamanio int) {
	nuevaTabla := make([]celdaHash[K, V], tamanio)
	var tablaVieja []celdaHash[K, V]
	tablaVieja, hash.tabla = hash.tabla, nuevaTabla

	for _, celda := range tablaVieja {
		if celda.estado == OCUPADO {
			posicion := hash.buscador(celda.clave)
			nuevaTabla[posicion].clave = celda.clave
			nuevaTabla[posicion].dato = celda.dato
			nuevaTabla[posicion].estado = OCUPADO
		}
	}
	hash.borrados = 0
}
