package utils

import "fmt"

func RenderOutput(header, content string) {
	fmt.Print("+")
	for i := 0; i <= len(header)+1; i++ {
		fmt.Print("-")
	}
	fmt.Printf("+\n| %v |\n+", header)
	for i := 0; i <= len(header)+1; i++ {
		fmt.Print("-")
	}
	fmt.Printf("+\n%v\n", content)
}
