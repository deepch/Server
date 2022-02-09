/*
	Autor Andrey Semochkin
*/

package device

import (
	"time"

	"github.com/deepch/Server/internal/channel"
	"github.com/sasha-s/go-deadlock"
	"github.com/sirupsen/logrus"
)

//ServiceDevice - basic service structure
type ServiceDevice struct {
	ServiceName string
	Logger      *logrus.Logger
}

//Device - device structure
type Device struct {
	Name         string                      `xml:"Name,omitempty" yaml:"Name,omitempty" json:"Name,omitempty" groups:"api,save"`
	ChannelsMode string                      `xml:"ChannelsMode,omitempty" yaml:"ChannelsMode,omitempty" json:"ChannelsMode,omitempty" groups:"api,save"`
	OnvifURL     string                      `xml:"OnvifURL,omitempty" yaml:"OnvifURL,omitempty" json:"OnvifURL,omitempty" groups:"api,save"`
	Channels     map[string]*channel.Channel `xml:"Channels,omitempty" yaml:"Channels,omitempty" json:"Channels,omitempty" groups:"api,save"`
	session      sessionGuard
}

//sessionGuard global in-memory storage array
type sessionGuard struct {
	lockGuard deadlock.Mutex
	funcName  string
	startTime time.Time
}

//New - the function is executed when the service is initialized
func (service *ServiceDevice) New() error {

	return nil
}

//Name - returns the service name
func (service *ServiceDevice) Name() string {

	return service.ServiceName

}

//Close - executed at the end of the service
func (service *ServiceDevice) Close() error {

	return nil

}
