package main

import (
	"fmt"

	"github.com/jimorc/fyne-print/print"
)

func main() {
	p := print.NewPrinters()
	for _, pr := range *p {
		fmt.Println(pr)
	}
}
