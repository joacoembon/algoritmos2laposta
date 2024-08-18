package diccionario_test

import (
	"cmp"
	"fmt"
	"math/rand"
	"strings"
	TDADiccionario "tdas/diccionario"
	"testing"

	"github.com/stretchr/testify/require"
)

var TAMS_VOLUMEN = []int{12500, 25000, 50000, 100000}

var FUNCION_STRINGS = strings.Compare
var FUNCION_INT = cmp.Compare[int]

func TestDiccionarioVacio(t *testing.T) {
	t.Log("Comprueba que Diccionario vacio no tiene claves")
	dic := TDADiccionario.CrearABB[string, string](FUNCION_STRINGS)
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("A") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar("A") })
}

func TestDiccionarioClaveDefault(t *testing.T) {
	t.Log("Prueba sobre un Hash vacío que si justo buscamos la clave que es el default del tipo de dato, " +
		"sigue sin existir")
	dic := TDADiccionario.CrearABB[string, string](FUNCION_STRINGS)
	require.False(t, dic.Pertenece(""))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar("") })

	dicNum := TDADiccionario.CrearABB[int, string](FUNCION_INT)
	require.False(t, dicNum.Pertenece(0))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicNum.Obtener(0) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicNum.Borrar(0) })
}

func TestUnElement(t *testing.T) {
	t.Log("Comprueba que Diccionario con un elemento tiene esa Clave, unicamente")
	dic := TDADiccionario.CrearABB[string, int](FUNCION_STRINGS)
	dic.Guardar("A", 10)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece("A"))
	require.False(t, dic.Pertenece("B"))
	require.EqualValues(t, 10, dic.Obtener("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("B") })
}

func TestValorNulo(t *testing.T) {
	t.Log("Probamos que el valor puede ser nil sin problemas")
	dic := TDADiccionario.CrearABB[string, *int](FUNCION_STRINGS)
	clave := "Pez"
	dic.Guardar(clave, nil)
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, (*int)(nil), dic.Obtener(clave))
	require.EqualValues(t, (*int)(nil), dic.Borrar(clave))
	require.False(t, dic.Pertenece(clave))
}

func TestIteradorInternoClaves(t *testing.T) {
	t.Log("Valida que todas las claves sean recorridas (y una única vez) con el iterador interno")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	claves := []string{clave1, clave2, clave3}
	dic := TDADiccionario.CrearABB[string, *int](FUNCION_STRINGS)
	dic.Guardar(claves[0], nil)
	dic.Guardar(claves[1], nil)
	dic.Guardar(claves[2], nil)

	cs := []string{"", "", ""}
	cantidad := 0
	cantPtr := &cantidad
	anterior := clave1
	ptrAnterior := &anterior

	dic.Iterar(func(clave string, dato *int) bool {
		if FUNCION_STRINGS(*(ptrAnterior), clave) > 0 {
			return false
		}
		*ptrAnterior = clave
		cs[cantidad] = clave
		*cantPtr = *cantPtr + 1
		return true
	})

	require.EqualValues(t, 3, cantidad)
	require.NotEqualValues(t, cs[0], cs[1])
	require.NotEqualValues(t, cs[0], cs[2])
	require.NotEqualValues(t, cs[2], cs[1])
}

func TestIteradorInternoValores(t *testing.T) {
	t.Log("Valida que los datos sean recorridas correctamente (y una única vez) con el iterador interno")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	clave4 := "Burrito"
	clave5 := "Hamster"

	dic := TDADiccionario.CrearABB[string, int](FUNCION_STRINGS)
	dic.Guardar(clave1, 6)
	dic.Guardar(clave2, 2)
	dic.Guardar(clave3, 3)
	dic.Guardar(clave4, 4)
	dic.Guardar(clave5, 5)

	factorial := 1
	ptrFactorial := &factorial
	dic.Iterar(func(_ string, dato int) bool {
		*ptrFactorial *= dato
		return true
	})

	require.EqualValues(t, 720, factorial)
}

