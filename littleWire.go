package gobot-digispark

/*
#include <littleWire.h>
*/
import "C"

//typedef usb_dev_handle littleWire;

//int littlewire_search();
func littleWireSearch() int {
  return int(C.littlewire_search())
}

//littleWire* littlewire_connect_byID(int desiredID);
func littleWireConnectByID(desiredID int) *littleWire {
  return C.littlewire_connect_byID(desiredID).(littleWire)
}

//littleWire* littlewire_connect_bySerialNum(int mySerial);
func littleWireConnectBySerialNum(mySerial int) *littleWire {
  return C.littlewire_connect_bySerialNum(mySerial).(littleWire)
}

//littleWire* littleWire_connect();
func littleWireConnect() *littleWire {
  return C.littleWire_connect().(littleWire)
}

//unsigned char readFirmwareVersion(littleWire* lwHandle);
func readFirmwareVersion(lwHandle littleWire*) uint8 {
  return uint8(C.readFirmwareVersion(lwHandle))
}

//void changeSerialNumber(littleWire* lwHandle,int serialNumber);
func changeSerialNumber(lwHandle *littleWire, serialNumber int) {
  C.changeSerialNumber(lwHandle, serialNumber)
}

//int customMessage(littleWire* lwHandle,unsigned char* receiveBuffer,unsigned char command,unsigned char d1,unsigned char d2, unsigned char d3, unsigned char d4);
func customMessage(lwHandle *littleWire, receiveBuffer *uint8, command uint8, d1 uint8, d2 uint8, d3 uint8, d4 uint8) int {
  return int(C.customMessage(lwHandle, receiveBuffer, command, d1, d2, d3, d4))
}

//int littleWire_error ();
func littleWireError() int {
  return int(C.littleWire_error())
}

//char *littleWire_errorName ();
//func *littleWire_errorName() {
//  
//}

//void digitalWrite(littleWire* lwHandle, unsigned char pin, unsigned char state);
func digitalWrite(lwHandle *littleWire, pin uint8, state uint8) {
  C.digitalWrite(lwHandle, pin, state)
}

//void pinMode(littleWire* lwHandle, unsigned char pin, unsigned char mode);
func pinMode(lwHandle *littleWire, pin uint8, mode uint8) {
  C.pinMode(lwHandle, pin, mode)
}

//unsigned char digitalRead(littleWire* lwHandle, unsigned char pin);
func digitalRead(lwHandle *littleWire, pin uint8) uint8 {
  return uint8(C.digitalRead(lwHandle, pin))
}

//void internalPullup(littleWire* lwHandle, unsigned char pin, unsigned char state);
func internalPullup(lwHandle *littleWire , pin uint8, state uint8) {
  C.internalPullup(lwHandle, pin, state)
}

//void analog_init(littleWire* lwHandle, unsigned char voltageRef);
func analogInit(lwHandle *littleWire, voltageRef uint8) {
  C.analog_init(lwHandle, voltageRef)
}

//unsigned int analogRead(littleWire* lwHandle, unsigned char channel);
func analogRead(lwHandle *littleWire, channel uint8) uint {
  return uint(analogRead(lwHandle, channel))
}

//void pwm_init(littleWire* lwHandle);
func pwmInit(lwHandle *littleWire) {
  C.pwm_init(lwHandle)
}

//void pwm_stop(littleWire* lwHandle);
func pwmStop(lwHandle *littleWire) {
  C.pwm_stop(lwHandle)
}

//void pwm_updateCompare(littleWire* lwHandle, unsigned char channelA, unsigned char channelB);
func pwmUpdateCompare(lwHandle *littleWire, channelA uint8, channelB uint8) {
  C.pwm_updateCompare(lwHandle, channelA, channelB)
}

//void pwm_updatePrescaler(littleWire* lwHandle, unsigned int value);
func pwmUpdatePrescaler(lwHandle *littleWire, value uint) {
  C.pwm_updatePrescaler(lwHandle, value)
}

//void spi_init(littleWire* lwHandle);
func spiInit(lwHandle *littleWire) {
  C.spi_init(lwHandle)
}

//void spi_sendMessage(littleWire* lwHandle, unsigned char * sendBuffer, unsigned char * inputBuffer, unsigned char length ,unsigned char mode);
func spiSendMessage(lwHandle *littleWire, sendBuffer *uint8, inputBuffer *uint8, length uint8, mode uint8) {
  C.spi_sendMessage(lwHandle, sendBuffer, inputBuffer, length, mode)
}

