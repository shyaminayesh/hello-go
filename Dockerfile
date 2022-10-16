FROM alpine AS build

RUN apk add --no-cache go make

WORKDIR /app

COPY . .

RUN make dist

FROM scratch
COPY --from=build /app/hello /bin/hello
EXPOSE 8080/tcp
CMD [ "/bin/hello" ]
