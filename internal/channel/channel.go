/*
	Autor Andrey Semochkin
*/

package channel

import (
	"time"

	"github.com/deepch/vdk/av/pubsub"
	"github.com/sirupsen/logrus"

	"github.com/deepch/Server/internal/spectator"
	"github.com/deepch/vdk/av"
	"github.com/deepch/vdk/format/rtspv2"
	"github.com/sasha-s/go-deadlock"
)

/*

	channel

*/

//Muxers - interface
type muxers interface {
	WriteHeader([]byte) error
	WriteFrame(*av.Packet) error
	Close() error
}

//sessionGuard global sessionGuard
type sessionGuard struct {
	lockGuard deadlock.Mutex
	funcName  string
	startTime time.Time
}

//Channel - element
type Channel struct {
	Name                         string  `yaml:"Name,omitempty" json:"Name,omitempty" groups:"api,configuration,save"`
	DataURL                      string  `yaml:"DataURL,omitempty" json:"DataURL,omitempty" groups:"api,configuration,save"`
	SnapshotURL                  string  `yaml:"SnapshotURL,omitempty" json:"SnapshotURL,omitempty" groups:"api,configuration,save"`
	VOD                          bool    `yaml:"VOD,omitempty" json:"VOD,omitempty" groups:"api,configuration,save"`
	Record                       bool    `yaml:"Record,omitempty" json:"Record,omitempty" groups:"api,configuration,save"`
	RecordTime                   uint32  `yaml:"RecordTime,omitempty" json:"RecordTime,omitempty" groups:"api,configuration,save"`
	Audio                        bool    `yaml:"Audio,omitempty" json:"Audio,omitempty" groups:"api,configuration,save"`
	RTPProxy                     bool    `yaml:"RTPProxy,omitempty" json:"RTPProxy,omitempty" groups:"api,configuration,save"`
	Debug                        bool    `yaml:"Debug,omitempty" json:"Debug,omitempty" groups:"api,configuration,save"`
	Status                       uint32  `yaml:"Status,omitempty" json:"Status,omitempty" groups:"api"`
	Bitrate                      float32 `yaml:"Bitrate,omitempty" json:"Bitrate,omitempty" groups:"api"`
	DebugRaw                     bool    `yaml:"DebugRaw,omitempty" json:"DebugRaw,omitempty" groups:"api,configuration,save"`
	logger                       *logrus.Logger
	spectators                   map[string]*spectator.Spectator
	decoderConfRecordBytes, meta []byte
	deviceUniqID, channelUniqID  string
	cursorFrame, cursorRTP       *pubsub.Queue
	message                      chan int
	needStop                     bool
	muxers                       []muxers
	session                      sessionGuard
	source                       *rtspv2.RTSPClient
	stage                        int
}
