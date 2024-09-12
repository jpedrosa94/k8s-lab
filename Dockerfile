# Stage 1: Build Stage
# Use a specific version of the official Golang image as the base image
FROM golang:1.22-bullseye AS build

# Create a non-root user for running the application
RUN useradd -u 1001 nonroot

# Set the working directory inside the container
WORKDIR /app

# Copy the entire application source code
COPY app/ .

# Compile the application during build and statically link the binary
RUN go build \
 -ldflags="-linkmode external -extldflags -static" \
 -tags netgo -o noahsarc main.go

# Stage 2: Deployable Image
# Use a minimal scratch image as the base image for the final image
FROM scratch

# Copy the /etc/passwd file from the build stage to provide non-root user information
COPY --from=build /etc/passwd /etc/passwd

# Copy the compiled application binary from the build stage to the final image
COPY --from=build /app/noahsarc .

# Use the non-root user created in the build stage
USER nonroot

# Expose the application port
EXPOSE 3000

# Command to run the application
CMD ["./noahsarc"]

# Add a health check to ensure the app is running correctly
HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
CMD curl --fail http://localhost:3000 || exit 1