//unsigned char debugSpi(littleWire* lwHandle, unsigned char message);
func debugSpi(lwHandle *littleWire, message uint8) uint8 {
  return uint8(C.debugSpi(lwHandle, message))
}

//void spi_updateDelay(littleWire* lwHandle, unsigned int duration);
func spiUpdateDelay(lwHandle *littleWire, duration uint) {
  C.spi_updateDelay(lwHandle, duration)
}

//void i2c_init(littleWire* lwHandle);
func i2cInit(lwHandle *littleWire) {
  C.i2c_init(lwHandle)
}

//unsigned char i2c_start(littleWire* lwHandle, unsigned char address7bit, unsigned char direction);
func i2cStart(lwHandle *littleWire, address7bit uint8, direction uint8) uint8 {
  return uint8(i2c_start(lwHandle, address7bit, direction))
}

//void i2c_write(littleWire* lwHandle, unsigned char* sendBuffer, unsigned char length, unsigned char endWithStop);
func i2cWrite(lwHandle *littleWire, sendBuffer *uint8, length uint8, endWithStop uint8) {
  C.i2c_write(lwHandle, sendBuffer, length, endWithStop)
}

//void i2c_read(littleWire* lwHandle, unsigned char* readBuffer, unsigned char length, unsigned char endWithStop);
func i2cRead(lwHandle *littleWire, readBuffer *uint8, length uint8, endWithStop uint8) {
  C.i2c_read(lwHandle, readBuffer, length, endWithStop)
}

//void i2c_updateDelay(littleWire* lwHandle, unsigned int duration);
func i2cUpdateDelay(lwHandle *littleWire, duration uint) {
  C.i2c_updateDelay(lwHandle, duration)
}

//void onewire_sendBit(littleWire* lwHandle, unsigned char bitValue);
func oneWireSendBit(lwHandle *littleWire, bitValue uint8) {
  C.onewire_sendBit(lwHandle, bitValue)
}

//void onewire_writeByte(littleWire* lwHandle, unsigned char messageToSend);
func oneWireWriteByte(lwHandle *littleWire, messageToSend uint8) {
  C.onewire_writeByte(lwHandle, messageToSend)
}

//unsigned char onewire_readByte(littleWire* lwHandle);
func oneWireReadByte(lwHandle *littleWire) uint8 {
  return uint8(C.onewire_readByte(lwHandle))
}

//unsigned char onewire_readBit(littleWire* lwHandle);
func oneWireReadBit(lwHandle *littleWire) uint8 {
  return uint8(C.onewire_readBit(lwHandle))
}

//unsigned char onewire_readBit(littleWire* lwHandle);
func oneWireReadBit(lwHandle *littleWire) uint8 {
  return uint8(C.onewire_readBit(lwHandle))
}

//unsigned char onewire_resetPulse(littleWire* lwHandle);
func oneWireResetPulse(lwHandle *littleWire) uint8 {
  return uint8(C.onewire_resetPulse(lwHandle))
}

//int onewire_firstAddress(littleWire* lwHandle);
func oneWireFirstAddress(lwHandle *littleWire) int {
  return int(C.onewire_firstAddress(lwHandle))
}

//int onewire_nextAddress(littleWire* lwHandle);
func oneWireNextAddress(lwHandle *littleWire) int {
  return int(C.onewire_nextAddress(lwHandle))
}

//void softPWM_state(littleWire* lwHandle,unsigned char state);
func softPWMState(lwHandle *littleWire, state uint8) {
  C.softPWM_state(lwHandle, state)
}

//void softPWM_write(littleWire* lwHandle,unsigned char ch1,unsigned char ch2,unsigned char ch3);
func softPWM_write(lwHandle *littleWire, ch1 uint8, ch2 uint8, ch3 uint8) {
  C.softPWM_write(lwHandle, ch1, ch2, ch3)
}

//void ws2812_write(littleWire* lwHandle, unsigned char pin,unsigned char r,unsigned char g,unsigned char b);
func ws2812Write(lwHandle *littleWire, pin uint8, r uint8, g uint8, b uint8) {
  C.ws2812_write(lwHandle, pin, r, g, b)
}

//void ws2812_flush(littleWire* lwHandle, unsigned char pin);
func ws2812Flush(lwHandle *littleWire, pin uint8) {
  C.ws2812_flush(lwHandle, pin)
}

//void ws2812_preload(littleWire* lwHandle, unsigned char r,unsigned char g,unsigned char b);
func ws2812Preload(lwHandle *littleWire, r uint8, g uint8 , b uint8) {
  C.ws2812_preload(lwHandle, r, g, b)
}
