// 1. import library and package
// 2. define function priority operation
// 3. define function manual calculation
// 4. define main function
// 5. run program

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//read input
// console input
// read input

//manipulate string

// 2. define function priority operation
func priorityOperation(operation string) int {
	switch operation {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	}
	return 0
}

// check if operation
func isOperation(operation string) bool {
	if operation == "+" || operation == "-" || operation == "*" || operation == "/" {
		return true
	}
	return false
}

func infixToPostfix(tokens []string) []string {
	var output []string
	var stack []string

	for _, token := range tokens {
		if isOperation(token) {
			for len(stack) > 0 && priorityOperation(stack[len(stack)-1]) >= priorityOperation(token) {
				output = append(output, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, token)
		} else if token == "(" {
			stack = append(stack, token)
		} else if token == ")" {
			for len(stack) > 0 && stack[len(stack)-1] != "(" {
				output = append(output, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1] // buang "("
		} else {
			output = append(output, token)
		}
	}

	for len(stack) > 0 {
		output = append(output, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	return output
}

func evalPostfix(tokens []string) float64 {
	var stack []float64

	for _, token := range tokens {
		if isOperation(token) {
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			var res float64
			switch token {
			case "+":
				res = a + b
			case "-":
				res = a - b
			case "*":
				res = a * b
			case "/":
				res = a / b
			}
			stack = append(stack, res)
		} else {
			val, _ := strconv.ParseFloat(token, 64)
			stack = append(stack, val)
		}
	}

	return stack[0]
}

func tokenize(expr string) []string {
	var tokens []string
	var number strings.Builder

	for _, ch := range expr {
		c := string(ch)
		if c >= "0" && c <= "9" {
			number.WriteRune(ch)
		} else {
			if number.Len() > 0 {
				tokens = append(tokens, number.String())
				number.Reset()
			}
			if strings.TrimSpace(c) != "" {
				tokens = append(tokens, c)
			}
		}
	}

	if number.Len() > 0 {
		tokens = append(tokens, number.String())
	}

	return tokens
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Kalkulator Kompleks (tanpa library)")
	fmt.Println("Contoh input: 3 + 5 * ( 2 - 1 )")
	fmt.Println("Ketik 'exit' untuk keluar")

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()

		if input == "exit" {
			fmt.Println("Keluar...")
			break
		}

		tokens := tokenize(input)
		postfix := infixToPostfix(tokens)
		result := evalPostfix(postfix)
		fmt.Printf("Hasil: %.2f\n", result)
	}
}
