FROM ubuntu:latest

# RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

RUN apt-get update && apt-get -y upgrade && apt-get -y install curl 
RUN curl -LO "https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl"
RUN mv kubectl /usr/local/bin/kubectl && chmod +x /usr/local/bin/kubectl
COPY bin/grgd-linux /usr/local/bin/iam

ENTRYPOINT ["iam"]
