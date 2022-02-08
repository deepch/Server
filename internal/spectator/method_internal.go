/*
	Autor Andrey Semochkin
*/

package spectator

import (
	"time"

	"github.com/deepch/Server/internal/constant"
)

//Make - make new internal
func (spectator *Spectator) make(clientAddress string, clientProto int) *Spectator {

	spectator.session = session{
		remoteAddress: clientAddress,
		protocol:      clientProto,
		startTime:     time.Now().Unix(),
		status:        constant.CONNECT,
	}

	return spectator

}

//setStatus - set session status
func (spectator *Spectator) setStatus(val int) {

	spectator.session.status = val

}
