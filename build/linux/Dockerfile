FROM portainer/base

COPY dist app/

VOLUME /data

WORKDIR /

EXPOSE 9000

ENTRYPOINT ["/app/dockm-linux-amd64"]