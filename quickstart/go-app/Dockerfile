FROM golang:alpine as build
RUN apk --no-cache add tzdata
WORKDIR /app
ADD . /app
RUN CGO_ENABLED=0 GOOS=linux go build -o demoapp .

FROM scratch as final
COPY --from=build /app/demoapp .
COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo
ENV TZ=Asia/Shanghai
CMD ["/demoapp"]