func TestIteradorInternoOrden(t *testing.T) {
	t.Log("Valida que los datos sean recorridas correctamente (y una única vez) con el iterador interno, sin recorrer datos borrados")
	clave0 := "Elefante"
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	clave4 := "Burrito"
	clave5 := "Hamster"

	dic := TDADiccionario.CrearABB[string, int](FUNCION_STRINGS)
	dic.Guardar(clave0, 7)
	dic.Guardar(clave1, 6)
	dic.Guardar(clave2, 2)
	dic.Guardar(clave3, 3)
	dic.Guardar(clave4, 4)
	dic.Guardar(clave5, 5)

	dic.Borrar(clave0)

	factorial := 1
	ptrFactorial := &factorial
	anterior := clave4
	ptrAnterior := &anterior
	dic.Iterar(func(clave string, dato int) bool {
		if FUNCION_STRINGS(*(ptrAnterior), clave) > 0 {
			return false
		}
		*ptrAnterior = clave
		*ptrFactorial *= dato
		return true
	})

	require.EqualValues(t, 720, factorial)
}

func TestIterarDiccionarioVacio(t *testing.T) {
	t.Log("Iterar sobre diccionario vacio es simplemente tenerlo al final")
	dic := TDADiccionario.CrearABB[string, int](FUNCION_STRINGS)
	iter := dic.Iterador()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestDiccionarioIterar(t *testing.T) {
	t.Log("Guardamos 3 valores en un Diccionario, e iteramos validando que las claves sean todas diferentes " +
		"pero pertenecientes al diccionario. Además los valores de VerActual y Siguiente van siendo correctos entre sí")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	dic := TDADiccionario.CrearABB[string, string](FUNCION_STRINGS)
	dic.Guardar(claves[1], valores[1])
	dic.Guardar(claves[0], valores[0])
	dic.Guardar(claves[2], valores[2])
	iter := dic.Iterador()

	require.True(t, iter.HaySiguiente())
	primero, _ := iter.VerActual()

	iter.Siguiente()
	segundo, _ := iter.VerActual()
	require.NotEqualValues(t, primero, segundo)
	require.True(t, iter.HaySiguiente())

	comparacion := FUNCION_STRINGS(primero, segundo)
	require.EqualValues(t, -1, comparacion)

	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	tercero, _ := iter.VerActual()
	require.NotEqualValues(t, primero, tercero)
	require.NotEqualValues(t, segundo, tercero)

	comparacion = FUNCION_STRINGS(segundo, tercero)
	require.EqualValues(t, -1, comparacion)
	iter.Siguiente()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestPruebaIterarTrasBorrados(t *testing.T) {
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"

	dic := TDADiccionario.CrearABB[string, string](FUNCION_STRINGS)
	dic.Guardar(clave1, "")
	dic.Guardar(clave2, "")
	dic.Guardar(clave3, "")
	dic.Borrar(clave1)
	dic.Borrar(clave2)
	dic.Borrar(clave3)
	iter := dic.Iterador()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	dic.Guardar(clave1, "A")
	iter = dic.Iterador()

	require.True(t, iter.HaySiguiente())
	c1, v1 := iter.VerActual()
	require.EqualValues(t, clave1, c1)
	require.EqualValues(t, "A", v1)
	iter.Siguiente()
	require.False(t, iter.HaySiguiente())
}

func TestVolumenIteradorCorte(t *testing.T) {
	t.Log("Prueba de volumen de iterador interno, para validar que siempre que se indique que se corte" +
		" la iteración con la función visitar, se corte")

	dic := TDADiccionario.CrearABB[int, int](FUNCION_INT)

	// Crea un arreglo de 0 a 10000
	arr := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		arr[i] = i
	}
	// Desordena el arreglo
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})

	/* Inserta intercaladamente 'n' parejas en el abb  */
	for _, clave := range arr {
		dic.Guardar(clave, clave)
	}

	seguirEjecutando := true
	siguioEjecutandoCuandoNoDebia := false

	anterior := 0
	ptrAnterior := &anterior
	orden := true

	dic.Iterar(func(c int, v int) bool {
		if FUNCION_INT(*(ptrAnterior), c) > 0 {
			orden = false
			return false
		}
		*ptrAnterior = c
		if !seguirEjecutando {
			siguioEjecutandoCuandoNoDebia = true
			return false
		}
		if c%100 == 0 {
			seguirEjecutando = false
			return false
		}
		return true
	})

	require.True(t, orden, "Se tendría que iterar en orden")
	require.False(t, seguirEjecutando, "Se tendría que haber encontrado un elemento que genere el corte")
	require.False(t, siguioEjecutandoCuandoNoDebia,
		"No debería haber seguido ejecutando si encontramos un elemento que hizo que la iteración corte")
}

