# Set the base image to use for building the container
FROM golang:latest

# Set metadata labels for the image
LABEL Authors="yjawad, malkhuza, aabdulhu, sayedalawi, yrahma"
LABEL Description="Container"
LABEL Version="Latest"

EXPOSE 5050

WORKDIR /app

RUN apt-get update && apt-get install -y sqlite3 libsqlite3-dev

COPY . .

RUN go build -o forum .

CMD ["./forum"]