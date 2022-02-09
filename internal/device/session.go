/*
	Autor Andrey Semochkin
*/

package device

import (
	"log"
	"time"

	"github.com/deepch/Server/internal/flag"
	"github.com/deepch/Server/internal/function"
)

//New - new database query session
func (session *sessionGuard) new() *sessionGuard {

	session.lockGuard.Lock()
	session.startTime = time.Now()
	session.funcName = function.GetCurrentFuncName()

	return session

}

//Close - close session database query
func (session *sessionGuard) close() {

	if flag.Debug {

		log.Println(session.funcName, time.Since(session.startTime))

	}

	session.lockGuard.Unlock()

}
