package main

import (
	"fmt"

	"github.com/jimorc/fyne-print/print"
)

func main() {
	printers := print.NewPrinters()
	defer printers.Close()
	for _, pr := range printers.Printers {
		fmt.Println(pr)
	}
}
