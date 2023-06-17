# Stage 1 - Build image
FROM golang:1.19 as build
WORKDIR /app
COPY . /app
RUN CGO_ENABLED=0 GOOS=linux go build -o cmd/api main.go

# Stage 2 - Serve API
FROM scratch
WORKDIR /app
COPY --from=build /app/cmd/api ./
EXPOSE 8000
CMD [ "./cmd/api" ]