FROM golang:latest
RUN mkdir /app
WORKDIR /app
ADD . /app
RUN go build -o build/main
EXPOSE 9090
ENTRYPOINT ["build/main"]