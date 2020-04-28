package usedesk

import (
	"fmt"
	"log"
	"testing"
)

func Test_Usedesk(t *testing.T) {
	t.Helper()

	api := New("usedesk token")

	m := Ticket{
		"subject": "Отзыв тест о мобильном приложении",
		"message": "Hello World",
	}

	ticketID, err := api.CreateTicket(m)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(ticketID)
}
