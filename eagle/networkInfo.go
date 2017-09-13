package eagle

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

// NetworkInfoMessage contains information about the zigbee network the
// EAGLE is connecting to
type NetworkInfoMessage struct {
	DeviceMacId  string
	CoordMacId   string
	Status       string
	Description  string
	ExtPanId     string
	Channel      int
	ShortAddr    string
	LinkStrength int
}

func (m *Message) NetworkInfo() (*NetworkInfoMessage, error) {
	if m.Type() != NetworkInfo {
		return nil, InvalidType
	}

	tmp := struct {
		DeviceMacId  string
		CoordMacId   string
		Status       string
		Description  string
		ExtPanId     string
		Channel      hexNumber
		ShortAddr    string
		LinkStrength hexNumber
	}{}

	if err := xml.Unmarshal(m.Inner, &tmp); err != nil {
		return nil, errors.Wrap(err, "XML Parse Error")
	}

	return &NetworkInfoMessage{
		DeviceMacId:  tmp.DeviceMacId,
		CoordMacId:   tmp.CoordMacId,
		Status:       tmp.Status,
		Description:  tmp.Description,
		ExtPanId:     tmp.ExtPanId,
		Channel:      int(tmp.Channel.Uint64()),
		ShortAddr:    tmp.ShortAddr,
		LinkStrength: int(tmp.LinkStrength.Uint64()),
	}, nil
}
