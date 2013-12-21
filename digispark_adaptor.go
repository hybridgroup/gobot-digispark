package gobotDigispark

import (
	"github.com/hybridgroup/gobot"
	"strconv"
)

type DigisparkAdaptor struct {
	gobot.Adaptor
	LittleWire *LittleWire
}

func (da *DigisparkAdaptor) Connect() bool {
	da.LittleWire = LittleWireConnect()
	da.Connected = true
	return true
}

func (da *DigisparkAdaptor) Reconnect() bool {
	return da.Connect()
}

func (da *DigisparkAdaptor) Finalize() bool   { return false }
func (da *DigisparkAdaptor) Disconnect() bool { return false }

func (da *DigisparkAdaptor) DigitalWrite(pin string, level byte) {
	p, _ := strconv.Atoi(pin)

	da.LittleWire.PinMode(uint8(p), 0)
	da.LittleWire.DigitalWrite(uint8(p), l)
}

func (da *DigisparkAdaptor) InitServo() {
	da.LittleWire.ServoInit()
}

func (da *DigisparkAdaptor) ServoWrite(pin string, angle uint8) {
	da.LittleWire.ServoUpdateLocation(angle, angle)
}
