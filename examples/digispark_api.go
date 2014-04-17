package main

import (
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot-digispark"
	"github.com/hybridgroup/gobot-gpio"
)

func main() {
	master := gobot.GobotMaster()
	gobot.Api(master)

	digispark := new(gobotDigispark.DigisparkAdaptor)
	digispark.Name = "digispark"

	led := gobotGPIO.NewLed(digispark)
	led.Name = "led"
	led.Pin = "0"

	master.Robots = append(master.Robots, &gobot.Robot{
		Name:        "digispark",
		Connections: []gobot.Connection{digispark},
		Devices:     []gobot.Device{led},
	})

	master.Start()
}
