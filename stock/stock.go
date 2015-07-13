package stock
import (
"github.com/hefju/MyStock/model"
    "github.com/hefju/MyStock/app/config"
//    "github.com/hefju/MyStock/app"
   // "fmt"
    "net/http"
    "io/ioutil"
    "strings"
    "strconv"
    "time"
    "fmt"
    "github.com/guotie/gogb2312"
    "github.com/donnie4w/go-logger/logger"
//    "log"
)



func Mytest(){
    stockprices:=model.GetStockPrice(59)
    fmt.Println(stockprices)
    return
    for _,v :=range stockprices{
//        fmt.Println("序号:",k)
        stockprice:=v
        if len(stockprice.Hq_str)<25{
            fmt.Println("Hq_str no value")
            return
        }
        list := strings.Split(stockprice.Hq_str, ",")
        if  len(list)<31{// s.S时间 = list[31]
            fmt.Println("list解释错误.")
            break
        }

        var s *model.StockPrice=stockprice
        s.J今日开盘价 = getfloat64(list[1])
        s.Z昨日收盘价 = getfloat64(list[2])
        s.D当前价格 = getfloat64(list[3])
        s.J今日最高价 = getfloat64(list[4])
        s.J今日最低价 = getfloat64(list[5])
        s.J竞买价 = getfloat64(list[6])
        s.J竞卖价 = getfloat64(list[7])
        s.C成交股票数 = getfloat64(list[8])
        s.C成交金额 = getfloat64(list[9])
        s.M买一数 = getfloat64(list[10])
        s.M买一价 = getfloat64(list[11])
        s.M买二数 = getfloat64(list[12])
        s.M买二价 = getfloat64(list[13])
        s.M买三数 = getfloat64(list[14])
        s.M买三价 = getfloat64(list[15])
        s.M买四数 = getfloat64(list[16])
        s.M买四价 = getfloat64(list[17])
        s.M买五数 = getfloat64(list[18])
        s.M买五价 = getfloat64(list[19])
        s.M卖一数 = getfloat64(list[20])
        s.M卖一价 = getfloat64(list[21])
        s.M卖二数 = getfloat64(list[22])
        s.M卖二价 = getfloat64(list[23])
        s.M卖三数 = getfloat64(list[24])
        s.M卖三价 = getfloat64(list[25])
        s.M卖四数 = getfloat64(list[26])
        s.M卖四价 = getfloat64(list[27])
        s.M卖五数 = getfloat64(list[28])
        s.M卖五价 = getfloat64(list[29])
        s.R日期 = list[30]
        s.S时间 = list[31]
        s.IsProcess=1//是否已经处理
//       fmt.Println(s)
    }
    //fmt.Println(stockprices)
   model.UpdateStockPrice(stockprices)
    }

//处理数据
func Processing(){
    stockprices:=model.GetStockPrice(0)
    for _,v :=range stockprices{
        //fmt.Println("序号:",k)
        stockprice:=v
        if len(stockprice.Hq_str)<25{//表示没有数据
            fmt.Println("Hq_str no value")
            return
        }
        list := strings.Split(stockprice.Hq_str, ",")
        if  len(list)<31{// s.S时间 = list[31]
            fmt.Println("list解析错误.")
            break
        }

        var s *model.StockPrice=stockprice
        s.J今日开盘价 = getfloat64(list[1])
        s.Z昨日收盘价 = getfloat64(list[2])
        s.D当前价格 = getfloat64(list[3])
        s.J今日最高价 = getfloat64(list[4])
        s.J今日最低价 = getfloat64(list[5])
        s.J竞买价 = getfloat64(list[6])
        s.J竞卖价 = getfloat64(list[7])
        s.C成交股票数 = getfloat64(list[8])
        s.C成交金额 = getfloat64(list[9])
        s.M买一数 = getfloat64(list[10])
        s.M买一价 = getfloat64(list[11])
        s.M买二数 = getfloat64(list[12])
        s.M买二价 = getfloat64(list[13])
        s.M买三数 = getfloat64(list[14])
        s.M买三价 = getfloat64(list[15])
        s.M买四数 = getfloat64(list[16])
        s.M买四价 = getfloat64(list[17])
        s.M买五数 = getfloat64(list[18])
        s.M买五价 = getfloat64(list[19])
        s.M卖一数 = getfloat64(list[20])
        s.M卖一价 = getfloat64(list[21])
        s.M卖二数 = getfloat64(list[22])
        s.M卖二价 = getfloat64(list[23])
        s.M卖三数 = getfloat64(list[24])
        s.M卖三价 = getfloat64(list[25])
        s.M卖四数 = getfloat64(list[26])
        s.M卖四价 = getfloat64(list[27])
        s.M卖五数 = getfloat64(list[28])
        s.M卖五价 = getfloat64(list[29])
        s.R日期 = list[30]
        s.S时间 = list[31]
        s.IsProcess=1//是否已经处理
    }
    //fmt.Println(stockprices)
    model.UpdateStockPrice(stockprices)//更新到数据库
}

