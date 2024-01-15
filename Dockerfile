# The base go-image
FROM golang:1.22rc1-bullseye
 
# Create a directory for the app
RUN mkdir /app
 
# Copy all files from the current directory to the app directory
COPY ./app /app
 
# Set working directory
WORKDIR /app

# install air
RUN go install github.com/cosmtrek/air@latest

# install gopls
RUN go install golang.org/x/tools/gopls@latest

# Run the server executable
ENTRYPOINT [ "air", "-c",".air.toml" ]