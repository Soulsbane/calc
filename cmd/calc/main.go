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
		"pi":        math.Pi,
		"round":     math.Round,
		"ceil":      math.Ceil,
		"floor":     math.Floor,
		"abs":       math.Abs,
		"sqrt":      math.Sqrt,
		"exp":       math.Exp,
		"exp2":      math.Exp2,
		"log":       math.Log,
		"log10":     math.Log10,
		"log1p":     math.Log1p,
		"log2":      math.Log2,
		"logb":      math.Logb,
		"pow":       math.Pow,
		"sin":       math.Sin,
		"cos":       math.Cos,
		"tan":       math.Tan,
		"asin":      math.Asin,
		"acos":      math.Acos,
		"atan":      math.Atan,
		"atan2":     math.Atan2,
		"max":       math.Max,
		"min":       math.Min,
		"mod":       math.Mod,
		"modf":      math.Modf,
		"remainder": math.Remainder,
		"cbrt":      math.Cbrt,
		"fma":       math.FMA,
		"hypot":     math.Hypot,
	}

	if len(args) > 0 {
		code := args[0]
		program, err := expr.Compile(code, expr.Env(env))

		if err != nil {
			fmt.Println(err)
		} else {
			output, runErr := expr.Run(program, env)

			if runErr != nil {
				fmt.Println(runErr)
			} else {
				fmt.Println(output)
			}
		}
	} else {
		fmt.Println("Usage: calc <expression>")
	}
}
