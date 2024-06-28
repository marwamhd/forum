# Set the base image to use for building the container
FROM golang:latest
# Set metadata labels for the image
LABEL Authors="Yousif and Marwa"
LABEL Description="Container"
LABEL Version="Latest"
# Set the working directory inside the container
WORKDIR /playground
# Copy the current directory contents into the container's working directory
COPY . .
# Define the command to run when the container starts
CMD ["go", "run", "."]
## docker run -p 8080:8080 dockerimageandcontainer