package main

import (
	"fmt"

	"github.com/jimorc/fyne-print/print"
)

func main() {
	p := print.NewPrinters()
	for i, pr := range *p {
		fmt.Printf("Printer %d:\n", i)
		fmt.Println(pr.String())
	}
}
