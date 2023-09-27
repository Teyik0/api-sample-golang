FROM golang:latest

WORKDIR /go/app

COPY . .

RUN go install
RUN go build -o /go/app/bin/main .

ENTRYPOINT [ "/go/app/bin/main" ]