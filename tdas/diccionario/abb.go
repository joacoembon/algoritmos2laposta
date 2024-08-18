package diccionario

import (
	"tdas/pila"
)

type nodoAbb[K comparable, V any] struct {
	izquierdo *nodoAbb[K, V]
	derecho   *nodoAbb[K, V]
	clave     K
	dato      V
}

type abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cantidad int
	cmp      func(K, K) int
}

type iteradorAbbRango[K comparable, V any] struct {
	desde    *K
	hasta    *K
	pila     pila.Pila[*nodoAbb[K, V]]
	comparar func(K, K) int
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{raiz: nil, cantidad: 0, cmp: funcion_cmp}
}

func crearNodoAbb[K comparable, V any](clave K, dato V) *nodoAbb[K, V] {
	return &nodoAbb[K, V]{izquierdo: nil, derecho: nil, clave: clave, dato: dato}
}

func (abb *abb[K, V]) buscar(clave K) **nodoAbb[K, V] {
	var nodo **nodoAbb[K, V] = &abb.raiz
	return abb.buscarNodoPtr(clave, nodo)
}
func (abb *abb[K, V]) buscarNodoPtr(clave K, nodo **nodoAbb[K, V]) **nodoAbb[K, V] {
	//llama recursivamente hasta llegar al caso base que es el espacio donde esta el nodo o donde deberia estar
	//devolviendo un puntero al puntero del mismo.
	if *nodo == nil {
		return nodo
	}
	comparacion := abb.cmp(clave, (*nodo).clave)
	if comparacion < 0 {
		return abb.buscarNodoPtr(clave, &(*nodo).izquierdo)
	} else if comparacion > 0 {
		return abb.buscarNodoPtr(clave, &(*nodo).derecho)
	}
	return nodo
}

func (abb *abb[K, V]) Guardar(clave K, dato V) {
	//Lo primero que hace es verificar si encontro el nodo y le asigna su nuevo valor, si no lo hace
	//pasa a la siguiente opcion que crea un nuevo nodo y lo asigna al lugar que corresponde.
	nodoEncontrado := abb.buscar(clave)

	if nodoEncontrado != nil && *nodoEncontrado != nil {
		(*nodoEncontrado).dato = dato
	} else {
		nuevoNodo := crearNodoAbb(clave, dato)
		*nodoEncontrado = nuevoNodo
		abb.cantidad++
	}
}

func (abb *abb[K, V]) Pertenece(clave K) bool {
	return (*abb.buscar(clave)) != nil
}

func (abb *abb[K, V]) Obtener(clave K) V {
	nodo := abb.buscar(clave)
	if (*nodo) == nil {
		panic("La clave no pertenece al diccionario")
	}
	return (*nodo).dato
}

func (abb *abb[K, V]) Borrar(clave K) V {
	//Contempla los 3 casos bordes. Borrar un nodo sin hijos, con uno o con dos.
	//Cuando se trata de la raiz, es otro caso borde por el cual implica un if.
	nodo := abb.buscar(clave)
	if (*nodo) == nil {
		panic("La clave no pertenece al diccionario")
	}
	direccion := nodo
	dato := (*direccion).dato
	if (*direccion).izquierdo == nil && (*direccion).derecho == nil {
		(*direccion) = nil
	} else if (*direccion).izquierdo != nil && (*direccion).derecho == nil {
		(*direccion) = (*direccion).izquierdo
	} else if (*direccion).izquierdo == nil && (*direccion).derecho != nil {
		(*direccion) = (*direccion).derecho
	} else {
		direccionNodoReemplazo := abb.buscarNodoReemplazo(&(*nodo).izquierdo)
		nodoReemplazo := (*direccionNodoReemplazo)
		datoNuevo := (*nodoReemplazo).dato
		claveNueva := (*nodoReemplazo).clave

		(*direccionNodoReemplazo) = (*direccionNodoReemplazo).izquierdo

		(*direccion).dato = datoNuevo
		(*direccion).clave = claveNueva
	}
	abb.cantidad--
	return dato
}

