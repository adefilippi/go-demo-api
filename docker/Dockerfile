ARG GO_VERSION=1.23

FROM golang:${GO_VERSION}-alpine
ENV GIN_MODE=release
COPY . ./
RUN go mod download
RUN go build -ldflags="-s -w" -trimpath -o /app
RUN apk add upx && upx --best /app
EXPOSE 8080
CMD [ "/app" ]