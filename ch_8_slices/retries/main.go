package main

import (
	"errors"
)

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {
	if plan == planFree {
		return messages[:2], nil
	} else if plan == planPro {
		return messages[:], nil
	} else {

		return nil, errors.New("unsupported plan")
	}
}

func gcdOfStrings(str1 string, str2 string) string {

	//create a slice from string a
	// slice2:= []string{str2}
	// slice1:=[]string{str1}

	//return the largest string that divides both
	str3 := []rune{}
	//given two strings "abc" and "abcd"
	//find the greatest divisible denominator
	// in our example we check if str1[0] is equal to str2[0] if it is we add it to str3 if not we skip
	// if we have equality , then we move forward str[1]

	if len(str2) >= len(str1) {
		return ""
	}

	for _, char1 := range str1 {

		for _, char2 := range str2 {
			if char1 == char2 {
				str3 = append(str3, char2)
			} else {
				continue
			}
		}
		continue
	}
	return string(str3)

}
