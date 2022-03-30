FROM golang:1.13.1-alpine
WORKDIR /app
RUN export GO111MODULE=on
COPY go.mod go.sum ./

# install modules
RUN go mod download

COPY . .

# Command to run the executable
CMD [ "./fetch" ]