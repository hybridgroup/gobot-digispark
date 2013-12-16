package gobotDigispark

import (
	"github.com/hybridgroup/gobot"
	"strconv"
)

type DigisparkAdaptor struct {
	gobot.Adaptor
	LittleWire *LittleWire
}

func (da *DigisparkAdaptor) Connect() {
	da.LittleWire = LittleWireConnect()
}

func (da *DigisparkAdaptor) DigitalWrite(pin string, level string) {
	p, _ := strconv.Atoi(pin)
	l, _ := strconv.Atoi(level)
	da.LittleWire.PinMode(uint8(p), 0)
	da.LittleWire.DigitalWrite(uint8(p), uint8(l))
}
func (da *DigisparkAdaptor) InitServo() {
	da.LittleWire.ServoInit()
}

func (da *DigisparkAdaptor) ServoWrite(pin string, angle uint8) {
	da.LittleWire.ServoUpdateLocation(angle, angle)
}

func (da *DigisparkAdaptor) Disconnect() {
}
