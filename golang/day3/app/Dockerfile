FROM golang

WORKDIR /go/src/github.com/kodran/workshop/day3/app
EXPOSE 40000 1541

RUN go get github.com/derekparker/delve/cmd/dlv
ADD . .

CMD [ "dlv", "debug", "github.com/kodran/workshop/day3/app", "--listen=:40000", "--headless=true", "--api-version=2", "--log" ]