package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFileName = "balance.txt"

func main() {
	accountBalance, err := getBalanceFromFile()
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("----------")
		panic("Can't continue, sorry!")
	}
	var choise int
	fmt.Println("Welcome to Go Bank!")
	for {

		fmt.Println("what do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Witdraw money")
		fmt.Println("4. Exit")

		fmt.Print("Your choice: ")
		fmt.Scan(&choise)

		if choise == 1 {
			fmt.Println("Your balance is: ", accountBalance)
		}

		if choise == 2 {
			var depositAmount float64
			fmt.Print("Enter amount to deposit: ")
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid amount!")
				continue
			}
			accountBalance += depositAmount
			fmt.Println("Your balance is: ", accountBalance)
			writeBalanceToFile(accountBalance)
		}

		if choise == 3 {
			var withdrawAmount float64
			fmt.Print("Enter amount to withdraw: ")
			fmt.Scan(&withdrawAmount)
			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount!, Must be greater than 0")
				continue
			}
			if withdrawAmount > accountBalance {
				fmt.Println("You don't have enough funds")
				continue
			} else {
				accountBalance -= withdrawAmount
				fmt.Println("Your balance is: ", accountBalance)
				writeBalanceToFile(accountBalance)
			}
		}

		if choise == 4 {
			fmt.Println("Goodbye")
			break
		} else {
			fmt.Println("Invalid choice!")
			continue
		}
	}
	fmt.Println("Thanks for using Go Bank!")
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFileName, []byte(balanceText), 0644)
}

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFileName)
	if err != nil {
		return 1000, errors.New("Failed to find balance file.")
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("Failed to parse stored balance value.")
	}
	return balance, nil
}
