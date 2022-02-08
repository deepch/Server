/*
	Autor Andrey Semochkin
*/

package function

/*

	Basic set of functions and value type convectors

*/

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"io/ioutil"
	"net/url"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/deepch/Server/internal/constant"
	uuid1 "github.com/google/uuid"
	uuid2 "github.com/satori/go.uuid"
)

//SubStrByShowLen preparing a string for a license
func SubStrByShowLen(val, suffix string, l int) string {

	if len(val) <= l {

		return val

	}

	ss, sl, rl, rs := "", 0, 0, []rune(val)

	prefixes := len(suffix)

	for _, r := range rs {

		rent := int(r)

		if rent < 128 {

			rl = 1

		} else {

			rl = 2

		}

		if sl+rl+prefixes > l {

			break

		}

		sl += rl

		ss += string(r)

	}

	if sl < prefixes {

		return ss

	}

	return ss + suffix

}

//CreateUniqID - google + satori random generator uuid
func CreateUniqID() string {

	u1 := uuid1.New().String()

	u2 := uuid2.NewV4().String()

	uuid := strings.Replace(u1+u2, "-", "", -1)

	return uuid

}

//StringConvertToInt - convert string to number without error handling
func StringConvertToInt(val string) int {

	i, err := strconv.Atoi(val)

	if err != nil {

		return 0

	}

	return i

}

//BoolConvertToStringSwitch - translates the value into в on / off
func BoolConvertToStringSwitch(val bool) string {

	if val {
		return constant.ConstON
	}

	return constant.ConstOFF
}

//FileExists - does the file exist
func FileExists(val string) (bool, error) {

	_, err := os.Stat(val)

	if err == nil {

		return true, nil

	}

	if errors.Is(err, os.ErrNotExist) {

		return false, nil

	}

	return false, err
}

//ParsURLTwoVal -
//fetches parameters from URL
func ParsURLTwoVal(val *url.URL) (string, string, error) {

	if val == nil {
		return "", "", constant.ErrorBadURL
	}

	//Split the path by deleting the left one /
	splicedString := strings.Split(strings.TrimLeft(val.Path, constant.ConstSlash), constant.ConstSlash)

	if len(splicedString) == 1 {

		return splicedString[0], constant.ConstDefaultChannel, nil

	} else if len(splicedString) == 2 {

		return splicedString[0], splicedString[1], nil

	} else {

		return "", "", constant.ErrorBadURL

	}

}

//URLDecode - decodes a line from base64
func URLDecode(val string) string {

	data, err := base64.StdEncoding.DecodeString(val)

	if err != nil {

		return val

	}

	return string(data)

}

//SignalToError - signal in error
func SignalToError(val int) error {

	switch val {

	case constant.MessageStop:

		return constant.ErrorDeviceChannelReceiveCloseSignal

	}

	return nil

}

//DeleteLastSlashes - removes an extra slash at the end of a line
func DeleteLastSlashes(val string) string {

	if len(val) > 1 {

		if runtime.GOOS == constant.ConstOSWindows {

			val = strings.TrimRight(val, constant.ConstSlashBack)

		} else {

			val = strings.TrimRight(val, constant.ConstSlash)

		}

	}

	return val

}

//CreateServicePatch - correct path generator
func CreateServicePatch(deviceUUID, channelUUID, prefix string) string {

	return constant.SystemRecordDirectory + swapSlashes(constant.ConstSlash+deviceUUID+constant.ConstSlash+prefix+constant.ConstSlash+channelUUID+constant.ConstSlash)

}

//CreateFileByPatch - creates a file at the specified path and continues writing there
func CreateFileByPatch(val string) (*os.File, error) {

	return os.OpenFile(val, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

}

//SwapSlashes - reverses slashes if OS Windows
func swapSlashes(val string) string {

	if runtime.GOOS == constant.ConstOSWindows {

		return strings.Replace(val, constant.ConstSlash, constant.ConstSlashBack, -1)

	}

	return val

}

//CreatePatch - creates path if it doesn't exist
func CreatePatch(val string) error {

	if _, err := os.Stat(val); os.IsNotExist(err) {

		err := os.MkdirAll(val, os.ModePerm)

		if err != nil {

			return err

		}

	}

	return nil

}

//TimeRFC3339toTime - parses the line into int64 unix nano
func TimeRFC3339toTime(val string) (int64, error) {

	t, err := time.Parse(time.RFC3339, val)

	if err != nil {

		return 0, err

	}

	return t.UnixNano(), nil

}

//GetMD5Hash - getting md5 string
func GetMD5Hash(val string) string {

	hash := md5.Sum([]byte(val))

	return hex.EncodeToString(hash[:])

}

//FindConfigurationRead - searches and reads configuration from disk
func FindConfigurationRead(flagConfigurationPatch string) ([]byte, error) {

	var patch string

	//If the path of the configuration file is specified through the flags module, use it
	if flagConfigurationPatch != "" {

		patch = flagConfigurationPatch

	} else {

		//We are looking for a configuration in the directory specified in the constant ConstDefaultConfigurationPatch
		if ok, err := FileExists(constant.ConstDefaultConfigurationPatch); ok && err == nil {

			patch = constant.ConstDefaultConfigurationPatch

			//We are looking for a configuration in the directory specified in the constant	ConstBackupConfigurationPatch
		} else if ok, err = FileExists(constant.ConstBackupConfigurationPatch); ok && err == nil {

			patch = constant.ConstBackupConfigurationPatch

		} else {

			return nil, err

		}

	}

	//Reading the config file from disk
	data, err := ioutil.ReadFile(patch)

	if err != nil {

		return nil, err

	}

	return data, nil

}

//FindConfigurationWrite - searches and writes configuration from disk
func FindConfigurationWrite(flagConfigurationPatch string, val []byte) error {

	var patch string

	//If the path of the configuration file is specified through the flags module, use it
	if flagConfigurationPatch != "" {

		patch = flagConfigurationPatch

	} else {

		//We are looking for a configuration in the directory specified in the constant ConstDefaultConfigurationPatch
		if ok, err := FileExists(constant.ConstDefaultConfigurationPatch); ok && err == nil {

			patch = constant.ConstDefaultConfigurationPatch

			//We are looking for a configuration in the directory specified in the constant	ConstBackupConfigurationPatch
		} else if ok, err = FileExists(constant.ConstBackupConfigurationPatch); ok && err == nil {

			patch = constant.ConstBackupConfigurationPatch

		} else {

			return err

		}

	}

	err := ioutil.WriteFile(patch, val, 0644)

	if err != nil {

		return err

	}

	return nil

}

//GetCurrentFuncName - returns the function name
func GetCurrentFuncName() string {

	pc, _, _, _ := runtime.Caller(2)

	return runtime.FuncForPC(pc).Name()

}
