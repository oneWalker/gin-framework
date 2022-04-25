package rabbitmq

import (
	"fmt"
	"github.com/pkg/errors"
	"strconv"
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

// global
var (
	// Client .
	Client *MessageClient
	once   sync.Once
)

// Init method
func Init() {

	once.Do(func() {
		Client = new(MessageClient)

		// amqp://用户名:密码@地址:端口号
		connectStr := fmt.Sprintf("amqp://%s:%s@%s:%d",
			viper.GetString("amqp.rabbitmq.username"),
			viper.GetString("amqp.rabbitmq.password"),
			viper.GetString("amqp.rabbitmq.host"),
			viper.GetInt("amqp.rabbitmq.port"),
		)
		Client.ConnectToBroker(connectStr)

		logrus.Info("rabbitmq connect successfully")
	})
}

// Close method
func Close() {
	if Client != nil {
		Client.Close()
		logrus.Info("rabbitmq connect closed")
	}
}

// MessageClient is our real implementation, encapsulates a pointer to an amqp.Connection
type MessageClient struct {
	conn *amqp.Connection
}

// ConnectToBroker connects to an AMQP broker using the supplied connectionString.
func (m *MessageClient) ConnectToBroker(connectionString string) {
	if connectionString == "" {
		panic("Cannot initialize connection to broker, connectionString not set. Have you initialized?")
	}

	var err error
	m.conn, err = amqp.Dial(fmt.Sprintf("%s/", connectionString))
	if err != nil {
		panic("Failed to connect to AMQP compatible broker at: " + connectionString)
	}
}

// Publish publishes a message to the named exchange.
func (m *MessageClient) Publish(body []byte, exchangeName string, exchangeType string) error {
	if m.conn == nil {
		return errors.New("Tried to send message before " +
			"connection was initialized.")
	}
	ch, err := m.conn.Channel() // Get a channel from the connection
	if err != nil {
		return fmt.Errorf("Failed to open a channel: %v", err)
	}
	defer ch.Close()
	err = ch.ExchangeDeclare(
		exchangeName, // name of the exchange
		exchangeType, // type
		true,         // durable
		false,        // delete when complete
		false,        // internal
		false,        // noWait
		nil,          // arguments
	)
	if err != nil {
		return fmt.Errorf("Failed to register an Exchange: %v", err)
	}

	queue, err := ch.QueueDeclare( // Declare a queue that will be created if not exists with some args
		"",    // our queue name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return fmt.Errorf("Failed to register an Queue: %v", err)
	}

	err = ch.QueueBind(
		queue.Name,   // name of the queue
		exchangeName, // bindingKey
		exchangeName, // sourceExchange
		false,        // noWait
		nil,          // arguments
	)
	if err != nil {
		return fmt.Errorf("Failed to Bind: %v", err)
	}

	err = ch.Publish( // Publishes a message onto the queue.
		exchangeName, // exchange
		exchangeName, // routing key      q.Name
		false,        // mandatory
		false,        // immediate
		buildMessage(body))
	if err != nil {
		return fmt.Errorf("Failed to publish: %v", err)
	}

	logrus.Infof("A message was sent: %v\n", string(body))
	return nil
}

// PublishOnQueue publishes the supplied body onto the named queue, passing the context.
func (m *MessageClient) PublishOnQueue(body []byte, queueName string) error {
	if m.conn == nil {
		return errors.New("Tried to send message before " +
			"connection was initialized.")
	}
	ch, err := m.conn.Channel() // Get a channel from the connection
	if err != nil {
		return fmt.Errorf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	queue, err := ch.QueueDeclare( // Declare a queue that will be created if not exists with some args
		queueName, // our queue name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return fmt.Errorf("Failed to register an Queue: %v", err)
	}

	// Publishes a message onto the queue.
	err = ch.Publish(
		"",         // exchange
		queue.Name, // routing key
		false,      // mandatory
		false,      // immediate
		buildMessage(body))
	if err != nil {
		return fmt.Errorf("Failed to publish: %v", err)
	}

	logrus.Infof("A message was sent to queue %v: %v\n", queueName, string(body))
	return nil
}

// PublishOnDLX publishes the supplied body onto the named queue, passing the context.
func (m *MessageClient) PublishOnDLX(body []byte, queueName string, dlxName string, ttl int) error {
	if m.conn == nil {
		return errors.New("Tried to send message before " +
			"connection was initialized.")
	}
	ch, err := m.conn.Channel() // Get a channel from the connection
	if err != nil {
		return fmt.Errorf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	queue, err := ch.QueueDeclare( // Declare a queue that will be created if not exists with some args
		queueName, // our queue name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		amqp.Table{"x-dead-letter-exchange": dlxName}, // arguments
	)
	if err != nil {
		return fmt.Errorf("Failed to register an Queue: %v", err)
	}

	// Publishes a message onto the queue.
	err = ch.Publish(
		"",         // exchange
		queue.Name, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			Expiration:  strconv.Itoa(ttl),
			ContentType: "application/json",
			Body:        body, // Our JSON body as []byte
		})
	if err != nil {
		return fmt.Errorf("Failed to publish: %v", err)
	}

	logrus.Infof("A message was sent to queue %v: %v\n", queueName, string(body))
	return nil
}

// Subscribe registers a handler function for a given exchange.
func (m *MessageClient) Subscribe(exchangeName string, exchangeType string, consumerName string, handlerFunc func(amqp.Delivery)) error {
	ch, err := m.conn.Channel()
	if err != nil {
		return fmt.Errorf("Failed to open a channel: %v", err)
	}

	err = ch.ExchangeDeclare(
		exchangeName, // name of the exchange
		exchangeType, // type
		true,         // durable
		false,        // delete when complete
		false,        // internal
		false,        // noWait
		nil,          // arguments
	)
	if err != nil {
		return fmt.Errorf("Failed to register an Exchange: %v", err)
	}

	logrus.Infof("declared Exchange, declaring Queue (%s)\n", "")
	queue, err := ch.QueueDeclare(
		"",    // name of the queue
		false, // durable
		false, // delete when usused
		false, // exclusive
		false, // noWait
		nil,   // arguments
	)
	if err != nil {
		return fmt.Errorf("Failed to register an Queue: %v", err)
	}

	logrus.Infof("declared Queue (%d messages, %d consumers), binding to Exchange (key '%s')\n",
		queue.Messages, queue.Consumers, exchangeName)
	err = ch.QueueBind(
		queue.Name,   // name of the queue
		exchangeName, // bindingKey
		exchangeName, // sourceExchange
		false,        // noWait
		nil,          // arguments
	)
	if err != nil {
		return fmt.Errorf("Queue Bind Failed: %s", err)
	}

	msgs, err := ch.Consume(
		queue.Name,   // queue
		consumerName, // consumer
		true,         // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)
	if err != nil {
		return fmt.Errorf("Failed to register a consumer: %v", err)
	}

	go consumeLoop(msgs, handlerFunc)
	return nil
}

// SubscribeToQueue registers a handler function for the named queue.
func (m *MessageClient) SubscribeToQueue(queueName string, consumerName string, handlerFunc func(amqp.Delivery)) error {
	ch, err := m.conn.Channel()
	if err != nil {
		return fmt.Errorf("Failed to open a channel: %v", err)
	}

	logrus.Infof("Declaring Queue (%s)\n", queueName)
	queue, err := ch.QueueDeclare(
		queueName, // name of the queue
		false,     // durable
		false,     // delete when usused
		false,     // exclusive
		false,     // noWait
		nil,       // arguments
	)
	if err != nil {
		return fmt.Errorf("Failed to register an Queue: %v", err)
	}

	msgs, err := ch.Consume(
		queue.Name,   // queue
		consumerName, // consumer
		true,         // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)
	if err != nil {
		return fmt.Errorf("Failed to register an consumer: %v", err)
	}

	go consumeLoop(msgs, handlerFunc)
	return nil
}

// SubscribeToDLX registers a handler function for a given exchange.
func (m *MessageClient) SubscribeToDLX(dlxQueueName string, dlxName string, handlerFunc func(amqp.Delivery)) error {
	ch, err := m.conn.Channel()
	if err != nil {
		return fmt.Errorf("Failed to open a channel: %v", err)
	}

	err = ch.ExchangeDeclare(
		dlxName,  // name of the exchange
		"fanout", // type
		true,     // durable
		false,    // delete when complete
		false,    // internal
		false,    // noWait
		nil,      // arguments
	)

	if err != nil {
		return fmt.Errorf("Failed to register an Exchange: %v", err)
	}

	logrus.Infof("declared Exchange, declaring Queue (%s)\n", "")
	queue, err := ch.QueueDeclare(
		dlxQueueName, // name of the queue
		false,        // durable
		false,        // delete when usused
		false,        // exclusive
		false,        // noWait
		nil,          // arguments
	)
	if err != nil {
		return fmt.Errorf("Failed to register an Queue: %v", err)
	}

	logrus.Infof("declared Queue (%d messages, %d consumers), binding to Exchange (key '%s')\n",
		queue.Messages, queue.Consumers, dlxName)
	err = ch.QueueBind(
		queue.Name, // name of the queue
		"#",        // bindingKey
		dlxName,    // sourceExchange
		false,      // noWait
		nil,        // arguments
	)
	if err != nil {
		return fmt.Errorf("Queue Bind Failed: %s", err)
	}

	msgs, err := ch.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	if err != nil {
		return fmt.Errorf("Failed to register a consumer: %v", err)
	}

	go consumeLoop(msgs, handlerFunc)
	return nil
}

// Close closes the connection to the AMQP-broker, if available.
func (m *MessageClient) Close() {
	if m.conn != nil {
		logrus.Info("Closing connection to AMQP broker\n")
		m.conn.Close()
	}
}

func buildMessage(body []byte) amqp.Publishing {
	return amqp.Publishing{
		ContentType: "application/json",
		Body:        body, // Our JSON body as []byte
	}
}

func consumeLoop(deliveries <-chan amqp.Delivery, handlerFunc func(d amqp.Delivery)) {
	for d := range deliveries {
		// Invoke the handlerFunc func we passed as parameter.
		handlerFunc(d)
	}
}
