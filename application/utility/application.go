package utility

import (
	"encoding/json"
	tempmail "github.com/hikouki1111/tempmail-wrapper"
	"log"
)

var Data *Userdata

func AddAccount() {
	acc, err := tempmail.NewAccount()
	if err != nil {
		log.Println(err)
	}
	Data.Accounts = append(Data.Accounts, *acc)
}

func DeleteAccount(token string) {
	for i, acc := range Data.Accounts {
		if token == acc.Token {
			err := acc.Delete()
			if err != nil {
				log.Println(err)
			}
			Data.Accounts = append(Data.Accounts[:i], Data.Accounts[i+1:]...)
			return
		}
	}
}

func GetAccounts() []map[string]interface{} {
	accounts := make([]map[string]interface{}, 0)
	for _, acc := range Data.Accounts {
		accounts = append(accounts, map[string]interface{}{"email": acc.Email, "token": acc.Token})
	}

	return accounts
}

func GetMailbox(email string) []string {
	var mailbox []tempmail.Mail
	var mails []string
	for _, acc := range Data.Accounts {
		if acc.Email == email {
			var err error
			mailbox, err = acc.GetMailbox()
			if err != nil {
				log.Println(err)
				return nil
			}
			break
		}
	}

	if mailbox != nil {
		for _, mail := range mailbox {
			jsonData, err := json.Marshal(mail)
			if err != nil {
				log.Println(err)
				return nil
			}
			mails = append(mails, string(jsonData))
		}
	}

	return mails
}

func GetAttachments(email, id string) string {
	var mailbox []tempmail.Mail
	var attachments string
	for _, acc := range Data.Accounts {
		if acc.Email == email {
			var err error
			mailbox, err = acc.GetMailbox()
			if err != nil {
				log.Println(err)
				return ""
			}
			break
		}
	}

	if mailbox != nil {
		for _, mail := range mailbox {
			if mail.ID == id {
				mailAttachments := make([]tempmail.Attachment, 0)
				for _, a := range mail.Attachments {
					mailAttachments = append(mailAttachments, a)
				}

				jsonData, err := json.Marshal(mailAttachments)
				if err != nil {
					log.Println(err)
					return ""
				}
				attachments = string(jsonData)
				break
			}
		}
	}

	return attachments
}
