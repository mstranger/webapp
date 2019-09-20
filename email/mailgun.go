package email

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/mailgun/mailgun-go"
)

const (
	welcomeSubject = "Welcome to website.com!"
	resetSubject   = "Instructions for resetting your password."
	resetBaseURL   = "https://www.example.com/reset"

	welcomeText = `Hi there!

	Welcome to Website.com! We really hope you enjoy using
	our application!

	Best,
	Name
	`

	welcomeHTML = `Hi there!<br>
	<br>
	Welcome to
	<a href="https://www.website.com">website.com</a>!
	We really hope you enjoy using our application!<br>
	<br>
	Best,<br>
	Name
	`
)

const resetTextTmpl = `Hi there!

It appears that you have have request a password reset.
If this was you, please follow the link below to update
your password:

%s

If you are asked for a token, please use the following value:

%s

If you didn't request a password reset you can safely ignore
this email and your account will not be changed.

Best,
Name Support
`

const resetHTMLTmpl = `Hi there!<br>
<br>
It appears that you have have request a password reset.
If this was you, please follow the link below to update
your password:<br>
<br>
<a href="%s">%s</a><br>
<br>
If you are asked for a token, please use the following value:<br>
<br>
%s<br>
<br>
If you didn't request a password reset you can safely ignore
this email and your account will not be changed.<br>
<br>
Best,<br>
Name Support<br>
`

// func WithMailgun(domain, apiKey, publicKey string) ClientConfig {
func WithMailgun(domain, apiKey string) ClientConfig {
	return func(c *Client) {
		// mg := mailgun.NewMailgun(domain, apiKey, publicKey)
		mg := mailgun.NewMailgun(domain, apiKey)
		c.mg = mg
	}
}

func WithSender(name, email string) ClientConfig {
	return func(c *Client) {
		c.from = buildEmail(name, email)
	}
}

type ClientConfig func(*Client)

func NewClient(opts ...ClientConfig) *Client {
	client := Client{
		// Set a default from email address...
		from: "support@website.com",
	}
	for _, opt := range opts {
		opt(&client)
	}

	return &client
}

type Client struct {
	from string
	mg   mailgun.Mailgun
}

func (c *Client) Welcome(toName, toEmail string) error {
	message := c.mg.NewMessage(c.from, welcomeSubject, welcomeText, buildEmail(toName, toEmail))
	message.SetHtml(welcomeHTML)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	_, _, err := c.mg.Send(ctx, message)
	return err
}

func (c *Client) ResetPw(toEmail, token string) error {
	v := url.Values{}
	v.Set("token", token)
	resetUrl := resetBaseURL + "?" + v.Encode()
	resetText := fmt.Sprintf(resetTextTmpl, resetUrl, token)
	message := c.mg.NewMessage(c.from, resetSubject, resetText, toEmail)
	resetHTML := fmt.Sprintf(resetHTMLTmpl, resetUrl, resetUrl, token)
	message.SetHtml(resetHTMLTmpl)
	_, _, err := c.mg.Send(ctx, message)
	return err
}

func buildEmail(name, email string) string {
	if name == "" {
		return email
	}
	return fmt.Sprintf("%s <%s>", name, email)
}
