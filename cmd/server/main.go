/*
	Autor Andrey Semochkin
*/

package main

import (
	"github.com/deepch/Server/internal/clients"
	"github.com/deepch/Server/internal/constant"
	"github.com/deepch/Server/internal/device"
	"github.com/deepch/Server/internal/flag"
	"github.com/deepch/Server/internal/logging"
	"github.com/deepch/Server/internal/memory"
	"github.com/deepch/Server/internal/pkg"
	"github.com/deepch/Server/internal/servers/serverGRPC"
	"github.com/deepch/Server/internal/servers/serverHTTP"
	"github.com/deepch/Server/internal/servers/serverRTMP"
	"github.com/deepch/Server/internal/servers/serverRTP"
	"github.com/deepch/Server/internal/signal"
	"github.com/deepch/Server/internal/task"
)

func main() {

	//init service's
	pkgList := []memory.Service{
		&flag.ServiceFlag{ServiceName: constant.ServiceFlag, Logger: logging.Log},
		&memory.ServiceConfigurations{ServiceName: constant.ServiceConfigurations, Logger: logging.Log},
		&logging.ServiceLogging{ServiceName: constant.ServiceLogging},
		&device.ServiceDevice{ServiceName: constant.ServiceDevice, Logger: logging.Log},
		&clients.ServiceClients{ServiceName: constant.ServiceClients, Logger: logging.Log},
		&serverGRPC.ServiceGRPC{ServiceName: constant.ServiceGRPC, Logger: logging.Log, WaitTaskFromSystem: signal.WaitTaskFromSystem},
		&serverRTMP.ServiceRTMP{ServiceName: constant.ServiceRTMP, Logger: logging.Log, WaitTaskFromSystem: signal.WaitTaskFromSystem},
		&task.ServiceTask{ServiceName: constant.ServiceTask, Logger: logging.Log},
		&serverHTTP.ServiceHTTP{ServiceName: constant.ServiceHTTP, Logger: logging.Log, WaitTaskFromSystem: signal.WaitTaskFromSystem},
		&serverRTP.ServiceRTPProxy{ServiceName: constant.ServiceRTPProxy, Logger: logging.Log, WaitTaskFromSystem: signal.WaitTaskFromSystem},
	}

	//logging server start
	logging.Log.Debug(constant.ConstLoadingSystemService)

	//init pkg system
	err := pkg.New(pkgList)

	if err != nil {
		//exit if packet init error
		logging.Log.Fatalln(constant.ConstInitializationPkgServiceFail, err)
	}

	logging.Log.Debug(constant.ConstInitializationAllPkgServiceSuccess)

	//wait service exit signals
	signal.SystemMainControlWait()

	//correct close all pkg
	err = pkg.Close(pkgList)

	if err != nil {
		//log  packet exit error
		logging.Log.Error(constant.ConstFinalizeAllPkgServiceFail, err)
	}

	//logging server finish
	logging.Log.Debug(constant.ConstSystemCorrectOffline)

}
