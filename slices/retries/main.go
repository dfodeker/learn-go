package main

import (
	"errors"
)

//primes := [6]int{2, 3, 5, 7, 11, 13}
// fmt.Println(primes)
// mySlice := primes[1:5]
// fmt.Println(mySlice)

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {
	//?
	//if the plan is a pro plan return all the strings from the message input in a slice
	//if the plan is a free plan, return the first 2 strings from the messages input ina slice
	//if the plan isnt either of those return a nil slice and an error that says unsupported plan

	if plan == planFree {
		return messages[:2], nil
	} else if plan == planPro {
		return messages[:], nil
	} else {

		return nil, errors.New("unsupported plan")
	}
}
