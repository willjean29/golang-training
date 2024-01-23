package mains

import (
	"fmt"
	"trainig/fileops"

	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFileName = "balance.txt"

func main() {
	accountBalance, err := fileops.GetBalanceFromFile(accountBalanceFileName)
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("----------")
		panic("Can't continue, sorry!")
	}
	var choise int
	fmt.Println("Welcome to Go Bank!", randomdata.FirstName(0))
	for {
		presentOptions()
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
			fileops.WriteBalanceToFile(accountBalance, accountBalanceFileName)
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
				fileops.WriteBalanceToFile(accountBalance, accountBalanceFileName)
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
