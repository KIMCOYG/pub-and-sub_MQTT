package main

import (
	"encoding/json"
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	// "mqtt-pub/pkg/publish"
)

type Device struct {
	// Sensor map[string]string `json:"sensor"`
	// ID     string `json:"id"`
	// IP     string `json:"ip"`
	// Sensor string `json:"sensor"`
	Ip       string `json:"ip"`
	Location string `json:"location"`
	Server   string `json:"server"`
	Type     string `json:"type"`
}

type Sensor struct {
	// ID    string `json:"id"` //device
	// Type  string `json:"type"`
	// Value string `json:"value"`
	Id    string  `json:"id"`
	Ip    string  `json:"ip"`
	Type  string  `json:"type"`
	Value float64 `json:"value"`
}

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

const TOPIC_DEVICE = "device/value"
const TOPIC_SENSOR = "sensor/value"

func main() {
	/* go HandlePublish("1", "device/value")
	HandlePublish("pub2", "sensor/value") */
	var broker = "10.0.6.37"
	var port = 31883
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID("PUB")
	opts.SetUsername("pub")
	opts.SetPassword("public")
	opts.SetAutoReconnect(true)
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler

	device := Device{"127.0.0.0:4322", "http://127:0.0.0:4322", "Linux", "uuid"}
	sensor := Sensor{"127.0.0.0:4322", "1", "temperature", 30}

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	dev, _ := json.Marshal(device)
	token := client.Publish(TOPIC_DEVICE, 0, false, dev)
	token.Wait()
	time.Sleep(time.Second)

	sen, _ := json.Marshal(sensor)
	sensorToken := client.Publish(TOPIC_SENSOR, 0, false, sen)
	sensorToken.Wait()
	time.Sleep(time.Second)

	client.Disconnect(250)
}

/* func publish(client mqtt.Client) {
	num := 10
	for i := 0; i < num; i++ {
		text := fmt.Sprintf("Message %d", i)
		token := client.Publish("topic/test", 0, false, text)
		token.Wait()
		time.Sleep(time.Second)
	}
}

func sub(client mqtt.Client) {
	// topic := "topic/test"
	token := client.Subscribe(TOPIC_SENSOR, 1, nil)
	token.Wait()
	fmt.Printf("Subscribed to topic: %s", TOPIC_SENSOR)
} */
