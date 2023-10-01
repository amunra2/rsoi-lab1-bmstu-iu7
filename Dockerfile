# Build
FROM golang:alpine AS builder
LABEL maintainer="@amunra2 - Ivan Tsvetkov"

RUN mkdir -p /src/backend
WORKDIR /src/backend

RUN apk add --no-cache make

ADD . .

RUN cd src && make build && mv bin/person-service /bin

# Run
FROM golang:alpine

RUN apk add --no-cache bash ca-certificates make

COPY --from=builder /src/backend/src/Makefile /src/Makefile
COPY --from=builder /bin/person-service /bin/
COPY --from=builder /src/backend/src/migrations /src/migrations

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

WORKDIR /src

CMD ["/bin/bash", "-c", "make -f /src/Makefile migrate-up && /bin/person-service"]
