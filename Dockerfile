# Set the base image to use for building the container
FROM golang:latest

# Set metadata labels for the image
LABEL Authors="yjawad, malkhuza, aabdulhu, sayedalawi, yrahma"
LABEL Description="Container"
LABEL Version="Latest"

# Set the working directory inside the container
WORKDIR /Forum

# Copy the current directory contents into the container's working directory
COPY . .

# Install SQLite and its development libraries
RUN apt-get update && apt-get install -y sqlite3 libsqlite3-dev

# Install Node.js and npm
RUN apt-get install -y nodejs

# Expose port 9191 for your Go application
EXPOSE 1010

# Define the command to run when the container starts
CMD ["go", "run", "."]
