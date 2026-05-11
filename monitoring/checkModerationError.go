package monitoring

import (
	"complaint-service/mail"
	"complaint-service/repository"
	"database/sql"
	"fmt"
	"log"
)

func CheckModerationError(db *sql.DB, mailer *mail.Mailer) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()

	count, err := repository.GetModerationErrorsNumber(db)
	if err != nil {
		log.Printf("Error in GetModerationErrorsNumber: %v", err)
		return
	}

	if count > 5 {
		log.Println("Sending email to admin")
		err := mail.SendAdminModerationNotification(mailer)
		if err != nil {
			log.Printf("Error in SendAdminModerationNotification: %v", err)
			return
		}
	}
}
