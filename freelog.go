/*=============================================================================
#     FileName: writer.go
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
    "os"
    "runtime"
    "strings"
    "time"
)

//default writer variable
var Std *LoggerExt
var writer IWriter
var CallDepth int = 2

func init() {
    reader := &iniReader{}
    reader.InitBytes([]byte(`
[ConsoleAdapter]
level = Trace
    `))

    _Init(reader)
}

func Start(inifile *string) {
    if *inifile == "" {
        return
    }

    reader := &iniReader{}
    if err := reader.Init(inifile); err != nil {
        fmt.Println("***writer config reader init error:\n    %q", err)
        return
    }

    _Init(reader)
}

func _Init(reader IConfigReader) {
    writer = &freeWriter{}
    writer.Init(1000, reader)
    fmt.Println("freelog is starting")

    flag := Ldefault
    Std = NewLoggerExt(writer, "", flag)
}

func RegisterAdapter(adapterName string, adapter GetAdapter) {
    if adapter == nil {
        fmt.Printf("日志适配器注册失败：%q \n", adapterName)
        return
    }
    adapterName = strings.ToLower(adapterName)

    if _, ok := registeredAdapters[adapterName]; ok {
        fmt.Printf("日志适配器注册失败，系统已存在日志：%q \n", adapterName)
        return
    } else {
        registeredAdapters[adapterName] = adapter
    }
}

// SetWriter sets the writer destination for the standard writer.
func SetWriter(w IWriter) {
    Std.mu.Lock()
    defer Std.mu.Unlock()
    Std.out = w
}

// Flags returns the writer flags for the standard writer.
func Flags() int {
    return Std.Flags()
}

// SetFlags sets the writer flags for the standard writer.
func SetFlags(flag int) {
    Std.SetFlags(flag)
}

func sout(v ...interface{}) string {
    s, ok := v[0].(string)
    if !ok || strings.Index(s, "%") == -1 {
        return fmt.Sprint(v...)
    } else {
        return fmt.Sprintf(s, v[1:]...)
    }
}

// -----------------------------------------

// Print calls Output to print to the standard writer.
// Arguments are handled in the manner of fmt.Print.
func Print(v ...interface{}) {
    Std.Output(LevelInfo, CallDepth, sout(v...))
}

// Printf calls Output to print to the standard writer.
// Arguments are handled in the manner of fmt.Printf.
func Printf(format string, v ...interface{}) {
    Std.Output(LevelInfo, CallDepth, fmt.Sprintf(format, v...))
}

// Println calls Output to print to the standard writer.
// Arguments are handled in the manner of fmt.Println.
func Println(v ...interface{}) {
    Std.Output(LevelInfo, CallDepth, fmt.Sprintln(v...))
}

// -----------------------------------------

func Tracef(format string, v ...interface{}) {
    Std.Output(LevelTrace, CallDepth, fmt.Sprintf(format, v...))
}

func Trace(v ...interface{}) {
    Std.Output(LevelTrace, CallDepth, sout(v...))
}

// -----------------------------------------

func Debugf(format string, v ...interface{}) {
    Std.Output(LevelDebug, CallDepth, fmt.Sprintf(format, v...))
}

func Debug(v ...interface{}) {
    Std.Output(LevelDebug, CallDepth, sout(v...))
}

// -----------------------------------------

func Infof(format string, v ...interface{}) {
    Std.Output(LevelInfo, CallDepth, fmt.Sprintf(format, v...))
}

func Info(v ...interface{}) {
    Std.Output(LevelInfo, CallDepth, sout(v...))
}

// -----------------------------------------

func Warnf(format string, v ...interface{}) {
    Std.Output(LevelWarn, CallDepth, fmt.Sprintf(format, v...))
}

func Warn(v ...interface{}) { Std.Output(LevelWarn, CallDepth, sout(v...)) }

// -----------------------------------------

func Errorf(format string, v ...interface{}) {
    Std.Output(LevelError, CallDepth, fmt.Sprintf(format, v...))
}

func Error(v ...interface{}) {
    Std.Output(LevelError, CallDepth, sout(v...))
}

// -----------------------------------------

// Fatal is equivalent to Print() followed by a call to os.Exit(1).
func Fatal(v ...interface{}) {
    Std.Output(LevelFatal, CallDepth, sout(v...))
    os.Exit(1)
}

// Fatalf is equivalent to Printf() followed by a call to os.Exit(1).
func Fatalf(format string, v ...interface{}) {
    Std.Output(LevelFatal, CallDepth, fmt.Sprintf(format, v...))
    os.Exit(1)
}

// Fatalln is equivalent to Println() followed by a call to os.Exit(1).
func Fatalln(v ...interface{}) {
    Std.Output(LevelFatal, CallDepth, fmt.Sprintln(v...))
    os.Exit(1)
}

// -----------------------------------------

// Panic is equivalent to Print() followed by a call to panic().
func Panic(v ...interface{}) {
    s := sout(v...)
    Std.Output(LevelPanic, CallDepth, s)
    panic(s)
}

// Panicf is equivalent to Printf() followed by a call to panic().
func Panicf(format string, v ...interface{}) {
    s := fmt.Sprintf(format, v...)
    Std.Output(LevelPanic, CallDepth, s)
    panic(s)
}

// Panicln is equivalent to Println() followed by a call to panic().
func Panicln(v ...interface{}) {
    s := fmt.Sprintln(v...)
    Std.Output(LevelPanic, CallDepth, s)
    panic(s)
}

// -----------------------------------------

func Stack(v ...interface{}) {
    s := sout(v...)
    s += "\n"
    buf := make([]byte, 1024*1024)
    n := runtime.Stack(buf, true)
    s += string(buf[:n])
    s += "\n"
    Std.Output(LevelError, CallDepth, s)
}

func PrintPanicStack() {
    if x := recover(); x != nil {
        now := time.Now()
        for i := 0; i < 10; i++ {
            funcName, file, line, ok := runtime.Caller(i)
            if ok {
                bs := fmt.Sprintf("frame %v:[func:%v,file:%v,line:%v]\n", i, runtime.FuncForPC(funcName).Name(), file, line)
                writer.WriteLog(&now, LevelError , []byte(bs))
            }
        }
    }
}
