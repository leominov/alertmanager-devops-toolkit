FROM golang:1.11.3-alpine3.7 as build
ADD . /go/src/github.com/leominov/alertmanager-devops-toolkit
WORKDIR /go/src/github.com/leominov/alertmanager-devops-toolkit
RUN go build -o /bin/alertmanager-devops-toolkit .

FROM alpine:3.7
COPY --from=build /bin/alertmanager-devops-toolkit /usr/local/bin/
ENTRYPOINT ["alertmanager-devops-toolkit"]
