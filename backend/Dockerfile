FROM golang:1.23

WORKDIR /usr/lat.sh/backend
COPY . .
# Q: Why is everything connected by &&?
# A: Trying to save some layers
RUN go mod download && go mod verify
RUN go build ./cmd/backend
CMD "./backend"
