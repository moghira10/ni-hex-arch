FROM golang:1.11-alpine AS builder

RUN apk update && apk add --no-cache git
# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/ni
# Copy everything from the current directory to the present working directory inside the container
COPY . .

# Enable GO111MODULE
ENV GO111MODULE=on
# Download all the dependencies
# RUN go get -d -v ./...

# # Install the package
# RUN go install -v ./...

RUN go mod download


# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build  -ldflags="-w -s" -a -installsuffix cgo -o /ni main.go

# EXPOSE 8080


FROM scratch
# Copy our static executable.
COPY --from=builder /ni /ni
# Run
ENTRYPOINT ["/ni"]

# CMD ["ni"]