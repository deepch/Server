/*
	Autor Andrey Semochkin
*/

package flag

import (
	"flag"

	"github.com/deepch/Server/internal/constant"
	"github.com/deepch/Server/internal/function"
	"github.com/sirupsen/logrus"
)

/*

	Service - flags (arguments that are specified when starting the program)

*/

var (
	//Debug - enable disable debug mode
	Debug = false //Is debug mode enabled
	//debugLevel - debug level see documentation
	debugLevel = constant.ConstDebugLevelError
	//ConfigurationPatch - path to configuration file
	ConfigurationPatch = ""
)

/*

	Service for getting command line arguments

*/

//ServiceFlag basic service structure
type ServiceFlag struct {
	ServiceName string
	Logger      *logrus.Logger
}

//New - the function is executed when the service is initialized.
func (service *ServiceFlag) New() error {

	//Debug mode on/off
	flagDebug := flag.Bool(constant.ConstNameDebug, constant.ConstNameDebugDefault, constant.ConstNameDebugDescriptions)

	//Logging level to use in logrus
	flagLoglevel := flag.String(constant.ConstNameDebugLevel, constant.ConstNameDebugLevelDefault, constant.ConstNameDebugLevelDescriptions)

	//Path to the configuration file
	flagConfigurationPatch := flag.String(constant.ConstNameConfigurationPatch, constant.ConstNameConfigurationPatchDefault, constant.ConstNameConfigurationPatchDescriptions)

	//Reading the flags
	flag.Parse()

	Debug = *flagDebug
	debugLevel = *flagLoglevel
	ConfigurationPatch = *flagConfigurationPatch

	//We log the received values in the debug log
	service.Logger.Debug(constant.ConstFlagNameDebug, function.BoolConvertToStringSwitch(Debug), constant.ConstFlagNameDebugLevel, debugLevel, constant.ConstFlagConfigurationPatch, ConfigurationPatch)

	return nil
}

//Name - returns the service name
func (service *ServiceFlag) Name() string {
	return service.ServiceName
}

//Close - executed at the end of the service
func (service *ServiceFlag) Close() error {
	return nil
}
