FROM golang:1.19-alpine AS build

# create a working directory inside the image
WORKDIR /app

# copy Go modules and dependencies to image
COPY go.mod ./

# download Go modules and dependencies
RUN go mod download

# copy directory files i.e all files ending with .go
COPY . ./

# compile application
RUN CGO_ENABLED=0 GOOS=linux go build -o /ezserver
#IMPORTANT: `CGO_ENABLED=0` else if turned on then binary dynamic links to libc/libmusl

##
## STEP 2 - DEPLOY
##
# FROM alpine:3.14
FROM busybox

WORKDIR /

COPY --from=build /ezserver /
#copy from /ezserver from build to root of this container

EXPOSE 5000
#basically just comments ezserver uses port 5000

CMD ["/ezserver"]