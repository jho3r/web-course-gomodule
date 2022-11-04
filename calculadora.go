package web_course_gomodule

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Calc struct {
}

func (Calc) Operate(entrada string, funcion string) int {
	entradaLimpia := strings.Split(entrada, funcion)
	operador1 := parsear(entradaLimpia[0])
	operador2 := parsear(entradaLimpia[1])
	var resultado int
	switch funcion {
	case "+":
		resultado = operador1 + operador2
	case "-":
		resultado = operador1 - operador2
	case "*":
		resultado = operador1 * operador2
	case "/":
		resultado = operador1 / operador2
	default:
		fmt.Println("Operador", funcion, "no soportado")
		resultado = 0
	}
	return resultado
}

func LeerEntrada(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func parsear(entrada string) int {
	operador, err := strconv.Atoi(entrada)
	if err != nil {
		fmt.Println(err)
	}
	return operador
}
