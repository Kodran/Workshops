FROM golang:latest

WORKDIR /go/src/app

COPY . .

# The -d flag instructs get to stop after downloading the packages; that is,
# it instructs get not to install the packages.

# The -t flag instructs get to also download the packages required to build
# the tests for the specified packages.

# The -v flag enables verbose progress and debug output.
RUN go get -d -t  -v ./...

# The -v flag enables verbose progress and debug output.
# The -i flag will save dependency artifacts in the pkg folders.
RUN go install -i -v ./...

CMD ["app"]