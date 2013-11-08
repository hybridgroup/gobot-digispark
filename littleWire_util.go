package gobot-digispark

/*
#include <littleWire_util.h>
*/
import "C"

//void delay(unsigned int duration);
func delay(duration uint) {
  C.delay(duration)
}

