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
	levelAll = iota
	levelTrace
	levelDebug
	levelInfo
	levelWarn
	levelError
    levelPanic
	levelFatal
	levelOff
)

const (
	AdapterConsole = "consolelogger"
	AdapterFile    = "filelogger"
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
		return levelAll, nil
	case "TRACE":
		fallthrough
	case "1":
		return levelTrace, nil
	case "DEBUG":
		fallthrough
	case "2":
		return levelDebug, nil
	case "INFO":
		fallthrough
	case "3":
		return levelInfo, nil
	case "WARN":
		fallthrough
	case "4":
		return levelWarn, nil
	case "ERROR":
		fallthrough
	case "5":
		return levelError, nil
	case "FATAL":
		fallthrough
	case "6":
		return levelFatal, nil
	default:
		return levelOff, errors.New("日志等级无效！将采用最低等级！")
	}

}
