ARG GO_VERSION=1.23

FROM golang:${GO_VERSION}-alpine
ENV GIN_MODE=release
WORKDIR /build
COPY . ./
RUN go mod download
RUN go install github.com/swaggo/swag/cmd/swag@latest && swag -v && swag init --parseDependency --parseInternal --exclude="pkg"
RUN go build -ldflags="-s -w" -trimpath -o /app
RUN apk add upx && upx --best /app
EXPOSE 8080
CMD [ "/app" ]