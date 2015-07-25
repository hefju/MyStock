package task
import (
    "time"
    "github.com/hefju/GoTools/net/mail"
    "fmt"
    "github.com/hefju/MyStock/app/config"
    "github.com/hefju/MyStock/model"
    "encoding/json"
    "net/http"
    "bytes"
    "github.com/donnie4w/go-logger/logger"
    "io/ioutil"
    "github.com/hefju/MyStock/stock"
    "github.com/hefju/GoTools/DateTime"
)
type Tasker interface {   //任务接口
    DoWork(t time.Time)   //执行任务,  注意复杂的工作一定要开一个goroutine来执行, 否则影响到ticker正常的循环
}
//数据下载程序
type StockDownload struct{
    LasttimeExec time.Time //上次执行的时间, 有些任务一天只需要执行一次,或者一个小时执行一次.通过这个时间来判断是否需要执行
    WorkStatus int//0表示未开始, 1表示正在运行.
}
//发送当前计算机的报告
type SystemReport struct{

}


//工作,执行任务.  这个需要考虑在ticker间隔内, 我的任务超时未执行完,会不会再启动多一次任务.
//例如我是下载任务, 网络不好了, 本来10秒完成的事情变成了1小时才完成, ticker间隔可能重复执行了这个下载任务.
//如果下载任务失败了, 那么任务会不会重新启动任务? 或者进行报告?
func (stockDownload *StockDownload) DoWork(t time.Time){

    if !stockDownload.checkWorkStatus(t){ //检测执行状态,是否需要执行 //每天执行一次
        return
    }
    t1:=time.Now()//计算运行时间

    if stockDownload.WorkStatus==1{//表示任务正在执行
        return
    }else {//WorkStatus=0
        stockDownload.WorkStatus=1//如果任务还未执行,就设置为执行状态
    }
    //do something
    done:=make(chan int)//完成标志
    go   downString(done)
    <-done

    stockDownload.WorkStatus=0//任务执行完毕,任务状态变为未执行
    //计算执行任务花费的时间
    ElapsedTime:="Elapsed Time:"+fmt.Sprintf("%s", time.Now().Sub(t1))

    //操作完毕,设置上次执行时间
    stockDownload.LasttimeExec=time.Now()

    //发送执行报表
    title:="StockDownload:"+time.Now().Format("2006.1.2 15:04:05")
    content:=ElapsedTime+""
    stockDownload.sendReport(title,content)
}
func downString(done chan int){//执行任务
    stock.GetMainIndex()
    stock.DownloadSlowly()
    done<-1
}
//检测执行状态, 如果一天内执行过就不要再执行了.(当然,你可以设置一小时执行一次, 或者1分钟执行一次)
func (stockDownload *StockDownload) checkWorkStatus(t time.Time) bool {
    //星期一至星期五, 并且是下午3点,
    if !DateTime.IsWorkday(t)|| t.Hour()!=15{
        return false
    }

    // 检查任务是否正在行过, 如果定时器ticker跟触发的操作处于相同的goroutine,那么定时器的间隔会被延长
     last:=stockDownload.LasttimeExec
    if last.Year() == t.Year() &&    last.Month() == t.Month() &&    last.Day() == t.Day() {
        return false
    } else {
        return true
    }
}
//发送执行报告
func (stockDownload *StockDownload) sendReport(title,content string){
    //如果不发送email, 就发送到消息中心,消息中心根据优先级别再发送email
  go  mail.SendEmail(title,content)
//    mail.InitEmail("","","","","")
//    fmt.Println( "mail.SendEmail ",title," ",content)
}

//发送当前计算机状体
func (systemReport *SystemReport) DoWork(t time.Time) {
    //发送报告不需要检测是否需要执行, ???这样好吗? 如果定时器太快, 这个方法需要自己的发送间隔啊.
//    fmt.Println("SystemReport")
//    return
    url :=config.ReportAddr
    report:=model.StatusReport{From:config.AppName,FromTime:time.Now().Unix(),Title:"状态报告",Content:"I'm still alive"}
    jsonStr, _ := json.Marshal(report)

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
    req.Header.Set("X-Custom-Header", "myvalue")
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        //panic(err)//发送数据失败, 程序不要死啊. panic就完啦
        defer func() {
            if r := recover(); r != nil { //r的信息有什么用?还不如直接输出err
                logger.Error("状态报告失败,", err)//  fmt.Println("发送失败,", r, err)
                //需要发送到消息中心.
                // jutool.SendEmail("状态报告失败","报告时间:"+t.Format("2006-01-02 15:04:05")+" "+err.Error())//上传失败也发个email通知我
            }
        }()
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    logger.Info("状态报成功,",t.Format("2006-01-02 15:04:05"), string(body))
}
