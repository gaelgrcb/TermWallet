package cli

import "fmt"

func Input() int {
	for {
		var op int
		_, err := fmt.Scanln(&op)

		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			var clean string
			fmt.Scanln(&clean)
			continue
		}

		return op
	}
}
