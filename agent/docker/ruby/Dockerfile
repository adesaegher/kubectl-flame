FROM golang:1.14-buster as agentbuild
WORKDIR /go/src/github.com/adesaegher/kubectl-flame
ADD . /go/src/github.com/adesaegher/kubectl-flame
RUN go get -d -v ./...
RUN cd agent && go build -o /go/bin/agent

FROM rust:1.50 AS rbspybuild
WORKDIR /
RUN git clone --depth 1 --branch v0.8.1 https://github.com/rbspy/rbspy
RUN cd rbspy && cargo build --release

FROM bitnami/minideb:stretch
RUN mkdir /app
COPY --from=agentbuild /go/bin/agent /app/agent
COPY --from=rbspybuild /rbspy/target/release/rbspy /app/rbspy

CMD [ "/app/agent" ]
