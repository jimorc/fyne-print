package main

import (
	"fmt"

	"github.com/jimorc/fyne-print/print"
)

func main() {
	printers := print.NewPrinters()
	defer printers.Close()
	for i, pr := range printers.Printers {
		fmt.Printf("Printer %d:\n", i)
		fmt.Printf("    Name: %s\n", pr.Name())
		fmt.Printf("    Instance: %s\n", pr.Instance())
	}
}
