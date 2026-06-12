package main

import "fmt"

func main() {
	for {
	var accountBalance = 1000.00
	fmt.Println("Welcome to the bank")
	fmt.Println("What do you want to do?? ")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	var choice int 
	fmt.Scanln(&choice)
	fmt.Println("You chose option", choice)
	

	if choice == 1{
		fmt.Println("Your account balance is:", accountBalance)


	}else if choice == 2{
		fmt.Println("Enter the amount to deposit:")
		var depositAmount float64

		fmt.Scan(&depositAmount)
		if(depositAmount <= 0){
			fmt.Println("Invalid amount. Please enter a positive number.")
			continue
		}
		accountBalance += depositAmount
		fmt.Println("Your account balance is now:", accountBalance)

	}else if choice == 3{
		fmt.Println("Enter Amount to withdraw:")
		var withdrawAmount float64
		fmt.Scan(&withdrawAmount)
		if(withdrawAmount > accountBalance){
			fmt.Println("Insufficient funds")
			continue
		}else{
			accountBalance -= withdrawAmount
			fmt.Println("Your account balance is now:", accountBalance)
		}
	} else{
		fmt.Println("Thank you for using the bank. Goodbye!")
		break;
	}

	}
	
}
