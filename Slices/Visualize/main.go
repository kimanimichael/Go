package main

import pretty "github.com/inancgumus/prettyslice"

func main() {
	nums := []int{1, 9, 3, 6, 2, 4}
	odds := nums[:3]
	evens := nums[3:]

	pretty.Show("nums", nums)
	pretty.Show("odds", odds)
	pretty.Show("evens", evens)
}