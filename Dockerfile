FROM golang:1.26-alpine AS build-stage

WORKDIR /app
COPY go.mod ./
RUN go mod download

COPY *.go ./
COPY *.html ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/tautulli-scoreboard

FROM gcr.io/distroless/base-debian12 AS release-stage
WORKDIR /
COPY --from=build-stage /app/tautulli-scoreboard /app/tautulli-scoreboard
USER nonroot:nonroot

CMD ["/app/tautulli-scoreboard"]