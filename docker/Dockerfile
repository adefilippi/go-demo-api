ARG GO_VERSION=1.23

FROM golang:${GO_VERSION}-alpine
COPY go.mod ./
ENV GIN_MODE=release
RUN go mod download
COPY . ./
RUN go build -o /app
EXPOSE 8080
CMD [ "/app" ]