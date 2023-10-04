package campaign

type MailerInterface interface {
	SendMail(campaign *Campaign) error
}
