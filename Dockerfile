FROM golang:1.25-alpine

WORKDIR /app

# Install air
RUN go install github.com/air-verse/air@latest

# Dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

EXPOSE 8111

CMD ["air"]
