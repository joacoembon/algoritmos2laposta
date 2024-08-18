package pila_test

import (
	"tdas/pila"
	"testing"

	"github.com/stretchr/testify/require"
)

const limite int = 1000

func TestPilaVacia(t *testing.T) {
	pila := pila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia(), "La pila no está vacía después de ser creada")
}

func TestApilarDesapilar(t *testing.T) {
	p := pila.CrearPilaDinamica[int]()
	p.Apilar(1)
	p.Apilar(2)
	require.Equal(t, 2, p.Desapilar(), "Desapilado incorrecto")
	require.Equal(t, 1, p.Desapilar(), "Desapilado incorrecto")
}

func TestPilaVolumen(t *testing.T) {
	pila := pila.CrearPilaDinamica[int]()
	for i := 0; i < limite; i++ {
		pila.Apilar(i)
	}
	for i := limite - 1; i >= 0; i-- {
		require.Equal(t, i, pila.Desapilar(), "Desapilado incorrecto")
	}
}

func TestCondicionBordeDesapilarVacia(t *testing.T) {
	pila := pila.CrearPilaDinamica[int]()
	require.Panics(t, func() { pila.Desapilar() }, "Desapilar en pila vacia no produce pánico")
}

func TestCondicionBordeVerTopeVacia(t *testing.T) {
	pila := pila.CrearPilaDinamica[int]()
	require.Panics(t, func() { pila.VerTope() }, "VerTope en pila vacia no produce pánico")
}

func TestCondicionBordeEstaVacia(t *testing.T) {
	pila := pila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia(), "Pila no está vacía después de desapilar en pila vacia")
}

func TestCondicionBordeInvalidoDesapilar(t *testing.T) {
	pila := pila.CrearPilaDinamica[int]()
	pila.Apilar(1)
	pila.Desapilar()
	require.Panics(t, func() { pila.Desapilar() }, "Desapilar en pila vacia no produce pánico")
}

func TestCondicionBordeInvalidoVerTope(t *testing.T) {
	pila := pila.CrearPilaDinamica[int]()
	pila.Apilar(1)
	pila.Desapilar()
	require.Panics(t, func() { pila.VerTope() }, "VerTope en pila vacia no produce pánico")
}

func TestApilarDiferentesTipos(t *testing.T) {
	pilaStr := pila.CrearPilaDinamica[string]()
	pilaInt := pila.CrearPilaDinamica[int]()

	pilaStr.Apilar("Hello")
	pilaInt.Apilar(42)

	require.Equal(t, "Hello", pilaStr.Desapilar(), "Desapilado incorrecto en pila de cadenas")
	require.Equal(t, 42, pilaInt.Desapilar(), "Desapilado incorrecto en pila de enteros")
}
