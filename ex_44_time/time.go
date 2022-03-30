package main

import (
	"fmt"
	"time"
)

const (
	lay1 = "2006年01月02日 15时04分05秒"
	lay2 = "2006-01-02 15:04:05"
	lay3 = "2006-01-02 15:04:05.999999"
)

func main() {
	// Now() 获取的 Time 和当前系统的时区一样
	now := time.Now()
	fmt.Println(now)

	// Month 返回的是英文的
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())

	// Unix() 返回的是 1970 年到现在的时间戳
	fmt.Println(now.Unix())
	// 毫秒
	fmt.Println(now.UnixMilli())
	// 纳秒
	fmt.Println(now.UnixNano())

	// RFC 3339 时间格式
	fmt.Println(now.Format(time.RFC3339))

	// 自定义格式化时间
	dt1 := now.Format(lay1)
	dt2 := now.Format(lay2)
	dt3 := now.Format(lay3)
	fmt.Println(dt1)
	fmt.Println(dt2)
	fmt.Println(dt3)

	fmt.Println("---")

	// 返回 now 所在的时区和偏移
	zone, offset := now.Zone()
	now.Local().Zone()
	fmt.Println(zone)
	fmt.Println(offset)

	// 时区设置，当前 Now() 获取的是本地时间，所以结果是一样的
	// time.Local 会返回当前时区的常量
	fmt.Println(now.Format(lay1))
	cstZone := time.FixedZone("CST", 8*3600)
	now.In(cstZone).Format(lay1)

	fmt.Println("---")

	// 字符串解析
	// 当前时区是 CST +8
	// 注意：Parse() 解析是返回的是 UTC 时区的 Time，而 Format() 会自动转换为当前时区打印，所以打印结果是对的
	// 而 time.Unix() 的参数会以当前时区的时间戳构造 Time（实际就是把 UTC 时间戳当成本地的时间戳），所以再次打印结果是错的
	dtStr := "2017年04月07日 16时30分47秒"
	utcTime, _ := time.Parse(lay1, dtStr)
	utcUnix := time.Unix(utcTime.Unix(), 0)
	fmt.Println(utcTime.Format(lay2))
	fmt.Println(utcTime.Unix())
	fmt.Println(utcUnix.Format(lay2))

	// 正确的姿势是使用 ParseInLocation()
	cstTime, _ := time.ParseInLocation(lay1, dtStr, time.Local)
	cstUnix := time.Unix(cstTime.Unix(), 0)
	fmt.Println(cstTime.Format(lay2))
	fmt.Println(cstTime.Unix())
	fmt.Println(cstUnix.Format(lay2))

	fmt.Println("---")

	// 字符串与时间戳互转，1609527845 -> 2021年01月02日 03时04分05秒
	// 同样需要使用 ParseInLocation
	unixstamp := 1609527845
	uTime := time.Unix(int64(unixstamp), 0)
	fmt.Println(uTime.Format(lay1))

	uStr := uTime.Format(lay1)
	us, _ := time.ParseInLocation(lay1, uStr, time.Local)
	fmt.Println(us.Format(lay1))
	fmt.Println(us.Unix())
}
