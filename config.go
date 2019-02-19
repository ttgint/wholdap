package main

var config = struct {
	URL        string
	BaseDN     string
	Attributes map[string]string
}{
	"url",
	"basedn",
	map[string]string{
		"username":   "sAMAccountName",
		"mail":       "mail",
		"name":       "displayName",
		"first":      "givenName",
		"last":       "sn",
		"company":    "company",
		"title":      "title",
		"location":   "l",
		"department": "department",
	},
}
