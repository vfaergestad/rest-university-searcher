FROM golang:1.17-alpine AS builder

LABEL maintainer="vegarfae@stud.ntnu.no"
LABEL stage=builder

WORKDIR /go/src/app/cmd

COPY ./cmd /go/src/app/cmd
COPY ./internal /go/src/app/internal
COPY ./go.mod /go/src/app/go.mod
COPY ./go.sum /go/src/app/go.sum
COPY ./serviceAccountKey.json /go/src/app/serviceAccountKey.json

RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o main


FROM scratch

LABEL maintainer="vegarfae@stud.ntnu.no"

WORKDIR /

COPY --from=builder /go/src/app/cmd/main .

CMD ["./main"]