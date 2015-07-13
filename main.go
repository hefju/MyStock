package main
import (
    "github.com/hefju/MyStock/stock"
//    "github.com/hefju/MyStock/app"
//    "github.com/hefju/GoTools"
    "fmt"
)
func main(){
//    app.Run()
   // fmt.Println(model.GetStockTestName())// model.GetTestList())

//    stock :=& task.StockDownload{}
//    stock.DoWork(t1)
    stock.GetMainIndex()
        stock.Download()
    //  stock.Processing()//7-9000
//    GoTools.Hello()

//    stock.Mytest()
//    t1:=time.Now()
//    fmt.Println("run time:",time.Now().Sub(t1))

    fmt.Println("end")
}

