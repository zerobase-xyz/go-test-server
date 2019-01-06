FROM nginx:1.14.0-alpine

ENV APP_HOST 127.0.0.1
ENV APP_PORT 8000

RUN apk --no-cache add gettext

COPY ./nginx.conf /tmp/nginx.conf.template
COPY ./entrypoint.sh /usr/local/bin/entrypoint.sh
RUN chmod +x /usr/local/bin/entrypoint.sh

ENTRYPOINT ["entrypoint.sh"]
