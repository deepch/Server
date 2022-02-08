/*
	Autor Andrey Semochkin
*/

package util

import (
	"github.com/deepch/Server/internal/record"
	"github.com/deepch/vdk/av"
)

//CopyFrame - copy frame in mux
func CopyFrame(muxer av.Muxer, demuxer *record.ModuleRecordDemuxer, syncTime bool, staticTS int64) error {

	streams, err := demuxer.ReadHeader()

	if err != nil {

		return err

	}

	err = muxer.WriteHeader(streams)

	if err != nil {

		return err

	}

	for _, i2 := range demuxer.ReadAvPacket(syncTime, staticTS) {

		err := muxer.WritePacket(*i2)

		if err != nil {

			return err

		}
	}

	err = muxer.WriteTrailer()

	if err != nil {

		return err

	}

	return nil

}
