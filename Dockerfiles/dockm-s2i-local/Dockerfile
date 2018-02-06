FROM click2cloud/centos-git

COPY s2i /usr/local/bin

WORKDIR click2cloud-dockm/

COPY dist app/

VOLUME /click2cloud-dockm/data

EXPOSE 9000

ENTRYPOINT ["app/dockm-linux-amd64", "-d", "data", "-a", "app", "--no-analytics"]