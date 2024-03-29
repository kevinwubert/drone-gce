FROM google/cloud-sdk:alpine
RUN mkdir /certs
ENV CLOUDSDK_CONFIG /config/mygcloud
ADD config/mygcloud /config/mygcloud
ADD bin/drone-gce-linux-amd64 /bin/drone-gce

ENTRYPOINT ["/bin/drone-gce"]
