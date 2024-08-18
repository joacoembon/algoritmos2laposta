package calculadora

import (
	"fmt"
	"strings"
	"tdas/pila"
	"tp1/operaciones"
)

// Función para realizar operaciones en notación posfija
func Calculadora(expresion string) string {
	pila := pila.CrearPilaDinamica[int64]() // Creamos una nueva pila para almacenar los números
	var resultado int64                     // Variable para almacenar el resultado final
	var contador int                        // Contador
	var oper bool

	tokens := strings.Fields(expresion) // Dividimos la expresión en tokens

	for _, token := range tokens {
		oper, contador = operaciones.RealizarOperaciones(pila, token, contador)
		if !oper {
			return "ERROR"
		}
	}

	// Verificamos si hay un único resultado en la pila
	if pila.EstaVacia() || contador != 1 {
		return "ERROR"

	}
	resultado = pila.Desapilar()

	// Devolvemos el resultado como una cadena
	return fmt.Sprintf("%d", resultado)

}
