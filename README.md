
## Work In Progress [WIP]
WARNING: `euka` is still in progress with limited feature and unexpected bugs, although you still can use this and we have working apps (euka-cli) as a example of using `euka`.

Current feature:

 - [x] List Kafka Topic
 - [ ] List Consumer Group
 - [x] Produce Messages
 - [ ] Consume Messages
 - [ ] Other Kafka feature

** Working feature also applied to `euka-cli` .
# euka
Interact with Kafka using Go and Sarama based tools.
>Euka provide you both Go packages for your Go Apps, also ready-to-go tools (euka-cli & euka-web) to interact with your Kafka cluster.

Products:
 - [euka-go-packages](#euka-go-packages)
 - [euka-cli](#euka-cli)
 - [euka-web](euka-web)

## euka-go-packages

Using euka for your Go apps is easy. First, use `go get` to install the latest version of `euka` its dependencies:

    go get -u github.com/eumnDev/euka/pkg/euka

Next, include Cobra in your application:

    import "github.com/eumnDev/euka/pkg/euka"

Example:
**[WIP]**
## euka-cli
Kafka command-line handy utility tools.

Usage:

    euka-cli [command]


Available Commands:

      help        Help about any command
      list        List any kafka data (topic)
      produce     Send anything to Kafka Topic
      version     Print the version number of Euka-CLI

Flags:

      -b, --bootstrap-server string   Simple provide bootstrap-server without config file.
      -c, --config string             config file (default "config.yaml")
      -h, --help                      help for euka-cli

Use `euka-cli [command] --help` for more information about a command.

Configuration file example can be seen on examples folder.

## euka-web
**[WIP]**
