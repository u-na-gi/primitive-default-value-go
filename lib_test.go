package primitivedefaultvaluego

import (
	"go/ast"
	"go/parser"
	"log"
	"testing"
)

type Person2 struct {
	Name  string
	Age   int
	Shark struct {
		Teeth int
	}
}

// parseExpr is a helper function that parses a string expression into an ast.Expr.
// It simplifies the creation of ast.Expr instances for this example.
func parseExpr(expr string) ast.Expr {
	parsedExpr, err := parser.ParseExpr(expr)
	if err != nil {
		log.Fatalf("Failed to parse expression '%s': %v", expr, err)
	}
	return parsedExpr
}

func TestXxx(t *testing.T) {
	t.Run("TestDefaultValue", func(t *testing.T) {
		// var a string
		// var b int
		// var c float32
		// var d bool
		// var e complex64

		// var f []int
		// var g map[string]int

		tests := []struct {
			name string
			arg  ast.Expr
			want string
		}{
			{"string", parseExpr("string"), `""`},
			{"int", parseExpr("int"), "0"},
			{"float32", parseExpr("float32"), "0.0"},
			{"bool", parseExpr("bool"), "false"},
			{"complex64", parseExpr("complex64"), "0 + 0i"},
			{"slice", parseExpr("[]int"), "nil"},
			{"map", parseExpr("map[string]int"), "nil"},
			{"struct", parseExpr("*Person2"), "nil"},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got := DefaultValue(tt.arg)
				if got != tt.want {
					t.Errorf("DefaultValue() = %v, want %v", got, tt.want)
				}
			})
		}
	})

}
