/*=============================================================================
#     FileName: const.go
#       Author: sunminghong, allen.fantasy@gmail.com, http://weibo.com/5d13
#         Team: http://1201.us
#   LastChange: 2013-12-14 15:50:25
#      History:
=============================================================================*/


package freelog

import (
	"errors"
    "strings"
)

const (
	LevelAll = iota
	LevelTrace
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
    LevelPanic
	LevelFatal
	LevelOff
)

func CheckAdapter(adaptername string) (string, error) {
	switch strings.ToLower(adaptername) {
	case AdapterConsole:
		return AdapterConsole, nil
	case AdapterFile:
		return AdapterFile, nil
	default:
		return adaptername, errors.New("非系统指定适配器！")
	}
}

func CheckLevel(level string) (int, error) {
	switch strings.ToUpper(level) {
	case "ALL":
		fallthrough
	case "0":
		return LevelAll, nil
	case "TRACE":
		fallthrough
	case "1":
		return LevelTrace, nil
	case "DEBUG":
		fallthrough
	case "2":
		return LevelDebug, nil
	case "INFO":
		fallthrough
	case "3":
		return LevelInfo, nil
	case "WARN":
		fallthrough
	case "4":
		return LevelWarn, nil
	case "ERROR":
		fallthrough
	case "5":
		return LevelError, nil
	case "FATAL":
		fallthrough
	case "6":
		return LevelFatal, nil
	default:
		return LevelOff, errors.New("日志等级无效！将采用最低等级！")
	}

}
