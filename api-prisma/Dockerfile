FROM golang:latest

WORKDIR /go/app

COPY . .

RUN go install
RUN go run github.com/steebchen/prisma-client-go db push
RUN go run github.com/steebchen/prisma-client-go generate

RUN go build -o /go/app/bin/main .

ENTRYPOINT [ "/go/app/bin/main" ]