func TestIterarTodosLosElementosEnRango(t *testing.T) {
	dos := 2
	cinco := 5

	otra := TDADiccionario.CrearABB[int, int](FUNCION_INT)
	otra.Guardar(5, 5)
	otra.Guardar(1, 1)
	otra.Guardar(2, 2)
	otra.Guardar(3, 3)
	otra.Guardar(8, 8)
	otra.Guardar(4, 4)
	otra.Guardar(6, 6)
	otra.Guardar(7, 7)
	perro := 0
	anterior := 0
	ptrAnterior := &anterior
	orden := true

	otra.IterarRango(&dos, &cinco, func(clave int, dato int) bool {
		if FUNCION_INT(*(ptrAnterior), clave) > 0 {
			orden = false
			return false
		}
		*ptrAnterior = clave
		perro = perro + dato
		return true
	})
	require.True(t, orden, "Se tendría que iterar en orden")
	require.EqualValues(t, 14, perro)

	gato := 0
	anterior = 0
	otra.IterarRango(&dos, &cinco, func(clave int, dato int) bool {
		if FUNCION_INT(*(ptrAnterior), clave) > 0 {
			orden = false
			return false
		}
		gato = gato + dato
		return clave != 4
	})
	require.True(t, orden, "Se tendría que iterar en orden")
	require.EqualValues(t, 9, gato)
}

func TestIteradorConRango(t *testing.T) {
	desde := 3
	hasta := 9
	ultimo := 9
	primero := 3
	la_verdad := true

	dicc := TDADiccionario.CrearABB[int, int](FUNCION_INT)
	dicc.Guardar(6, 6)
	dicc.Guardar(1, 1)
	dicc.Guardar(3, 3)
	dicc.Guardar(2, 2)
	dicc.Guardar(8, 8)
	dicc.Guardar(10, 10)
	dicc.Guardar(9, 9)
	dicc.Guardar(4, 4)
	dicc.Guardar(5, 5)
	dicc.Guardar(7, 7)
	dicc.Guardar(11, 11)

	iter := dicc.IteradorRango(&desde, &hasta)
	for iter.HaySiguiente() {
		_, dato := iter.VerActual()
		if dato < desde || dato > hasta {
			la_verdad = false
		}
		iter.Siguiente()
	}
	require.True(t, la_verdad)

	iter_2 := dicc.IteradorRango(&desde, nil)
	for iter_2.HaySiguiente() {
		_, dato := iter_2.VerActual()
		if dato < desde {
			la_verdad = false
		}
		ultimo = dato
		iter_2.Siguiente()
	}
	require.True(t, la_verdad)
	require.EqualValues(t, 11, ultimo)

	iter_3 := dicc.IteradorRango(nil, &hasta)
	for iter_3.HaySiguiente() {
		_, dato := iter_3.VerActual()
		if dato > hasta {
			la_verdad = false
		}
		if dato < primero {
			primero = dato
		}
		ultimo = dato
		iter_3.Siguiente()
	}
	require.True(t, la_verdad)
	require.EqualValues(t, hasta, ultimo)
	require.EqualValues(t, 1, primero)

	iter_4 := dicc.IteradorRango(&desde, &hasta)
	for i := desde; iter_4.HaySiguiente(); i++ {
		_, dato := iter_4.VerActual()
		if i != dato {
			la_verdad = false
		}
		iter_4.Siguiente()
	}
	require.True(t, la_verdad)
}
func ejecutarPruebaVolumen(b *testing.B, n int) {
	dic := TDADiccionario.CrearABB[string, int](FUNCION_STRINGS)

	claves := make([]string, n)
	valores := make([]int, n)

	// Crea un arreglo de 0 a n
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	// Desordena el arreglo
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})

	/* Inserta intercaladamente 'n' parejas en el abb  */
	for _, clave := range arr {
		claves[clave] = fmt.Sprintf("%08d", clave)
		valores[clave] = clave
		dic.Guardar(claves[clave], valores[clave])
	}

	require.EqualValues(b, n, dic.Cantidad(), "La cantidad de elementos es incorrecta")

	/* Verifica que devuelva los valores correctos */
	ok := true
	for i := 0; i < n; i++ {
		ok = dic.Pertenece(claves[i])
		if !ok {
			break
		}
		ok = dic.Obtener(claves[i]) == valores[i]
		if !ok {
			break
		}
	}

	require.True(b, ok, "Pertenece y Obtener con muchos elementos no funciona correctamente")
	require.EqualValues(b, n, dic.Cantidad(), "La cantidad de elementos es incorrecta")

	/* Verifica que borre y devuelva los valores correctos */
	for i := 0; i < n; i++ {
		ok = dic.Borrar(claves[i]) == valores[i]
		if !ok {
			break
		}
		ok = !dic.Pertenece(claves[i])
		if !ok {
			break
		}
	}

	require.True(b, ok, "Borrar muchos elementos no funciona correctamente")
	require.EqualValues(b, 0, dic.Cantidad())
}

