package services


type IEmailService interface {
	SendWelcomeEmail(email string) error
	SendReferralLinkAccess(email string) error
}

type EmailConfig struct {
	SMTPEmail string
	SMTPHost    string
	SMTPPort    int
	SMTPUser    string
	SMTPPassword   string
}

type EmailService struct {
	EmailConfig EmailConfig
}
