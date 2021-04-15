package publish

import (
	"encoding/json"
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Device struct {
	// Sensor map[string]string `json:"sensor"`
	ID     string `json:"id"`
	IP     string `json:"ip"`
	Sensor string `json:"sensor"`
}

type Sensor struct {
	ID    string `json:"id"` //device
	Type  string `json:"type"`
	Value string `json:"value"`
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

func HandlePublish(id string, topic string) {
	var broker = "10.0.6.37"
	var port = 31883
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID(id)
	opts.SetUsername(id)
	opts.SetPassword("public")
	opts.SetAutoReconnect(true)
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler

	device := Device{"5", "0.0.0.4", "temperature2"}
	sensor := Sensor{"5", "humidity3", "30"}

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if topic == "device/value" {
		dev, _ := json.Marshal(device)
		token := client.Publish(TOPIC_DEVICE, 0, false, dev)
		token.Wait()
		time.Sleep(time.Second)
	} else {
		sen, _ := json.Marshal(sensor)
		sensorToken := client.Publish(TOPIC_SENSOR, 0, false, sen)
		sensorToken.Wait()
		time.Sleep(time.Second)
	}

	client.Disconnect(250)
}

func publish(client mqtt.Client) {
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
}
