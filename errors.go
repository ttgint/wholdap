package main

import "errors"

var (
	errCouldNotConnect    = errors.New("ldap: could not establish connection")
	errInvalidCredentials = errors.New("ldap: invalid username or password")
	errCouldNotBind       = errors.New("ldap: could not perform initial bind")
	errNoResults          = errors.New("ldap: No result")
	errSearch             = errors.New("ldap: error performing search")
)
