package eagle

import (
	"encoding/xml"
	"time"

	"github.com/pkg/errors"
)

type CurrentSummationMessage struct {
	DeviceMacId        string
	MeterMacId         string
	Timestamp          time.Time
	SummationDelivered int
	SummationReceived  int
	multiplier         float32
}

// KHw returns how many KWh have been delivered and received by the utility
func (c *CurrentSummationMessage) KWh() (float32, float32) {
	return float32(c.SummationDelivered) * c.multiplier,
		float32(c.SummationReceived) * c.multiplier
}

func (m *Message) CurrentSummation() (*CurrentSummationMessage, error) {
	if m.Type() != CurrentSummationDelivered {
		return nil, InvalidType
	}

	tmp := struct {
		DeviceMacId        string
		MeterMacId         string
		TimeStamp          timestamp
		SummationDelivered hexNumber
		SummationReceived  hexNumber
		Multiplier         hexNumber
		Divisor            hexNumber
	}{}

	if err := xml.Unmarshal(m.Inner, &tmp); err != nil {
		return nil, errors.Wrap(err, "XML Parse Error")
	}

	multiplier := float32(1)
	if m := tmp.Multiplier.Uint64(); m > 1 {
		multiplier = float32(m)
	}

	if d := tmp.Divisor.Uint64(); d > 1 {
		multiplier = multiplier / float32(d)
	}

	msg := &CurrentSummationMessage{
		DeviceMacId:        tmp.DeviceMacId,
		MeterMacId:         tmp.MeterMacId,
		Timestamp:          tmp.TimeStamp.Time(),
		SummationDelivered: int(tmp.SummationDelivered.Uint64()),
		SummationReceived:  int(tmp.SummationReceived.Uint64()),
		multiplier:         multiplier,
	}

	return msg, nil
}
