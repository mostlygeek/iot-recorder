package eagle

import (
	"bytes"

	"github.com/pkg/errors"
)

var (
	InvalidType        = errors.New("Invalid Message Type")
	UnsupportedMessage = errors.New("Unsupported Message Type")
)

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
