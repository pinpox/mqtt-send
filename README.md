# mqtt-send

Simple CLI tool to send MQTT messages to a topic. The message is read from
stdin, connection parameters can be provided via flags.

## Usage

```
Usage of ./mqtt-send:
  -broker string
    	[BROKER] Broker address (default "birne")
  -cliend-id string
    	[CLIEND_ID] Client ID to send as (default "mqtt-send")
  -password string
    	[PASSWORD] Password for the broker connection (default "mosquitto")
  -port int
    	[PORT] Broker port (default 1883)
  -topic string
    	[TOPIC] Topic to send to (default "awtrix/custom/mqtt-send")
  -user string
    	[USER] User for the broker connection (default "pinpox")
```

Settings can be specified by flags or the shown `[ENVIRONMENT_VAR]`.

### Example

```sh
echo '{ "text": "this is a test form mqtt-send", "icon": 92, "color": "ff0000" }' | ./mqtt-send
```
