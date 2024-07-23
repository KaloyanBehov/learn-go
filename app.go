package main

import "fmt"

func main() {
	accountBlance := 1000.0
	fmt.Println("Welcome to go bank")
	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")
		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println("Your account balance is :", accountBlance)
		} else if choice == 2 {
			fmt.Println("enter the deposit amount")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			accountBlance += depositAmount
			fmt.Println("Balance updated new balance: ", accountBlance)
		} else if choice == 3 {
			fmt.Println("how much do you want to witdhraw")
			var witdhrawAmount float64
			fmt.Scan(&witdhrawAmount)
			if witdhrawAmount <= accountBlance {
				accountBlance -= witdhrawAmount
				fmt.Print("new account balance", accountBlance)
			}
		} else {
			fmt.Println("Goodbye")
			break
		}
	}
	fmt.Print("Thanks for using this program")

}
