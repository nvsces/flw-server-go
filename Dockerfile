FROM golang:latest
RUN mkdir app
WORKDIR /app
COPY . .
RUN go build -o main .
RUN chmod +x ./main
CMD ["./main"]