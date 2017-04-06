FROM alpine
RUN apk add --update ca-certificates
COPY ./alexa-parrot /alexa-parrot
EXPOSE 8080
ENTRYPOINT ["/alexa-parrot"]
