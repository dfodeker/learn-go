package main

func getNameCounts(names []string) map[rune]map[string]int {
	//magic here apparently iono what all that is, ahh emojis
	//nested structures
	nested := make(map[rune]map[string]int)
	for i := 0; i < len(names); i++ {
		//get first character
		name := []rune(names[i])
		initial := name[0]
		//fmt.Println(name)

		if _, ok := nested[initial]; !ok {
			nested[initial] = make(map[string]int)
		}
		nested[initial][names[i]] += 1 //this should increnemt

	}

	return nested
}

/*
assignment,
complete the getNameCounts function , it takes a slice of strings (names)
and returns a nested map, the parent maps keys are all the uniques
 first characters of the name , the nested map keys are all the names themselves

 the value is the count of each name
 for example
 billy
 billy
 bob
 joe

 should create the following nested map
 b:{
	billy:2
	bob:1
 },
 j:{
	joe:1
 }
	the test suite is not printing the map, directly but instead checking
	a few specific keys
	we can conver a string into a slice of runes as follows
	word:=hello
	runes:=[]rune(word)
	[h,e,l,l,o]


*/
