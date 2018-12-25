package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mostlygeek/iot-recorder/eagle"
)

type Message struct {
}

func main() {
	http.ListenAndServe(":8081", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		/*
			fmt.Println("---------------------------------------------------")
			fmt.Printf("%s %s %s\n", req.RemoteAddr, req.Method, req.RequestURI)
				fmt.Println("Headers:")
				for key, vals := range req.Header {
					fmt.Printf("  - %s: %s\n", key, vals[0])
				}
		*/
		fmt.Println(req.Method, req.URL.Path)
		if req.Body != nil {
			body, _ := ioutil.ReadAll(req.Body)

			v := eagle.Message{}
			err := xml.Unmarshal([]byte(body), &v)
			if err != nil {
				fmt.Println("Could not parse XML")
			} else {
				mType := v.Type()
				if mType == eagle.InstantaneousDemand {
					message, _ := v.InstantaneousDemand()
					fmt.Printf("Demand: %s, %d watts\n", message.Timestamp, message.Demand)
				} else if mType == eagle.NetworkInfo {
					message, _ := v.NetworkInfo()
					fmt.Printf("Zigbee Status: %s\n", message.Status)
				} else if mType == eagle.DeviceInfo {
					message, _ := v.DeviceInfo()
					fmt.Printf("Device Info, firmware: %s\n", message.FWVersion)
				} else if mType == eagle.CurrentSummationDelivered {
					message, _ := v.CurrentSummation()
					del, rec := message.KWh()
					fmt.Printf("Power used: %.3fKWh, generated: %.3fKWh\n", del, rec)
				} else if mType == eagle.MessageCluster {
					message, _ := v.MessageCluster()
					if message.Id != 0 { // is it a eagle bug?
						fmt.Printf("Message #%d: (%s) %s\n",
							message.Id, message.Priority, message.Text)
					}
				} else if mType == eagle.PriceCluster {
					m, _ := v.PriceCluster()
					fmt.Printf("Price: tier(%d), curr: %d - price: %0.3f\n",
						m.Tier, m.Currency, m.Price)

				} else {
					fmt.Printf("Unhandled type: %d\n", v.Type())
				}
			}
		} else {
			fmt.Println("Body: <empty>")
		}

		w.WriteHeader(http.StatusOK)
	}))
}
