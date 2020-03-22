FROM golang:latest
COPY . $GOPATH/app
WORKDIR $GOPATH/app
RUN go build -o runner
CMD [ "./runner" ]