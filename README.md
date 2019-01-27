# <img src="buf-io.png" width="200">

Storage management platform connected with your favourite products.

# Demo
![](buf-io.gif)

## Support

### Data storage
* [DropBox](dropbox.md)

### Integrations
* [Slack](slack.md)

## Development

### Prerequisites
* Docker with Docker Compose
* Your favourite IDE with remote debug. (I highly recommends Goland with these setup: https://mikemadisonweb.github.io/2018/06/14/go-remote-debug/)
* Data storage and integrations supported by buf-io
* buf-io configure file (buf-io.yaml)

## buf-io.yaml
This is a configuration file. Currently structure looks like this:

```yaml
config:
  storage:
    - example storage e.g dropbox
  integrations:
    example integration e.g slack:
```

If you need details for particular storage/integrations read their .md files.
