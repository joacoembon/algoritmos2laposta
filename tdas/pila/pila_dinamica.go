package pila

/* Definición del struct pila proporcionado por la cátedra. */

type pilaDinamica[T any] struct {
	datos     []T
	cantidad  int
	capacidad int
}

const capacidadInicial int = 10 // Capacidad inicial de la pila
const redimension int = 2       // Redimension a elegir, en este caso 2

// CrearPilaDinamica crea una nueva pila dinámica vacía.
func CrearPilaDinamica[T any]() Pila[T] {
	return &pilaDinamica[T]{
		datos:     make([]T, capacidadInicial),
		cantidad:  0,
		capacidad: capacidadInicial,
	}
}

func (p *pilaDinamica[T]) EstaVacia() bool {
	return p.cantidad == 0
}

func (p *pilaDinamica[T]) VerTope() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}
	return p.datos[p.cantidad-1]
}

func (p *pilaDinamica[T]) Apilar(valor T) {
	if p.cantidad == p.capacidad {
		p.redimensionarCapacidad(p.capacidad * redimension)
	}
	p.datos[p.cantidad] = valor
	p.cantidad++
}

func (p *pilaDinamica[T]) Desapilar() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}

	// Reducir la capacidad si la cantidad es menor o igual a la capacidad dividida por 4.
	if p.cantidad <= p.capacidad/(redimension*redimension) {
		p.redimensionarCapacidad(p.capacidad / redimension)
	}
	elemento := p.datos[p.cantidad-1]
	p.cantidad--

	return elemento
}

func (p *pilaDinamica[T]) redimensionarCapacidad(nuevaCapacidad int) {
	nuevoArray := make([]T, nuevaCapacidad)
	copy(nuevoArray, p.datos)
	p.datos = nuevoArray
	p.capacidad = nuevaCapacidad
}
