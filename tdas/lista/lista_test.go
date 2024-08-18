package lista_test

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

const NUM0 int = 0
const NUM1 int = 1
const NUM2 int = 2
const NUM3 int = 3
const NUM4 int = 4
const NUM5 int = 5
const NUM6 int = 6
const NUM12 int = 12
const VOLUMEN int = 10000

func TestListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())
}

func TestListaConElementosNoEstaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	palabras := []string{"lo", "que", "se", "me", "ocurrio"}
	for _, palabra := range palabras {
		lista.InsertarPrimero(palabra)
	}
	require.False(t, lista.EstaVacia())
}

func TestLargo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	numeros := []int{NUM1, NUM2, NUM3}
	for i := range numeros {
		lista.InsertarPrimero(numeros[i])
	}
	require.Equal(t, lista.Largo(), len(numeros))
}

func TestVolumen(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	volumen := make([]int, VOLUMEN)
	for i := range volumen {
		lista.InsertarPrimero(i)
	}
	require.Equal(t, len(volumen)-1, lista.VerPrimero())
	require.Equal(t, volumen[0], lista.VerUltimo())
	require.Equal(t, len(volumen), lista.Largo())
	for i := range volumen {
		dato := lista.BorrarPrimero()
		require.Equal(t, len(volumen)-1-i, dato)
	}
	require.True(t, lista.EstaVacia())
	for i := range volumen {
		lista.InsertarUltimo(i)
	}
	require.Equal(t, volumen[0], lista.VerPrimero())
	require.Equal(t, len(volumen)-1, lista.VerUltimo())
	require.Equal(t, len(volumen), lista.Largo())
	for range volumen {
		lista.BorrarPrimero()
	}
	for i := range volumen {
		if i%2 == 0 {
			lista.InsertarUltimo(i)
			require.Equal(t, i, lista.VerUltimo())
		} else {
			lista.InsertarPrimero(i)
			require.Equal(t, i, lista.VerPrimero())
		}
	}
	require.Equal(t, len(volumen), lista.Largo())
}

func TestMirarElPrimeroYElUltimo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	numeros := []int{NUM1, NUM2, NUM3}
	lista.InsertarPrimero(numeros[0])
	require.Equal(t, numeros[0], lista.VerPrimero())
	require.Equal(t, numeros[0], lista.VerUltimo())
	lista.InsertarPrimero(numeros[1])
	require.Equal(t, numeros[1], lista.VerPrimero())
	require.Equal(t, numeros[0], lista.VerUltimo())
	lista.InsertarUltimo(numeros[2])
	require.Equal(t, numeros[1], lista.VerPrimero())
	require.Equal(t, numeros[2], lista.VerUltimo())
}

func TestListaVaciaLevantaPanic(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() }, "No se puede ver el primero de una lista vacia")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() }, "No se puede ver el ultimo de una lista vacia")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() }, "No se puede ver el ultimo de una lista vacia")
}

func TestIteradorRecorreLista(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(NUM0)
	lista.InsertarPrimero(NUM1)
	lista.InsertarPrimero(NUM2)
	lista.InsertarPrimero(NUM3)
	lista.InsertarPrimero(NUM4)
	iter := lista.Iterador()
	for i := 0; iter.HaySiguiente(); i++ {
		require.Equal(t, NUM4-i, iter.VerActual())
		iter.Siguiente()
	}
}

// Al insertar un elemento en la posición en la que se crea el iterador, efectivamente se inserta al principio.
func TestInsertarAlPrincipio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(NUM3)
	iter := lista.Iterador()
	iter.Insertar(NUM5)
	require.Equal(t, NUM5, lista.VerPrimero(), "No se inserto correctamente")
	require.Equal(t, NUM2, lista.Largo(), "El largo no es el correcto")
}

// Insertar un elemento cuando el iterador está al final efectivamente es equivalente a insertar al final.
func TestInsertarAlFinal(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(NUM3)
	lista.InsertarPrimero(NUM2)
	lista.InsertarPrimero(NUM6)
	lista.InsertarPrimero(NUM4)
	iter := lista.Iterador()

	for iter.HaySiguiente() {
		iter.Siguiente()
	}
	iter.Insertar(NUM5)
	require.Equal(t, NUM5, lista.VerUltimo(), "No se inserto correctamente")
	require.Equal(t, NUM5, lista.Largo(), "El largo no es el correcto")
}

// Insertar un elemento en el medio se hace en la posición correcta.
func TestInsertarEnMedio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(NUM3)
	lista.InsertarPrimero(NUM2)
	lista.InsertarPrimero(NUM6)
	lista.InsertarPrimero(NUM4)
	iter := lista.Iterador()

	for iter.VerActual() != NUM2 {
		iter.Siguiente()
	}
	iter.Insertar(NUM5)
	lista.BorrarPrimero()
	lista.BorrarPrimero()
	require.Equal(t, NUM5, lista.VerPrimero(), "No se inserto correctamente")
	require.Equal(t, NUM3, lista.Largo(), "El largo no es el correcto")

}

