FROM golang:1.15
WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 8080
ENV GOOGLE_APPLICATION_CREDENTIALS /go/src/app/gcloudFsKey/weather-calend-fires-2a048fef69e5.json

CMD ["dbService"]