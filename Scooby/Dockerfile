
FROM golang:1.18-bullseye
#RUN go install github.com/spf13/viper@latest


ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor
ENV SCOOBYCALLSCOUNT=100
ENV APP_HOME /go/src/scooby

RUN mkdir -p "${APP_HOME}"
ADD "./src" "${APP_HOME}"

#RUN go get github.com/spf13/viper


WORKDIR "${APP_HOME}"
RUN go mod vendor
RUN go install
RUN go build -o main
EXPOSE 10000
CMD ["/go/src/scooby/main"]