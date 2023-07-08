FROM golang:1.20-bullseye

WORKDIR /code
ADD . .

WORKDIR /code/src
RUN go build pollywog.go; chmod +x pollywog

WORKDIR /
RUN cp /code/src/pollywog .
COPY pollywog.yml .

EXPOSE 9999
CMD ["/pollywog", "pollywog.yml"]
