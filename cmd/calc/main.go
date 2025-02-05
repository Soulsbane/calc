package main

import (
	"fmt"
	"math"
	"os"

	"github.com/expr-lang/expr"
)

func main() {
	args := os.Args[1:]
	env := map[string]interface{}{
		"pi": math.Pi,
	}

	if len(args) > 0 {
		code := args[0]
		program, err := expr.Compile(code, expr.Env(env))

		if err != nil {
			fmt.Println(err)
		} else {
			output, err := expr.Run(program, env)

			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(output)
			}
		}
	} else {
		fmt.Println("Usage: calc <expression>")
	}
}
