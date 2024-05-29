package facade

import (
	"errors"

	"github.com/rehok/go-message"
	"github.com/rehok/go-message/drivers"
)

type Message struct {
	Active string  `json:"active"`
	Driver *Driver `json:"driver"`
}

func NewMessage() *Message {
	return &Message{
		Active: "sendge",
		Driver: NewDriver(),
	}
}

func (s *Message) ActiveDriver() (message.Message, error) {
	return s.Driver.Get(s.Active)
}

type Driver struct {
	Sendge *drivers.Sendge `json:"sendge"`
}

func NewDriver() *Driver {
	return &Driver{
		Sendge: &drivers.Sendge{},
	}
}

func (d *Driver) Items() []message.Message {
	return []message.Message{
		d.Sendge,
	}
}

func (d *Driver) Get(id string) (message.Message, error) {
	if id == "" {
		return nil, errors.New("id undefined")
	}
	switch id {
	case "sendge":
		return d.Sendge, nil
	}
	return nil, errors.New("driver not found")
}
