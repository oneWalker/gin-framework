package schedule

import (
	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"
)

func InitCron() {
	i := 0
	c := cron.New()

	// AddFunc
	spec := "*/5 * * * * ?"

	//add the function directly
	c.AddFunc(spec, func() {
		logrus.Println("cron running:", i)
	})

	// AddJob方法
	//通过struct的初始化去运行里面的每一个任务
	c.AddJob(spec, new(TestJob1))
	//c.AddJob(spec, TestJob2{})

	// 启动计划任务
	c.Start()
	// 关闭着计划任务, 但是不能关闭已经在执行中的任务.
	defer c.Stop()

	//查询语句，保持程序运行，这里等同于for{}
	select {}
}

type TestJob1 struct {
}

func (tj1 TestJob1) Run() {
	logrus.Info("testJob1...")
}

type TestJob2 struct {
}

func (tj2 TestJob2) Run() {
	logrus.Info("testJob2...")
}
