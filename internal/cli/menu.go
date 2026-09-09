package cli

import (
	"fmt"
)

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
		op := Input()

		switch op {
		case 1:
			addIncome()
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

func addIncome() {
	fmt.Println("What type of income do you want to add?\n" +
		"(1) Nominal Income\n" +
		"(2) External Income\n" +
		"(3) Other\n" +
		"(4) Return to home")

Loop:
	for {
		op := Input()

		switch op {
		case 1:
		case 2:
		case 3:
		case 4:
			Start()
			break Loop
		default:
			invalid := fmt.Sprintf("Your choice: %v is not available, try again\n", op)
			fmt.Println(invalid)
		}
	}
}

func addExpense() {
	fmt.Println("What type of expense do you want to add?\n" +
		"(1) Nominal Income\n" +
		"(2) External Income\n" +
		"(3) Other\n" +
		"(4) Return to home")

Loop:
	for {
		op := Input()

		switch op {
		case 1:
		case 2:
		case 3:
		case 4:
			Start()
			break Loop
		default:
			invalid := fmt.Sprintf("Your choice: %v is not available, try again\n", op)
			fmt.Println(invalid)
		}
	}
}
