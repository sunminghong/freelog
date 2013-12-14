/*=============================================================================
#     FileName: baselog.go
#       Author: sunminghong, allen.fantasy@gmail.com, http://weibo.com/5d13
#         Team: http://1201.us
#   LastChange: 2013-12-14 10:18:33
#      History:
=============================================================================*/

/*
一个通用的、多输出、多级别的go 日志库，并且每个输出可以分别定义级别；可以方便自定义输出类
*/
package freelog

import (
    "fmt"
    "sync"
    "time"
)

type baseLog struct {
    registeredLoggers map[string]ILogger
    lowestLevel       int
    msgChannel        chan *LogMsg
    lock              sync.Mutex
}

func (this *baseLog) Init(channelLength int64, configReader IConfigReader) {
    this.registeredLoggers = make(map[string]ILogger)
    this.msgChannel = make(chan *LogMsg, channelLength)
    this.lowestLevel = levelOff

    this.setLoggers(configReader)

    go this.runLog()
}

func (this *baseLog) setLoggers(configReader IConfigReader) {
    adps := configReader.GetAdapters()
    if adps == nil {
        panic("logger config reader error!")
    }

    for _, adp := range adps {
        if adp, err := CheckAdapter(adp); err == nil {
            this.AddLogger(adp, configReader)
        } else {
            fmt.Printf("日志适配器没有注册：%q \n", adp)
        }
    }

    return
}

func (this *baseLog) AddLogger(
    adpName string, configReader IConfigReader) error {

    this.lock.Lock()
    defer this.lock.Unlock()

    getAdapter, ok := registeredAdapters[adpName]
    if !ok {
        fmt.Printf("未注册的日志适配器：%q \n", adpName)
        return fmt.Errorf("未注册的日志适配器：%q", adpName)
    }

    logger := getAdapter()
    level, err := logger.Init(configReader)
    if err != nil {
        fmt.Printf("日志输出器初始化失败：%q,%v \n", adpName, err)
        return err
    }
    if level == levelOff {
        return nil
    }

    this.registeredLoggers[adpName] = logger

    if this.lowestLevel > level {
        this.lowestLevel = level
    }

    return nil
}

func (this *baseLog) WriteLog(t *time.Time, level int, msg []byte) {
    if this.lowestLevel > level {
        return
    }

    logmsg := new(LogMsg)
    logmsg.Timestamp = t
    logmsg.Level = level
    logmsg.Msg = msg
    this.msgChannel <- logmsg

    return
}

func (this *baseLog) Close() {
    for _, logger := range this.registeredLoggers {
        logger.Close()
    }
}

func (this *baseLog) runLog() {
    for {
        select {
        case logmsg := <-this.msgChannel:
            for _, logger := range this.registeredLoggers {
                //fmt.Println("logger outer:", i,logmsg.Level,string(logmsg.Msg))
                logger.Write(logmsg)
            }
        }
    }
}
