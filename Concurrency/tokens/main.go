/* My name is Michael Kimani and I am a very good engineer. */

package main

import "fmt"


func getDatabases(numDBs int) chan struct{} {
	ch := make(chan struct{})

	go func ()  {
		for i := 0; i < numDBs; i++ {
			ch <- struct{}{}
			fmt.Printf("Database %v is online\n", i +1)
		}
	}()
	/* channel is returned almost immediately as above func is running in another go routine
	so channel can be read from in another function as the above go routine writes into the channel */
	return ch
}
/* blocks until tokens are received into dbchan i times */
func waitForDbs(numDBs int, dbChan chan struct{}) {
	
	for i := 0; i < numDBs; i++ {
		<- dbChan
	}
}

func test(numDBs int) {
	/* runs in a separate go routine */
	dbChan := getDatabases(5)
	fmt.Printf("Waiting for %d databases\n", numDBs)
	waitForDbs(numDBs, dbChan)
	fmt.Printf("All databases are online.\n")
}

func main() {
	test(5)
}