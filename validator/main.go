package main

import (
	"fmt"
	"os"

	"github.com/google/cel-go/cel"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("❌ Uso: go run validator/main.go '<espressione>'")
		os.Exit(1)
	}
	expr := os.Args[1]
	env, _ := cel.NewEnv(cel.Variable("request.auth.claims.exp", cel.IntType))

	_, iss := env.Compile(expr)
	if iss.Err() != nil {
		fmt.Printf("❌ ERRORE: %v\n", iss.Err())
		os.Exit(1)
	}
	fmt.Println("✅ SUCCESSO! L'espressione CEL è valida e pronta per Kubernetes.")
}
