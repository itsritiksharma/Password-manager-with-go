package consoleprompts

import (
	"fmt"
	"strconv"
)

func prompt(message string, requireMultiplePrompts bool, invalidPromptReceivedCondition bool, invalidPromptMessage string) string {

	var selectedOption string
	var readingPrompt string = "reading-prompt"

	if requireMultiplePrompts {
		for readingPrompt != "not-reading-prompt" {
			if readingPrompt == "reading-prompt" {
				fmt.Print(message)
				fmt.Scan(&selectedOption)
				if invalidPromptReceivedCondition {
					fmt.Println(invalidPromptMessage)
					readingPrompt = "reading-prompt"
				} else {
					readingPrompt = "not-reading-prompt"
				}
			}
		}
	} else {
		fmt.Print(message)
		fmt.Scan(&selectedOption)
	}

	return string(selectedOption)
}

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
	fmt.Println("6. Show JSON file data.")
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

			if intOption <= 0 || intOption >= 7 {
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
