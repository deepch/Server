/*
	Autor Andrey Semochkin
*/

package channel

import (
	"time"

	"github.com/sirupsen/logrus"

	"github.com/deepch/vdk/codec/h264parser"

	"github.com/deepch/Server/internal/constant"
	"github.com/deepch/Server/internal/function"
	"github.com/deepch/Server/internal/record"
	"github.com/deepch/Server/internal/spectator"
	"github.com/deepch/vdk/av"
	"github.com/deepch/vdk/av/pubsub"
	"github.com/deepch/vdk/format/rtspv2"
)

//Make - Creates an empty channel
func (channel *Channel) make(deviceUniqID string, channelUniqID string, logger *logrus.Logger) {

	channel.deviceUniqID = deviceUniqID
	channel.channelUniqID = channelUniqID
	channel.cursorRTP = pubsub.NewQueue()
	channel.cursorFrame = pubsub.NewQueue()
	channel.spectators = make(map[string]*spectator.Spectator)
	channel.logger = logger
	channel.message = make(chan int, 100)
	channel.setStatus(constant.UNDEFINED)

}

//Start - launches a channel
func (channel *Channel) start() {

	go channel.loop()

}

//Stop - stops the channel
func (channel *Channel) stop() {

	channel.setNeedStop(true)
	channel.message <- constant.MessageStop

	err := channel.cursorRTP.Close()

	if err != nil {

		channel.logger.Warn(err)

	}
	err = channel.cursorFrame.Close()

	if err != nil {

		channel.logger.Warn(err)

	}

}

//receiver - source handler
func (channel *Channel) receiver() error {

	//We update the status that we started a connection attempt
	channel.setStatus(constant.CONNECT)

	//Let's prepare the source
	err := channel.prepareSource()

	if err != nil {

		return err

	}

	//When finished, close the source
	defer channel.closeSource()

	//Refreshing on connection held
	channel.setStatus(constant.CONNECTED)

	//Initializing muxers
	err = channel.initMuxers(channel.deviceUniqID, channel.channelUniqID)

	if err != nil {

		return err

	}

	//When closing, end the muxers
	defer channel.closeMuxers()

	//We translate the status into online
	channel.setStatus(constant.ONLINE)

	//Switching to streaming mode
	return channel.stream()

}

//selectMeta - returns the stream's meta data
func (channel *Channel) selectMeta() ([]byte, error) {

	return channel.meta, nil

}

//selectRecordTime - returns the recording time
func (channel *Channel) selectRecordTime() uint32 {

	return channel.RecordTime

}

//selectSnapshotURL - returns the path to the device image
func (channel *Channel) selectSnapshotURL() string {

	return channel.SnapshotURL

}

//Registered - registers a channel viewer
func (channel *Channel) registered(clientAddress string, clientProto int) (*pubsub.QueueCursor, string, error) {

	clientUniqID := function.CreateUniqID()

	channel.spectators[clientUniqID] = spectator.Make(clientAddress, clientProto)

	if clientProto == constant.RTSP {

		return channel.cursorRTP.Latest(), clientUniqID, nil

	}

	return channel.cursorFrame.Latest(), clientUniqID, nil

}

//UnRegistered - cancels client registration
func (channel *Channel) unRegistered(clientUniqID string) {

	delete(channel.spectators, clientUniqID)

}

//Stream - reading stream data
func (channel *Channel) stream() error {

	for {
		err := channel.streamSelect()

		if err != nil {

			return err

		}

	}

}

//frameChan - queue processing
func (channel *Channel) frameChan(frame *[]byte) error {

	return channel.cursorRTP.WritePacket(av.Packet{Data: *frame, IsKeyFrame: true})

}

//setStatus - sets the status of a stream
func (channel *Channel) setStatus(val uint32) {

	channel.Status = val

}

//muxerSwitch - muxers
func (channel *Channel) muxerSwitch(pktAV *av.Packet) error {

	//Mixing for archive
	for _, muxer := range channel.muxers {

		//Give the frame to all available mixers
		err := muxer.WriteFrame(pktAV)

		if err != nil {

			return err

		}

	}

	return nil

}

//openSource - opening video source
func (channel *Channel) openSource() error {

	//Initializing the Client
	source, err := rtspv2.Dial(
		rtspv2.RTSPClientOptions{
			URL:              function.URLDecode(channel.DataURL),
			DisableAudio:     !channel.Audio,
			DialTimeout:      constant.ConstDeviceChannelDialTimeout,
			ReadWriteTimeout: constant.ConstDeviceChannelReadWriteTimeout,
			Debug:            channel.Debug,
			OutgoingProxy:    channel.RTPProxy,
		})

	if err != nil {
		return err
	}

	//Putting source in storage
	channel.source = source

	return nil

}

