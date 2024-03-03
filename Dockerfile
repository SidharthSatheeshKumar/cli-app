FROM golang:latest 

WORKDIR /app

COPY . /app/

RUN go mod download

# CGO_ENABLED=0 - taking statically linked go binary file without any dependencies
# GOOS=linux - set the environment is linux
RUN CGO_ENABLED=0 GOOS=linux go build -o cli-app

ENTRYPOINT  [ "./cli-app" ]
