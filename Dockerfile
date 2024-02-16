FROM golang:1.19 AS build
WORKDIR /go/src
COPY internal/route53dns ./internal/route53dns
COPY main.go .
COPY go.sum .
COPY go.mod .

ENV CGO_ENABLED=0

RUN go build -o route53dns .

FROM scratch AS runtime
COPY --from=build /go/src/route53dns ./
EXPOSE 8080/tcp
ENTRYPOINT ["./route53dns"]
