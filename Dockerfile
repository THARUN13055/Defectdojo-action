# Use a specific version of Go for building the application
FROM golang:1.23.0-alpine AS build

# Labels for the image
LABEL Created="Tharun G" \
      Maintainer="TharunDevops1@outlook.com"

# Set the working directory
WORKDIR /defectdojo

# Copy go.mod and go.sum files first to leverage Docker cache
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY main.go ./

CMD ["go","run","main.go"]
# Build the application
# RUN go build -o app main.go

# # Create a new image using a distroless base
# FROM gcr.io/distroless/static-debian12

# # Copy the built application from the build stage
# COPY --from=build /defectdojo/app /

# # Set the command to run the application
# CMD ["/app"]