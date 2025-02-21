package main

type cost struct {
	day   int
	value float64
}

func getDayCosts(costs []cost, day int) []float64 {
	// ?

	slice := make([]float64, 0)
	askDay := day
	for i := 0; i < len(costs); i++ {
		//if this day
		currentDay := costs[i].day
		currentValue := costs[i].value

		sameday := currentDay == askDay
		if sameday {
			if currentValue == 0.0 {
			} else {
				slice = append(slice, currentValue)
			}

		}
	}

	return slice

}

/*
Editorial
Append
The built-in append function is used to dynamically add elements to a slice:

func append(slice []Type, elems ...Type) []Type

If the underlying array is not large enough, append() will create a new underlying array and point the returned slice to it.

Notice that append() is variadic, the following are all valid:

slice = append(slice, oneThing)
slice = append(slice, firstThing, secondThing)
slice = append(slice, anotherSlice...)

*/

/*
We've been asked to add a feature that extracts costs for a given day.

Complete the getDayCosts() function using the append() function. It accepts a slice of cost structs and a day int, and it returns a float64 slice containing that day's costs.

If there are no costs for that day, return an empty slice.
*/
