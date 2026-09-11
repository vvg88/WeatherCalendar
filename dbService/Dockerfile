FROM golang:1.27.1
WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN go install -v ./...

EXPOSE 8080
ENV GOOGLE_APPLICATION_CREDENTIALS /go/src/app/gcloudFsKey/weather-calend-fires-2a048fef69e5.json

CMD ["dbService"]