package main

type User struct {
	Name string
	Membership
}

type Membership struct {
	Type             string
	MessageCharLimit int
}

func newUser(name string, membershipType string) User {
	//
	charLimit := 100
	if membershipType == "premium" {
		charLimit = 1000
	}
	u := User{
		Name: name,
		Membership: Membership{
			Type:             membershipType,
			MessageCharLimit: charLimit,
		},
	}
	return u
}
