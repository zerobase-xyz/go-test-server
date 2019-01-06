#!/bin/sh

envsubst '$$APP_HOST$$APP_PORT' \
         < /tmp/nginx.conf.template \
         > /etc/nginx/nginx.conf

nginx -g "daemon off;"
