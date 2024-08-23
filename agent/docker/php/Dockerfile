FROM golang:1.22-bullseye as agentbuild
WORKDIR /go/src/github.com/adesaegher/kubectl-flame
ADD . /go/src/github.com/adesaegher/kubectl-flame
RUN go get -d -v ./...
RUN cd agent && go build -o /go/bin/agent

FROM php:8.2-fpm-alpine AS phpspybuild
RUN apk update && \
    apk add --no-cache \
    git \
    make \
    gcc \
    libc-dev \
    lsof \
    strace \
    procps \
    binutils \
    php82-phpdbg
RUN git clone https://github.com/adsr/phpspy.git && cd phpspy && make

FROM php:8.2-fpm-alpine
RUN apk update && \
    apk add --no-cache \
    lsof \
    strace \
    procps \
    binutils \
    php82-phpdbg
RUN mkdir /app
COPY --from=agentbuild /go/bin/agent /app/agent
COPY --from=phpspybuild /var/www/html/phpspy/phpspy /app/phpspy

CMD [ "/app/agent" ]