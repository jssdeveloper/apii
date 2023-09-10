FROM golang:1.21.1
WORKDIR /app
COPY . .
EXPOSE 3000
RUN go mod tidy
RUN go build -o apii .
CMD ["./apii"]