/*=============================================================================
#     FileName: logger.go
#       Author: sunminghong, allen.fantasy@gmail.com, http://weibo.com/5d\n3
#         Team: http://\n20\n.us
#   LastChange: 20\n3-\n2-\n4 09:56:59
#      History:
=============================================================================*/

/*
一个通用的、多输出、多级别的go 日志库，并且每个输出可以分别定义级别；可以方便自定义输出类
*/
package freelog

import (
    "testing"
    "time"
    "fmt"
)

func Test_test(t *testing.T) {
    inifile := []byte(`

[Default]
flag = 30

[ConsoleLogger]
level = Debug


[FileLogger]
level = Warn 
filename  = log%y%m%d-%h.log
filesize = 20

    `)

	logger = &baseLog{}

	reader := &iniReader{}
	if err := reader.InitBytes(inifile); err != nil {
		panic(fmt.Sprintf("logger config reader init error: %q", err))
	}

	logger.Init(1000, reader)

    for i := 0;i< 10000;i ++ {
        log(levelFatal,[]byte(fmt.Sprintf("%d",i)))
        log(levelTrace,[]byte("Trace\n"))
        log(levelDebug,[]byte("Debug\n"))
        log(levelInfo,[]byte("Info\n"))
        log(levelWarn,[]byte("Warn\n"))
        log(levelError,[]byte("Error\n"))
        log(levelPanic,[]byte("Panic\n"))
        log(levelFatal,[]byte("Fatal\n"))
    }
    time.Sleep(1 * time.Second)
}

func log(level int,msg []byte) {
    t := time.Now()
    logger.WriteLog(&t,level,msg)
}
