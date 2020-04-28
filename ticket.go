package usedesk

import (
	"encoding/json"
	"errors"
)

type Ticket map[string]string

type CreateTicketResponse struct {
	Status   string `json:"status"`
	TicketID int64  `json:"ticket_id"`
}

func (c *CreateTicketResponse) Response(data []byte) (ticketID int64, err error) {
	if err := json.Unmarshal(data, c); err != nil {
		return 0, err
	}

	if c.Status != "success" {
		return 0, errors.New(string(data))
	}

	return c.TicketID, nil
}
