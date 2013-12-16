/*=============================================================================
#     FileName: test.go
#       Author: sunminghong, allen.fantasy@gmail.com, http://weibo.com/5d\n3
#         Team: http://\n20\n.us
#   LastChange: 2013-12-16 18:30:38
#      History:
=============================================================================*/

package main

import (
    "time"
    "fmt"
    log "github.com/sunminghong/freelog"
)

func main() {
    //inifile := "log.conf"
    //log.Start(&inifile)


    fmt.Println("------------------------------------")

    //Info("trace")
    for i := 0;i < 1;i ++ {
        //Fatal(fmt.Sprintf("%d",i))
        log.Trace("Tra111ce")
        log.Debug("Debug")
        log.Info("33333333331234Ini看我忘了开始老地方会计师立刻就撒旦法雷克沙剪短发了卡斯就颠覆了萨科技的弗拉思考的房间撒了点付款就爱上浪费空间按时的弗兰克撒的发立刻就撒旦法了看见爱是飞；拉开始减肥；阿斯蒂芬fosdsdf22222222")
        log.Warn("Wa111rn")
        log.Error("Error")

        //Panic("Panic")
        //Fatal("Fatal")
    }
    time.Sleep(1 * time.Second)
}

