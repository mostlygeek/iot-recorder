package eagle

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

type DeviceInfoMessage struct {
	DeviceMacId  string
	InstallCode  string
	LinkKey      string
	FWVersion    string
	HWVersion    string
	ImageType    string
	Manufacturer string
	ModelId      string
	DateCode     string
	Port         string
}

func (m *Message) DeviceInfo() (*DeviceInfoMessage, error) {
	if m.Type() != DeviceInfo {
		return nil, InvalidType
	}

	payload := &DeviceInfoMessage{}

	if err := xml.Unmarshal(m.Inner, &payload); err != nil {
		return nil, errors.Wrap(err, "XML Parse Error")
	} else {
		return payload, nil
	}
}
