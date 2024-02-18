FROM golang:1.21-alpine as builder

WORKDIR /app

# cli args for private repo pull
ARG ACCESS_TOKEN

RUN apk add --no-cache make gcc musl-dev linux-headers git jq bash

# allow private repo pull
RUN git config --global url."https://${ACCESS_TOKEN}:x-oauth-basic@github.com/Exca-DK".insteadOf "https://github.com/Exca-DK"

# precatche downloads
COPY ./go.mod .
COPY ./go.sum .
RUN go mod download

# copy pkgs
COPY ./service /app/service
COPY ./core /app/core
COPY ./cmd /app/cmd
COPY ./gen /app/gen
COPY ./contracts /app/contracts
COPY ./mocks /app/mocks

COPY ./Makefile .

RUN make service

FROM alpine:latest

COPY --from=builder /app/build/bin/piggy /usr/local/bin

CMD ["piggy"]
ENTRYPOINT ["piggy"]