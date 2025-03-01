package main

import "fmt"

//insert element: m[key] = elem

//get an element: elem= m[key]

//delete: delete(m, key)

//check if key exists:  elem, ok := m[key]
/*
*if key is in m, then ok is true and elem is the value as expected
if the key is not in the map then ok is flase and elem is the zero value for the maps element type

Assignment
complete the deleteIf necessary function , the user struct has a scheduled for deletion field that determines
if they are scheduled for deletion or not

if the user doesnt exist in the map , return the error "not found"
if the user exists but are not schedueld fore deletion return deleted as false with no errors
if they exist and are shceduled for deletion then return deleted as true with no errors and delete their record
from the map

...

like slices maps are also passed by refernece into functions , this meanas that when a map is passed
into a function we write we can make changes to the orgiginal we dong have a copy
*/

func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {

	//check if element exists
	elem, ok := users[name]
	if ok == false {
		return false, fmt.Errorf("not found")
	}
	//check for schedule
	if elem.scheduledForDeletion == false {
		return false, nil
	}

	delete(users, name)

	return true, nil
	//if the exist and are scheduled , return deleted as true and remove from map

}

type user struct {
	name                 string
	number               int
	scheduledForDeletion bool
}
