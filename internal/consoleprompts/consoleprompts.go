package consoleprompts

import (
	"fmt"
	"strconv"
)

/**
 * Welcome to password manager.
 * What would you like to do?
 * 1. Create a new vault.
 * 2. Signin to an existing vault.
 * 3. Delete vault.
 * 4. Fetch a credential from a vault.
 * 5. Fetch all credentials from a vault.
 * Quit (enter q or quit or press "ctrl+c")
 */
func PromptUserInput() int64 {
	var selectedOption string
	var finalOption int64
	var gettingInput int = 1

	fmt.Println("What would you like to do?")
	fmt.Println("1. Create a new vault.")
	fmt.Println("2. Signin to an existing vault.")
	fmt.Println("3. Delete vault.")
	fmt.Println("4. Fetch all credentials from a vault.")
	fmt.Println("5. Fetch a credential from a vault.")
	fmt.Println("Quit (enter q or quit or press \"ctrl+c\")")

	for gettingInput != 0 {
		if gettingInput == 1 {

			fmt.Scan(&selectedOption)

			if selectedOption == "q" || selectedOption == "quit" {
				fmt.Println("Exiting")
				gettingInput = 1
				return 0
			}
			intOption, err := strconv.ParseInt(selectedOption, 10, 32)
			if err != nil {
				fmt.Println("Please enter a valid option.")
				gettingInput = 1
				continue
			}

			if intOption <= 0 || intOption >= 6 {
				fmt.Println("Please enter a valid option.")
				gettingInput = 1
			} else {
				finalOption = intOption
				gettingInput = 0
			}
		}
	}

	return finalOption

}
