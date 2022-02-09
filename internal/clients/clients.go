/*
	Autor Andrey Semochkin
*/

package clients

import (
	"github.com/sirupsen/logrus"
)

//ServiceClients - basic service structure
type ServiceClients struct {
	ServiceName     string
	Logger          *logrus.Logger
	CheckErrorCount int
}

//Writer - for clients
var Writer *ServiceClients

//New - the function is executed when the service is initialized
func (service *ServiceClients) New() error {

	Writer = service

	return nil

}

//Name - returns the service name
func (service *ServiceClients) Name() string {

	return service.ServiceName

}

//Close - executed at the end of the service
func (service *ServiceClients) Close() error {

	/*
		Tasks that can be performed when the server is terminated
	*/

	return nil

}
