# --- Build stage: compile the app --- #
# Set the image
FROM golang:1.25 AS build-stage

WORKDIR /app

# Copy from build context into /app
COPY go.mod ./
RUN go mod download

COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /go-health-api

# --- Final stage: Create the image with just the absolute necesary to run the app. No building dependencies = smaller, secure image. --- #
# Docker Hardened Image -- Secure, minimal, distroless
FROM dhi.io/static:20230311-debian12 AS build-release-stage

# Copy the binary built in the last stage
COPY --from=build-stage /go-health-api /go-health-api

# Container port 
EXPOSE 8080

# safe unpriviledged user
USER 65532:65532

ENTRYPOINT ["/go-health-api"]
