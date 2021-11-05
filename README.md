# cenats
CLI for publishing/receiving CloudEvent messages through NATS

## Use Case
cenats is designed to make testing services that generate
or consume events easy. The syntax mirrors the official
nats cli. Use "cenats pub" to publish a message located
from a json file and "cenats sub" to listen on a topic
and print the payload of any cloudevents that are sent there.

## Installation
```bash
go install github.com/ehenry2/cenats
```

## Usage
```bash
Usage:
  cenats [command]

Available Commands:
  completion  generate the autocompletion script for the specified shell
  help        Help about any command
  pub         Publish a message in CloudEvents format to NATS
  sub         A brief description of your command

Flags:
      --config string    config file (default is $HOME/.cenats.yaml)
  -h, --help             help for cenats
  -s, --subject string   NATS subject to publish to. In NATS, subjects scope messages into streams or topics. (required)
  -u, --url string       URL of the NATS server (default "nats://127.0.0.1:4222")

Use "cenats [command] --help" for more information about a command.

```