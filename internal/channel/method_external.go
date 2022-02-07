/*
	Autor Andrey Semochkin
*/

package channel

import (
	"github.com/deepch/Server/internal/constant"

	"github.com/deepch/vdk/av/pubsub"
	"github.com/sirupsen/logrus"
)

//Make - channel
func (channel *Channel) Make(deviceUniqID string, channelUniqID string, logger *logrus.Logger) {

	channel.make(deviceUniqID, channelUniqID, logger)

}

//Start - channel start
func (channel *Channel) Start() {

	//write session
	session := channel.session.new()
	defer session.close()

	channel.start()

}

//Stop - channel
func (channel *Channel) Stop() {

	//Запись информации о сессии
	session := channel.session.new()
	defer session.close()

	channel.stop()

}

//SelectMeta - select meta data
func (channel *Channel) SelectMeta() ([]byte, error) {

	if channel == nil {

		return nil, constant.ErrChannelNotFound

	}

	//create new session
	session := channel.session.new()
	defer session.close()

	return channel.selectMeta()

}

//SelectRecordTime - return record time
func (channel *Channel) SelectRecordTime() uint32 {

	//new session
	session := channel.session.new()
	defer session.close()

	return channel.selectRecordTime()

}

//SelectSnapshotURL - return snapshot url
func (channel *Channel) SelectSnapshotURL() string {

	//new session
	session := channel.session.new()
	defer session.close()

	return channel.selectSnapshotURL()

}

//Registered - registerer viewer
func (channel *Channel) Registered(clientAddress string, clientProto int) (*pubsub.QueueCursor, string, error) {

	//check channel exist
	if channel == nil {

		return nil, "", constant.ErrChannelNotFound

	}

	//new session
	session := channel.session.new()
	defer session.close()

	return channel.registered(clientAddress, clientProto)

}

//UnRegistered - unregister viewer
func (channel *Channel) UnRegistered(clientUniqID string) error {

	//check channel exist
	if channel == nil {

		return constant.ErrChannelNotFound

	}

	//new session
	session := channel.session.new()
	defer session.close()

	err := channel.selectSpectatorIsSet(clientUniqID)

	if err != nil {

		return err

	}

	channel.unRegistered(clientUniqID)

	return nil

}
