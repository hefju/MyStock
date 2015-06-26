package model
import (
    "github.com/go-xorm/core"
    "github.com/go-xorm/xorm"
    _ "github.com/mattn/go-sqlite3"
    "log"
    "fmt"
)

func GetTestList()[]Stock{
    stocks:=make([]Stock,0)
   // err:=engine.Where("id in (2266,2524)").Find(&stocks)
    err:=engine.Where("stype='sh'").Find(&stocks)
    if err!=nil{
        log.Println(err)
    }
    return stocks
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
func Insert(list []Stock) {
    afr, err := engine.Insert(&list)
    if err != nil {
        log.Println("Insert err:", err)
    }
    log.Println("Insert result:", afr)
}
func Insert3(list []StockIndex) {
  afr, err := engine.Insert(&list)
    if err != nil {
        log.Println("Insert err:", err)
    }
    log.Println("Insert result:", afr)
}
func Insert4(list []StockPrice) {
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
//    afr, err := engine.Insert(&list)
//    if err != nil {
//        log.Println("Insert err:", err)
//    }
//    log.Println("Insert result:", afr)
//    for _,v:=range list{
//        InsertOne(v)
//    }
}
func Insert2(list interface{}) {
    afr, err := engine.Insert(&list)
    if err != nil {
        log.Println("Insert err:", err)
    }
    log.Println("Insert result:", afr)
}
func InsertOne(obj StockPrice){
//    afr, err := engine.Insert(&obj)
    afr, err := engine.InsertOne(&obj)
    if err != nil {
        log.Println("Insert err:", err)
    }
    log.Println("Insert result:", afr)
}

//func Insert5(list []Price) {
//    afr, err := engine.Insert(&list)
//    if err != nil {
//        log.Println("Insert err:", err)
//    }
//    log.Println("Insert result:", afr)
//}
//func Insert6(list []User) {
//    afr, err := engine.Insert(&list)
//    if err != nil {
//        log.Println("Insert err:", err)
//    }
//    log.Println("Insert result:", afr)
//}


type Stock struct {
    Id    int64
    Sname string
    Scode string
    Stype string
    Status int
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
// s_sh000001
// s_sz399001
type StockIndex struct {
    Id     int64
    STime  string
    SCode  string
    SName  string
//    CurrentPoint float64
//    CurrentPrice float64
//    SPercent float64
//    TradingVolume float64
//    Turnover float64
    D当前点数  float64
    D当前价格  float64
    Z涨跌率   float64
    C成交量手  float64
    C成交额万元 float64
    CreatedAt  int64  `xorm:"created"`
    UpdatedAt  int64  `xorm:"updated"`
}

var engine *xorm.Engine

func init() {
    var err error
    // engine, err = xorm.NewEngine("odbc", "driver={SQL Server};Server=192.168.1.200; Database=JXC; uid=sa; pwd=123;")
    //engine, err = xorm.NewEngine("odbc", "driver={SQL Server};server=.;database=charge;integrated security=SSPI;")
    engine, err = xorm.NewEngine("sqlite3", "./mystock.db")

    if err != nil {
        log.Fatalln("xorm create error", err)
    }
    //engine.ShowSQL = true
    engine.SetMapper(core.SameMapper{})
    // engine.CreateTables(new(tp_charge_billing))

    err = engine.Sync2(new(Stock),new(StockIndex),new(StockPrice)) //, new(Group)) ,new(StockPrice)
    if err != nil {
        log.Fatalln("xorm sync error", err)
    }
}

