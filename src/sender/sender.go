package sender

import (
	"github.com/clanbeat/mailer/Godeps/_workspace/src/github.com/sendgrid/sendgrid-go"
	"html/template"
	"log"
)

type SendGridConf struct {
	Username      string
	Secret        string
	TemplateCache map[string]*template.Template
}

var Config *SendGridConf

const fromEmail = "Clanbeat <hello@clanbeat.com>"

func SendgridSetup(username, secret string) {
	Config = &SendGridConf{
		Username: username,
		Secret:   secret,
	}
}

func CacheTemplates(path string) error {
	if cache, err := BuildTemplateCache(path); err != nil {
		return err
	} else {
		Config.TemplateCache = cache
	}
	return nil
}

func TemplateExists(name string) bool {
	return Config.TemplateCache[name] != nil
}

func Send(em *Mailable) error {
	if err := em.validate(); err != nil {
		return err
	}
	sg := sendgrid.NewSendGridClient(Config.Username, Config.Secret)
	message := sendgrid.NewMail()
	message.SetFrom(em.From)
	message.AddTo(em.To)
	message.SetSubject(em.Subject)

	msg, err := em.content(Config.TemplateCache)
	if err != nil {
		return err
	}

	message.SetHTML(msg)
	log.Println("Email sent to", em.To)
	if err := sg.Send(message); err == nil {
		return nil
	} else {
		log.Println("no email sent")
		log.Println(err)
		return err
	}
}
