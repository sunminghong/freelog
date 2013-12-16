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
    //"runtime"
)

func Test_test(t *testing.T) {
    inifile := []byte(`
[Default]
flag = 30

[ConsoleAdapter]
level = Trace

[FileAdapter]
level = Info
filename  = log%y%m%d-%h.log
filesize = 20
    `)

	writer = &freeWriter{}
    reader := &iniReader{}
    if err := reader.InitBytes(inifile); err != nil {
        panic(fmt.Sprintf("writer config reader init error: %q", err))
    }

    writer.Init(1000, reader)

    //writer = &wwriter{}

    flag := Ldefault
    Std = NewLoggerExt(writer,"", flag)


//////////////////////////////////////////////////////////////////////////
/*
    for i := 0;i< 10;i ++ {
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

    for i :=0;i<3;i++ {
        pc, file, line, ok := runtime.Caller(i)
        if ok {
            fmt.Println(pc)
            fmt.Println(file)
            fmt.Println(line)
            fmt.Println(ok)
            f := runtime.FuncForPC(pc)
            fmt.Println(f.Name())
        }
    }
*/

    fmt.Println("------------------------------------")
/*
    s := "\n"
	buf := make([]byte, 1024 * 1024)
	n := runtime.Stack(buf, true)
	s += string(buf[:n])
	s += "\n"
	fmt.Println(s)
*/

    //Info("trace")
    for i := 0;i < 1;i ++ {
        //Fatal(fmt.Sprintf("%d",i))
        Trace("Tra111ce")
        Debug("Debug")
        Info("33333333331234Ini看我忘了开始老地方会计师立刻就撒旦法雷克沙剪短发了卡斯就颠覆了萨科技的弗拉思考的房间撒了点付款就爱上浪费空间按时的弗兰克撒的发立刻就撒旦法了看见爱是飞；拉开始减肥；阿斯蒂芬fosdsdf22222222")
        Warn("Wa111rn")
        Error("Error")

        //Panic("Panic")
        //Fatal("Fatal")
    }
    time.Sleep(1 * time.Second)

    writer.SetLevel(AdapterFile,levelError)

    //Info("trace")
    for i := 0;i < 1;i ++ {
        //Fatal(fmt.Sprintf("%d",i))
        Trace("Tra111ce")
        Debug("Debug")
        Info("33333333331234Ini看我忘了开始老地方会计师立刻就撒旦法雷克沙剪短发了卡斯就颠覆了萨科技的弗拉思考的房间撒了点付款就爱上浪费空间按时的弗兰克撒的发立刻就撒旦法了看见爱是飞；拉开始减肥；阿斯蒂芬fosdsdf22222222")
        Warn("Wa111rn")
        Error("Error")

        //Panic("Panic")
        //Fatal("Fatal")
    }


    time.Sleep(1 * time.Second)
}

func log(level int,msg []byte) {
    t := time.Now()
    writer.WriteLog(&t,level,msg)
}

type wwriter struct { }

func (self *wwriter) Init(channelLength int64,reader IConfigReader) {

}

func (self *wwriter) AddLogger(name string, reader IConfigReader) error {
    return nil
}

func (self *wwriter) WriteLog(t *time.Time, level int, msg []byte) {
    fmt.Print(string(msg))
}

