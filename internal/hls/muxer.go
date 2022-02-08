/*
	Autor Andrey Semochkin
*/

package hls

import (
	"bytes"
	"strconv"
	"time"

	"github.com/deepch/Server/internal/function"

	"github.com/deepch/Server/internal/constant"
	"github.com/sirupsen/logrus"
)

/*

	Микшер playlist hls

*/

//HLS - main structure
type HLS struct {
	Logger           *logrus.Logger
	allBlockDuration time.Duration
	header, body     *bytes.Buffer
	sequence         int
}

//Muxer - create muxer hls
func Muxer() *HLS {

	return &HLS{header: bytes.NewBuffer([]byte{}), body: bytes.NewBuffer([]byte{})}

}

//WriteBlock - hls block write
func (service *HLS) WriteBlock(blockDuration time.Duration, blockStart int64, blockEnd int64, fileName string) {

	service.body.WriteString(constant.ConstExtINFTag)
	service.body.WriteString(strconv.FormatFloat(blockDuration.Seconds(), 'f', 6, 64))
	service.body.WriteString(constant.ConstComma)
	service.body.WriteString(constant.ConstStringBreak)
	service.body.WriteString(constant.ConstTSFragmentName)
	service.body.WriteString(constant.ConstSlash)
	service.body.WriteString(fileName)
	service.body.WriteString(constant.ConstSlash)
	service.body.WriteString(strconv.FormatInt(blockStart, 10))
	service.body.WriteString(constant.ConstSlash)
	service.body.WriteString(strconv.FormatInt(blockEnd, 10))
	service.body.WriteString(constant.ConstSlash)
	service.body.WriteString(strconv.FormatInt(service.allBlockDuration.Milliseconds(), 10))
	service.body.WriteString(constant.ConstSlash)
	service.body.WriteString(function.CreateUniqID())
	service.body.WriteString(constant.ConstTSFileMime)
	service.body.WriteString(constant.ConstStringBreak)
	service.allBlockDuration += blockDuration

}

//WriteTrailer - end writer
func (service *HLS) WriteTrailer() []byte {

	service.header.WriteString(constant.ConstExtM3U)
	service.header.WriteString(constant.ConstStringBreak)
	service.header.WriteString(constant.ConstExtPlayListType)
	service.header.WriteString(constant.ConstVOD)
	service.header.WriteString(constant.ConstStringBreak)
	service.header.WriteString(constant.ConstExtTargetduration)
	service.header.WriteString(strconv.Itoa(int(service.allBlockDuration.Seconds())))
	service.header.WriteString(constant.ConstStringBreak)
	service.header.WriteString(constant.ConstExtM3UVersion)
	service.header.WriteString(constant.ConstExtM3UVersionDefault)
	service.header.WriteString(constant.ConstStringBreak)
	service.header.WriteString(constant.ConstExtMediaSequence)
	service.header.WriteString(strconv.Itoa(service.sequence))
	service.header.WriteString(constant.ConstStringBreak)
	service.header.Write(service.body.Bytes())
	service.header.WriteString(constant.ConstEndM3U8ListTag)

	return service.header.Bytes()

}
