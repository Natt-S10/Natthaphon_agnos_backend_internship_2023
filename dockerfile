FROM golang:bullseye AS build

WORKDIR /src/app
COPY . . 

ENV GO111MODULE=on
RUN go mod download

RUN go build -o server .

FROM golang:bullseye AS server

ENV PORT=4000
WORKDIR /app
COPY --from=build /src/app .
# COPY --from=build /src/app/server .

EXPOSE 4000

CMD ./server