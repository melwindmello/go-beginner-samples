package bankingapp

import "fmt"

func StartBankingApplication() {

	fmt.Println("******* WELCOME TO ABC BANK *******\n\n")
	account1 := Account{AccountNumber: "1001", Name: "John Doe", AccountBalance: 1001.00,
		AccountType: "Savings", MinimumBalance: 100.00, InterestRate: 7.75}

	fmt.Println(account1) //equivalent to fmt.Println(account1.String()) assuming String() is implemented

	var selectedOption int32 = ShowMenuAndGetUserSelection()

	for selectedOption != 0 {
		switch selectedOption {
		case 1:
			account1.WithdrawAmount()
		case 2:
			account1.DepositAmount()
		case 3:
			fmt.Println(account1)
		}

		selectedOption = ShowMenuAndGetUserSelection()
	}

	fmt.Println("\n\n******* THANK YOU FOR USING OUR SERVICES *******")

}
