package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	log.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	log.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	log.Printf("Connect lost: %v", err)
}

var broker string
var port int
var clientID string
var username string
var password string
var topic string
var message string

func init() {

	clientIDFlag := flag.String("cliend-id", "mqtt-send", "Client ID to send as")
	passwordFlag := flag.String("password", "mosquitto", "Password for the broker connection")
	usernameFlag := flag.String("user", "mosquitto", "User for the broker connection")
	brokerFlag := flag.String("broker", "birne", "Broker address")
	portFlag := flag.Int("port", 1883, "Broker port")
	topicFlag := flag.String("topic", "awtrix/custom/mqtt-send", "Topic to send to")

	flag.Parse()

	if envPass := os.Getenv("MQTT_PASSWORD"); envPass != "" {
		password = envPass
	} else {
		password = *passwordFlag
	}

	broker = *brokerFlag
	port = *portFlag
	clientID = *clientIDFlag
	username = *usernameFlag
	topic = *topicFlag

	message = getPiped()

}

func main() {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID(clientID)
	opts.SetUsername(username)
	opts.SetPassword(password)
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	log.Println("Sending", message)
	publish(client, message)

	client.Disconnect(250)

	// TODO implement subbing
	// sub(client)
}

func publish(client mqtt.Client, message string) {
	// topic string, qos byte, retained bool, payload interface{}
	token := client.Publish(topic, 0, false, message)
	token.Wait()
	// time.Sleep(time.Second)
}

// func sub(client mqtt.Client) {
// 	topic := "topic/test"
// 	token := client.Subscribe(topic, 1, nil)
// 	token.Wait()
// 	fmt.Printf("Subscribed to topic: %s", topic)
// }

func getPiped() (message string) {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		var stdin []byte
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			stdin = append(stdin, scanner.Bytes()...)
		}
		if err := scanner.Err(); err != nil {
			return ""
		}
		return fmt.Sprintf("%s", stdin)
	}

	return ""
}
