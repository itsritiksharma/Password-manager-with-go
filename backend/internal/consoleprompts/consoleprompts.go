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
func PromptUserInput(signedInToVault bool) int64 {
	var selectedOption string
	var finalOption int64
	var gettingInput bool = true

	fmt.Println("----------------------------------------------")
	fmt.Println("What would you like to do?")
	if !signedInToVault {
		fmt.Println("1. Create a new vault.")
		fmt.Println("2. Signin to an existing vault.")
		fmt.Println("3. Fetch all credentials from a vault.")
		fmt.Println("4. Fetch a credential from a vault.")
	} else {
		fmt.Println("1. Fetch all credentials from vault.")
		fmt.Println("2. Fetch a credential from vault.")
		fmt.Println("3. Delete vault.")
		fmt.Println("4. Signout")
	}
	fmt.Println("Quit (enter q or quit or press \"ctrl+c\")")

	for gettingInput {
		fmt.Scan(&selectedOption)

		if selectedOption == "q" || selectedOption == "quit" {
			fmt.Println("Exiting")
			return 0
		}
		intOption, err := strconv.ParseInt(selectedOption, 10, 32)
		if err != nil {
			fmt.Println("Please enter a valid option.")
			continue
		}

		if intOption <= 0 || intOption >= 5 {
			fmt.Println("Please enter a valid option.")
			continue
		} else {
			finalOption = intOption
			gettingInput = false
		}

	}

	return finalOption

}
