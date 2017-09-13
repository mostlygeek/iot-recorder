package eagle

import (
	"encoding/xml"
	"time"

	"github.com/pkg/errors"
)

type MessageClusterMessage struct {
	DeviceMacId          string
	Id                   int
	MeterMacId           string
	Timestamp            time.Time
	Text                 string
	Priority             string
	ConfirmationRequired bool
	Confirmed            bool
	Queue                string
}

func (m *Message) MessageCluster() (*MessageClusterMessage, error) {
	if m.Type() != MessageCluster {
		return nil, InvalidType
	}

	tmp := struct {
		DeviceMacId          string
		MeterMacId           string
		TimeStamp            timestamp
		Id                   hexNumber
		Text                 string
		Priority             string
		ConfirmationRequired string
		Confirmed            string
		Queue                string
	}{}

	if err := xml.Unmarshal(m.Inner, &tmp); err != nil {
		return nil, errors.Wrap(err, "XML Parse Error")
	}

	mc := &MessageClusterMessage{
		DeviceMacId: tmp.DeviceMacId,
		MeterMacId:  tmp.MeterMacId,
		Id:          int(tmp.Id.Uint64()),
		Timestamp:   tmp.TimeStamp.Time(),
		Text:        tmp.Text,
		Priority:    tmp.Priority,
		Queue:       tmp.Queue,
	}

	if tmp.ConfirmationRequired == "Y" {
		mc.ConfirmationRequired = true
	}
	if tmp.Confirmed == "Y" {
		mc.Confirmed = true
	}

	return mc, nil
}
