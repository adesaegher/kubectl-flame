FROM golang:1.14-buster as agentbuild
WORKDIR /go/src/github.com/adesaegher/kubectl-flame
ADD . /go/src/github.com/adesaegher/kubectl-flame
RUN go get -d -v ./...
RUN cd agent && go build -o /go/bin/agent

FROM python:3.8-alpine AS phpspybuild
RUN apk update && \
    apk add --no-cache \
    git \
    make \
    gcc \
    libc6-dev \
    lsof \
    strace \
    procps
RUN git clone https://github.com/adsr/phpspy.git && cd phpspy && make && cp phpspy /

FROM python:3.8-alpine
RUN mkdir /app
COPY --from=agentbuild /go/bin/agent /app/agent
COPY --from=phpspybuild /phpspy /app/phpspy

CMD [ "/app/agent" ]