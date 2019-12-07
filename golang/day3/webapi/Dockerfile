FROM golang

WORKDIR /go/src/github.com/kodran/workshop/day3/webapi
EXPOSE 40000 1541

# The -d flag instructs get to stop after downloading the packages; that is,
# it instructs get not to install the packages.

# The -t flag instructs get to also download the packages required to build
# the tests for the specified packages.

# The -v flag enables verbose progress and debug output.
RUN go get -d -t  -v ./...

RUN go get github.com/derekparker/delve/cmd/dlv
ADD . .

CMD [ "dlv", "debug", "github.com/kodran/workshop/day3/webapi", "--listen=:40000", "--headless=true", "--api-version=2", "--log" ]