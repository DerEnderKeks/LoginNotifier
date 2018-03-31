# LoginNotifier [![Build Status](https://travis-ci.org/DerEnderKeks/LoginNotifier.svg?branch=master)](https://travis-ci.org/DerEnderKeks/LoginNotifier)

This program notifies you of successful SSH login attempts. It parses the `/var/log/auth.log` file and sends you an alert via the supported platforms containing the user and source IP address.

# Installation

> TODO

# Configuration

The configuration file is by default located in `/etc/loginnotifier/config.json`. To specify a different location use the `--config` parameter.
In case the file doesn't exist it will be created during startup.

## Alerts

Currently supported platforms:
- Slack
- Discord

To enable alerts for a specific platform set the config option `alerts.<platform>.enabled` to `true`.

### Slack

To use Slack alerts you have to specify the [Webhook URL](https://my.slack.com/services/new/incoming-webhook/), the channel name, the username for the alerts and an icon emoji. You can find more information about Slack Webhooks [here](https://api.slack.com/incoming-webhooks).

### Discord

To use Discord alerts you have to specify the Webhook URL, the username for the alerts and an avatar url. You can find more information about Discord Webhooks [here](https://support.discordapp.com/hc/en-us/articles/228383668-Intro-to-Webhooks).

## Log file

By default this program uses the `/var/log/auth.log` file to detect new sessions. You can specify a different file with the config option `source_log`.

# License

[GNU Affero General Public License 3.0](https://www.gnu.org/licenses/agpl-3.0.en.html)