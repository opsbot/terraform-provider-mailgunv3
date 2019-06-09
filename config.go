package main

import (
	"log"

	mailgun "github.com/mailgun/mailgun-go/v3"
)

// Config struct
type Config struct {
	APIKey  string
	BaseURL string
}

// Client returns a new client for accessing mailgun.
func (c *Config) Client() *mailgun.MailgunImpl {
	domain := "" // We don't set a domain right away
	client := mailgun.NewMailgun(domain, c.APIKey)

	if c.BaseURL != "" {
		client.SetAPIBase(c.BaseURL)
	}

	log.Printf("[INFO] Mailgun Client configured ")

	return client
}
