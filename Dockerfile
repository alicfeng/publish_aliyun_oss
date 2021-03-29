FROM alpine:3.7
LABEL Maintainer="AlicFeng <a@samego.com>" \
      Description="oss publish plugin based on golang"

COPY release/publish_cli /usr/local/sbin/publish_cli

RUN chmod a+x /usr/local/sbin/publish_cli

CMD ["/usr/local/sbin/publish_cli"]