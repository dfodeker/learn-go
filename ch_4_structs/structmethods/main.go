package main

import (
	"fmt"
)

type authenticationInfo struct {
	username string
	password string
}

func (r authenticationInfo) getBasicAuth() string {
	s := fmt.Sprintf("Authorization: Basic %v:%v", r.username, r.password)
	return s
}
