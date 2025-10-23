package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/atotto/clipboard"
)

func main() {
	clipText, err := clipboard.ReadAll();
	if err!=nil{
		fmt.Println("Failed read from the clipboard:", err)
		return
	}

	phoneRegex := regexp.MustCompile(`(\+91[\-\s]?|0)?[6-9]\d{9}`)
	emailRegex := regexp.MustCompile(`[a-zA-Z0-9-.+%_]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)

	phoneNumbers := phoneRegex.FindAllString(clipText, -1)
	emailAddresses := emailRegex.FindAllString(clipText, -1)
	allResults := append(phoneNumbers, emailAddresses...)
	
	if len(allResults) > 0{
		newClipText := strings.Join(allResults, "\n")
		err := clipboard.WriteAll(newClipText)
		if err!=nil{
			fmt.Println("Failed to write to the clipboard", err)
			return
		}
		fmt.Println("Clipboard updated with the phone numbers and e-mail addresses")
	}else{
		fmt.Println("No phone numbers or e-mail addresses found")
	}

}
