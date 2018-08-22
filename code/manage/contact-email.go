package manage

import (
	"net/smtp"

	"broadcastle.co/code/crm/code/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Email sends a email to a contact.
func Email(cmd *cobra.Command, args []string) {

	Init()
	defer Close()

	// id, err := utils.UfS(args[0])
	// if err != nil {
	// 	logrus.Fatal(err)
	// }

	// contact := db.Contact{}
	// contact.ID = id

	// if err := contact.Query(); err != nil {
	// 	logrus.Fatal(err)
	// }

	contacts, err := utils.Contacts(cmd, args)
	if err != nil {
		logrus.Fatal(err)
	}

	for _, contact := range contacts {

		header, err := utils.Input("subject", "")
		if err != nil {
			logrus.Fatal(err)
		}

		body, err := utils.Input("message", "")
		if err != nil {
			logrus.Fatal(err)
		}

		send, err := utils.InputBool("send", false)
		if err != nil {
			logrus.Fatal(err)
		}

		if !send {
			logrus.Info("email aborted")
			return
		}

		auth := smtp.PlainAuth("", viper.GetString("email.username"), viper.GetString("email.password"), viper.GetString("email.server"))

		emailserver := viper.GetString("email.server") + ":" + viper.GetString("email.port")

		message := "To: " + contact.Email + "\r\n" + "Subject: " + header + "\r\n\r\n" + body + "\r\n"

		if err := smtp.SendMail(emailserver,
			auth,
			viper.GetString("email.from"),
			[]string{contact.Email},
			[]byte(message)); err != nil {
			logrus.Fatal(err)
		}

	}

}
