FROM jenkins/jenkins:lts

USER root

RUN apt-get update && \
    apt-get install -y make git golang && \
    rm -rf /var/lib/apt/lists/*

USER jenkins