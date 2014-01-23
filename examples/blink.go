package main

import (
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot-digispark"
	"github.com/hybridgroup/gobot-gpio"
)

func main() {

	digispark := new(gobotDigispark.DigisparkAdaptor)
	digispark.Name = "Digispark"

	led := gobotGPIO.NewLed(digispark)
	led.Name = "led"
	led.Pin = "0"

	work := func() {
		gobot.Every("0.5s", func() {
			led.Toggle()
		})
	}

	robot := gobot.Robot{
		Connections: []gobot.Connection{digispark},
		Devices:     []gobot.Device{led},
		Work:        work,
	}

	robot.Start()
}
