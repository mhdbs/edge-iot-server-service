FROM alpine As builder 

MAINTAINER Mohammed Bilal

WORKDIR /go/src/mhdbs/edge-iot-server-service

COPY . /go/src/mhdbs/edge-iot-server-service

RUN apk add --no-cache git make musl-dev go

ENV GOROOT /usr/lib/go
ENV GOPATH /go
ENV PATH "/usr/local/go/bin:$PATH"

RUN go mod vendor
# RUN /go/bin/dep ensure -v

RUN go build -o edge-iot-server-service

FROM alpine

COPY --from=builder \
    /go/src/mhdbs/edge-iot-server-service/edge-iot-server-service \
    /go/src/mhdbs/edge-iot-server-service

COPY --from=builder \
    /go/src/mhdbs/edge-iot-server-service/config/edge-iot-server-service.json \
    /var/edge-iot-server-service.json

EXPOSE 8080

ENTRYPOINT [ "./go/src/mhdbs/edge-iot-server-service" ]