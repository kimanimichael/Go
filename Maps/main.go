package main

import (
	"fmt"
	"errors"
)


type user struct {
	name string
	number int
}

func getUserMap (names []string, phoneNumbers[]int) (map[string] user, error){
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("Names not equal to phone numbers")
	}
	userMap := make(map[string]user)
	for i := 0; i < len(names); i++ {
		name := names[i]
		phoneNumber := phoneNumbers[i]
		var siteUser user
		siteUser.name = name
		siteUser.number = phoneNumber
		userMap[name] = siteUser
	}
	return userMap, nil
}

func main () {
	ages := make(map[string]int)
	ages["MK"] = 24
	ages["Brian G"] = 27

	cars := map[string]int {
		"Kim": 1,
		"MK": 0 }

	fmt.Println(len(ages))
	fmt.Println(len(cars))

	names := make([]string, 3)
	names = append(names, "MK")
	names = append(names, "Brian G")

	numbers := make([]int, 3)
	numbers = append(numbers, 5)
	numbers = append(numbers, 6)

	webUser, _ := getUserMap(names, numbers)
	fmt.Println(len(webUser))
}