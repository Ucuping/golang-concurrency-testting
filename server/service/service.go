package service

import "go-concurrency-testting/server/provider"

type usernameCheck struct{}

type usernameService interface {
	UsernameCheck(urls []string) []string
}

var (
	// this is where the usernameCheck struct implemets the usernameService interface
	UsernameService usernameService = &usernameCheck{}
)

func (u *usernameCheck) UsernameCheck(urls []string) []string {
	c := make(chan string)
	var links []string
	matchingLink := []string{}

	for _, url := range urls {
		go provider.Checker.CheckUrl(url, c)
	}
	for i := 0; i < len(urls); i++ {
		links = append(links, <-c)
	}
	// remove the "no_match" and "cant_access-resource" values from the links array
	for _, v := range links {
		if v == "cant_access_resource" {
			continue
		}
		if v == "no_match" {
			continue
		}
		matchingLink = append(matchingLink, v)
	}
	return matchingLink
}
