package mailer

import (
    "bytes"
    "fmt"
    "net/smtp"
    "os"
    "text/template"
)

var (
    p, _          = os.Getwd()
    templatesPath = p + "/templates/mail/"
)

type Mail struct {
    From      string
    Template  string
    Subject   string
    Variables interface{}
}

func (mail *Mail) Send(addressee string) error {
    body, err := mail.body()
    if err != nil {
        return err
    }
    err = smtp.SendMail(
        os.Getenv("MAIL_HOST")+":"+os.Getenv("MAIL_PORT"),
        mail.auth(),
        mail.From,
        []string{addressee},
        body)
    if err != nil {
        return err
    }
    return nil
}

func (mail *Mail) auth() smtp.Auth {
    return smtp.PlainAuth("", os.Getenv("MAIL_USERNAME"), os.Getenv("MAIL_PASSWORD"), os.Getenv("MAIL_HOST"))
}

func (mail *Mail) body() ([]byte, error) {
    t, err := template.ParseFiles(templatesPath + mail.Template + ".html")
    if err != nil {
        return nil, err
    }
    var body bytes.Buffer
    mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
    body.Write([]byte(fmt.Sprintf("Subject: "+mail.Subject+" \n%s\n\n", mimeHeaders)))
    err = t.Execute(&body, mail.Variables)
    if err != nil {
        return nil, err
    }
    return body.Bytes(), nil
}
