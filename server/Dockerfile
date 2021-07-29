FROM golang:1.16.6-buster as builder

WORKDIR /src/server

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -v -ldflags "-s -w" -o /server

FROM scratch
COPY --from=builder /server /server

EXPOSE 8080

ENTRYPOINT ["/server"]
CMD [""]
