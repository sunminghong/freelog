/*=============================================================================
#     FileName: logger.go
#       Author: sunminghong, allen.fantasy@gmail.com, http://weibo.com/5d13
#         Team: http://1201.us
#   LastChange: 2013-12-14 09:56:59
#      History:
=============================================================================*/


/*
一个通用的、多输出、多级别的go 日志库，并且每个输出可以分别定义级别；可以方便自定义输出类
*/
package freelog

import (
    "time"
)

type IConfigReader interface {
    Init(file *string) error
    InitBytes(conf []byte) error
    GetAdapters() (adapterNames []string)
    GetBool(adapterName string, option string) (value bool, err error)
    GetFloat64(adapterName string, option string) (value float64, err error)
    GetString(adaterName string, option string) (value string, err error)
}

type IWriter interface {
	Init(channelLength int64,reader IConfigReader)
	AddAdapter(name string, reader IConfigReader) error
    SetLevel(name string,level int)
    WriteLog(t *time.Time, level int, msg []byte)
}

type IAdapter interface {
    Init(config IConfigReader) (level int, err error)
    GetLevel() (level int)
    SetLevel(level int)
    Write(msg *LogMsg) (err error)
    Close()
}

type GetAdapter func() IAdapter
var registeredAdapters map[string]GetAdapter = make(map[string]GetAdapter)


type LogMsg struct {
    Timestamp *time.Time
    Level int
    Msg   []byte
}


