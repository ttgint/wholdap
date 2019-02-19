package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"strings"

	ldap "gopkg.in/ldap.v2"
)

func main() {
	username := flag.String("u", "", "BaseUser for search")
	password := flag.String("p", "", "Password for BaseUser")
	flag.Parse()

	lookup := flag.Args()[0]

	result, err := Find(lookup, "sAMAccountName", *username, *password)

	if err == nil {
		j, _ := json.MarshalIndent(result, "", "")
		log.Println(string(j))
		return
	}

	result, err = Find(lookup, "mail", *username, *password)

	if err == nil {
		j, _ := json.MarshalIndent(result, "", "")
		log.Println(string(j))
		return
	}
	log.Println(err)
}

// Connect connects and returns the connection handle to the specified LDAP server
func Connect(url string) (*ldap.Conn, error) {
	l, err := ldap.Dial("tcp", url)
	if err != nil {
		err = errors.New("Cannot contact LDAP server")
	}

	return l, err
}

// Find from ldap and return entry fields
func Find(lookup, attr, username, password string) (map[string]string, error) {
	user := map[string]string{}

	if len(lookup) < 1 {
		log.Println(errInvalidCredentials)
		return user, errInvalidCredentials
	}

	l, err := Connect(config.URL)

	if err != nil {
		log.Println(err)
		return user, errCouldNotConnect
	}

	defer l.Close()

	err = l.Bind(username, password)

	if err != nil {
		log.Println(err)
		return user, errCouldNotBind
	}

	searchRequest := ldap.NewSearchRequest(
		config.BaseDN, // The base dn to search
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0, 0, false,
		fmt.Sprintf("(%v=%v)", attr, lookup),
		getAttrs(), // A list attributes to retrieve
		nil,
	)

	searchResult, err := l.Search(searchRequest)

	if err != nil {
		log.Println(err)
		return user, errSearch
	}

	if len(searchResult.Entries) < 1 {
		log.Println(errNoResults)
		return user, errNoResults
	}

	return parseEntry(searchResult.Entries[0]), nil
}

func getAttrs() []string {
	attrs := make([]string, 0, len(config.Attributes))
	for _, val := range config.Attributes {
		attrs = append(attrs, val)
	}

	return attrs
}

func parseEntry(entry *ldap.Entry) map[string]string {
	user := map[string]string{}

	for key, attr := range config.Attributes {
		user[key] = strings.ToLower(entry.GetAttributeValue(attr))
	}

	return user
}
