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
    "runtime"
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

	output = &freeOutput{}

	reader := &iniReader{}
	if err := reader.InitBytes(inifile); err != nil {
		panic(fmt.Sprintf("output config reader init error: %q", err))
	}

	output.Init(1000, reader)

    flag := Ldefault
    Std = NewLoggerext(output,"", flag)


//////////////////////////////////////////////////////////////////////////

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
    fmt.Println("------------------------------------")


    s := "\n"
	buf := make([]byte, 1024 * 1024)
	n := runtime.Stack(buf, true)
	s += string(buf[:n])
	s += "\n"
	fmt.Println(s)


    for i := 0;i< 10;i ++ {
        Fatal(fmt.Sprintf("%d",i))
        Trace("Trace\n")
        Debug("Debug\n")
        Info("Info\n")
        Warn("Warn\n")
        Error("Error\n")
        Panic("Panic\n")
        Fatal("Fatal\n")
    }
    time.Sleep(1 * time.Second)



}

func log(level int,msg []byte) {
    t := time.Now()
    output.WriteLog(&t,level,msg)
}
