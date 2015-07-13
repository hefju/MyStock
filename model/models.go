package model
import (
    "github.com/go-xorm/core"
    "github.com/go-xorm/xorm"
    _ "github.com/mattn/go-sqlite3"
    "log"
    "fmt"
)

//报告类
type StatusReport struct  {
    Id int64
    From string      //发送人
    FromTime int64  //发送的时间
    Title string //标题(分类: 健康,统计的,)
    Content string //详细内容
}


func GetTestList()[]Stocks{//获取代号,目前只获取sh系列的,
    stocks:=make([]Stocks,0)
   // err:=engine.Where("id in (2266,2524)").Find(&stocks)
  //  err:=engine.Where("stype='sh' and status>-2").Find(&stocks)
    err:=engine.Where(" status>-2").Find(&stocks)
    if err!=nil{
        log.Println(err)
    }
    return stocks
}

func GetStockPrice(id int64)[]*StockPrice{
    sp:=make([]*StockPrice,0)
    engine.Where("id>7 and id<9000").Find(&sp)
//    sp:=new (StockPrice)
//    engine.Id(id).Get(sp)
    return sp
}

func UpdateStockPrice(list []*StockPrice){
    session:=engine.NewSession()
    defer session.Close()
    err:=session.Begin()
    if err!=nil{
        fmt.Println(err)
    }
    count:=0
    for _,item:=range list{
        _,err:=session.Id(item.Id).Update(item)
        if err !=nil{
            session.Rollback()
            fmt.Println("Update err:",err)
            return
        }
        count++
    }
    err=session.Commit()
    if err!=nil{
        return
    }
    fmt.Println("Update result:",count)
}

func UpdateStockUnuse(scodes []string){//将获取不到信息的票设置未-2
  //直接在数据库执行下面的语句
//    update stock set status=-2 where scode in(
//    select substr(scode,3,30) from stockprice where length(hq_str)=24
//    group by scode
//    )
}

func GetStockTestName()[]string{
    snames:=make([]string,0)
    stocks:=GetTestList()
    for _,v:=range stocks{
        item:=v.Stype+v.Scode
        snames=append(snames,item)
    }
    return snames
}

func InsertStockIndex(list []StockIndex) {
  afr, err := engine.Insert(&list)
    if err != nil {
        log.Println("Insert err:", err)
    }
    log.Println("Insert result:", afr)
}
//
func InsertStockPrice(list []StockPrice) {
    session:=engine.NewSession()
    defer session.Close()
    err:=session.Begin()
    if err!=nil{
        fmt.Println(err)
    }
    count:=0
    for _,item:=range list{
        _,err:=session.Insert(item)
        if err !=nil{
            session.Rollback()
            fmt.Println("Insert err:",err)
            return
        }
        count++
    }
    err=session.Commit()
    if err!=nil{
        return
    }
    fmt.Println("Insert result:",count)
}




//*********************************************//
var engine *xorm.Engine
func init() {
    var err error
    // engine, err = xorm.NewEngine("odbc", "driver={SQL Server};Server=192.168.1.200; Database=JXC; uid=sa; pwd=123;")
    //engine, err = xorm.NewEngine("odbc", "driver={SQL Server};server=.;database=charge;integrated security=SSPI;")
    engine, err = xorm.NewEngine("sqlite3", "./mystock.db")

    if err != nil {
        log.Fatalln("xorm create error", err)
    }
   // engine.ShowSQL = true
    engine.SetMapper(core.SameMapper{})
    // engine.CreateTables(new(tp_charge_billing))

  //  err = engine.Sync2(new(Stocks),new(StockIndex),new(StockPrice)) //, new(Group)) ,new(StockPrice)
    if err != nil {
        log.Fatalln("xorm sync error", err)
    }
}


//type Stock struct {
//    Id    int64
//    Sname string
//    Scode string
//    Stype string
//    Status int  //-2表示没有用的取不到信息的,-1表示未设置,0表示正常,
//}

//这个是最正确的名单数据
type Stocks struct {
    Id     int64
    Scode  string //纯数字
    Scode2 string //字母+数字
    Sname  string//中文名
    Stype  string//字母
    Status int //-2表示没有用的取不到信息的,-1表示未设置,0表示正常,
}

type StockPrice struct {
    Id     int64
    SName  string
    SCode  string
    J今日开盘价 float64
    Z昨日收盘价 float64
    D当前价格  float64
    J今日最高价 float64
    J今日最低价 float64
    J竞买价   float64
    J竞卖价   float64
    C成交股票数 float64
    C成交金额  float64
    M买一数   float64
    M买一价   float64
    M买二数   float64
    M买二价   float64
    M买三数   float64
    M买三价   float64
    M买四数   float64
    M买四价   float64
    M买五数   float64
    M买五价   float64
    M卖一数   float64
    M卖一价   float64
    M卖二数   float64
    M卖二价   float64
    M卖三数   float64
    M卖三价   float64
    M卖四数   float64
    M卖四价   float64
    M卖五数   float64
    M卖五价   float64

    R日期    string
    S时间    string
    Hq_str string //下载的字符串
    IsProcess int//是否处理,因为处理可能会出错, 程序崩溃,所以后期再处理
    CreatedAt  int64  `xorm:"created"`
    UpdatedAt  int64  `xorm:"updated"`
}

//大市 s_sh000001  s_sz399001
type StockIndex struct {
    Id     int64
    STime  string
    SCode  string
    SName  string
    D当前点数  float64
    Z涨跌价格  float64
    Z涨跌率   float64
    C成交量手  float64
    C成交额万元 float64
    CreatedAt  int64  `xorm:"created"`
    UpdatedAt  int64  `xorm:"updated"`
    //    CurrentPoint float64
    //    CurrentPrice float64
    //    SPercent float64
    //    TradingVolume float64
    //    Turnover float64
}
