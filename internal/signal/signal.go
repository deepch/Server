/*
	Autor Andrey Semochkin
*/

package signal

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/deepch/Server/internal/constant"
	"github.com/deepch/Server/internal/logging"
)

/*
	The termination system call trapping functionality requires the use of an error condition to trap.
*/

//WaitTaskFromSystem  - Global variable waiting for a signal
var WaitTaskFromSystem = make(chan bool, 1)

//SystemMainControlWait - the function catches the process termination command ctr+c or pid kill etc
func SystemMainControlWait() {

	waitFromInSystemSignalChannel := make(chan os.Signal, 1)
	waitTaskOutSignalChannel := make(chan bool, 1)

	signal.Notify(waitFromInSystemSignalChannel,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGABRT,
		syscall.SIGKILL,
		syscall.SIGQUIT)

	//Let's start an additional thread to wait and process the signal
	go systemMainControlWaitReader(waitFromInSystemSignalChannel, waitTaskOutSignalChannel)

	select {

	//External signal from the system
	case <-waitTaskOutSignalChannel:

	//Internal signal
	case <-WaitTaskFromSystem:

		logging.Log.Info(constant.ConstInternalSignalAccepted)

	}

}

//systemMainControlWaitReader - read logging send if necessary add a handler.
func systemMainControlWaitReader(waitFromInSystemSignalChannel chan os.Signal, waitTaskOutSignalChannel chan bool) {

	//Blocks the stream, read the channel and wait for a signal from the system
	receivedSystemSignal := <-waitFromInSystemSignalChannel

	logging.Log.Info(constant.ConstSystemSignalCopyToInternalSignal, receivedSystemSignal)

	//We forward the message to the adjacent channel that is waiting in select
	waitTaskOutSignalChannel <- true

}
