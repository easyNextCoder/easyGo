package main

import (
	"fmt"
	"github.com/robfig/cron"
	"time"
)

/*
Field name   | Mandatory? | Allowed values  | Allowed special characters
----------   | ---------- | --------------  | --------------------------
Seconds      | Yes        | 0-59            | * / , -
Minutes      | Yes        | 0-59            | * / , -
Hours        | Yes        | 0-23            | * / , -
Day of month | Yes        | 1-31            | * / , - ?
Month        | Yes        | 1-12 or JAN-DEC | * / , -
Day of week  | Yes        | 0-6 or SUN-SAT  | * / , - ?

https://pkg.go.dev/github.com/robfig/cron
*/

var cronTab = cron.New()

var specs = []string{
	"* * * * * *",      // 每一秒执行一次
	"5-40/8 * * * * *", // 第5秒-第40秒区间，每隔8秒种打印一次
	"4/5 * * * * *",    // 第4秒开始-第59秒，每隔5秒打印一次(4/5 = 4-59/5)
	"0 0 5 * * 0,1,2",  // 每周日，周一，周二的凌晨5:00:00打印一次
	"*/9 29 0 1 * *",   // 每个月的一号，0点第29分钟，每隔9秒打印一次
	"0 30 * * * *",     // 每天每个小时的30分打印一次
	"0 0 5 * * 1",      // 每周一上午5点整打印一次
	"3 0 0 * * *",      // 每天0点的3秒打印一次
}

func cronWork() {

	cronTab.AddFunc("3 0 0 * * *", func() {
		fmt.Println(time.Now())
	})
	//cronTab.Start()

	cronTab.Run()

}

func scheduleWork(thisSpec string, nextN int) {
	schedule, err := cron.Parse(thisSpec)
	if err != nil {
		fmt.Printf("scheduleWork err(%s)", err)
		return
	}

	now := time.Now()

	fmt.Println("now time is:", now)

	for nextN > 0 {

		nextN--

		next := schedule.Next(now)

		str := fmt.Sprintf("next time is:%v\n", next)
		if nextN == 0 {
			str += "\n"
		}

		fmt.Printf(str)

		now = next
	}

}

func main() {

	for _, spec := range specs {
		fmt.Printf("spec is %s\n", spec)
		scheduleWork(spec, 10)
	}

}
