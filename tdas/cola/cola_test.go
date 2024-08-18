package cola_test

import (
	"tdas/cola"
	"testing"

	"github.com/stretchr/testify/require"
)

const volumen int = 1000

func TestCrearColaVacia(t *testing.T) {
	c := cola.CrearColaEnlazada[int]()
	require.True(t, c.EstaVacia(), "La cola no está vacía después de ser creada")
}

func TestEncolarDesencolar(t *testing.T) {
	c := cola.CrearColaEnlazada[int]()
	c.Encolar(1)
	c.Encolar(2)
	require.Equal(t, 1, c.Desencolar(), "Desencolado incorrecto")
	require.Equal(t, 2, c.Desencolar(), "Desencolado incorrecto")
}

func TestVolumen(t *testing.T) {
	c := cola.CrearColaEnlazada[int]()
	for i := 0; i < volumen; i++ {
		c.Encolar(i)
	}
	for i := 0; i < volumen; i++ {
		require.Equal(t, i, c.Desencolar(), "Desencolado incorrecto")
	}
}

func TestCondicionBordeVerPrimeroVacia(t *testing.T) {
	c := cola.CrearColaEnlazada[int]()
	require.Panics(t, func() { c.VerPrimero() }, "VerPrimero en cola vacía no produce pánico")
}

func TestCondicionBordeEstaVacia(t *testing.T) {
	c := cola.CrearColaEnlazada[int]()
	require.True(t, c.EstaVacia(), "Cola no está vacía después de desencolar en cola vacía")
}

func TestCondicionBordeInvalidoDesencolar(t *testing.T) {
	c := cola.CrearColaEnlazada[int]()
	c.Encolar(1)
	c.Desencolar()
	require.Panics(t, func() { c.Desencolar() }, "Desencolar en cola vacía no produce pánico")
}

func TestCondicionBordeInvalidoVerPrimero(t *testing.T) {
	c := cola.CrearColaEnlazada[int]()
	c.Encolar(1)
	c.Desencolar()
	require.Panics(t, func() { c.VerPrimero() }, "VerPrimero en cola vacía no produce pánico")
}

func TestEncolarDiferentesTipos(t *testing.T) {
	colaStr := cola.CrearColaEnlazada[string]()
	colaInt := cola.CrearColaEnlazada[int]()

	colaStr.Encolar("Hola")
	colaInt.Encolar(42)

	require.Equal(t, "Hola", colaStr.Desencolar(), "Desencolado incorrecto en cola de cadenas")
	require.Equal(t, 42, colaInt.Desencolar(), "Desencolado incorrecto en cola de enteros")
}