//closeSource - closing video source
func (channel *Channel) closeSource() {

	channel.source.Close()
	channel.stage = 0

}

//prepareSource - preparing to connect the source
func (channel *Channel) prepareSource() error {

	err := channel.openSource()

	if err != nil {

		return err

	}

	//Update meta data
	err = channel.metaUpdate()

	if err != nil {

		channel.logger.Warn(err)

	}

	return nil
}

//metaUpdate - checks if the stream has a viewer
func (channel *Channel) metaUpdate() error {

	//Save stream meta data
	channel.meta = channel.source.SDPRaw

	if channel.source.CodecData != nil && len(channel.source.CodecData) > 0 && channel.source.CodecData[0].Type() == av.H264 {

		channel.decoderConfRecordBytes = channel.source.CodecData[0].(h264parser.CodecData).AVCDecoderConfRecordBytes()
		return channel.cursorFrame.WriteHeader(channel.source.CodecData)

	}

	return nil

}

//selectSpectatorIsSet - checks if the stream has a viewer
func (channel *Channel) selectSpectatorIsSet(clientUniqID string) error {

	if _, ok := channel.spectators[clientUniqID]; !ok {

		return constant.ErrSpectatorNotFound

	}

	return nil

}

//setNeedStop - set Stop
func (channel *Channel) setNeedStop(val bool) {

	channel.needStop = val

}

//selectNeedStop - set Need Stop?
func (channel *Channel) selectNeedStop() bool {

	return channel.needStop

}

//loop - loop work
func (channel *Channel) loop() {

	for {

		//If you want to stop the thread
		if channel.selectNeedStop() {

			//Exit
			return
		}

		//Trying to work with the flow
		err := channel.receiver()

		//Handling the error
		switch err {

		//If you received a forced closure
		case constant.ErrorDeviceChannelReceiveCloseSignal:

			//Change status to error
			channel.setStatus(constant.ERROR)

			//We leave
			return

		default:

			//Change status to offline
			channel.setStatus(constant.DISCONNECTED)

		}

		//Change status to pending
		channel.setStatus(constant.WAIT)

		//Expect before according to constants
		time.Sleep(constant.ConstDeviceChannelReconnectionTimer)

	}

}

//framesChan - queue processing
func (channel *Channel) framesChan(frame *av.Packet) error {

	if channel.stage == 0 && frame.IsKeyFrame {

		channel.stage = 1

	}

	if channel.stage == 0 {

		return nil

	}

	err := channel.cursorFrame.WritePacket(*frame)
	if err != nil {

		return err

	}

	err = channel.muxerSwitch(frame)

	if err != nil {

		return err

	}

	return nil

}

//initMuxers - the function is executed when the mixers are initialized
func (channel *Channel) initMuxers(deviceUniqID, channelUniqID string) error {

	//Recording mixer if enabled globally and locally on a channel
	if channel.Record {

		//Create a Recording Muxer
		recordMuxer, err := record.Muxer(function.CreateServicePatch(deviceUniqID, channelUniqID, constant.ConstRecord), channel.logger)

		if err != nil {

			return err

		}

		//Writing Headlines
		err = recordMuxer.WriteHeader(channel.decoderConfRecordBytes)

		if err != nil {

			return err

		}

		//Putting the muxer in the archive
		channel.muxers = append(channel.muxers, recordMuxer)

	}

	return nil

}

//streamSelect - selecting data from channels
func (channel *Channel) streamSelect() error {
	select {

	case sig := <-channel.message:

		return function.SignalToError(sig)

	case f := <-channel.source.OutgoingProxyQueue:

		err := channel.frameChan(f)

		if err != nil {

			return err

		}

	case f := <-channel.source.OutgoingPacketQueue:

		err := channel.framesChan(f)

		if err != nil {

			return err

		}

	}

	return nil

}

//closeMuxers - Closing the channel mixers
func (channel *Channel) closeMuxers() {

	//Close all mixers
	for _, muxer := range channel.muxers {

		err := muxer.Close()

		if err != nil {

			channel.logger.Warning(constant.ErrorDeviceChannelCloseMuxer, err)

		}

	}

	channel.muxers = nil

}
