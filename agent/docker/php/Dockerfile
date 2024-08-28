FROM golang:1.22-bullseye as agentbuild
WORKDIR /go/src/github.com/adesaegher/kubectl-flame
ADD . /go/src/github.com/adesaegher/kubectl-flame
RUN go get -d -v ./...
RUN cd agent && go build -o /go/bin/agent

FROM desaegher/php-debug:8.1 AS phpspybuild
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
RUN git clone https://github.com/brendangregg/FlameGraph


FROM desaegher/php-debug:bullseye-8.1
RUN apt update && \
    apt install -y \
    lsof \
    strace \ 
    procps \
    binutils
RUN mkdir /app
COPY --from=agentbuild /go/bin/agent /app/agent
COPY --from=phpspybuild /var/www/html/phpspy/phpspy /app/phpspy
COPY --from=phpspybuild /var/www/html/phpspy/stackcollapse-phpspy.pl /app/stackcollapse-phpspy.pl
COPY --from=phpspybuild /var/www/html/FlameGraph /app/FlameGraph
COPY --from=phpspybuild /usr/local/bin/php /usr/local/bin/php

CMD [ "/app/agent" ]
