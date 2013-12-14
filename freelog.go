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
    "strings"
)

var logger ILog

func init() {
	registeredAdapters = make(map[string]GetAdapter)

    //file := "./logconfig.ini"
	//Reload(&file)
}

func Start(inifile *string) {

	logger = &baseLog{}

	reader := &iniReader{}
	if err := reader.Init(inifile); err != nil {
		panic(fmt.Sprintf("logger config reader init error: %q", err))
	}

	logger.Init(1000, reader)
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
