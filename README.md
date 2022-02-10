## "Notify New Release" Action for GitHub Actions

Sends notification with information about a new version available with identification and changelog via slack.

## Usage

```yaml
name: example
on: [push]

jobs:
  hello_world_job:
    runs-on: ubuntu-latest
    name: A job to say hello
    steps:
      - name: Hello world action step
        uses: adrianolaselva/action-release-notify@master
        env:
          API_TOKEN_GITHUB: ${{ secrets.API_TOKEN_GITHUB }} 
          OWNER: owner-name
          REPOSITORY: repository-name
          TAG: '5.3.0.104'
          SLACK_WEBHOOK: ${{ secrets.SLACK_WEBHOOK }}
          SLACK_CHANNEL: squad-backoffice-reviews
          SLACK_COLOR: '#347dba'
          SLACK_TITLE: Nova Versão disponível
          SLACK_ICON: https://upload.wikimedia.org/wikipedia/en/3/39/R2-D2_Droid.png
          SLACK_USERNAME: 'R2-D2'
```