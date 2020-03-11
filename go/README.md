# Webhook Samples in Go

> Webhook Samples in Go.

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

- Written in Go, go installation you can use `brew install go`

### Clone

- Clone this repo to your desktop or server using:

``` bash
sudo mkdir -p /data/webhook
cd /data/webhook
sudo git clone https://github.com/ExinOne/webhook-samples
```

### Setup

Search `7000000012` in [Mixin Messenger](https://mixin.one/messenger) and add **[Webhook](https://mixin.one/codes/4d792128-1db8-4baf-8d90-d0d8189a4a7e)** as contact.

Invite Webhook and somebody who want to receive monitor message to a small group in Mixin Messenger. Open Webhook in the group, you can see the access token.

> Note: The access token is only available for the owner of the group.

Copy `conf.template.json` to `conf.json` and put your `AccessToken` in the `conf.json`.

``` bash
{
    "LogFile": "webhook.log",
    "WebhookUrl": "https://webhook.exinwork.com/api/send?access_token=",
    "AccessToken": ""
}
```

Then run `go build` you can get a executable file named `go`. And you can exectue `./go` in the terminal to test it works or not.

## Features

- Provide Webhook samples in Go
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