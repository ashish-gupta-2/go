FROM docker.hub/alpine:2.0

# install all depedencies in the container
RUN apk add --no-cache libc6-compat

# copy the local binary into the container
COPY ./out/bin/goapi /usr/local/bin/goapi

# expose required ports
EXPOSE 7171

# set entrypoint for the container
ENTRYPOINT ["/usr/local/bin/goapi"]
