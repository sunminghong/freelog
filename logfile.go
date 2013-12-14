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
    "os"
    "time"
    "strconv"
    "strings"
    "sync"
    "fmt"
)

type FileLogger struct {
    mwFile    *MutexWriter
    level     int
    filename  string
    //shortFile bool
    maxSize   int64
    nowSize   int64
    check     sync.Mutex

    fileFormat string
}

//文件操作结构体
type MutexWriter struct {
    sync.Mutex
    file *os.File
}

func init() {
    RegisterAdapter(AdapterFile, NewFileLogger)
}

func NewFileLogger() ILogger {
    flog := new(FileLogger)
    flog.level = levelInfo
    flog.filename = ""
    flog.mwFile = new(MutexWriter)
    return flog
}

func (this *MutexWriter) Write(msg []byte) (int, error) {
    this.Lock()
    defer this.Unlock()

    return this.file.Write(msg)
}

func (this *MutexWriter) SetFile(file *os.File) {
    if this.file != nil {
        this.file.Close()
    }

    this.file = file
}

func (this *FileLogger) GetLevel() (level int) {
    return this.level
}

func (this *FileLogger) Init(ini IConfigReader) (level int, err error) {

    logfile, err := ini.GetString(AdapterFile, "filename")
    if err != nil {
        return
    }
    this.filename = logfile

    if strings.Contains(logfile, "%") {
        ////f ="20060102-150405"
        //f := "2006010215"
        logfile = strings.Replace(logfile, "%y", "2006", -1)
        logfile = strings.Replace(logfile, "%m", "01", -1)
        logfile = strings.Replace(logfile, "%d", "02", -1)

        logfile = strings.Replace(logfile, "%h", "15", -1)
        logfile = strings.Replace(logfile, "%H", "15", -1)
        logfile = strings.Replace(logfile, "%M", "04", -1)

        this.fileFormat = logfile
        fmt.Println("///////////////////////",logfile)

        t := time.Now()
        this.filename = this.getFileName(&t)
    }

    //shortFile, err := ini.GetBool(AdapterFile, "shortFile")
    //if err == nil {
    //    this.shortFile = shortFile
    //}

    filesize, err := ini.GetFloat64(AdapterFile, "filesize")
    if err != nil {
        filesize = 100
    }
    this.maxSize = int64(filesize) * 1024 * 1024

    lev, err := ini.GetString(AdapterFile, "level")
    if err != nil {
        this.level = levelWarn
    } else {
        this.level, err = CheckLevel(lev)
    }

    err = this.initFile()
    return this.level, nil
}

func (this *FileLogger) Write(msg *LogMsg) (err error) {
    if this.level > msg.Level {
        return nil
    }

    msgsize := int64(len(msg.Msg))
    this.nowSize += msgsize

    this.checkLogFile(msg.Timestamp)

    this.mwFile.Write(msg.Msg)

    if 0 == this.nowSize {
        this.nowSize += msgsize
    }

    return nil
}

func (this *FileLogger) createLlogFile() (*os.File, error) {
    return os.OpenFile(
        this.filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0660)
}

func (this *FileLogger) initFile() error {
    file, err := this.createLlogFile()
    if err != nil {
        return err
    }

    stat, err := file.Stat()
    if err != nil {
        this.nowSize = 0
    } else {
        this.nowSize = stat.Size()
    }

    this.mwFile.SetFile(file)

    return nil
}

func (this *FileLogger) getFileName(t *time.Time) string {
    return t.Local().Format(this.fileFormat)
}

func (this *FileLogger) checkLogFile(t *time.Time) {
    reInit := false
    if this.nowSize >= this.maxSize {
        reInit = true
        this.check.Lock()
        this.Close()

        for i := 0; ; i++ {
            newname := this.filename + strconv.Itoa(i)
            err := os.Rename(this.filename, newname)
            if err != nil {
                continue
            } else {
                break
            }
        }
    }

    if this.fileFormat != "" {
        newfile := this.getFileName(t)
        if newfile == this.fileFormat {
            return
        }
        if !reInit {
            reInit = true

            this.check.Lock()
            this.Close()
        }
        this.filename = newfile
    }

    if reInit {
        this.initFile()
        this.check.Unlock()
    }
}

func (this *FileLogger) Close() {
    this.mwFile.file.Close()
}
