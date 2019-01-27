## Getting Started
* Create app under https://www.dropbox.com/developers
* Put printed OAuth Access Token (or go to https://www.dropbox.com/developers/apps and generate new one) in environment variables following these steps: 
```code
export DROPBOX_TOKEN=your_oauth_token
```
or create .env file and add `DROPBOX_TOKEN=your_oauth_token`
* configure storage
 

## buf-io.yaml
This is a configuration file for DropBox.

### Example:
```yaml
config:
  storage:
    - dropbox
  integrations:
    example integration e.g slack:
```