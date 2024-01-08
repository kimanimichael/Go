package mystrings

/* First letter for Reverse func needs to be capitalized to allow for importation to other packages */
func Reverse(s string) string {
	result := ""

	for _, v := range s {
		result = string(v) + result
	}
	return result
}