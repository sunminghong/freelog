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
    "fmt"
)

const (
	AdapterConsole = "consoleadapter"
)

type ConsoleAdapter struct {
    level     int

    section string
}

func init() {
    RegisterAdapter(AdapterConsole, NewConsoleAdapter)
}

func NewConsoleAdapter() IAdapter {
    clog := new(ConsoleAdapter)
    clog.section = AdapterConsole
    clog.level = levelWarn
    return clog
}

func (this *ConsoleAdapter) GetLevel() int {
    return this.level
}

func (this *ConsoleAdapter) SetLevel(level int) {
    this.level = level
}

func (this *ConsoleAdapter) Init(ini IConfigReader) (level int, err error) {
    lev, err := ini.GetString(this.section, "level")
    if err != nil {
        this.level = levelWarn
    } else {
        this.level, err = CheckLevel(lev)
    }

    return this.level, nil
}

func (this *ConsoleAdapter) Write(msg *LogMsg) (err error) {
    if this.level > msg.Level {
        return nil
    }
    //s := "\n"
	//buf := make([]byte, 1024 * 1024)
	//n := runtime.Stack(buf, true)
	//s += string(buf[:n])
	//s += "\n"
    //fmt.Println(s)

    fmt.Print(string(msg.Msg))

    return nil
}

func (this *ConsoleAdapter) Close() {

}
