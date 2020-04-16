FROM golang:1.14-alpine AS build

RUN apk add -U --no-cache make git ca-certificates

WORKDIR /go/src/github.com/banknovo/configurator
COPY . .

RUN make linux

FROM scratch AS run

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /go/src/github.com/banknovo/configurator/dist/configurator /configurator

ENTRYPOINT ["/configurator"]
