package utils

// RemoveDuplicate function to remove duplicate data from array string
// Params:
// slice: array string
// Returns duplicate-removed array string
func RemoveDuplicate(slice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// InArrayExist function to check if value exists in array string
// Params:
// val: string value
// array: array string to check
// Returns boolean
func InArrayExist(val string, array []string) (exists bool) {
	exists = false
	for i := 0; i < len(array); i++ {
		if val == array[i] {
			exists = true
		}
	}
	return
}
