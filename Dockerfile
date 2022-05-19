FROM golang AS build

WORKDIR /app

COPY . . 

RUN GOPRIVATE='github.com/psvehla/*' go mod tidy -compat=1.18
RUN GOPRIVATE='github.com/psvehla/*' go mod download 
RUN CGO_ENABLED=0 go build -o app ./cmd/hello

##
## Deploy
##
FROM alpine:latest  

WORKDIR /root/

COPY --from=build /app/app ./

EXPOSE 8000

CMD ["./app", "--domain=0.0.0.0"]
