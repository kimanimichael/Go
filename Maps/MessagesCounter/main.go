package main

import (
	"fmt"
)
/* gets current count per user ID, increments it and stores it */
func getCounts (userIDs []string) map[string]int {
	counts := make(map[string] int)
	for _, userID := range userIDs {
		if userID == "" {
			continue
		}
		count := counts[userID]
		count ++
		counts[userID] = count
	}
	return counts
}

func getAlphabetNameCounts (userIDs[]string) map[rune]map[string]int {
	counts := make(map[rune]map[string]int)
	for _, userID := range userIDs {
		if userID == "" {
			continue
		}
		firstChar := rune(userID[0])
		_, ok := counts[firstChar]
		if !ok {
			counts[firstChar] = make(map[string]int)
		}
		counts[firstChar][userID]++
	}
	return counts
}

func main() {
	users := make([]string, 18)
	users = append(users, "MK")
	users = append(users, "Brian G")
	users = append(users, "Monicah K")
	users = append(users, "John K")
	users = append(users, "Maureen Bi")
	users = append(users, "Imma K")
	users = append(users, "MK")
	users = append(users, "John K")
	users = append(users, "Joe N")
	users = append(users, "Maureen Bi")
	users = append(users, "Brian G")
	users = append(users, "MK")
	users = append(users, "Maureen Bi")
	users = append(users, "Imma K")
	users = append(users, "MK")
	users = append(users, "Maureen Bi")
	users = append(users, "John K")

	messageCounts := getCounts(users)
	
	for userName, messageCount := range messageCounts {
		fmt.Print(userName,":")
		fmt.Println(messageCount)
	}
	/* use nested maps */
	fmt.Println("Nested map tests")
	messageAlphabetCounts := getAlphabetNameCounts(users)

	for firstLetter, messageAlphabetCountsmessageCount := range messageAlphabetCounts {
		fmt.Print(string(firstLetter),":")
		fmt.Println(messageAlphabetCountsmessageCount)
	}

}