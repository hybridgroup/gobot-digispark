package gobot-digispark

/*
#include <littleWire_servo.h>
*/
import "C"

//void servo_init(littleWire* lwHandle);
func servoInit(lwHandle littleWire* ) {
  C.servo_init(lwHandle)
}

//void servo_updateLocation(littleWire* lwHandle,unsigned char locationChannelA,unsigned char locationChannelB);
func servoUpdateLocation(lwHandle *littleWire, locationChannelA uint8, locationChannelB uint8) {
  C.servo_updateLocation(lwHandle, locationChannelA, locationChannelB)
}
