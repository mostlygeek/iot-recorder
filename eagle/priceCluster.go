package eagle

import (
	"encoding/xml"
	"math"
	"time"

	"github.com/pkg/errors"
)

type PriceClusterMessage struct {
	DeviceMacId string
	MeterMacId  string
	Timestamp   time.Time
	Price       float32
	Currency    int
	Tier        int
	TierLabel   string
	RateLabel   string
}

func (m *Message) PriceCluster() (*PriceClusterMessage, error) {
	if m.Type() != PriceCluster {
		return nil, InvalidType
	}

	tmp := struct {
		DeviceMacId    string
		MeterMacId     string
		TimeStamp      timestamp
		Price          hexNumber
		Currency       hexNumber
		TrailingDigits hexNumber
		Tier           hexNumber
		TierLabel      string
		RateLabel      string
	}{}

	if err := xml.Unmarshal(m.Inner, &tmp); err != nil {
		return nil, errors.Wrap(err, "XML Parse Error")
	}

	pc := &PriceClusterMessage{
		DeviceMacId: tmp.DeviceMacId,
		MeterMacId:  tmp.MeterMacId,
		Timestamp:   tmp.TimeStamp.Time(),
		Currency:    int(tmp.Currency.Uint64()),
		Tier:        int(tmp.Tier.Uint64()),
		TierLabel:   tmp.TierLabel,
		RateLabel:   tmp.RateLabel,
	}

	digits := int(tmp.TrailingDigits.Uint64())
	pc.Price = float32(tmp.Price.Uint64()) / float32(math.Pow10(digits))

	return pc, nil
}
