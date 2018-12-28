FROM acoshift/go-alpine

RUN mkdir app_log

VOLUME /app_log

ADD postgres-service /

CMD ./postgres-service
