package main

import (
    "encoding/base64"
    "fmt"
    "os"
    "time"
)

const (
    Duration_M3S = 1000
    Duration_MS = 1000000 // Milli
    Duration_S = 1000000000 // Second
)


const (
    DATE_FMT = "2006-01-02"
    DATE_FMT_NO_X = "20060102"
    DATETIME_FMT = "2006-01-02 15:04:05"
)

var loc *time.Location
func init(){
    loc, _ = time.LoadLocation("Asia/shanghai")
}

func basic(){
    var t time.Time = time.Now()
    fmt.Println(t.Unix())
    fmt.Println(t.UnixMilli())
    fmt.Println(t.UnixMicro())
    fmt.Println(t.UnixNano())
    time.Sleep(3 * time.Second)

    useTime := time.Since(t)
    fmt.Println(useTime.Seconds())
}

func time_fmt(){
    // 时间转字符串
    _data_time := time.Now()
    fmt.Println(_data_time)
    fmt.Println(_data_time.Format(DATETIME_FMT))
    fmt.Println(_data_time.Format(DATE_FMT))
    fmt.Println(_data_time.Format(DATE_FMT_NO_X))

    // 字符串转时间。
    t, err := time.ParseInLocation(DATETIME_FMT, time.Now().Format(DATETIME_FMT), loc)
    if err == nil{
        fmt.Println(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
    }else{
        fmt.Println(err.Error())
    }

}


func base64_t(){
    fin,_ := os.Open("/Users/yixiang/Desktop/liushui1.jpeg")
    defer fin.Close()
    bs := make([]byte, 1 << 21)
    n, _ := fin.Read(bs)
    str := base64.StdEncoding.EncodeToString(bs[:n])
    fmt.Println("一共读取多少：", n)

    bbb, _ := base64.StdEncoding.DecodeString(str) // 解码

    fin2, _ := os.OpenFile("liushui_p.jpeg", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
    fin2.Write(bbb)
    defer fin2.Close()
}



func main(){
    base64_t()
    //basic()
    //time_fmt()
}