func (abb *abb[K, V]) buscarNodoReemplazo(nodo **nodoAbb[K, V]) **nodoAbb[K, V] {
	//Funcion que busca el nodo que va a reemplazar al nodo borrado (es el caso en el que
	//se quiere borrar un nodo con dos hijos)
	if (*nodo).derecho == nil {
		return nodo
	}
	return abb.buscarNodoReemplazo(&(*nodo).derecho)
}

func (abb *abb[K, V]) Cantidad() int {
	return abb.cantidad
}

func (abb *abb[K, V]) Iterar(funcion func(clave K, dato V) bool) {
	abb.IterarRango(nil, nil, funcion)
}

func (abb *abb[K, V]) Iterador() IterDiccionario[K, V] {
	return abb.IteradorRango(nil, nil)
}

func (abb *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	if abb.raiz == nil {
		return
	}
	abb.iterarRango(abb.raiz, desde, hasta, visitar)
}

func (abb *abb[K, V]) iterarRango(nodo *nodoAbb[K, V], desde *K, hasta *K, visitar func(clave K, dato V) bool) bool {

	compararDesde, compararHasta := compararDesdeYHasta[K, V](nodo, desde, hasta, abb.cmp)

	if compararDesde > 0 && nodo.izquierdo != nil {
		if !abb.iterarRango(nodo.izquierdo, desde, hasta, visitar) {
			return false
		}
	}
	if compararDesde >= 0 && compararHasta <= 0 {
		if !visitar(nodo.clave, nodo.dato) {
			return false
		}
	}
	if compararHasta < 0 && nodo.derecho != nil {
		if !abb.iterarRango(nodo.derecho, desde, hasta, visitar) {
			return false
		}
	}
	return true
}

func (abb *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	pila := pila.CrearPilaDinamica[*nodoAbb[K, V]]()
	iter := &iteradorAbbRango[K, V]{desde: desde, hasta: hasta, pila: pila, comparar: abb.cmp}
	iter.buscarPrimero(abb.raiz)
	return iter
}

func (iter *iteradorAbbRango[K, V]) HaySiguiente() bool {
	return !iter.pila.EstaVacia()
}

func (iter *iteradorAbbRango[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return iter.pila.VerTope().clave, iter.pila.VerTope().dato
}

func (iter *iteradorAbbRango[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	nodoActual := iter.pila.Desapilar()
	iter.buscarIzquierdos(nodoActual.derecho)
}

func compararDesdeYHasta[K comparable, V any](nodo *nodoAbb[K, V], desde, hasta *K, comparar func(K, K) int) (int, int) {
	compararDesde := 1
	compararHasta := -1
	if desde != nil {
		compararDesde = comparar(nodo.clave, *(desde))
	}
	if hasta != nil {
		compararHasta = comparar(nodo.clave, *(hasta))
	}
	return compararDesde, compararHasta
}

func (iter *iteradorAbbRango[K, V]) buscarIzquierdos(nodo *nodoAbb[K, V]) {
	if nodo == nil {
		return
	}
	compararDesde, compararHasta := compararDesdeYHasta(nodo, iter.desde, iter.hasta, iter.comparar)

	if compararDesde >= 0 && compararHasta <= 0 {
		iter.pila.Apilar(nodo)
	}
	if compararDesde < 0 {
		iter.buscarIzquierdos(nodo.derecho)
	}
	iter.buscarIzquierdos(nodo.izquierdo)
}

func (iter *iteradorAbbRango[K, V]) buscarPrimero(nodo *nodoAbb[K, V]) {
	if nodo == nil {
		return
	}
	compararDesde, _ := compararDesdeYHasta[K, V](nodo, iter.desde, iter.hasta, iter.comparar)

	if compararDesde >= 0 {
		iter.buscarIzquierdos(nodo)
	} else if compararDesde < 0 {
		iter.buscarPrimero(nodo.derecho)
	}
}
