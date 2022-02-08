/*
	Autor Andrey Semochkin
*/

package channel

import (
	"log"
	"time"

	"github.com/deepch/Server/internal/flag"
	"github.com/deepch/Server/internal/function"
)

//New - session storage
func (session *sessionGuard) new() *sessionGuard {

	session.lockGuard.Lock()
	session.startTime = time.Now()
	session.funcName = function.GetCurrentFuncName()

	return session

}

//Close - session storage
func (session *sessionGuard) close() {

	//logger
	if flag.Debug {

		log.Println(session.funcName, time.Since(session.startTime))

	}

	session.lockGuard.Unlock()

}
