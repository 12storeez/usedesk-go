package usedesk

import (
	"github.com/go-resty/resty/v2"
)

const (
	BASE_URL      = "https://api.usedesk.ru"
	CREATE_TICKET = "/create/ticket"
)

type usedesk struct {
	defaults map[string]string
	client   *resty.Client
}

func New(apiToken string) *usedesk {
	return &usedesk{
		defaults: map[string]string{"api_token": apiToken},
		client:   resty.New().SetHostURL(BASE_URL),
	}
}

func (u *usedesk) CreateTicket(t Ticket) (ticketID int64, err error) {
	params := appendMaps(u.defaults, t)
	u.client.SetQueryParams(params)
	resp, err := u.client.R().Post(CREATE_TICKET)
	if err != nil {
		return 0, err
	}
	defer resp.RawBody().Close()

	var result CreateTicketResponse

	return result.Response(resp.Body())
}

func appendMaps(source, target map[string]string) map[string]string {
	for key, value := range source {
		target[key] = value
	}

	return target
}
