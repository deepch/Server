/*
	Autor Andrey Semochkin
*/

package clients

import (
	"github.com/deepch/Server/internal/constant"
	"github.com/deepch/vdk/format/rtmp"
	"github.com/deepch/vdk/format/rtspv2"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/websocket"
)

//ClientRTMP - send from queue to RTMP clients
func (service *ServiceClients) ClientRTMP(deviceUniqID, channelUniqID string, conn *rtmp.Conn) error {

	return service.clientRTMP(deviceUniqID, channelUniqID, conn)

}

//ClientRTSP - send from queue to RTSP clients
func (service *ServiceClients) ClientRTSP(deviceUniqID, channelUniqID string, conn *rtspv2.ProxyConn) error {

	return service.clientRTSP(deviceUniqID, channelUniqID, conn)

}

//ClientWebSocket - dispatch from the queue to WebSocket clients
func (service *ServiceClients) ClientWebSocket(client *gin.Context) {

	websocket.Handler(func(ws *websocket.Conn) {

		service.clientWebSocket(client.Param(constant.ConstDeviceUniqID), client.Param(constant.ConstChannelUniqID), ws)

	}).ServeHTTP(client.Writer, client.Request)

}
