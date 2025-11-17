FROM golang:1.20-alpine AS build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Ajout important : go mod tidy pour générer go.sum si absent
RUN go mod tidy

RUN go build -o payroll-api

FROM alpine:3.17
WORKDIR /app
COPY --from=build /app/payroll-api .
EXPOSE 8080
CMD ["./payroll-api"]

EXPOSE 10000
