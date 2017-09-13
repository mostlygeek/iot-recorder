package eagle

import (
	"encoding/xml"
	"time"

	"github.com/pkg/errors"
)

type InstantDemandMessage struct {
	Timestamp   time.Time
	DeviceMacId string
	MeterMacId  string
	Demand      int
}

func (m *Message) InstantaneousDemand() (*InstantDemandMessage, error) {
	if m.Type() != InstantaneousDemand {
		return nil, InvalidType
	}

	tmp := struct {
		DeviceMacId string
		MeterMacId  string
		TimeStamp   timestamp
		Demand      hexNumber
		Multiplier  hexNumber
		Divisor     hexNumber
	}{}

	if err := xml.Unmarshal(m.Inner, &tmp); err != nil {
		return nil, errors.Wrap(err, "XML Parse Error")
	}

	msg := &InstantDemandMessage{
		DeviceMacId: tmp.DeviceMacId,
		MeterMacId:  tmp.MeterMacId,
		Timestamp:   tmp.TimeStamp.Time(),
		Demand:      int(tmp.Demand.Uint64()),
	}

	return msg, nil
}
