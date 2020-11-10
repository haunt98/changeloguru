FROM golang:1.15-buster AS build

WORKDIR /go/src/changeloguru

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o /go/bin/changeloguru

FROM gcr.io/distroless/base-debian10

COPY --from=build /go/bin/changeloguru /

ENTRYPOINT ["/changeloguru"]
