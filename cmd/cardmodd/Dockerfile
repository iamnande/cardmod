# api: base layer
FROM scratch

# api: copy over binary
COPY build/cardmodd /cardmodd

# api: start the service
USER 6710
ENTRYPOINT [ "./cardmodd" ]