//下载数据.
func Download(){
    // fmt.Println(model.GetStockTestName())// model.GetTestList()
    getstocks(model.GetStockTestName())
    fmt.Println("Download end!")
}
//根据代号来获取数据, 因为有些代号是没有用的, 所以传入参数前就处理掉.
func getstocks(names []string){
    stocks:=make([]model.StockPrice,0)//数据, 用来插入到数据库的, 由于处理可能导致程序出错, 所以先保存到数据库,以后再处理.
    for k,scode:=range names{
        info:= httpget(config.BaseAddr+scode)
        var s model.StockPrice
        s.SCode = scode
        s.Hq_str=info //下载的内容
        s.IsProcess=-1//是否已经处理
        stocks=append(stocks,s)
        if k==-1{//测试用的. -1表示不限制, 0表示只执行一条,1表示只执行两条
            break
        }
    }
    model.InsertStockPrice(stocks)//xorm如何插入必须要有struct类型.
}
//发送请求的时候, 不能太频繁, 否则会本机socket缓冲会满,远程会拒绝访问
func DownloadSlowly() {
   stocks:=make([]model.StockPrice,0)//数据, 用来插入到数据库的, 由于处理可能导致程序出错, 所以先保存到数据库,以后再处理.

    codes:=model.GetStockTestName()
    arrayCount:=len(codes)
    index:=0
    over:=false
    for {
        max:=index+10
        if(max>arrayCount){
            max=arrayCount
            over=true
        }
        list:=codes[index:max]
        for _,scode:=range list{
           // log.Println(scode)

            info:= httpget(config.BaseAddr+scode)
            var s model.StockPrice
            s.SCode = scode
            s.Hq_str=info //下载的内容
            s.IsProcess=-1//是否已经处理,-1未处理,1已处理
            stocks=append(stocks,s)
        }
        if over{
            break
        }
//        log.Println("")
        index=max
        time.Sleep(time.Millisecond*100)
    }
   model.InsertStockPrice(stocks)//xorm如何插入必须要有struct类型.
}

//获取上证指数和深圳成指
func GetMainIndex(){
    //  maincode:=make([]string,0)
    maincode:=[]string{"s_sh000001","s_sz399001"}
    maincode[0]=  config.BaseAddr+  maincode[0]
    maincode[1]=  config.BaseAddr+  maincode[1]

    GetMainIndexString(maincode)//获取,分析,插入数据
    fmt.Println("GetMainIndex end!")
}
func GetMainIndexString(urls []string) {
     mainindex:=make([]model.StockIndex,0)
    for k,v:=range urls{
       info:= httpget(v)
       // strings.Repeat(info,`";`)
        info=info[0:len(info)-3]
        list := strings.Split(info, ",")
        var i model.StockIndex

        i.SCode=v[len(v)-10:]
        i.SName=GetIndexName( i.SCode)
        i.STime=time.Now().Format("2006-01-02")//这是获取数据的日期,未必是数据发生的日期.例如星期六日和节假日

        i.D当前点数=getfloat64( list[1])
        i.Z涨跌价格=getfloat64( list[2])
        i.Z涨跌率=getfloat64( list[3])
        i.C成交量手=getfloat64( list[4])
        i.C成交额万元=getfloat64( list[5])
        mainindex=append(mainindex,i)
        if k==10{//测试用的, 不用理
            break
        }
    }
    model.InsertStockIndex(mainindex)//由于xorm需要反射获取struct信息, 所以传入参数不能是interface{}
   // fmt.Println(mainindex)
}
//根据代码设置指数名称
func GetIndexName(name string)string{
    sname:=""
    switch name {
        case "s_sh000001":
    sname="上证指数"
        case "s_sz399001":
        sname="深证成指"
        default:
        sname="unknow"
    }
    return sname
}
//数据类型转换
func getfloat64(str string)float64{
    f,err:=strconv.ParseFloat(str,64)
    if err!=nil{
        //app.Log.Fatal("getfloat64:",str,"-",err)
        logger.Fatal("getfloat64:",str,"-",err)
        return 0
    }
    return f
}

//网络请求
func httpget(url string)string{
    resp,err:=http.Get(url)
    if err!=nil{
      // app.Log.Println("httpget:",err)
        logger.Error("httpget:",err)
    }
    defer resp.Body.Close()
    body,err:=ioutil.ReadAll(resp.Body)
    if err!=nil{
        //app.Log.Println("httpget-ioutil.ReadAll:",err)
        logger.Error("httpget-ioutil.ReadAll:",err)
    }
    bodystr := string(handleGb2312(body))//目标网站使用的是gb2312的编码,有点麻烦
    return bodystr
}
func handleGb2312(body []byte)[]byte{
    output, err, _, _ := gogb2312.ConvertGB2312(body)
    if err != nil {
        fmt.Println(err)
    }
    return output
}