// Al remover el elemento cuando se crea el iterador, cambia el primer elemento de la lista.
func TestRemoverPrimerElementoConIterador(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(NUM3)
	lista.InsertarPrimero(NUM2)
	lista.InsertarPrimero(NUM6)
	lista.InsertarPrimero(NUM4)
	iter := lista.Iterador()

	require.Equal(t, NUM4, lista.Largo(), "El largo no es el correcto")

	iter.Borrar()

	require.Equal(t, NUM6, lista.VerPrimero(), "No se inserto correctamente")
	require.Equal(t, NUM3, lista.Largo(), "El largo no es el correcto")

}

// Remover el último elemento con el iterador cambia el último de la lista.
func TestRemoverUltimoElementoConIterador(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(NUM3)
	lista.InsertarUltimo(NUM2)
	lista.InsertarUltimo(NUM6)
	lista.InsertarUltimo(NUM4)
	iter := lista.Iterador()

	for iter.HaySiguiente() && iter.VerActual() != lista.VerUltimo() {
		iter.Siguiente()
	}
	iter.Borrar()
	require.Equal(t, NUM6, lista.VerUltimo(), "No se inserto correctamente")
	require.Equal(t, NUM3, lista.Largo(), "El largo no es el correcto")
}

// Verificar que al remover un elemento del medio, este no está.
func TestRemoverElementoMedioConIterador(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(NUM3)
	lista.InsertarPrimero(NUM2)
	lista.InsertarPrimero(NUM5)
	lista.InsertarPrimero(NUM6)
	lista.InsertarPrimero(NUM4)
	iter := lista.Iterador()

	for iter.HaySiguiente() && iter.VerActual() != NUM5 {
		iter.Siguiente()
	}
	iter.Borrar()
	lista.BorrarPrimero()
	lista.BorrarPrimero()
	require.Equal(t, NUM2, lista.VerPrimero(), "No se inserto correctamente")
	require.Equal(t, NUM2, lista.Largo(), "El largo no es el correcto")

}

// Casos del iterador interno, incluyendo casos con corte (la función visitar devuelve false eventualmente).
func TestFuncionVisitarConIteradorInterno(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(NUM3)
	lista.InsertarPrimero(NUM2)
	lista.InsertarPrimero(NUM5)
	lista.InsertarPrimero(NUM6)
	lista.InsertarPrimero(NUM4)

	suma, cont := 0, 0
	lista.Iterar(func(v int) bool {
		if v%2 == 0 {
			suma += v
			cont += 1
		}
		if cont == 3 {
			return false
		}
		return true
	})
	require.Equal(t, NUM3, cont, "La funcion no devuelve nunca false")
	require.Equal(t, NUM12, suma, "La funcion no devuelve nunca false")
}

func TestVolumenIterarSinCorte(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	volumen := make([]int, VOLUMEN)
	for i := range volumen {
		lista.InsertarPrimero(i)
	}
	lista.Iterar(func(v int) bool {
		require.Equal(t, v, len(volumen)-(len(volumen)-v))
		return true
	})
}

func TestVolumenIterarConCorte(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	volumen := make([]int, VOLUMEN)
	for i := range volumen {
		lista.InsertarUltimo(i)
	}
	suma, cont := NUM0, NUM0
	lista.Iterar(func(v int) bool {
		if v%2 == NUM0 {
			suma++
		}
		cont++
		return cont != VOLUMEN/10
	})
	require.Equal(t, VOLUMEN/20, suma)
}

func TestVolumenIteradorSinCorte(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	volumen := make([]int, VOLUMEN)
	for i := range volumen {
		lista.InsertarUltimo(i)
	}
	iter := lista.Iterador()
	i := NUM0
	for iter.HaySiguiente() {
		require.Equal(t, iter.VerActual(), i)
		iter.Siguiente()
		i++
	}
}

func TestIteradorInsertarEnListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	require.False(t, iter.HaySiguiente())
	iter.Insertar(NUM5)

	require.True(t, iter.HaySiguiente())
	require.EqualValues(t, NUM5, lista.VerUltimo())
	require.EqualValues(t, NUM5, lista.VerPrimero())

	iter.Insertar(NUM1)
	iter.Insertar(NUM4)
	iter.Insertar(NUM3)
	iter.Insertar(NUM5)
	iter.Insertar(NUM6)
	require.EqualValues(t, NUM5, lista.VerUltimo())
	require.EqualValues(t, NUM6, lista.VerPrimero())
	require.EqualValues(t, NUM6, iter.VerActual())

	posicion := NUM0
	for iter.VerActual() != NUM4 && iter.HaySiguiente() {
		posicion++
		iter.Siguiente()
	}
	require.EqualValues(t, NUM4, iter.VerActual())
	require.EqualValues(t, NUM3, posicion)
}
