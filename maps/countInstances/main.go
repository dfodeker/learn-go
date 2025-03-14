package main

/* Ecah time the user is  sent a message, their user name is logged in a slice , we want a more efifcient way to count how many
messages each user recieved .

implement the getsCounts function it takes a slice of strings messagedUsers and a map of string -> int validUsers
*/

func getCounts(messagedUsers []string, validUsers map[string]int) {

	// for i:=0; i<len(validUsers);i++{
	// 	if _, ok:=messagedUsers
	// }
	for _, user := range messagedUsers {
		if _, key := validUsers[user]; key {
			validUsers[user]++
		}
	}

	//so in this problem we have an array of strings
	/* messaged users and valid users
	 * we need to update the valid users map with the number of times each
	 * user recieved  a message
	 * []string{"Cersei","Jamie","Cersai"}
	 * map[string]int{"Cersei":2,"Jamie:1"}
	 * indicating that cersei recieved 2 messages a jamie recieved one
	 * with that being clear, we need can
	 * go through the messaged users and when its valid increment the matching by 1
	 * I think this may actually be a really good way to solve this
	 *
	 */

}
