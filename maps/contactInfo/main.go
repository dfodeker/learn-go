package main

import "fmt"

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	// ?
	userMap := make(map[string]user)
	if len(names) != len(phoneNumbers) {
		return nil, fmt.Errorf("invalid sizes")
	}

	for i := 0; i < len(names); i++ {
		userMap[names[i]] = user{
			name:        names[i],
			phoneNumber: phoneNumbers[i],
		}

	}
	return userMap, nil

}

type user struct {
	name        string
	phoneNumber int
}

// We can speed up our contact-info lookups by using a map!

// Key-based map lookup: O(1)
// Slice brute-force search: O(n)
// Complete the getUserMap function. It takes a slice of names and a slice of phone numbers, and returns a map of name -> user structs and an error.
// A user struct just contains a user's name and phone number. The first name in the names slice pairs with the first phone number, and so on.

// If the length of names and phoneNumbers is not equal, return an error with the string "invalid sizes".

// ages := make(map[string]int)
// ages["John"] = 37
// ages["Mary"] = 24
// ages["Mary"] = 21 // overwrites 24

// ages = map[string]int{
// 	"John": 37,
// 	"Mary": 21,
//   }
