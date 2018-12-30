############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git


# Create appuser
RUN adduser -D -g '' appuser


COPY . /tmp/src/mypackage/myapp/
WORKDIR /tmp/src/mypackage/myapp/

# Using go mod with go 1.11
RUN go mod download

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /admissioncontroller

############################
# STEP 2 build a small image
############################
FROM scratch

# Copy our static executable.
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /admissioncontroller /admissioncontroller

# Use an unprivileged user.
USER appuser

# Run the hello binary.
ENTRYPOINT ["/admissioncontroller"]
