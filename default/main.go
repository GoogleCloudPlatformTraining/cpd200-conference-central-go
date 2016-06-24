package main

import (
)

func init() {
	http.HandleFunc("/crons/set_announcement", SetAnnouncementHandler)
	http.HandleFunc("/tasks/send_confirmation_email", SendConfirmationEmailHandler)
}
