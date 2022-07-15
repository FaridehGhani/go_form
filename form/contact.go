package form

import (
	"regexp"
	"strings"
)

var rxEmail = regexp.MustCompile(".+@.+\\..+")

type Contact struct {
	Email   string
	Content string
	Errors  map[string]string
}

func (msg *Contact) Validate() bool {
	msg.Errors = make(map[string]string)

	match := rxEmail.Match([]byte(msg.Email))
	if match == false {
		msg.Errors["Email"] = "please enter a valid email address"
	}

	if strings.TrimSpace(msg.Content) == "" {
		msg.Errors["Content"] = "please enter a message"
	}

	return len(msg.Errors) == 0
}
