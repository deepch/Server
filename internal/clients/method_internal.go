/*
	Autor Andrey Semochkin
*/

package clients

import (
	"time"

	"github.com/deepch/vdk/format/mp4f"

	"github.com/deepch/Server/internal/constant"
	"github.com/deepch/Server/internal/memory"
	"github.com/deepch/vdk/av/pubsub"
	"github.com/deepch/vdk/format/rtmp"
	"github.com/deepch/vdk/format/rtspv2"
	"golang.org/x/net/websocket"
)

//clientRTMPLoop - send from queue to RTSP clients
func (service *ServiceClients) clientRTMPLoop(queueCursor *pubsub.QueueCursor, conn *rtmp.Conn) error {

	//We send only packets from the key frame to the client
	var played bool

	//Read the packet queue and send to the client
	for {
		//Getting a packet from the queue
		pkt, err := queueCursor.ReadPacket()

		if err != nil {

			return err

		}

		//If a key frame arrives, we allow the transfer to the client
		if pkt.IsKeyFrame {

			played = true

		}

		//if the stream has not yet started, wait for a keyframe
		if !played {

			continue
		}

		//Sending a packet to the socket to the client
		err = conn.WritePacket(pkt)

		if err != nil {

			return err

		}

	}

}

//clientRTMPPrepare - send from queue to RTMP clients
func (service *ServiceClients) clientRTMPPrepare(queueCursor *pubsub.QueueCursor, conn *rtmp.Conn) error {

	//Waiting for stream codecs
	streams, err := queueCursor.Streams()

	if err != nil {

		return err

	}

	//Write headers to the client
	err = conn.WriteHeader(streams)

	if err != nil {

		return err

	}

	return nil

}

//clientRTMP - send from queue to RTMP clients
func (service *ServiceClients) clientRTMP(deviceUniqID, channelUniqID string, conn *rtmp.Conn) error {

	//Registering a viewer get a pointer to the broadcast queue
	queueCursor, clientUniqID, err := memory.TnxMemory.SelectDevice(deviceUniqID).SelectChannel(channelUniqID).Registered(conn.NetConn().RemoteAddr().String(), constant.RTMP)

	if err != nil {

		return err

	}

	//When the client left, he deleted his session
	defer func() {

		err := memory.TnxMemory.SelectDevice(deviceUniqID).SelectChannel(channelUniqID).UnRegistered(clientUniqID)

		if err != nil {

			service.Logger.Warn(err)

		}

	}()

	//write headlines
	err = service.clientRTMPPrepare(queueCursor, conn)

	if err != nil {

		return err

	}

	//Going into a cycle
	return service.clientRTMPLoop(queueCursor, conn)

}

//clientRTSPLoop - send from queue to RTSP clients
func (service *ServiceClients) clientRTSPLoop(queueCursor *pubsub.QueueCursor, conn *rtspv2.ProxyConn) error {

	//We send only packets from the key frame to the client
	var played bool

	//Read the packet queue and send to the client

	for {

		//Getting a packet from the queue
		pkt, err := queueCursor.ReadPacket()

		if err != nil {

			return err

		}

		//If a key frame arrives, we allow the transfer to the client
		if pkt.IsKeyFrame {

			played = true

		}

		//If the video is not ready
		if !played {

			continue

		}
		//Sending a packet to the socket to the client
		err = conn.WritePacket(&pkt.Data)

		if err != nil {

			return err

		}

	}

}

//clientRTSP - send from queue to RTSP clients
func (service *ServiceClients) clientRTSP(deviceUniqID, channelUniqID string, conn *rtspv2.ProxyConn) error {

	//Registering a viewer get a pointer to the broadcast queue
	queueCursor, clientUniqID, err := memory.TnxMemory.SelectDevice(deviceUniqID).SelectChannel(channelUniqID).Registered(conn.NetConn().RemoteAddr().String(), constant.RTSP)

	if err != nil {

		return err

	}

	//When the client left, he deleted his session
	defer func() {

		err := memory.TnxMemory.SelectDevice(deviceUniqID).SelectChannel(channelUniqID).UnRegistered(clientUniqID)

		if err != nil {

			service.Logger.Warn(err)

		}

	}()

	return service.clientRTSPLoop(queueCursor, conn)

}

//clientWebsocketPrepare - send from queue to RTMP clients
func (service *ServiceClients) clientWebsocketPrepare(queueCursor *pubsub.QueueCursor, conn *websocket.Conn) (*mp4f.Muxer, error) {

	//Waiting for stream codecs
	streams, err := queueCursor.Streams()

	if err != nil {

		return nil, err

	}

	//create a mixer instance
	muxer := mp4f.NewMuxer(nil)

	//write information about codecs
	err = muxer.WriteHeader(streams)

	if err != nil {

		return nil, err

	}

	//We get init
	meta, fist := muxer.GetInit(streams)

	//Sending the meta
	err = websocket.Message.Send(conn, meta)

	if err != nil {

		return nil, err

	}

	//Sending int
	err = websocket.Message.Send(conn, fist)

	if err != nil {

		return nil, err

	}

	return muxer, nil

}

//clientWebsocketLoop - send from queue to RTSP clients
func (service *ServiceClients) clientWebsocketLoop(muxer *mp4f.Muxer, queueCursor *pubsub.QueueCursor, conn *websocket.Conn) {

	//We send only packets from the key frame to the client
	var played bool

	//Read the packet queue and send to the client
	for {
		//Getting a packet from the queue
		frame, err := queueCursor.ReadPacket()

		if err != nil {

			return

		}

		//If a key frame arrives, we allow the transfer to the client
		if frame.IsKeyFrame {

			err = conn.SetWriteDeadline(time.Now().Add(10 * time.Second))

			if err != nil {

				return

			}

			played = true

		}

		//If the video is not ready
		if !played {

			continue

		}

		//Sending a packet to the socket to the client
		gotFrame, buffer, err := muxer.WritePacket(frame, false)

		if err != nil {

			return

		}

		if gotFrame {

			err = websocket.Message.Send(conn, buffer)

			if err != nil {

				return

			}

		}

	}

}

//clientWebSocket - sending from a queue to Websocket clients
func (service *ServiceClients) clientWebSocket(deviceUniqID, channelUniqID string, conn *websocket.Conn) {

	//Registering a viewer get a pointer to the broadcast queue
	queueCursor, clientUniqID, err := memory.TnxMemory.SelectDevice(deviceUniqID).SelectChannel(channelUniqID).Registered(conn.RemoteAddr().String(), constant.RTMP)

	if err != nil {

		return

	}

	//When the client left, he deleted his session
	defer func() {

		err := memory.TnxMemory.SelectDevice(deviceUniqID).SelectChannel(channelUniqID).UnRegistered(clientUniqID)

		if err != nil {

			service.Logger.Warn(err)

		}

	}()

	muxer, err := service.clientWebsocketPrepare(queueCursor, conn)

	if err != nil {

		return

	}

	service.clientWebsocketLoop(muxer, queueCursor, conn)

}
