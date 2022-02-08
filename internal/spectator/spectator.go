/*
	Autor Andrey Semochkin
*/

package spectator

/*

	spectator structs

*/

//Spectator - basic struct
type Spectator struct {
	session session //users session
}

//Session descriptions
type session struct {
	status        int    //current status
	remoteAddress string //remote client address
	protocol      int    //proto connection
	startTime     int64  //start session time
}
