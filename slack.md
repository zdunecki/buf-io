## Getting Started
* Follow https://api.slack.com/slack-apps steps
* Find OAuth Access Token and and put them in environment variables following these steps:
```code
export SLACK_TOKEN=your_oauth_token
```
or create .env file and add `SLACK_TOKEN=your_oauth_token`
* configure storage
* setup **Permissions**, **Event Subscriptions** and **Interactive Components**  

### Permissions
* Go to https://api.slack.com/apps/`{YOUR_APP}`/oauth
* Add needed Scopes: `channels:history`, `channels:read`, `chat:write:bot`, `files:read`

### Event Subscription
* Go to https://api.slack.com/apps/`{YOUR_APP}`/event-subscriptions
* Fill Request URL with:http://`{PATH_TO_YOUR_APP}`/slack/events
* Add needed Workspace Events: `file_created`, `message.channels`

`{PATH_TO_YOUR_APP}` - I highly recommend to use https://ngrok.com/ if you are in dev mode.

### Interactive Components
* Go to https://api.slack.com/apps/`{YOUR_APP}`/interactive-messages
* Fill Request URL with:http://`{PATH_TO_YOUR_APP}`/slack/interactive-components

## buf-io.yaml
This is a configuration file for Slack.

### Structure
```yaml
config:
  storage:
    - example storage e.g dropbox
  integrations:
    slack:
      namespace: "NAME-SPACE-TO-FILE"
      noack:
        - 3X4MPL3ID
```

### Variables
Inside buf-io.yaml you are able to get predefined variables from buf-io code. Syntax for following variables looks like this: `{{AVAILABLE_VARIABLE_NAME}}`

### Description
* `namespace` - File path where **storage** will save a file. Available variables inside namespace:
    * `CHANNEL_NAME` - Slack channel name
    * `CHANNEL_ID` - Slack channel id
    * `FILE_NAME` - uploaded file name

* `noack` - Defaults every time when you upload a file you will be prompted about agreement for upload.
If you want prevent this behaviour just add room id there.

### Example:
```yaml
config:
  storage:
    - example storage e.g dropbox
  integrations:
    slack:
      namespace: "{{CHANNEL_NAME}}({{CHANNEL_ID}})"
      noack:
        - 3X4MPL3ID
```