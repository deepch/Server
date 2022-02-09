/*
	Autor Andrey Semochkin
*/

package device

import (
	"github.com/deepch/Server/internal/channel"
	"github.com/deepch/Server/internal/constant"
	"github.com/sirupsen/logrus"
)

//make - selects a channel
func (device *Device) make(deviceUniqID string, Logger *logrus.Logger) {

	for channelUniqID, c := range device.channels() {

		c.Make(deviceUniqID, channelUniqID, Logger)

	}

}

//Channels - get a list of all channels
func (device *Device) channels() map[string]*channel.Channel {

	return device.Channels

}

//start - start a channel
func (device *Device) start() {

	for _, c := range device.channels() {

		c.Start()

	}

}

//stop - stop a channel
func (device *Device) stop() {

	for _, c := range device.channels() {

		c.Stop()

	}

}

//selectChannels - returns the channels element
func (device *Device) selectChannels() map[string]*channel.Channel {

	return device.Channels

}

//selectChannel - returns the channel element
func (device *Device) selectChannel(channelUniqID string) *channel.Channel {

	return device.Channels[channelUniqID]

}

//selectSummary - flow information
func (device *Device) selectSummary() (*Device, error) {

	if device == nil {

		return nil, constant.ErrorDeviceNotFound

	}

	return device, nil

}

//selectStatus - returns the status of all threads
func (device *Device) selectStatus() (map[string]uint32, error) {

	if device != nil {

		return nil, constant.ErrorDeviceNotFound

	}

	res := make(map[string]uint32)

	if device.Channels == nil {

		return nil, constant.ErrChannelNotFound

	}

	for i, c := range device.channels() {

		res[i] = c.Status

	}

	return res, nil

}
