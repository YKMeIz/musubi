<img align="right" src="https://raw.githubusercontent.com/nrechn/musubi/master/logo.png">

# Musubi Message Server
[![License](https://img.shields.io/badge/license-GPL--3.0-red.svg?style=flat-square)](https://github.com/nrechn/musubi/blob/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/nrechn/musubi?style=flat-square)](https://goreportcard.com/report/github.com/nrechn/musubi)


Musubi is a message server powered by [Akari Message Framework](https://github.com/nrechn/akari). It follows KISS (Keep it simple, stupid) design principle, and is designed for IoT communication and notification push from *nix side to any device.

## Installation

> Note: The whole installation process requires `root` privilege.

#### Download and install source

```sh
$ wget -O /bin/musubi https://github.com/nrechn/musubi/releases/download/v0.1/musubi
```

Give `musubi` execute permission:

```sh
$ chmod 755 /bin/musubi
```

#### Create the configuration file

```sh
$ mkdir /etc/musubi/
$ vim /etc/musubi/config.yml
```

Add following content to your `config.yml` :

```yml
# Domain name and port for Musubi message server.
domainName: localhost
portNumber: 8080

# Server certificate chain and private key of domain name.
# Leave them empty if you don't want to enable TLS/SSL
certChain: /path/to/your/cert/chain
certKey: /path/to/your/cert/key

# Relative path for handling HTTP POST requests.
messageRelativePath: /nc

# Relative path for providing websocket service.
websocketRelativePath: /wc

# Path to SQLite database file.
databasePath: /path/to/your/database/file

# Pushbullet Token.
pushbullet:
  token: your.pushbullet.token

```

#### (Optional) Create systemd service unit configuration

A systemd service unit configuration could be utilized to manage `musubi daemon`.

```sh
$ wget -O /usr/lib/systemd/system/musubi.service https://github.com/nrechn/musubi/raw/master/musubi.service
```

Reload systemd, scanning for new or changed units:

```sh
$ systemctl daemon-reload
```

## Usage

> Assume you have installed SQLite in your host system.

#### Create a new user for Musubi message server

In the example below, we want to create new user `musubiUser`, and output new token for `musubiUser`:

```sh
$ musubi register musubiUser
Create User: musubiUser
musubiUser's token is: 42ff87b6de39282bb233b7aed5f658b1
```

#### Run Musubi message server

If you have systemd service unit configuration:

```sh
$ systemctl start musubi
```

OR,

If you prefer start `musubi daemon` manually:

```sh
$ musubi daemon
```

#### Send a message to Musubi message server

```sh
curl -H "Content-Type: application/json" \
-X POST \
-d '{"Source": "42ff87b6de39282bb233b7aed5f658b1", "Destination": ["42ff87b6de39282bb233b7aed5f658b1"], "Data": {}}' \
http://localhost:8080/nc
```

If you get `{"Status":"error! Source musubiUser is offline."}`, that means Musubi message server works.

> Note: Up to here is only about server side usage. Client side and interaction are still under developing.

## How it works

Musubi message server serves a HTTP POST API to receive HTTP requests, and serves a websocket service to communicate with any device supports websocket. It receives messages from HTTP POST request or websocket. Then pushing messages to target destination(s).

Musubi message server also supports to broadcast messages; send message to third party services, such as Pushbullet.

## Features

### Mandatory Identity Authentication

Musubi message server identifies every device by token. Each message must have correct token of `Source` and `Destination`. When a device try to connect websocket service, it needs to provide its identity in 30 second in order to register itself. Otherwise, websocket service will reject and close the connection.

However, for human readable purpose, token is stored with `name`:

```json
{
   "Name":"musubiUser",
   "Token":"42ff87b6de39282bb233b7aed5f658b1"
}
```

### Unified Message Format

All messages sent to or sent by Musubi message server has an unified format. It means all messages transferred with Musubi message server must follow this Json format:

```json
{
   "Source":"example token",
   "Destination":[
      "example token or special command",
      "example token"
   ],
   "Data":{
      "example 1":"example",
      "example 2":"example"
   }
}
```

Musubi message server reads and check `Source` and `Destination` to determine where the message is from and where the message is going. `Data` is utilized for users to exchange information.

### Broadcast

If `Destination` set as `["BROADCAST"]`, Musubi based server will broadcast this message to every device registered as online.

```json
{
   "Source":"example token",
   "Destination":[
      "BROADCAST"
   ],
   "Data":{
      "example 1":"example",
      "example 2":"example"
   }
}
```

### Pushbullet Support

Musubi message server supports sending notification via Pushbullet. Set `"Destination":["PUSHBULLET"]` to send a message to Pushbullet service.
> Currently, only support sending "push" notification in type "note".

```json
{
   "Source":"example token",
   "Destination":[
      "PUSHBULLET",
      "example token",
      "example token"
   ],
   "Data":{
      "Type":"note",
      "Title":"push a note",
      "Body":"note body",
      "AccessToken":"your Pushbullet token"
   }
}
```
If you set multiple destinations, Musubi based server will try to send the message to the destinations following `"HPUSHBULLET"`. If one of those destinations is offline, Musubi based server will send the message to Pushbullet. This method could be seen as adding an alternative destination for receiving notification.
> Note: `"Data"` should have same format as example above. Otherwise, Pushbullet notification would fail to send.
