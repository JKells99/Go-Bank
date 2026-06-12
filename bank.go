package main

import "fmt"
import "os"
import "strconv"

const balanceFile = "balance.txt"
func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(balanceFile,[]byte(balanceText), 0644)


}

func readBalanceFromFile() float64 {
	data ,_ := os.ReadFile(balanceFile)
	balanceText := string(data)
	balance, _ := strconv.ParseFloat(balanceText, 64)

	return balance
}

func main() {
	var accountBalance float64 = readBalanceFromFile()
	fmt.Println("Welcome to the bank")
	for {
	
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
		writeBalanceToFile(accountBalance)


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
		writeBalanceToFile(accountBalance)

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
			writeBalanceToFile(accountBalance)
		}
	} else{
		fmt.Println("Thank you for using the bank. Goodbye!")
		break;
	}

	}
	
}
