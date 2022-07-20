package form

import (
	"errors"
	"regexp"
	"strings"

	"github.com/jinzhu/gorm"

	"github.com/FaridehGhani/go_form/infra/mysql"
)

// Contact form model
type Contact struct {
	gorm.Model
	Email   string `json:"email"`
	Content string `json:"content"`
	Errors  string `json:"errors"`
}

var rxEmail = regexp.MustCompile(".+@.+\\..+")

// Validate contact form
func (c *Contact) Validate() bool {
	match := rxEmail.Match([]byte(c.Email))
	if match == false {
		c.Errors = "please enter a valid email address"
		return false
	}

	if strings.TrimSpace(c.Content) == "" {
		c.Errors = "please enter a message"
		return false
	}

	return true
}

/* database service*/

func (c *Contact) CreateContent() (*Contact, error) {
	contact := c
	if err := mysql.Connect().Create(&contact).Error; err != nil {
		return nil, err
	}

	return contact, nil
}

func (c *Contact) GetContact(email string) (*Contact, error) {
	var content Contact
	if mysql.Connect().Where("email = ?", email).Find(&content).RecordNotFound() {
		return nil, errors.New("contact form not found")
	}

	return &content, nil
}
