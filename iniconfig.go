/*=============================================================================
#     FileName: iniconfig.go
#       Author: sunminghong, allen.fantasy@gmail.com, http://weibo.com/5d13
#         Team: http://1201.us
#   LastChange: 2013-12-14 11:57:05
#      History:
=============================================================================*/

/*
一个通用的、多输出、多级别的go 日志库，并且每个输出可以分别定义级别；可以方便自定义输出类
*/
package freelog

import (
    iniconf "github.com/sunminghong/iniconfig"
)

type iniReader struct {
    c *iniconf.ConfigFile
}

func (ini *iniReader) Init(inifile *string) error {
    c, err := iniconf.ReadConfigFile(*inifile)
    if err != nil {
        return err
    }
    ini.c = c
    return nil
}

func (ini *iniReader) InitBytes(inibytes []byte) error {
    c, err := iniconf.ReadConfigBytes(inibytes)
    if err != nil {
        return err
    }
    ini.c = c
    return nil
}

func (ini *iniReader) GetAdapters() (adapterNames []string) {
    adapterNames = ini.c.GetSections()
    return
}

func (ini *iniReader) GetBool(adapterName string, option string) (value bool, err error) {
    return ini.c.GetBool(adapterName, option)
}

func (ini *iniReader) GetFloat64(adapterName string, option string) (value float64, err error) {
    return ini.c.GetFloat64(adapterName, option)
}

func (ini *iniReader) GetString(adapterName string, option string) (value string, err error) {
    return ini.c.GetString(adapterName, option)
}
