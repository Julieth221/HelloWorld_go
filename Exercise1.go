package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func imprimir() {
	fmt.Println("texto desde la funcion imprimir")
}

func hola_string(s string) string {
	return "Hola" + s
}

func sumar(a int, b int) int {
	return a + b
}

func comparar_bool() {
	var a bool = false
	b := true

	if a == b {
		fmt.Println(" A y B son iguales")
	} else {
		fmt.Println(" A y B son diferentes")
	}
}

func array() {
	var aprendices [5]string

	aprendices[0] = "Juan"
	aprendices[1] = "Maria"
	aprendices[2] = "Pedro"
	aprendices[3] = "Luis"
	aprendices[4] = "Daniela"

	fmt.Println(aprendices)
}

func tipo_dato() {
	var texto string = "julieth"
	var numero int = 1000
	var decimal float64 = 0.001
	fmt.Println(reflect.TypeOf(texto))
	fmt.Println(reflect.TypeOf(numero))
	fmt.Println(reflect.TypeOf(decimal))
}

func converstrtobool() {
	var palabra string = "true"
	boolean, err := strconv.ParseBool(palabra)

	fmt.Println(boolean, reflect.TypeOf(boolean))
	fmt.Print("este es el error", err)

}

func converboootostr() {
	var palabraboll bool = true
	strbool := strconv.FormatBool(palabraboll)
	fmt.Println(strbool, reflect.TypeOf(strbool))
}

// ejercicio 15
func variables() {
	var (
		nombre    string = "sofia"
		edad      int    = 21
		mayorEdad bool   = true
	)
	fmt.Println("Nombre: ", nombre)
	fmt.Println("Edad: ", edad)
	fmt.Println("Es mayor de edad: ", mayorEdad)
}

// ejercicio 16
func DefaultValue() {
	var (
		nombre    string
		edad      int
		mayorEdad bool
	)
	fmt.Println("Nombre: ", nombre)
	fmt.Println("Edad: ", edad)
	fmt.Println("Es mayor de edad: ", mayorEdad)
}

// ejercicio 17
func operador() {
	nombre := "sofia"
	edad := 21
	mayorEdad := true

	fmt.Printf("Nombre: %s\nEdad: %d\nEs mayor de edad: %t ", nombre, edad, mayorEdad)
}

// ejercicio 18
var var1 = "este es el nivel 1"

func scope() {
	var var2 = "este es el nivel 2"
	{
		var var3 = "este es el nivel 3"
		fmt.Println(var1)
		fmt.Println(var2)
		fmt.Println(var3)
	}
}

// ejercicio 19
func punteros() {
	vehiculo1 := "rojo"
	fmt.Println("El vehiculo es de color ", vehiculo1)

	vehiculo2 := &vehiculo1
	fmt.Println("El vehiculo es de color ", *vehiculo2)

	vehiculo3 := &vehiculo1
	fmt.Println("El vehiculo es de color ", *vehiculo3)

	vehiculo1 = "gris"
	fmt.Println("El vehiculo es de color ", vehiculo1, &vehiculo1)
	fmt.Println("El vehiculo es de color ", *vehiculo2, vehiculo2)
	fmt.Println("El vehiculo es de color ", *vehiculo3, vehiculo3)

}

// ejercicio 20

func equivalenciaEnPies(altura *float32) float32 {
	*altura = *altura - 0.10
	return *altura
}

var altura float32 = 1.70

// ejercicio 22

func main() {
	//  fmt.Println("Texto desde la funcion main")
	// imprimir()
	// fmt.Println(hola_string(" Julieth"))
	// fmt.Println(sumar(5, 3))
	// comparar_bool()
	// array()
	// tipo_dato()
	// converstrtobool()
	// converboootostr()
	// variables()
	// DefaultValue()
	// operador()
	// scope()
	// punteros()
	fmt.Println("La altura es:", altura, "mts")
	fmt.Println("La altura es:", equivalenciaEnPies(&altura), " pies")
	fmt.Println("La nueva altura es:", altura, "mts")
}
