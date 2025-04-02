package main

func getLast[T any](s []T) T {
	if len(s) == 0 {
		var empty T
		return empty
	}
	len := len(s)
	last := s[len-1]
	return last
	// ?
}

//generic function that returns the last element from a slice
/* sorryyyy
usage// example
func splitAnySlice[T any](s []T) ([]T, []T) {
    mid := len(s)/2
    return s[:mid], s[mid:]
}
At Mailio we store all the emails for a campaign in memory as a slice. We store payments for a single user in the same way.

Complete the getLast() function. It should be a generic function that returns the last element from a slice, no matter the types stored in the slice. If the slice is empty, it should return the zero value of the type.

Tip: Zero Value of a Type
Creating a variable that's the zero value of a type is easy:

var myZeroInt int

It's the same with generics, we just have a variable that represents the type:

var myZero T
*/
