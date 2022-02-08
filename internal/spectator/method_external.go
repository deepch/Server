/*
	Autor Andrey Semochkin
*/

package spectator

//Make - new spectator
func Make(clientAddress string, clientProto int) *Spectator {

	var tmp Spectator

	return tmp.make(clientAddress, clientProto)

}
