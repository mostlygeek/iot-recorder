package main

import (
	"encoding/xml"
	"fmt"

	"github.com/mostlygeek/iot-recorder/eagle"
)

func main() {

	data := instantDemand
	fmt.Println(data)

	v := eagle.Message{}
	err := xml.Unmarshal([]byte(data), &v)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf(`"%s"`, string(v.Inner))
}

var instantDemand = `<?xml version="1.0"?><rainforest macId="0xd8d5b9001ff9" version="undefined" timestamp="1504803337s">
<InstantaneousDemand>
  <DeviceMacId>0xd8d5b900000038be</DeviceMacId>
  <MeterMacId>0x0007810000b236f0</MeterMacId>
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

var priceCluster = `<?xml version="1.0"?><rainforest macId="0xd8d5b9001ff9" version="undefined" timestamp="1504805582s">
<PriceCluster>
  <DeviceMacId>0xd8d5b900000038be</DeviceMacId>
  <MeterMacId>0x0007810000b236f0</MeterMacId>
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

var messageCluster = `<?xml version="1.0"?><rainforest macId="0xd8d5b9001ff9" version="undefined" timestamp="1504805491s">
<MessageCluster>
  <DeviceMacId>0xd8d5b900000038be</DeviceMacId>
  <MeterMacId>0x0007810000b236f0</MeterMacId>
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

var networkInfo = `<?xml version="1.0"?><rainforest macId="0xd8d5b9001ff9" version="undefined" timestamp="1504805859s">
<NetworkInfo>
  <DeviceMacId>0xd8d5b900000038be</DeviceMacId>
  <CoordMacId>0x0007810000b236f0</CoordMacId>
  <Status>Connected</Status>
  <Description>Successfully Joined</Description>
  <ExtPanId>0x0007810000b236f0</ExtPanId>
  <Channel>11</Channel>
  <ShortAddr>0x1751</ShortAddr>
  <LinkStrength>0x64</LinkStrength>
  <Port>/dev/ttySP0</Port>
</NetworkInfo>

</rainforest>`

var deviceInfo = `<?xml version="1.0"?><rainforest macId="0xd8d5b9001ff9" version="undefined" timestamp="1504805914s">
<DeviceInfo>
  <DeviceMacId>0xd8d5b900000038be</DeviceMacId>
  <InstallCode>0xe6de9971cc2c25b6</InstallCode>
  <LinkKey>0x0edf51b1b3750b7e8005f806c35cc438</LinkKey>
  <FWVersion>1.4.48 (6952)</FWVersion>
  <HWVersion>1.2.5</HWVersion>
  <ImageType>0x1301</ImageType>
  <Manufacturer>Rainforest Automation, Inc.</Manufacturer>
  <ModelId>Z109-EAGLE</ModelId>
  <DateCode>2014081323520712</DateCode>
  <Port>/dev/ttySP0</Port>
  <Port>/dev/ttySP0</Port>
</DeviceInfo>

</rainforest>`

var currentSummation = `<?xml version="1.0"?><rainforest macId="0xd8d5b9001ff9" version="undefined" timestamp="1504805742s">
<CurrentSummationDelivered>
  <DeviceMacId>0xd8d5b900000038be</DeviceMacId>
  <MeterMacId>0x0007810000b236f0</MeterMacId>
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
