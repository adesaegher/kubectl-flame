FROM golang:1.22-bullseye as agentbuild
WORKDIR /go/src/github.com/adesaegher/kubectl-flame
ADD . /go/src/github.com/adesaegher/kubectl-flame
RUN pwd && ls && ls ..
RUN go get -d -v ./...
RUN cd agent && go build -o /go/bin/agent

FROM alpine as builder
RUN apk add --update git
RUN git clone https://github.com/brendangregg/FlameGraph

FROM alpine
RUN apk add --update perf perl
RUN mkdir -p /app/FlameGraph
COPY --from=agentbuild /go/bin/agent /app/agent
RUN mv /usr/bin/perf /app/perf
COPY --from=builder /FlameGraph /app/FlameGraph

CMD [ "/app/agent" ]
