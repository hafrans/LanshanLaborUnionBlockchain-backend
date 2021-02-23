
ARG GO_VERSION=1.14

FROM golang:${GO_VERSION} as builder

COPY . /app/

WORKDIR /app/

ENV CGO_ENABLED=0

ENV GOOS=linux

ENV GO111MODULE=on

ENV GOARCH=amd64

ENV GOPROXY=https://goproxy.cn


RUN go mod download -x && go mod verify 

RUN go build -ldflags="-w -s" -o main ./main.go

FROM scratch

COPY ./docs/ /app/docs/

COPY --from=builder /app/main /app/

COPY ./conf/conf.ini.example /app/runtime/databases/table.lock

WORKDIR /app/

ENV CGO_ENABLED=0

ENV GOOS=linux

ENV GO111MODULE=on

ENV GOARCH=amd64

ENV GOPROXY=https://goproxy.cn

ENV TZ=Asia/Shanghai

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

VOLUME /app/conf/

VOLUME /app/runtime/static/

EXPOSE 8088/tcp

CMD ["/app/main"]

