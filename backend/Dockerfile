# Build the Go backend and create a small runtime image
FROM golang:1.20-alpine AS build
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o payroll-api

FROM alpine:3.17
WORKDIR /app
COPY --from=build /app/payroll-api .
EXPOSE 8080
CMD ["./payroll-api"]
