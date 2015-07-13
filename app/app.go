package app
import (
//    "fmt"
    "os"
    "github.com/donnie4w/go-logger/logger"
    "github.com/hefju/MyStock/app/config"
    "github.com/hefju/MyStock/task"
    "time"
)
//初始化日志
func initLogger(){
  //  fmt.Println("app-initLogger")
    os.MkdirAll(config.AppRootPath+"/log", 0777)
    logger.SetRollingDaily(config.AppRootPath+"/log", "test.log") //如果没有log文件夹, 需要新增文件夹
    logger.SetLevel(logger.DEBUG)
}
//运行程序
func Run(){
    //1.初始化config, 因为Logger的初始化目录要从config读取
    initLogger()//2.初始化日志

    //获取任务队列
    list:=getTaskList()

    //定时触发任务
    ticker:=time.NewTicker(time.Second*3)//time.Minute*10)
    for t:=range ticker.C{
        for _,item:=range list  {
            item.DoWork(t)
        }
    }

}

func getTaskList()[]task.Tasker{
    stockDownload:=&task.StockDownload{}//下载任务
    systemReport :=&task.SystemReport{}//状态报告任务
    list:=make([]task.Tasker,0)
    list=append(list,stockDownload,systemReport)
    return list
}
