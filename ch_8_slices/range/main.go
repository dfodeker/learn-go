package main

func indexOfFirstBadWord(msg []string, badWords []string) int {

	for i, word := range msg {
		//this would log our word and index here
		//lets compare our badwords now
		//only we dont know how to compare 2 slices yet
		//my best guess would be to create another range with either a j or not use it at all with _
		for _, badWord := range badWords {
			if word == badWord {
				return i
			}
		}
	}
	return -1
}

/*

go provides sytactic sugar to iterate easily over elements of a slice
for INDEX, ELEMENT:= range SLICE {
}
the element is a copy of the value at that index
fruits:= []string{"apple", "banana", "grape"}
for i, fruit:=range fruits{
fmt.Println(i,fruit)
}
*/
