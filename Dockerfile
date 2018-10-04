FROM golang:1.11.0-stretch as build
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go get gopkg.in/yaml.v2
RUN go build -o alertmanager-devops-toolkit .

FROM busybox:1
COPY --from=build /app/alertmanager-devops-toolkit /usr/local/bin/
RUN mkdir /work
WORKDIR /work
