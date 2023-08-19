FROM golang:1.19-alpine3.16 AS server_builder

RUN apk add gcc libc-dev

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . test-task-jungle-consulting/.

WORKDIR test-task-jungle-consulting
RUN go build -ldflags "-w -s -linkmode external -extldflags -static" -a cmd/app/main.go

FROM scratch
EXPOSE 8080
COPY --from=server_builder /app/test-task-jungle-consulting/main .
COPY --from=server_builder /app/test-task-jungle-consulting/.env .
CMD ["./main"]