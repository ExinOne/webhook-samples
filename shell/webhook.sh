#!/bin/bash
#
# Copyright © 2019 Exin <robin@exin.one>
#
# Distributed under terms of the MIT license.
#
# Desc: Webhook shell sample code.
# User: Robin@Exin
# Date: 2020-03-09
# Time: 18:15:36

# load the config library functions
source config.shlib

# load configuration
log_file="$(config_get LOG_FILE)"
webhook_url="$(config_get WEBHOOK_URL)"
access_token="$(config_get ACCESS_TOKEN)"

log="`date '+%Y-%m-%d %H:%M:%S'` UTC `hostname` `whoami` INFO Hello World!"
echo $log >> $log_file

# call webhook via curl
curl ${webhook_url}=${access_token} -XPOST -H 'Content-Type: application/json' -d '{"category":"PLAIN_TEXT","data":"'"$log"'"}' > /dev/null 2>&1
if [ $? -eq 0 ]
then
    log="`date '+%Y-%m-%d %H:%M:%S'` UTC `hostname` `whoami` INFO send mixin successfully."
    echo $log >> $log_file
else
    log="`date '+%Y-%m-%d %H:%M:%S'` UTC `hostname` `whoami` INFO send mixin failed."
    echo $log >> $log_file
fi