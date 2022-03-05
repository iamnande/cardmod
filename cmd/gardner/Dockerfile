# gardner: base layer
FROM scratch

# gardner: copy over binary
COPY build/gardner /gardner

# gardner: start the service
USER 6710
ENTRYPOINT [ "./gardner" ]