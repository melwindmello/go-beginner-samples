package bankingapp

import "fmt"
import "github.com/fatih/color" //https://github.com/fatih/color

type Account struct {
	AccountNumber     string
	Name              string
	AccountBalance    float32
	AccountType       string
	MinimumBalance    float32
	InterestRate      float32
	WithdrawalAmounts []float32
	DepositedAmount   []float32
}

func (act Account) String() string {
	stringToFormat := GetLineSeparator()
	stringToFormat += "\n Account number = %v\n Name = %v\n Balance = %v\n Account Type = %v\n" +
		" Min. Balance = %v\n Interest Rate = %v \n" +
		" Withdrawals = %v\n Deposits = %v\n"
	stringToFormat += GetLineSeparator()
	return fmt.Sprintf(stringToFormat,
		act.AccountNumber, act.Name, act.AccountBalance,
		act.AccountType, act.MinimumBalance, act.InterestRate,
		act.WithdrawalAmounts, act.DepositedAmount)
}

func (act *Account) WithdrawAmount() {
	var tranxAmount float32 = 0
	var validationErrors bool = false
	fmt.Print("\nEnter amount to withdraw: ")
	fmt.Scanf("%g\n", &tranxAmount)

	if tranxAmount < 1 {
		color.HiRed("Invalid amount entered. Try Again.")
		validationErrors = true
	}

	if (act.AccountBalance - tranxAmount) < act.MinimumBalance {
		color.HiRed("Cannot withdraw %g. Available Balance is %g. Minimum to be maintained is %g. Try Again.\n",
			tranxAmount, act.AccountBalance, act.MinimumBalance)
		validationErrors = true
	}

	if !validationErrors {
		act.AccountBalance -= tranxAmount
		act.WithdrawalAmounts = append(act.WithdrawalAmounts, tranxAmount)
		color.HiGreen("%g Withdrawn from account. Available Balance is %g\n", tranxAmount, act.AccountBalance)
	}

}

func (act *Account) DepositAmount() {
	var tranxAmount float32 = 0
	var validationErrors bool = false

	fmt.Print("Enter amount to deposit: ")
	fmt.Scanf("%g\n", &tranxAmount)

	if tranxAmount < 1 {
		color.HiRed("Invalid amount entered. Try Again.")
		validationErrors = true
	}

	if !validationErrors {
		act.AccountBalance += tranxAmount
		act.DepositedAmount = append(act.DepositedAmount, tranxAmount)
		color.HiGreen("%g Deposited into account. Balance is %g\n", tranxAmount, act.AccountBalance)
	}

}
