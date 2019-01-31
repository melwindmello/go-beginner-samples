package bankingapp

import "fmt"
import "github.com/fatih/color" //https://github.com/fatih/color

func ShowMenuAndGetUserSelection() int32 {
	var option int32 = 0
	color.HiCyan("\n================   SELECT AN OPTION ========================")
	color.HiCyan("1	Withdraw")
	color.HiCyan("2	Deposit")
	color.HiCyan("3	Print Accoutn Details")
	color.HiCyan("4	View ")
	color.HiCyan(GetLineSeparator())
	color.HiCyan("Enter your option:")
	//color.HiCyan("Prints text in cyan.")
	fmt.Scanf("%d\n", &option)

	if option == 1 || option == 2 || option == 3 {
		return option
	} else {
		return 0
	}

}

func GetLineSeparator() string {
	return "============================================================"
}
