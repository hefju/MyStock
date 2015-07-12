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
)
type Tasker interface {
    DoWork(t time.Time)   //执行任务
}
//数据下载程序
type StockDownload struct{
    LasttimeExec time.Time //上次执行的时间, 有些任务一天只需要执行一次,或者一个小时执行一次.通过这个时间来判断是否需要执行
}
//发送当前计算机的报告
type SystemReport struct{

}


//工作,执行任务
func (stockDownload *StockDownload) DoWork(t time.Time){
    fmt.Println("StockDownload")
    return
    if !stockDownload.checkWorkStatus(t){ //检测执行状态,是否需要执行
        return
    }
    t1:=time.Now()

    //do something
   //\ config.ReportAddr
   // fmt.Println(config.AppName)

    ElapsedTime:="Elapsed Time:"+fmt.Sprintf("%s", time.Now().Sub(t1))//计算执行任务花费的时间

    //操作完毕,设置上次执行时间
    stockDownload.LasttimeExec=time.Now()

    //发送执行报表
    title:="StockDownload:"+time.Now().Format("2006.1.2 15:04:05")
    content:=ElapsedTime+""
    stockDownload.sendReport(title,content)
}
//检测执行状态, 如果一天内执行过就不要再执行了.(当然,你可以设置一小时执行一次, 或者1分钟执行一次)
func (stockDownload *StockDownload) checkWorkStatus(t time.Time) bool {
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
  //go  mail.SendEmail(title,content)
    mail.InitEmail("","","","","")
    fmt.Println( "mail.SendEmail ",title," ",content)
}

//发送当前计算机状体
func (systemReport *SystemReport) DoWork(t time.Time) {
    fmt.Println("SystemReport")
    return
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
