# Base image consists of
#  FROM bitnami/kubectl
#  USER root
#  RUN apt-get update && apt-get install -y wget

# syntax=docker/dockerfile:1
FROM kubeshop/testkube-kubepug-executor:base
COPY kubepug /bin/runner
USER root
RUN mkdir -p $HOME/.kubepug  && \
    cd $HOME/.kubepug && \
    wget https://github.com/rikatz/kubepug/releases/download/v1.1.3/kubepug_linux_amd64.tar.gz && \
    tar -zxvf kubepug_linux_amd64.tar.gz && \
    rm -f kubepug_linux_amd64.tar.gz
ENV PATH="$HOME/.kubepug/:${PATH}"
WORKDIR /root/
USER 1001
ENTRYPOINT ["/bin/runner"]
