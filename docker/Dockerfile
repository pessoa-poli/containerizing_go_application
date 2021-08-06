FROM golang:1.14.4-alpine3.12

ENV somevar "this is a value brought from system env vars"

# Copies the good stuff (ps: Git is needed to build the beat)
COPY src src

#Build the beat and add to path
RUN cd src/
RUN go build -o=builtit src/*.go

CMD ./builtit

