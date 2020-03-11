# Webhook Samples in Shell

> Webhook Samples in Shell.

[![Build Status](http://img.shields.io/travis/badges/badgerbadgerbadger.svg?style=flat-square)](https://travis-ci.org/badges/badgerbadgerbadger) [![License](http://img.shields.io/:license-mit-blue.svg?style=flat-square)](http://badges.mit-license.org)

## Table of Contents

- [Installation](#installation)
- [Features](#features)
- [Contributing](#contributing)
- [Team](#team)
- [FAQ](#faq)
- [Support](#support)
- [License](#license)

## Installation

- Written in Shell

### Clone

- Clone this repo to your desktop or server using:

``` bash
sudo mkdir -p /data/webhook
cd /data/webhook
sudo git clone https://github.com/ExinOne/webhook-samples
cd shell
```

### Setup

Search `7000000012` in [Mixin Messenger](https://mixin.one/messenger) and add **[Webhook](https://mixin.one/codes/4d792128-1db8-4baf-8d90-d0d8189a4a7e)** as contact.

Invite Webhook and somebody who want to receive monitor message to a small group in Mixin Messenger. Open Webhook in the group, you can see the access token.

> Note: The access token is only available for the owner of the group.

Copy `config.cfg.defaults` to `config.cfg` and put your `ACCESS_TOKEN` in the `config.cfg`.

``` bash
LOG_FILE=webhook.log
WEBHOOK_URL=https://webhook.exinwork.com/api/send?access_token
ACCESS_TOKEN=YOUR_ACCESS_TOKEN
```

Then exectue `./webhook.sh` in the terminal to test it works or not.

## Features

- Provide Webhook samples in Shell
- Send example message via Webhook which based on Mixin API

## Contributing

To be continued.

## Team

@Exin

## FAQ

To be continued.

## Support

Reach out to us at one of the following places!

- Website at <a href="https://exin.one" target="_blank">`exin.one`</a>
- Twitter at <a href="http://twitter.com/ExinOne" target="_blank">`@ExinOne`</a>
- Email at `robin@exin.one`

## License

[![License](http://img.shields.io/:license-mit-blue.svg?style=flat-square)](http://badges.mit-license.org)

- **[MIT license](https://opensource.org/licenses/mit-license.php)**
- Copyright 2019 © <a href="https://exinpool.com" target="_blank">ExinPool</a>.