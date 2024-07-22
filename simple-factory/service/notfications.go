package service

import "fmt"

type Notification interface {
	Send(message string)
}

type EmailNotification struct {
	recipient string
}

func (e *EmailNotification) Send(message string) {
	fmt.Printf("Sending email to %s: %s\n", e.recipient, message)
}

type SMSNotification struct {
	recipient string
}

func (s *SMSNotification) Send(message string) {
	fmt.Printf("Sending SMS to %s: %s\n", s.recipient, message)
}

func New(notificationType string, recipient string) Notification {
	switch notificationType {
	case "email":
		return &EmailNotification{
			recipient: recipient,
		}
	case "sms":
		return &SMSNotification{
			recipient: recipient,
		}
	default:
		return nil
	}
}
