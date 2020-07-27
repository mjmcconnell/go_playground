// main.go

package main

import (
	"github.com/mjmcconnell/go_playground/base"
	"github.com/mjmcconnell/go_playground/products"
)

func main() {
	a := base.App{}
	a.Initialize()

	products.Initialize(&a)

	a.Run(":8080")
}
