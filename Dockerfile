# Stage 1: Build stage
FROM golang:1.25-alpine AS builder

WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o ./bin/api ./cmd/api/

RUN chmod +x ./entrypoint.sh
RUN chmod +x ./entrypoint-prod.sh

FROM alpine AS prod

ARG SCRIPT="entrypoint-prod.sh"
ENV SCRIPT=$SCRIPT

WORKDIR /app

COPY .env .env
COPY --from=builder /app/bin/api /app/bin/api
COPY --from=builder /app/entrypoint-prod.sh /app/entrypoint-prod.sh
CMD sh $SCRIPT

FROM builder AS dev

ARG SCRIPT="entrypoint.sh"
ENV SCRIPT=$SCRIPT

RUN go install github.com/air-verse/air@latest

CMD sh $SCRIPT
