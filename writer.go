/*=============================================================================
#     FileName: baselog.go
#       Author: sunminghong, allen.fantasy@gmail.com, http://weibo.com/5d13
#         Team: http://1201.us
#   LastChange: 2013-12-14 10:18:33
#      History:
=============================================================================*/

/*
一个通用的、多输出、多级别的go 日志库，并且每个输出可以分别定义级别；
可以方便自定义输出类
*/
package freelog

import (
    "fmt"
    "sync"
    "time"
)

type freeWriter struct {
    registeredAdapters map[string]IAdapter
    lowestLevel       int
    msgChannel        chan *LogMsg
    lock              sync.Mutex
}

func (this *freeWriter) Init(channelLength int64, configReader IConfigReader) {
    this.registeredAdapters = make(map[string]IAdapter)
    this.msgChannel = make(chan *LogMsg, channelLength)
    this.lowestLevel = LevelOff

    this.setAdapters(configReader)

    go this.runLog()
}

func (this *freeWriter) setAdapters(configReader IConfigReader) {
    adps := configReader.GetAdapters()
    if adps == nil {
        panic("logger config reader error!")
    }

    for _, adp := range adps {
        if adp, err := CheckAdapter(adp); err == nil {
            this.AddAdapter(adp, configReader)
        } else {
            fmt.Printf("日志适配器没有注册：%q \n", adp)
        }
    }

    return
}

func (this *freeWriter) SetLevel(name string, level int) {
    adp,ok := this.registeredAdapters[name]
    if !ok {
        return
    }

    adp.SetLevel(level)

    for _,adp := range this.registeredAdapters {
        if this.lowestLevel > adp.GetLevel() {
            this.lowestLevel = level
        }
    }
}

func (this *freeWriter) AddAdapter(
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
    if level == LevelOff {
        return nil
    }

    this.registeredAdapters[adpName] = logger

    if this.lowestLevel > level {
        this.lowestLevel = level
    }

    return nil
}

func (this *freeWriter) WriteLog(t *time.Time, level int, msg []byte) {
    if this.lowestLevel > level {
        return
    }

    bufs := make([]byte, len(msg))
    copy(bufs,msg)

    logmsg := new(LogMsg)
    logmsg.Timestamp = t
    logmsg.Level = level
    logmsg.Msg = bufs
    this.msgChannel <- logmsg

    return
}

func (this *freeWriter) Close() {
    for _, logger := range this.registeredAdapters {
        logger.Close()
    }
}

func (this *freeWriter) runLog() {
    for {
        select {
        case logmsg := <-this.msgChannel:
            for _, logger := range this.registeredAdapters {
                logger.Write(logmsg)
            }
        }
    }
}
