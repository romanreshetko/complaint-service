package mail

import (
	service_integrations "complaint-service/service-integrations"
	"fmt"
	"net/smtp"
	"os"
)

type Mailer struct {
	host string
	port string
	from string
}

func NewMailer(from string) *Mailer {
	host := os.Getenv("SMTP_HOST")
	port := os.Getenv("SMTP_PORT")
	return &Mailer{host: host, port: port, from: from}
}

func (m *Mailer) SendMail(to, subject, body string) error {
	user := os.Getenv("SMTP_USER")
	password := os.Getenv("SMTP_PASS")
	addr := fmt.Sprintf("%s:%s", m.host, m.port)
	auth := smtp.PlainAuth("", user, password, m.host)
	msg := []byte("To: " + to + "\r\n" +
		"From: " + m.from + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\"\r\n" +
		"\r\n" +
		body)

	return smtp.SendMail(addr, auth, m.from, []string{to}, msg)
}

func SendBlockNotification(mailer *Mailer, userID int64) error {
	userInfo, err := service_integrations.GetUserForEmail(userID)
	if err != nil {
		return err
	}

	body := fmt.Sprintf(
		"Здравствуйте, %s\n Ваш отзыв был заблокирован. Если вы считаете блокировку "+
			"ошибочной, Вы можете ее обжаловать в личном кабинете на сайте в течение 7 дней.\n"+
			"С уважением,\n Команда cityviewpoint",
		userInfo.Nickname,
	)

	return mailer.SendMail(userInfo.Email, "Уведомление о блокировке", body)
}
