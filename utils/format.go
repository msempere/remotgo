package utils

import "fmt"

func RenderOutput(header, content string, quite bool) {
	if !quite {
		fmt.Print("+")
		for i := 0; i <= len(header)+1; i++ {
			fmt.Print("-")
		}
		fmt.Printf("+\n| %v |\n+", header)
		for i := 0; i <= len(header)+1; i++ {
			fmt.Print("-")
		}
		fmt.Println("+")
	}
	fmt.Println(content)
}
