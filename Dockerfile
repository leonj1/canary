FROM alpine:3.5

ADD canary/

ENTRYPOINT [ "/canary" ]

