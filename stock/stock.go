package stock
import (
"github.com/hefju/MyStock/model"
    "github.com/hefju/MyStock/app"
   // "fmt"
    "net/http"
    "io/ioutil"
    "strings"
    "strconv"
    "time"
)
var BaseAddr="http://hq.sinajs.cn/list="//s_sh000001
func Download(){
   // fmt.Println(model.GetStockTestName())// model.GetTestList()
    getstocks(model.GetStockTestName())
}

func Mytest(){
    stocks:=make([]model.StockPrice,0)
    for i:=0 ;i<1300 ;i++  {
        var s model.StockPrice
        s.SCode =strconv.Itoa( i)
        stocks=append(stocks,s)
    }
  //  model.Insert4(stocks)
    model.Insert2(stocks)
}

func getstocks(names []string){
    stocks:=make([]model.StockPrice,0)
    for k,v:=range names{
        info:= httpget(BaseAddr+v)
        //list := strings.Split(info, ",")
        var s model.StockPrice
        s.SCode = v
//        s.J今日开盘价 = getfloat64(list[1])
//        s.Z昨日收盘价 = getfloat64(list[2])
//        s.D当前价格 = getfloat64(list[3])
//        s.J今日最高价 = getfloat64(list[4])
//        s.J今日最低价 = getfloat64(list[5])
//        s.J竞买价 = getfloat64(list[6])
//        s.J竞卖价 = getfloat64(list[7])
//        s.C成交股票数 = getfloat64(list[8])
//        s.C成交金额 = getfloat64(list[9])
//        s.M买一数 = getfloat64(list[10])
//        s.M买一价 = getfloat64(list[11])
//        s.M买二数 = getfloat64(list[12])
//        s.M买二价 = getfloat64(list[13])
//        s.M买三数 = getfloat64(list[14])
//        s.M买三价 = getfloat64(list[15])
//        s.M买四数 = getfloat64(list[16])
//        s.M买四价 = getfloat64(list[17])
//        s.M买五数 = getfloat64(list[18])
//        s.M买五价 = getfloat64(list[19])
//        s.M卖一数 = getfloat64(list[20])
//        s.M卖一价 = getfloat64(list[21])
//        s.M卖二数 = getfloat64(list[22])
//        s.M卖二价 = getfloat64(list[23])
//        s.M卖三数 = getfloat64(list[24])
//        s.M卖三价 = getfloat64(list[25])
//        s.M卖四数 = getfloat64(list[26])
//        s.M卖四价 = getfloat64(list[27])
//        s.M卖五数 = getfloat64(list[28])
//        s.M卖五价 = getfloat64(list[29])
//        s.R日期 = list[30]
//        s.S时间 = list[31]
        s.Hq_str=info //下载的内容
        s.IsProcess=-1//是否已经处理
        stocks=append(stocks,s)
        if k==-1{//-1表示不限制, 0表示只执行一条,1表示只执行两条
            break
        }
    }
    model.Insert4(stocks)
   // fmt.Println(stocks)
}

//获取上证指数和深圳成指
func GetMainIndex(){
  //  maincode:=make([]string,0)
    maincode:=[]string{"s_sh000001","s_sz399001"}
    maincode[0]=  BaseAddr+  maincode[0]
    maincode[1]=  BaseAddr+  maincode[1]

GetMainIndexString(maincode)
//    fmt.Println(maincode)
}
func GetMainIndexString(urls []string) {
     mainindex:=make([]model.StockIndex,0)
    for k,v:=range urls{
       info:= httpget(v)
       // strings.Repeat(info,`";`)
        info=info[0:len(info)-3]
        list := strings.Split(info, ",")
        var i model.StockIndex
//        fmt.Println("k:",k," ",list)
//        fmt.Println("0:",list[0])
//        fmt.Println("1:",list[1])
//        fmt.Println("2:",list[2])
//        fmt.Println("3:",list[3])
//        fmt.Println("4:",list[4])
//        fmt.Println("5:",list[5])
        i.SCode=v[len(v)-10:]
        i.SName=GetIndexName( i.SCode)
        i.STime=time.Now().Format("2006-01-02")

                i.D当前点数=getfloat64( list[1])
                i.D当前价格=getfloat64( list[2])
                i.Z涨跌率=getfloat64( list[3])
                i.C成交量手=getfloat64( list[4])
                i.C成交额万元=getfloat64( list[5])
        mainindex=append(mainindex,i)
        if k==10{
            break
        }
    }
    //InsertMainindex()
    model.Insert3(mainindex)
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
        app.Log.Fatal("getfloat64:",str,"-",err)
        return 0
    }
    return f
}

//网络请求
func httpget(url string)string{
    resp,err:=http.Get(url)
    if err!=nil{
       app.Log.Println("httpget:",err)
    }
    defer resp.Body.Close()
    body,err:=ioutil.ReadAll(resp.Body)
    if err!=nil{
        app.Log.Println("httpget-ioutil.ReadAll:",err)
    }
    bodystr := string(body)
    return bodystr
}