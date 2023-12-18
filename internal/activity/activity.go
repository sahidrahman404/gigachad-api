package activity

import (
	"bytes"
	"time"

	"github.com/sahidrahman404/gigachad-api/assets"
	"github.com/sahidrahman404/gigachad-api/internal/funcs"

	"github.com/go-mail/mail/v2"

	htmlTemplate "html/template"
	textTemplate "text/template"
)

const SendEmailTaskQueueName = "SEND_EMAIL_TASK_QUEUE"

type Mailer struct {
	dialer *mail.Dialer
	from   string
}

type MailDetails struct {
	Recipient string
	Data      map[string]interface{}
	Patterns  []string
}

func NewMailer(host string, port int, username, password, from string) *Mailer {
	dialer := mail.NewDialer(host, port, username, password)
	dialer.Timeout = 5 * time.Second

	return &Mailer{
		dialer: dialer,
		from:   from,
	}
}

func (m *Mailer) Send(data MailDetails) error {
	for i := range data.Patterns {
		data.Patterns[i] = "emails/" + data.Patterns[i]
	}

	msg := mail.NewMessage()
	msg.SetHeader("To", data.Recipient)
	msg.SetHeader("From", m.from)

	ts, err := textTemplate.New("").Funcs(funcs.TemplateFuncs).ParseFS(assets.EmbeddedFiles, data.Patterns...)
	if err != nil {
		return err
	}

	subject := new(bytes.Buffer)
	err = ts.ExecuteTemplate(subject, "subject", data.Data)
	if err != nil {
		return err
	}

	msg.SetHeader("Subject", subject.String())

	plainBody := new(bytes.Buffer)
	err = ts.ExecuteTemplate(plainBody, "plainBody", data.Data)
	if err != nil {
		return err
	}

	msg.SetBody("text/plain", plainBody.String())

	if ts.Lookup("htmlBody") != nil {
		ts, err := htmlTemplate.New("").Funcs(funcs.TemplateFuncs).ParseFS(assets.EmbeddedFiles, data.Patterns...)
		if err != nil {
			return err
		}

		htmlBody := new(bytes.Buffer)
		err = ts.ExecuteTemplate(htmlBody, "htmlBody", data.Data)
		if err != nil {
			return err
		}

		msg.AddAlternative("text/html", htmlBody.String())
	}

	err = m.dialer.DialAndSend(msg)

	return err
}
