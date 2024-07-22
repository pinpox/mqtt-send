# mqtt-send

Simple CLI tool to send MQTT messages to a topic. The message is read from
stdin, connection parameters can be provided via flags.

## Usage

```
Usage of mqtt-send:
  -broker string
    	Broker address (default "birne")
  -cliend-id string
    	Client ID to send as (default "mqtt-send")
  -password string
    	Password for the broker connection (default "mosquitto")
  -port int
    	Broker port (default 1883)
  -topic string
    	Topic to send to (default "awtrix/custom/mqtt-send")
  -user string
    	User for the broker connection (default "mosquitto")
```

The password can also be provided as environment variable `MQTT_PASSWORD`

### Example

```sh
echo '{ "text": "this is a test form mqtt-send", "icon": 92, "color": "ff0000" }' | ./mqtt-send
```
