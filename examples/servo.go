package main

import (
	"fmt"
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot-digispark"
	"github.com/hybridgroup/gobot-gpio"
)

func main() {

	digispark := new(gobotDigispark.DigisparkAdaptor)
	digispark.Name = "Digispark"

	servo := gobotGPIO.NewServo(digispark)
	servo.Name = "servo"
	servo.Pin = "0"

	work := func() {
		servo.InitServo()
		gobot.Every("1s", func() {
			i := int(gobot.Rand(180))
			fmt.Println("Turning", i)
			servo.Move(i)
		})
	}

	robot := gobot.Robot{
		Connections: []interface{}{digispark},
		Devices:     []interface{}{servo},
		Work:        work,
	}

	robot.Start()
}
