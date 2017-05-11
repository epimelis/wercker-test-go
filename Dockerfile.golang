FROM golang:latest
#FROM golang:onbuild
#FROM golang:alpine
MAINTAINER Epi Melis <epimelis1@gmail.com>
RUN mkdir /app
ADD . /app/ 
WORKDIR /app
RUN go  build -o zzz .
CMD ["/app/zzz"]
