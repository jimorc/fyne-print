//go:build !windoews

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
		fmt.Printf("    IsDefault: %t\n", pr.IsDefault())
		fmt.Println("    Options:")
		for k, v := range pr.Options() {
			fmt.Printf("        %s: %s\n", k, v)
		}
		caps := pr.Capabilities().AsStrings()
		for _, cap := range caps {
			fmt.Printf("        %s\n", cap)
		}
		fmt.Println(pr.MediaSizes().AsString())
	}
}
