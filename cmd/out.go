package cmd

import "fmt"

func Out(a ...any) {
	fmt.Print(a...)
}

func Outf(format string, a ...any) {
	fmt.Printf(format, a...)
}
