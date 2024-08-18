package operaciones

import (
	"fmt"
	"math"
	"strconv"

	"tdas/pila" // Importamos el paquete de la Pila
)

func RealizarOperaciones(pila pila.Pila[int64], token string, contador int) (bool, int) {

	switch token {
	case "+", "-", "*", "/", "^", "sqrt", "log", "?":

		// Verificamos si hay suficientes operandos en la pila
		if pila.EstaVacia() && token != "?" {
			return false, contador

			// Realizamos la operación correspondiente y apilamos
		} else if token == "sqrt" && contador >= 1 {
			posible_error := raiz(pila)
			if posible_error != nil {
				return false, contador
			}

		} else if token == "+" && contador >= 2 {
			suma(pila)
			contador--

		} else if token == "-" && contador >= 2 {
			resta(pila)
			contador--

		} else if token == "*" && contador >= 2 {
			multiplicacion(pila)
			contador--

		} else if token == "/" && contador >= 2 {
			posible_error := division(pila)

			if posible_error != nil {
				return false, contador

			}
			contador--

		} else if token == "^" && contador >= 2 {
			posible_error := potencia(pila)

			if posible_error != nil {
				return false, contador
			}
			contador--

		} else if token == "log" && contador >= 2 {
			posible_error := logaritmo(pila)

			if posible_error != nil {
				return false, contador

			}
			contador--

		} else if token == "?" && contador >= 3 {
			operadorTernario(pila)
			contador -= 2

		} else { //No cumple con la cantidad de numeros para dicha operacion
			return false, contador

		}

	default:
		// Convertimos el token a un número y lo colocamos en la pila
		posible_error := convertidorNumero(pila, token)
		if posible_error != nil {
			return false, contador

		}
		contador++

	}
	return true, contador
}

// Convertimos el token a un número y lo colocamos en la pila

func convertidorNumero(pila pila.Pila[int64], token string) error {
	num, err := strconv.Atoi(token)
	if err != nil {
		return fmt.Errorf("ERROR")

	}
	pila.Apilar(int64(num))
	return nil

}

//Necesita un numero:

func raiz(pila pila.Pila[int64]) error {
	a := pila.Desapilar()
	if a < 0 {
		return fmt.Errorf("ERROR")

	} else {
		pila.Apilar(int64(math.Sqrt(float64(a))))

	}
	return nil

}

//Necesita 2 numeros:

func suma(pila pila.Pila[int64]) {
	b := pila.Desapilar()
	a := pila.Desapilar()
	pila.Apilar(a + b)
}

func resta(pila pila.Pila[int64]) {
	b := pila.Desapilar()
	a := pila.Desapilar()
	pila.Apilar(a - b)
}

func multiplicacion(pila pila.Pila[int64]) {
	b := pila.Desapilar()
	a := pila.Desapilar()
	pila.Apilar(a * b)
}

func division(pila pila.Pila[int64]) error {
	b := pila.Desapilar()
	a := pila.Desapilar()

	if b == 0 {
		return fmt.Errorf("ERROR")

	} else {
		pila.Apilar(a / b)

	}
	return nil
}

func potencia(pila pila.Pila[int64]) error {
	b := pila.Desapilar()
	a := pila.Desapilar()

	if b < 0 {
		return fmt.Errorf("ERROR")

	} else {
		pila.Apilar(int64(math.Pow(float64(a), float64(b))))

	}
	return nil
}

func logaritmo(pila pila.Pila[int64]) error {
	b := pila.Desapilar()
	a := pila.Desapilar()
	if a < 1 || b <= 1 {
		return fmt.Errorf("ERROR")

	} else {
		pila.Apilar(int64(math.Log(float64(a)) / math.Log(float64(b))))

	}
	return nil
}

//Necesita 3 numeros:

func operadorTernario(pila pila.Pila[int64]) {
	c := pila.Desapilar()
	b := pila.Desapilar()
	a := pila.Desapilar()

	if a != 0 {
		pila.Apilar(b)

	} else {
		pila.Apilar(c)

	}
}
