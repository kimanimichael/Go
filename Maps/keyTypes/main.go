package main

import (
	"fmt"
)

func add (m map[string]map[string]int, path, country string) {
	/* checks whether m[path] (inner map) exists and assigns it to mm */
	mm, ok := m[path]
	if !ok {
		/* creates non-existent inner map and names it mm*/
		mm = make(map[string]int)
		/* assigns map created to m[path] */
		m[path] = mm
	}
	mm[country]++
}

func main() {
	hits := make(map[string]map[string]int)
	/* access elements in a nested map */
	n_australia := hits["doc/"]["au"]
	n_nairobi := hits["doc/"]["k"]
	n_australia ++
	fmt.Println("nested map: ",n_australia)
	fmt.Println("nested map: ",n_nairobi)

	/* mm assigned to the inner map but the inner map does not exist yet*/
	// mm := hits["doc/"]
	
	/* causes a panic as inner map does not exist */
	// mm["ke"]++

	add(hits, "doc/", "ke")
	/* struct that we can use as a countryHitsPath in a map */
	type countryHitsPath struct {
		path, country string
	}
	/* map with a struct as a countryHitsPath*/
	betterHits := make(map[countryHitsPath]int)

	/* instance of struct to be used as a map key */
	vietnamHits := countryHitsPath{
		"doc/",
		"vnam"}
	/* insert element with the vietnamHits key and assign it to vietHits*/
	vietHits := betterHits[vietnamHits]
	vietHits++
	fmt.Printf("Hits from Vietnam: %d\n", vietHits)
	/* increment and possibly create values inside the map*/
	betterHits[countryHitsPath{"doc/", "eth"}]++
	for i :=0; i< 20; i++ {
		betterHits[countryHitsPath{"doc/", "eth"}]++
	}
	
	ethiopianHits := betterHits[countryHitsPath{"doc/", "eth"}]
	fmt.Printf("Hits from Ethiopia: %d\n", ethiopianHits)
																																																																																																																																																																																																																											
}


