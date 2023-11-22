package main

import (
	"errors"
	"fmt"
)

type user struct{
	name string
	number int
	scheduledForDeletion bool
}

/* function which passes a map*/ 
func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {
	currentUser, ok:= users[name]
	if ok {
		if currentUser.scheduledForDeletion {
			delete(users, name)
			return true, nil
		}
		return false, nil
	}
	return false, errors.New("User does not exist")
}

func main() {
	vehicleNumbers := make(map[string]int)
	/* insert an element to a key */
	vehicleNumbers["MK"] = 0
	vehicleNumbers["Brain G"] = 0
	vehicleNumbers["John Kim"] = 1
	fmt.Println(vehicleNumbers)

	/* access specific element through a key */
	fmt.Println(vehicleNumbers["John Kim"])
	/* delete an element by passing the key*/
	delete(vehicleNumbers, "Brian G")
	fmt.Println(vehicleNumbers["Brian G"])

	dadInfo, ok := vehicleNumbers["John Kim"]
	if ok {
		fmt.Println("Dad's cars: ", dadInfo)
	}

	userMaurice := user{
		"Kirimi",
		123,
		false }

	userHannah := user{
		"Wambui",
		456,
		true }
	
	userChris := user{
		"Karuti",
		789,
		false }

	userXavier := user{
		"Kibet",
		789,
		true }

	truckUsers := make(map[string]user)
	truckUsers["Chris"] = userChris
	truckUsers["Hannah"] = userHannah
	truckUsers["Xavier"] = userXavier
	truckUsers["userMaurice"] = userMaurice


	_, deletionErros := deleteIfNecessary(truckUsers, "Xaier")
	fmt.Println(deletionErros)
}