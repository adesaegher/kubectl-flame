FROM golang:1.22-bullseye as agentbuild
WORKDIR /go/src/github.com/adesaegher/kubectl-flame
ADD . /go/src/github.com/adesaegher/kubectl-flame
RUN go get -d -v ./...
RUN cd agent && go build -o /go/bin/agent

FROM php:8.1.29-fpm AS phpspybuild
RUN apt update && \
    apt install -y \
    git \
    make \
    gcc \
    libc-dev \
    lsof \
    strace \
    procps \
    binutils
RUN git clone https://github.com/adsr/phpspy.git && cd phpspy && make

FROM php:8.1.29-fpm
RUN apt update && \
    apt install -y \
    lsof \
    strace \
    procps \
    binutils
RUN mkdir /app
COPY --from=agentbuild /go/bin/agent /app/agent
COPY --from=phpspybuild /var/www/html/phpspy/phpspy /app/phpspy

CMD [ "/app/agent" ]