FROM alpine:3.10 as app

RUN apk --no-cache upgrade && apk --no-cache add ca-certificates
ADD lesson3 /usr/local/bin/lesson3
WORKDIR /usr/local/bin/

CMD ["lesson3"]