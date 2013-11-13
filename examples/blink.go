package main

import (
  "github.com/hybridgroup/gobot"
  "github.com/hybridgroup/gobot-digispark"
  "github.com/hybridgroup/gobot-gpio"
  "fmt"
)

func main() {

  digispark := new(gobotDigispark.DigisparkAdaptor)
  digispark.Name = "Digispark"

  led := gobotGPIO.NewLed(&digispark)
  led.Driver = gobot.Driver{
    Name: "led",
    Pin: "2",
  }

  connections := []interface{} {
    digispark,
  }
  devices := []interface{} {
    led,
  }

  work := func(){
    gobot.Every("0.5s", func(){ 
      led.Toggle() 
      if led.IsOn() {
        fmt.Println("On")
      } else {
        fmt.Println("Off")
      }
    })
  }
  
  robot := gobot.Robot{
      Connections: connections, 
      Devices: devices,
      Work: work,
  }

  robot.Start()
}
