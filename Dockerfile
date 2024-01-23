FROM golang:1.20
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-adv-app
EXPOSE 8080
CMD ["/docker-adv-app"]


