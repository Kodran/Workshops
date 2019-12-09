####################################################
# Debug env
####################################################
FROM golang:latest AS debug

WORKDIR /go/src/github.com/kodran/workshop/day3/app

# Expose ports used for debugging
EXPOSE 40000 1541

# Install delve debug tool
RUN go get github.com/derekparker/delve/cmd/dlv

# Copy all files in the proyect
COPY . .

# Run delve tool for debugging
CMD [ "dlv", "debug", "github.com/kodran/workshop/day3/app", "--listen=:40000", "--headless=true", "--api-version=2", "--log" ]

####################################################
# Build application with his own go environment
####################################################
FROM golang:latest AS release
WORKDIR /go/src/github.com/kodran/workshop/day3/app
# Copy all files in the proyect
COPY . .

# Get go dependencies
RUN go get -d -t -v ./...

# Build Go application targetting to run on linux
RUN GOOS=linux go build -o app -a --ldflags '-extldflags "-static"' -tags netgeo -installsuffix netgeo .

####################################################
# Release only the application bin into the container
# Self contained application, builded only with linux runtime

FROM scratch 
WORKDIR /root/
COPY --from=release /go/src/github.com/kodran/workshop/day3/app .
CMD ["./app"]  