- Mailgun (mailgun.com)

```go
import mailgun "github.com/mailgun/mailgun-go"
from := "demo@sandboxd35fb954a4614a17b307669a656b0767.mailgun.org"
to := "John Doe <opponere@ukr.net>"
subject := "Welcome to Website.com!"
text := `Hi there!
Welcome to out awesome website!

Thanks for signing up!
ms`

html := `Hi there!<br>
Welcome to <a href="https://website.com">website.com</a>!
`

msg := mailgun.NewMessage(from, subject, text, to)
msg.SetHtml(html)

cfg := LoadConfig(...)
mgCfg := cfg.Mailgun
mgClient := mailgun.NewMailgun(mgCfg.Domain, mgCfg.APIKey, mgCfg.PublickAPIKey)

respMsg, id, err := mgClient.Send(msg)
```

Steps:

1. Create an account and get info
2. Config
3. How to create a message
4. Email address fromats
5. HTML messages
6. Sending with the Mailgun client
7. Creating an email package
8. Adding emailer to controllers
