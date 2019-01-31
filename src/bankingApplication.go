package main

import (
	"bankingapp"
)

func main() {
	bankingapp.StartBankingApplication()
}

//3.	Create another option to check last 5 withdrawals of last 5 deposits. Note: Use slice.
//4.	Create a map to store accountNo and available balance, Ask the user to provide accountNo and display available balance.
//3.	Add Interest calculation functionality for FD type of Account.
//4.	Use panic and recover for error handling.
//1.	Modify Banking application to add new account details.
//2.	Create an array of pointers to store account information.
//3.	Create a menu option to display account details.

//1.	Modify our Banking application to create an account structures, savingsAccount and CurrentAccount, ensure that we can perform all operations on it.
//2.	Create an interface for InterestCal which will have method calculateInterest() and implement it in above structures.

//1.	Create a package called Bank and create files for main, Account structures, Transactions.
//2.	Create go routines for withdraw and deposit, and store these transactions in an array.
//3.	View last ten transactions details.
//4.	Use channels to view above transaction details.
//5.	Create a Test file to test above withdraw and deposit operations
