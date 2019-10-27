FROM google/cloud-sdk:alpine
# RUN mkdir /certs
# ENV CLOUDSDK_CONFIG /config/mygcloud
# ADD config/mygcloud /config/mygcloud
RUN apk -Uuv add ca-certificates
ADD bin/drone-gce-linux-amd64 /bin/gce

ENTRYPOINT ./bin/gce
