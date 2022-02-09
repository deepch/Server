/*
	Autor Andrey Semochkin
*/

package constant

import (
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//Basic record parameters
var (
	SystemRecordDirectory = "record"
	SystemByteDelimiter   = []byte{255, 0, 255, 0, 255, 0}
	MachineID             = ""
)

/*

Module for describing standard variables and constants

*/

//Main service names
const (
	ServiceFlag           = "service flag"
	ServiceConfigurations = "service configurations"
	ServiceClients        = "service clients"
	ServiceLogging        = "service logging"
	ServiceDevice         = "service device"
	ServiceGRPC           = "service GRPC"
	ServiceRTMP           = "service RTMP"
	ServiceTask           = "service task"
	ServiceHTTP           = "service HTTP"
	ServiceRTPProxy       = "service RTP proxy"
)

//Default video frame types
const (
	notKeyFrame   = iota // 0
	KeyFrame             // 1
	MetaDataFrame        // 2
)

//Communication not editable settings
const (
	ConstDefaultConfigVersion                               = "2.0.0"
	ConstDefaultConfigMark                                  = "save"
	ConstContentJSON                                        = "application/json"
	ConstDeviceChannelReconnectionTimer                     = 2 * time.Second
	ConstDeviceChannelDialTimeout                           = 3 * time.Second
	ConstDeviceChannelReadWriteTimeout                      = 3 * time.Second
	ConstPoster                                             = "poster"
	ConstRecord                                             = "record"
	ConstLoggingLevel                                       = logrus.FatalLevel
	ConstDataFileMime                                       = ".udfv1"
	ConstMetaFileMime                                       = ".umfv1"
	ConstAnalyticsFileMime                                  = ".uafv1"
	ConstStatisticsFileMime                                 = ".usfv1"
	ConstTSFileMime                                         = ".ts"
	ConstTSFragmentName                                     = "fragment"
	ConstJPEGFileMime                                       = "image.jpeg"
	ConstEndM3U8ListTag                                     = "#EXT-X-ENDLIST"
	ConstExtINFTag                                          = "#EXTINF:"
	ConstExtM3U                                             = "#EXTM3U"
	ConstExtPlayListType                                    = "#EXT-X-PLAYLIST-TYPE:"
	ConstVOD                                                = "VOD"
	ConstExtTargetduration                                  = "#EXT-X-TARGETDURATION:"
	ConstExtM3UVersion                                      = "#EXT-X-VERSION:"
	ConstExtM3UVersionDefault                               = "3"
	ConstExtMediaSequence                                   = "#EXT-X-MEDIA-SEQUENCE:"
	ConstRecordHLSMinSegmentTime                            = 4
	ConstRecordTimeFormat                                   = "2006_01_02_15"
	ConstMaxRecordSelectHour                                = 24
	ConstMaxLicenseTry                                      = 6
	ConstOk                                                 = 1
	ConstLicenseEncoding                                    = "23$U$@"
	ConstStatusServer                                       = "https://localhost/api/device_status.php"
	ConstStatisticsServer                                   = "https://localhost/api/statistics_server.php"
	ConstLicenseServer                                      = "https://localhost/api/license_check_key.php"
	ConstHTTPContentType                                    = "Content-Type"
	ConstHTTPContentLength                                  = "Content-Length"
	ConstHTTPContentDisposition                             = "Content-Disposition"
	ConstHTTPContentAttachmentFilename                      = "attachment; filename="
	ConstHTTPMimeTS                                         = "video/MP2T"
	ConstHTTPMimeJPEG                                       = "image/jpeg"
	ConstHTTPMimeMP4                                        = "video/mp4"
	ConstHTTPMimeM3U8                                       = "application/vnd.apple.mpegurl"
	ConstMimeMP4                                            = ".mp4"
	ConstVersion                                            = "0.5"
	ConstSuccess                                            = "Success"
	ConstNone                                               = "None"
	ConstPong                                               = "Pong"
	ConstLoadingSystemService                               = "Loading system service"
	ConstSystemCorrectOffline                               = "System has finished working correctly"
	ConstInitializationPkgServiceFail                       = "Initialization pkg service fail:"
	ConstInitializationAllPkgServiceSuccess                 = "Initialization all pkg service success"
	ConstFinalizeAllPkgServiceFail                          = "Finalize all pkg service fail:"
	ConstInitializationModuleRTMPError                      = "Initialization module RTMP customError:"
	ConstInitializationModuleRTSPError                      = "Initialization module RTSP customError:"
	ConstInitializationModuleHTTPError                      = "Initialization module HTTP customError:"
	ConstInitializationModuleGRPCError                      = "Initialization module GRPC customError:"
	ConstInitializationModuleGRPCErrorFailedRegisterGateway = "Initialization module GRPC Failed Register Gateway customError:"
	ConstInitializationModuleGRPCErrorFailedDialServer      = "Initialization module GREP Failed Dial Server customError:"
	ConstAuthorization                                      = "authorization"
	ConstDefaultConfigurationPatch                          = "/etc/server/config.yaml"
	ConstBackupConfigurationPatch                           = "configs/config.yaml"
	ConstInternalSignalAccepted                             = "Internal signal accepted"
	ConstSystemSignalCopyToInternalSignal                   = "System signal copy to internal signal"
	ConstFlagNameDebug                                      = "FlagDebug"
	ConstFlagNameDebugLevel                                 = "FlagDebugLevel"
	ConstFlagConfigurationPatch                             = "FlagConfigurationPatch"
	ConstNameDebug                                          = "Debug"
	ConstNameDebugLevel                                     = "DebugLevel"
	ConstNameConfigurationPatch                             = "ConfigurationPatch"
	ConstNameDebugDescriptions                              = "enable / disable debug mode"
	ConstNameDebugDefault                                   = false
	ConstNameDebugLevelDefault                              = "customError"
	ConstNameDebugLevelDescriptions                         = "log level (string): trace, debug, info, waring, customError, fatal"
	ConstNameConfigurationPatchDescriptions                 = "full patch to yaml configuration file"
	ConstNameConfigurationPatchDefault                      = ""
	ConstValidateToken                                      = "server"
	ConstRunLevel                                           = "run level"
	ConstStopLevel                                          = "stop level"
	ConstDebugLevelError                                    = "customError"
	ConstSlash                                              = "/"
	ConstComma                                              = ","
	ConstSlashBack                                          = `/`
	ConstLogFormatText                                      = "text"
	ConstLogFormatJSON                                      = "json"
	ConstStringBreak                                        = "\r\n"
	ConstLogFormatNested                                    = "nested"
	ConstLogFormatNestedComponent                           = "component"
	ConstLogFormatNestedCategory                            = "category"
	ConstDefaultCameraLogin                                 = "admin"
	ConstOSWindows                                          = "windows"
	ConstDefaultChannel                                     = "0"
	ConstON                                                 = "ON"
	ConstOFF                                                = "OFF"
	ConstTracing                                            = "tracing"
	ConstDeviceUniqID                                       = "deviceUniqID"
	ConstChannelUniqID                                      = "channelUniqID"
	ConstTimeStart                                          = "timeStart"
	ConstTimeEnd                                            = "timeEnd"
	ConstSeekStart                                          = "seekStart"
	ConstSeekEnd                                            = "seekEnd"
	ConstTimeLine                                           = "timeLine"
	ConstRecordHeaderSize                                   = 50
	ConstRecordHeaderSplitSize                              = 46
	ConstDefaultIDX                                         = 0
)

const (
	ConstContentTypeBinary  = "application/octet-stream"
	ConstContentTypeForm    = "application/x-www-form-urlencoded"
	ConstContentTypeJSON    = "application/json"
	ConstContentTypeHTML    = "text/html; charset=utf-8"
	ConstContentTypeText    = "text/plain; charset=utf-8"
	ConstContentTypeMpegURL = "application/vnd.apple.mpegurl"
	ConstContentTypeMP2T    = "video/MP2T"
)

//cron settings section
const (
	ConstTaskDeleteFiles           = "0 10 * * * *" //Every 10 minutes of every hour (once an hour)
	ConstTaskGetImage              = "0 42 * * * *" //Every 42 minutes of every hour (once an hour)
	ConstTaskGetInstallationStatus = "0 30 * * * *" //Every 30 minutes of every hour (once an hour)
	ConstTaskSendStatus            = "0 * * * * *"  //Every 0 second (once per minute)
	ConstSendPlatformStatistics    = "30 * * * * *" //Every 30 seconds (once per minute)
)

//http server section
const (
	ConstHTTPGET                       = "GET"
	ConstHTTPTimeOut                   = 30 * time.Minute
	ConstHTTPTimeJSON                  = `{"code": -1, "msg":"http: Handler timeout"}`
	ConstHTTPTimeMsg                   = "timeout happen, url:"
	ConstHTTPWWWPrefix                 = "www."
	ConstHTTPCurrentVersion            = "/V1"
	ConstHTTPCurrentGRPCPatch          = "/GRPC"
	ConstHTTPMethodAny                 = "/*any"
	ConstHTTPMethodDemoWebsocket       = "/demo/websocket.html"
	ConstHTTPMethodDemoWebRTC          = "/demo/webrtc.html"
	ConstHTTPMethod404                 = "assets/web/404.html"
	ConstHTTPMethodFavicon             = "assets/ico/favicon.ico"
	ConstHTTPMethodSwagger             = "/swagger/*any"
	ConstHTTPMethodDoc                 = "doc/doc.json"
	ConstHTTPMethodV1PlayerWebsocket   = "/Player/WebSocket/Live/:deviceUniqID/:channelUniqID"
	ConstHTTPMethodV1PlayerJPEG        = "/Player/JPEG/Live/:deviceUniqID/:channelUniqID/:fileName"
	ConstHTTPMethodV1PlayerHLSRecord   = "/Player/HLS/Record/:deviceUniqID/:channelUniqID/:timeStart/:timeEnd/:fileName"
	ConstHTTPMethodV1PlayerHLSRecordTS = "/Player/HLS/Record/:deviceUniqID/:channelUniqID/:timeStart/:timeEnd/fragment/:fileName/:seekStart/:seekEnd/:timeLine/:fileName"
	ConstHTTPMethodV1PlayerMP4Record   = "/Player/MP4/Record/:deviceUniqID/:channelUniqID/:timeStart/:timeEnd/:fileName"
	ConstHTTPUriDoc                    = "/doc/doc.json"
	ConstHTTPPatchFavicon              = "./assets/ico/favicon.ico"
	ConstHTTPPatchDemoWebsocket        = "./assets/web/demo/websocket.html"
	ConstHTTPPatchDemoWebRTC           = "./assets/web/demo/webrtc.html"
	ConstHTTPPatchDocSwagger           = "assets/docs/server.swagger.json"
)

//Status of channels that can be
const (
	UNDEFINED = iota
	ONLINE
	WAIT
	CONNECT
	CONNECTED
	DISCONNECTED
	ERROR
)

//Description of basic types of delivery protocols
const (
	RTMP = iota
	RTSP
)

//MessageStop - Thread Processing Signals
const (
	MessageStop = iota
)

//errorCustom - error types
type errorCustom string

/*

	Basic server messages

*/

//Basic Messages
const (
	ErrorDeviceNotFound                  = errorCustom("deviceElement not found")
	ErrorDeviceChannelNotCodecInfo       = errorCustom("deviceElement channel not codec info")
	ErrorDeviceChannelReceiveCloseSignal = errorCustom("deviceElement channel receive close signal")
	ErrorDeviceChannelRouteAV            = errorCustom("deviceElement channel route AV")
	ErrorDeviceChannelCloseMuxer         = errorCustom("deviceElement channel close muxer")
	ErrorBadMachineID                    = errorCustom("bad machine ID")
	ErrorOnvifMediaUrlNotFound           = errorCustom("onvif media url not found")
	ErrorOnvifTokenNotFound              = errorCustom("onvif token not found")
	ErrorBadURL                          = errorCustom("bad url")
	ErrMetaDataNotFound                  = errorCustom("meta data not found")
)

/*

	Basic License Messages

*/

//License Messages
const (
	ErrorBadLicense       = errorCustom("validation bad validation")
	ErrorBadLicenseStatus = errorCustom("validation bad validation status")
)

//ErrorBadPortPort  - Server messages
const (
	ErrorBadPortPort = errorCustom("bad port number")
)

/*

	Status grpc

*/

//Mistakes from different places
var (
	ErrNoAuthHeader      = status.Error(codes.Unauthenticated, "authorization header required")
	ErrCouldNotAuthorize = status.Error(codes.Internal, "authorization failed")
	ErrIncorrectToken    = status.Error(codes.Unauthenticated, "authorization fail incorrect token")
	ErrIncorrectPeriod   = status.Error(codes.Unauthenticated, "customError incorrect period")
	ErrDeviceNotFound    = status.Error(codes.NotFound, "device not found")
	ErrChannelNotFound   = status.Error(codes.NotFound, "channel not found")
	ErrSpectatorNotFound = status.Error(codes.NotFound, "spectator not found")
)

/*

	Basic User Messages

*/

func (e errorCustom) Error() string { return string(e) }
