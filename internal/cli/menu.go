package cli

import "fmt"

func Start() {
	fmt.Print("-------------------------------------------------------\n" +
		"Welcome to the TermWallet \n" +
		"Which of these options would you like to start with?\n" +
		"(1) Add income \n" +
		"(2) Add expense \n" +
		"(3) See Balance and Transactions \n" +
		"(4) Exit\n\n" +
		"Select your choice with a number\n" +
		"-------------------------------------------------------\n")
Loop:
	for {
		var op int
		_, err := fmt.Scanln(&op)

		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			var clean string
			fmt.Scanln(&clean)
			continue
		}

		switch op {
		case 1:
		case 2:
		case 3:
		case 4:
			fmt.Println("Goodbye see you soon! ;)")
			break Loop
		default:
			invalid := fmt.Sprintf("Your choice: %v is not available, try again\n", op)
			fmt.Println(invalid)
		}
	}
}
