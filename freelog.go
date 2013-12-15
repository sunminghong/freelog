/*=============================================================================
#     FileName: output.go
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
    "strings"
    "os"
    "runtime"
)

//default output variable
var Std *Loggerext
var output IOutput

func init() {
	registeredAdapters = make(map[string]GetAdapter)

    //file := "./logconfig.ini"
	//Reload(&file)
}

func Start(inifile *string) {

	output = &freeOutput{}

	reader := &iniReader{}
	if err := reader.Init(inifile); err != nil {
		panic(fmt.Sprintf("output config reader init error: %q", err))
	}

	output.Init(1000, reader)

    flag := Ldefault
    Std = NewLoggerext(output,"", flag)
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

// SetOutput sets the output destination for the standard output.
func SetOutput(w iwriter) {
	Std.mu.Lock()
	defer Std.mu.Unlock()
	Std.out = w
}

// Flags returns the output flags for the standard output.
func Flags() int {
	return Std.Flags()
}

// SetFlags sets the output flags for the standard output.
func SetFlags(flag int) {
	Std.SetFlags(flag)
}

func sout(v ...interface{}) string {
    s,ok := v[0].(string)
    if !ok || strings.Index(s,"%") == -1 {
        return fmt.Sprint(v)
    } else {
        return fmt.Sprintf(s,v[1:]...)
    }
}

// -----------------------------------------

// Print calls Output to print to the standard output.
// Arguments are handled in the manner of fmt.Print.
func Print(v ...interface{}) {
	Std.Output(levelInfo, 2, sout(v...))
}

// Printf calls Output to print to the standard output.
// Arguments are handled in the manner of fmt.Printf.
func Printf(format string, v ...interface{}) {
	Std.Output(levelInfo, 2, fmt.Sprintf(format, v...))
}

// Println calls Output to print to the standard output.
// Arguments are handled in the manner of fmt.Println.
func Println(v ...interface{}) {
	Std.Output(levelInfo, 2, fmt.Sprintln(v...))
}

// -----------------------------------------

func Tracef(format string, v ...interface{}) {
	Std.Output(levelTrace, 2, fmt.Sprintf(format, v...))
}

func Trace(v ...interface{}) {
	Std.Output(levelTrace, 2, sout(v...))
}


// -----------------------------------------

func Debugf(format string, v ...interface{}) {
	Std.Output(levelDebug, 2, fmt.Sprintf(format, v...))
}

func Debug(v ...interface{}) {
	Std.Output(levelDebug, 2, sout(v...))
}

// -----------------------------------------

func Infof(format string, v ...interface{}) {
	Std.Output(levelInfo, 2, fmt.Sprintf(format, v...))
}

func Info(v ...interface{}) {
	Std.Output(levelInfo, 2, sout(v...))
}

// -----------------------------------------

func Warnf(format string, v ...interface{}) {
	Std.Output(levelWarn, 2, fmt.Sprintf(format, v...))
}

func Warn(v ...interface{}) { Std.Output(levelWarn, 2, sout(v...)) }

// -----------------------------------------

func Errorf(format string, v ...interface{}) {
	Std.Output(levelError, 2, fmt.Sprintf(format, v...))
}

func Error(v ...interface{}) {
    Std.Output(levelError, 2, sout(v...))
}

// -----------------------------------------

// Fatal is equivalent to Print() followed by a call to os.Exit(1).
func Fatal(v ...interface{}) {
	Std.Output(levelFatal, 2, sout(v...))
	os.Exit(1)
}

// Fatalf is equivalent to Printf() followed by a call to os.Exit(1).
func Fatalf(format string, v ...interface{}) {
	Std.Output(levelFatal, 2, fmt.Sprintf(format, v...))
	os.Exit(1)
}

// Fatalln is equivalent to Println() followed by a call to os.Exit(1).
func Fatalln(v ...interface{}) {
	Std.Output(levelFatal, 2, fmt.Sprintln(v...))
	os.Exit(1)
}

// -----------------------------------------

// Panic is equivalent to Print() followed by a call to panic().
func Panic(v ...interface{}) {
	s := sout(v...)
	Std.Output(levelPanic, 2, s)
	panic(s)
}

// Panicf is equivalent to Printf() followed by a call to panic().
func Panicf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	Std.Output(levelPanic, 2, s)
	panic(s)
}

// Panicln is equivalent to Println() followed by a call to panic().
func Panicln(v ...interface{}) {
	s := fmt.Sprintln(v...)
	Std.Output(levelPanic, 2, s)
	panic(s)
}

// -----------------------------------------

func Stack(v ...interface{}) {
	s := sout(v...)
	s += "\n"
	buf := make([]byte, 1024 * 1024)
	n := runtime.Stack(buf, true)
	s += string(buf[:n])
	s += "\n"
	Std.Output(levelError, 2, s)
}

func PrintPanicStack() {
    if x := recover(); x != nil {
            Printf("%v", x)
            for i := 0; i < 10; i++ {
                    funcName, file, line, ok := runtime.Caller(i)
                    if ok {
                            Printf("frame %v:[func:%v,file:%v,line:%v]\n", i, runtime.FuncForPC(funcName).Name(), file, line)
                    }
            }
    }
}
