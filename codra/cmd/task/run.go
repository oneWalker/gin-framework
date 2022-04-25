package task

import (
	"gin-practice/pkg/db/mongodb"
	"gin-practice/pkg/db/mysql"
	"gin-practice/pkg/rabbitmq"
	"gin-practice/pkg/tools"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

// var taskService *service.TaskService
// var examdataService *service.ExamdataService

// NewCommand .
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		DisableFlagParsing: true, // 不解析参数
		Use:                "task",
		Short:              "Task server",
		Run: func(_ *cobra.Command, _ []string) {

			mongodb.Init()

			mysql.Init()

			rabbitmq.Init()

			//taskService = service.NewTaskService()
			//examdataService = service.NewExamdataService()
			//examdataService.Timer()

			queueName := viper.GetString("amqp.rabbitmq.dlx_queue")
			exchange := viper.GetString("amqp.rabbitmq.dlx_exchange")

			err := rabbitmq.Client.SubscribeToDLX(queueName, exchange, onMessage)
			if err != nil {
				logrus.Errorf("rabbitmq.Client.SubscribeToDLX err %+v: \n", err)
			}

			tools.GracefulShutdownAndCleanup(func() {
				mongodb.Close()
				mysql.Close()
				rabbitmq.Close()
			})
		},
	}

	return cmd
}

// 从amqp中获取数据
func onMessage(delivery amqp.Delivery) {

	// var examroom model.Examroom
	// err := json.Unmarshal(delivery.Body, &examroom)
	// if err != nil {
	// 	logrus.Error(err)
	// 	return
	// }

	// logrus.WithField("method", "onMessage").Info(examroom)

	// switch examroom.ExamSystemType {
	// case model.SystemTypeAgora: // 声网
	// 	err := taskService.AgoraBillTask(&examroom)
	// 	if err != nil {
	// 		logrus.Errorf("Agora BillTask error: %s, roomid = %d\n", err, examroom.ID)
	// 	}
	// case model.SystemTypeYSX: // 云视讯
	// 	err := taskService.YSXBillTask(&examroom)
	// 	if err != nil {
	// 		logrus.Errorf("YSX BillTask error: %s, roomid = %d\n", err, examroom.ID)
	// 	}
	// case model.SystemTypeNull: // 不使用第三方系统
	// 	err := taskService.NullBillTask(&examroom)
	// 	if err != nil {
	// 		logrus.Errorf("SystemTypeNull BillTask error: %s, roomid = %d\n", err, examroom.ID)
	// 	}

	// default:
	// }

	// logrus.Info("examroom_over:", examroom)
	// if err := examdataService.Save(examroom.ID); err != nil {
	// 	logrus.Errorf("Examdata SaveTask error: %+v\n", err)
	// }

}
