/*
	Autor Andrey Semochkin
*/

package device

import (
	"github.com/deepch/Server/internal/channel"
	"github.com/sirupsen/logrus"
)

//Make - make new channel
func (device *Device) Make(deviceUniqID string, Logger *logrus.Logger) {

	//Recording session information
	session := device.session.new()
	defer session.close()

	//create channels
	device.make(deviceUniqID, Logger)

}

//Start - start a channel
func (device *Device) Start() {

	//Recording session information
	session := device.session.new()
	defer session.close()

	device.start()

}

//Stop - stop a channel
func (device *Device) Stop() {

	//Recording session information
	session := device.session.new()
	defer session.close()

	device.stop()

}

//SelectChannels - returns all channels of the device
func (device *Device) SelectChannels() map[string]*channel.Channel {

	//Recording session information
	session := device.session.new()
	defer session.close()

	return device.selectChannels()

}

//SelectChannel - returns all channel of the device
func (device *Device) SelectChannel(channelUniqID string) *channel.Channel {

	//If the device was removed
	if device == nil {

		return nil

	}
	//Recording session information
	session := device.session.new()
	defer session.close()

	return device.selectChannel(channelUniqID)

}

//SelectSummary - flow information
func (device *Device) SelectSummary() (*Device, error) {

	//Recording session information
	session := device.session.new()
	defer session.close()

	return device.selectSummary()

}

//SelectStatus - returns the status of all threads
func (device *Device) SelectStatus() (map[string]uint32, error) {

	//Recording session information
	session := device.session.new()
	defer session.close()

	return device.selectStatus()

}
