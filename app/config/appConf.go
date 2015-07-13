package config

import (
	"github.com/Unknwon/goconfig"
	//    "github.com/donnie4w/go-logger/logger"
	"fmt"
	"github.com/hefju/MyStock/app/config/myconfig"
	"log"
)

var (
	AppName     string //程序名
	AppRootPath string //程序的运行目录,如果win service就需要指定目录来存放文件
	ReportAddr  string //当前程序的名称,发送报告的地址
	BaseAddr    string //数据下载的基地址

    EmailHOST       string
    EmailSERVER_ADDR string
    EmailUSER      string
    EmailPASSWORD   string
EmailReceiver string
)

func init() {
	// fmt.Println("appConf.go init")
	// 创建并获取一个 ConfigFile 对象，以进行后续操作
	cfg, err := goconfig.LoadConfigFile("conf.ini")
	if err != nil {
		log.Fatal("无法加载配置文件：%s", err)
	}
	// 加载完成后所有数据均已存入内存，任何对文件的修改操作都不会影响到已经获取到的对象
	section, key := "System", "AppName"
	value, err := cfg.GetValue(section, key)
	if err != nil {
		msg := fmt.Sprintf("无法获取键值（%s-%s）：%s", section, key, err)
		panic(msg) //如果设置不正确, 就不能启动了
	}
	AppName = value

	section, key = "System", "AppRootPath"
	value, err = cfg.GetValue(section, key)
	if err != nil {
		msg := fmt.Sprintf("无法获取键值（%s-%s）：%s", section, key, err)
		panic(msg)
	}
	AppRootPath = value

    section, key = "Email", "Receiver"
    value, err = cfg.GetValue(section, key)
    if err != nil {
        msg := fmt.Sprintf("无法获取键值（%s-%s）：%s", section, key, err)
        panic(msg)
    }
    EmailReceiver = value//邮件接收人

	ReportAddr = myconfig.ReportAddr//状态报告地址
    BaseAddr=myconfig.BaseAddr      //数据下载地址

    EmailHOST=myconfig.EmailHOST
    EmailSERVER_ADDR =myconfig.EmailSERVER_ADDR
    EmailUSER    =myconfig.EmailUSER
    EmailPASSWORD  =myconfig.EmailPASSWORD
}
