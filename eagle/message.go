package eagle

import (
	"bytes"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

var (
	InvalidType        = errors.New("Invalid Message Type")
	UnsupportedMessage = errors.New("Unsupported Message Type")
)

type hexNumber string

func (h hexNumber) Uint64() uint64 {
	if v, err := strconv.ParseUint(string(h), 0, 64); err != nil {
		return 0
	} else {
		return v
	}
}

type timestamp string

// Seconds is the seconds since the unix epoch. Returns 0 if error
func (t timestamp) Time() time.Time {
	if v, err := strconv.ParseUint(string(t), 0, 64); err != nil {
		return time.Unix(0, 0)
	} else {
		// seconds Jan 1, 2000 00:00:00 since the unix epoch
		return time.Unix(int64(v+946598400), 0)
	}
}

// Message Types
type MessageType int

const (
	Unknown MessageType = iota
	DeviceInfo
	NetworkInfo
	InstantaneousDemand
	PriceCluster
	MessageCluster
	CurrentSummationDelivered
	TimeCluster
)

type Message struct {
	MacId     string `xml:"macId,attr"`
	Timestamp string `xml:"timestamp,attr"`

	// All the inner HTML, it could be anything
	Inner []byte `xml:",innerxml"`
}

func (m *Message) Type() MessageType {
	switch {
	case bytes.Contains(m.Inner, []byte("<InstantaneousDemand>")):
		return InstantaneousDemand
	case bytes.Contains(m.Inner, []byte("<NetworkInfo>")):
		return NetworkInfo
	case bytes.Contains(m.Inner, []byte("<DeviceInfo>")):
		return DeviceInfo
	case bytes.Contains(m.Inner, []byte("<MessageCluster>")):
		return MessageCluster
	case bytes.Contains(m.Inner, []byte("<PriceCluster>")):
		return PriceCluster
	case bytes.Contains(m.Inner, []byte("<TimeCluster>")):
		return TimeCluster
	case bytes.Contains(m.Inner, []byte("<CurrentSummationDelivered>")):
		return CurrentSummationDelivered
	default:
		return Unknown
	}
}
