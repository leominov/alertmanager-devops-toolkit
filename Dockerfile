FROM golang:1.14-alpine3.11 as build
ADD . /go/src/github.com/leominov/alertmanager-devops-toolkit
WORKDIR /go/src/github.com/leominov/alertmanager-devops-toolkit
RUN wget https://github.com/prometheus/alertmanager/releases/download/v0.21.0/alertmanager-0.21.0.linux-amd64.tar.gz && \
    tar zxvf alertmanager-0.21.0.linux-amd64.tar.gz && \
    mv alertmanager-0.21.0.linux-amd64/amtool /bin/amtool
RUN go build -o /bin/alertmanager-devops-toolkit .

FROM alpine:3.11
COPY --from=build /bin/alertmanager-devops-toolkit /usr/local/bin/
COPY --from=build /bin/amtool /usr/local/bin/
ENTRYPOINT ["alertmanager-devops-toolkit"]
