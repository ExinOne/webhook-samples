#!/usr/bin/env python
# -*- coding: utf-8 -*-
#
# Copyright © 2019 Exin <robin@exin.one>
#
# Distributed under terms of the MIT license.
#
# Desc: Webhook python sample.
# User: Robin@Exin
# Date: 2020-03-09
# Time: 18:24:19

import logging
import requests
import yaml

config = yaml.safe_load(open("config.yml"))

def log_config():
    log_file = config["log_file"]
    logging.basicConfig(filename=log_file,
                                filemode='a',
                                format='%(asctime)s.%(msecs)d %(name)s %(levelname)s %(message)s',
                                datefmt='%Y-%m-%d %H:%M:%S',
                                level=logging.DEBUG)

def send_mixin(content):
    value = {'category':'PLAIN_TEXT', 'data':content}
    webhook_url = config["webhook"]["webhook_url"]
    access_token = config["webhook"]["access_token"]
    response = requests.post(webhook_url.format(access_token), data=value)
    if response.status_code == 200:
        logging.info("Send mixin successfully.")
    else:
        logging.error("Send mixin failed.")

def test():
    send_mixin("Hello World")

def main():
    log_config()
    test()

if __name__ == "__main__":
    main()