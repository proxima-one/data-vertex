FROM golang:alpine
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -o -i main .
EXPOSE 4000
CMD ["./main"]
