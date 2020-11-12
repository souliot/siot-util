package sutil

import (
	"errors"

	"gopkg.in/gomail.v2"
)

var (
	ErrNoSender  = errors.New("The mail sender should not be nil")
	ErrNoSubject = errors.New("The mail subject should not be nil")
	ErrNoBody    = errors.New("The mail body should not be nil")
	ErrNoToers   = errors.New("The mail send toers should not be nil")
)

type Mail struct {
	ServerHost string
	ServerPort int
	From       string
	Alias      string
	FromPwd    string
}

func NewMail(host string, port int, username, alias, pwd string) *Mail {
	return &Mail{
		ServerHost: host,
		ServerPort: port,
		From:       username,
		Alias:      alias,
		FromPwd:    pwd,
	}
}

func (m *Mail) Send(msg *MailMessage, opts ...MessageOption) (err error) {
	if msg == nil {
		msg = new(MailMessage)
	}
	msg.apply(opts)
	err = m.checkMessage(msg)
	if err != nil {
		return
	}
	gm := gomail.NewMessage()
	// 设置发件人
	gm.SetHeader("From", gm.FormatAddress(m.From, m.Alias))

	// 设置收件人
	toers := make([]string, 0)
	for to, alias := range msg.Toers {
		toers = append(toers, gm.FormatAddress(to, alias))
	}
	gm.SetHeader("To", toers...)

	// 设置抄送人
	if len(msg.Ccers) > 0 {
		ccers := make([]string, 0)
		for cc, alias := range msg.Ccers {
			ccers = append(ccers, gm.FormatAddress(cc, alias))
		}
		gm.SetHeader("Cc", ccers...)
	}

	// 设置抄送人
	if len(msg.Bccers) > 0 {
		bccers := make([]string, 0)
		for bcc, alias := range msg.Bccers {
			bccers = append(bccers, gm.FormatAddress(bcc, alias))
		}
		gm.SetHeader("Bcc", bccers...)
	}

	// 设置邮件主题
	gm.SetHeader("Subject", msg.Subject)
	// 设置邮件内容
	gm.SetBody("text/html", msg.Body)
	err = gomail.NewDialer(m.ServerHost, m.ServerPort, m.From, m.FromPwd).DialAndSend(gm)
	return
}

func (m *Mail) checkMessage(msg *MailMessage) (err error) {
	if m.From == "" {
		return ErrNoSender
	}
	if msg.Subject == "" {
		return ErrNoSubject
	}
	if msg.Body == "" {
		return ErrNoBody
	}
	if len(msg.Toers) <= 0 {
		return ErrNoToers
	}
	return
}

type MailMessage struct {
	Subject string
	Body    string
	Files   map[string]string
	Toers   map[string]string
	Ccers   map[string]string
	Bccers  map[string]string
}

func (m *MailMessage) apply(opts []MessageOption) {
	for _, opt := range opts {
		opt(m)
	}
}

type MessageOption func(*MailMessage)

func WithSubject(subject string) MessageOption {
	return func(m *MailMessage) {
		m.Subject = subject
	}
}

func WithBody(body string) MessageOption {
	return func(m *MailMessage) {
		m.Body = body
	}
}

func WithFiles(files map[string]string) MessageOption {
	return func(m *MailMessage) {
		m.Files = files
	}
}

func WithToers(toers map[string]string) MessageOption {
	return func(m *MailMessage) {
		m.Toers = toers
	}
}

func WithCcers(ccers map[string]string) MessageOption {
	return func(m *MailMessage) {
		m.Ccers = ccers
	}
}

func WithBccers(bccers map[string]string) MessageOption {
	return func(m *MailMessage) {
		m.Bccers = bccers
	}
}
