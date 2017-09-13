package eagle

import (
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestType(t *testing.T) {
	assert := assert.New(t)
	tests := map[MessageType]string{
		InstantaneousDemand:       instantDemand,
		NetworkInfo:               networkInfo,
		DeviceInfo:                deviceInfo,
		MessageCluster:            messageCluster,
		CurrentSummationDelivered: currentSummation,
		PriceCluster:              priceCluster,
	}

	for expected, test := range tests {
		m := Message{}
		if !assert.NoError(xml.Unmarshal([]byte(test), &m)) {
			return
		}

		assert.Equal(expected, m.Type())
	}
}

func TestCurrentSummation(t *testing.T) {
	assert := assert.New(t)

	message := &Message{}
	if !assert.NoError(xml.Unmarshal([]byte(currentSummation), &message)) {
		return
	}

	payload, err := message.CurrentSummation()
	if !assert.NoError(err) {
		return
	}

	fmt.Println(payload)

}

// some testing data ...
var instantDemand = `<?xml version="1.0"?><rainforest macId="0xabcdef123456" version="undefined" timestamp="1504803337s">
<InstantaneousDemand>
  <DeviceMacId>0xabcdef1234567901</DeviceMacId>
  <MeterMacId>0xabcdef1234567901</MeterMacId>
  <TimeStamp>0x21443680</TimeStamp>
  <Demand>0x000cef</Demand>
  <Multiplier>0x00000001</Multiplier>
  <Divisor>0x000003e8</Divisor>
  <DigitsRight>0x03</DigitsRight>
  <DigitsLeft>0x06</DigitsLeft>
  <SuppressLeadingZero>Y</SuppressLeadingZero>
  <Port>/dev/ttySP0</Port>
</InstantaneousDemand>

</rainforest>`

var priceCluster = `<?xml version="1.0"?><rainforest macId="0xabcdef123456" version="undefined" timestamp="1504805582s">
<PriceCluster>
  <DeviceMacId>0xabcdef1234567901</DeviceMacId>
  <MeterMacId>0xabcdef1234567901</MeterMacId>
  <TimeStamp>0x21443f43</TimeStamp>
  <Price>0x00000507</Price>
  <Currency>0x007c</Currency>
  <TrailingDigits>0x04</TrailingDigits>
  <Tier>0x01</Tier>
  <StartTime>0x21443f43</StartTime>
  <Duration>0xffff</Duration>
  <RateLabel>Block 2</RateLabel>
  <Port>/dev/ttySP0</Port>
</PriceCluster>

</rainforest>`

var messageCluster = `<?xml version="1.0"?><rainforest macId="0xabcdef123456" version="undefined" timestamp="1504805491s">
<MessageCluster>
  <DeviceMacId>0xabcdef1234567901</DeviceMacId>
  <MeterMacId>0xabcdef1234567901</MeterMacId>
  <TimeStamp></TimeStamp>
  <Id></Id>
  <Text></Text>
  <Priority></Priority>
  <StartTime></StartTime>
  <Duration></Duration>
  <ConfirmationRequired>N</ConfirmationRequired>
  <Confirmed>N</Confirmed>
  <Queue>Active</Queue>
  <Port>/dev/ttySP0</Port>
</MessageCluster>

</rainforest>`

var networkInfo = `<?xml version="1.0"?><rainforest macId="0xabcdef123456" version="undefined" timestamp="1504805859s">
<NetworkInfo>
  <DeviceMacId>0xabcdef1234567901</DeviceMacId>
  <CoordMacId>0xabcdef1234567901</CoordMacId>
  <Status>Connected</Status>
  <Description>Successfully Joined</Description>
  <ExtPanId>0xabcdef1234567901</ExtPanId>
  <Channel>11</Channel>
  <ShortAddr>0x1751</ShortAddr>
  <LinkStrength>0x64</LinkStrength>
  <Port>/dev/ttySP0</Port>
</NetworkInfo>

</rainforest>`

var deviceInfo = `<?xml version="1.0"?><rainforest macId="0xabcdef123456" version="undefined" timestamp="1504805914s">
<DeviceInfo>
  <DeviceMacId>0xabcdef1234567901</DeviceMacId>
  <InstallCode>0x1234567890abcdef</InstallCode>
  <LinkKey>0x1234567890abcdef1234567890abcdef</LinkKey>
  <FWVersion>1.4.48 (6952)</FWVersion>
  <HWVersion>1.2.5</HWVersion>
  <ImageType>0x1301</ImageType>
  <Manufacturer>Rainforest Automation, Inc.</Manufacturer>
  <ModelId>Z109-EAGLE</ModelId>
  <DateCode>2014081323520712</DateCode>
  <Port>/dev/ttySP0</Port>
</DeviceInfo>

</rainforest>`

var currentSummation = `<?xml version="1.0"?><rainforest macId="0xabcdef123456" version="undefined" timestamp="1504805742s">
<CurrentSummationDelivered>
  <DeviceMacId>0xabcdef1234567901</DeviceMacId>
  <MeterMacId>0xabcdef1234567901</MeterMacId>
  <TimeStamp>0x21443fe2</TimeStamp>
  <SummationDelivered>0x00000000061cbd97</SummationDelivered>
  <SummationReceived>0x0000000000000054</SummationReceived>
  <Multiplier>0x00000001</Multiplier>
  <Divisor>0x000003e8</Divisor>
  <DigitsRight>0x01</DigitsRight>
  <DigitsLeft>0x06</DigitsLeft>
  <SuppressLeadingZero>Y</SuppressLeadingZero>
  <Port>/dev/ttySP0</Port>
</CurrentSummationDelivered>

</rainforest>`

var timeCluster = `<?xml version="1.0"?><rainforest macId="0xabcdef123456" version="undefined" timestamp="1504807618s">
<TimeCluster>
  <DeviceMacId>0xabcdef1234567901</DeviceMacId>
  <MeterMacId>0xabcdef1234567901</MeterMacId>
  <UTCTime>0x21444738</UTCTime>
  <LocalTime>0x2143d6b8</LocalTime>
  <Port>/dev/ttySP0</Port>
</TimeCluster>`