func BenchmarkDiccionario(b *testing.B) {
	b.Log("Prueba de stress del Diccionario. Prueba guardando distinta cantidad de elementos (muy grandes), " +
		"ejecutando muchas veces las pruebas para generar un benchmark. Valida que la cantidad " +
		"sea la adecuada. Luego validamos que podemos obtener y ver si pertenece cada una de las claves geeneradas, " +
		"y que luego podemos borrar sin problemas")
	for _, n := range TAMS_VOLUMEN {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebaVolumen(b, n)
			}
		})
	}
}

func ejecutarPruebasVolumenIterador(b *testing.B, n int) {
	dic := TDADiccionario.CrearABB[string, *int](FUNCION_STRINGS)

	claves := make([]string, n)
	valores := make([]int, n)

	// Crea un arreglo de 0 a n
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	// Desordena el arreglo
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})

	/* Inserta 'n' parejas en el dic */
	for _, clave := range arr {
		claves[clave] = fmt.Sprintf("%08d", clave)
		valores[clave] = clave
		dic.Guardar(claves[clave], &valores[clave])
	}

	// Prueba de iteración sobre las claves almacenadas.
	iter := dic.Iterador()
	require.True(b, iter.HaySiguiente())

	ok := true
	var i int
	var clave string
	var valor *int

	for i = 0; i < n; i++ {
		if !iter.HaySiguiente() {
			ok = false
			break
		}
		c1, v1 := iter.VerActual()
		clave = c1
		if clave == "" {
			ok = false
			break
		}
		valor = v1
		if valor == nil {
			ok = false
			break
		}
		*valor = n
		iter.Siguiente()
	}
	require.True(b, ok, "Iteracion en volumen no funciona correctamente")
	require.EqualValues(b, n, i, "No se recorrió todo el largo")
	require.False(b, iter.HaySiguiente(), "El iterador debe estar al final luego de recorrer")

	ok = true
	for i = 0; i < n; i++ {
		if valores[i] != n {
			ok = false
			break
		}
	}
	require.True(b, ok, "No se cambiaron todos los elementos")
}
func BenchmarkIterador(b *testing.B) {
	b.Log("Prueba de stress del Iterador del Diccionario. Prueba guardando distinta cantidad de elementos " +
		"(muy grandes) b.N elementos, iterarlos todos sin problemas. Se ejecuta cada prueba b.N veces para generar " +
		"un benchmark")
	for _, n := range TAMS_VOLUMEN {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebasVolumenIterador(b, n)
			}
		})
	}
}
