package config
import (
    "github.com/Unknwon/goconfig"
//    "github.com/donnie4w/go-logger/logger"
    "github.com/hefju/MyStock/app/config/myconfig"
    "fmt"
    "log"
)
var (
    AppName,AppRootPath,ReportAddr string //当前程序的名称,发送报告的地址
)
func init(){
    fmt.Println("appConf.go init")
    // 创建并获取一个 ConfigFile 对象，以进行后续操作
    // 文件名支持相对和绝对路径
    cfg, err := goconfig.LoadConfigFile("conf.ini")
    if err != nil {
        log.Fatal("无法加载配置文件：%s", err)
    }
    // 加载完成后所有数据均已存入内存，任何对文件的修改操作都不会影响到已经获取到的对象
    // 对默认分区进行普通读取操作
    section,key:="System", "AppName"
    value, err := cfg.GetValue(section, key)
    if err != nil {
        msg:=fmt.Sprintf("无法获取键值（%s-%s）：%s",section,key, err)
//        logger.Fatal(msg)
        panic(msg)//如果设置不正确, 就不能启动了
    }
    AppName=value

    section,key="System", "AppRootPath"
    value, err = cfg.GetValue(section, key)
    if err != nil {
        msg:=fmt.Sprintf("无法获取键值（%s-%s）：%s",section,key, err)
        panic(msg)
    }
    AppRootPath=value

    ReportAddr=myconfig.ReportAddr
   // log.Printf("%s > %s: %s", "System", "SystemVersion", value)
}

