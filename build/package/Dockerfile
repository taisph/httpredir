FROM golang:1.12-alpine AS build

RUN apk --no-cache add \
    git

COPY . /app
WORKDIR /app
RUN go build -v -o ./out/httpredir ./cmd/httpredir/...

FROM golang:1.12-alpine

COPY --from=build /app/out/httpredir /app/httpredir
ENTRYPOINT /app/httpredir
