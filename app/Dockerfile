FROM debian:latest

# Install necessary dependencies
RUN apt-get update && apt-get install -y \
    curl \
    git \
    wget

RUN wget https://go.dev/dl/go1.21.0.linux-amd64.tar.gz
RUN tar -xvf go1.21.0.linux-amd64.tar.gz
RUN mv go /usr/local

# Set environment variables
ENV GOROOT=/usr/local/go
ENV GOPATH=/go
ENV PATH=$GOPATH/bin:$GOROOT/bin:$PATH

# Set the working directory
WORKDIR /go/src/app

COPY code/ .

RUN go build

CMD ["./main"]
