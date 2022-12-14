package util

import (
	"strconv"
	"strings"
)

func IntSliceToString(intSlice []int, separator string) string {
	result := ""
	if len(intSlice) > 0 {
		stringSlice := make([]string, len(intSlice))
		for index, item := range intSlice {
			stringSlice[index] = strconv.Itoa(item)
		}
		result = strings.Join(stringSlice, separator)
	}
	return result
}

func IntInSlice(intSlice []int, value int) bool {
	for _, item := range intSlice {
		if item == value {
			return true
		}
	}
	return false
}

func MaskMail(email string) string {
	emailSlice := strings.Split(email, "@")
	if len(emailSlice) > 1 {
		mask := strings.Repeat("X", len(emailSlice[0]))
		return mask + "@" + emailSlice[1]
	} else {
		return email
	}
}
