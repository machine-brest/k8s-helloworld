# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.191.1/containers/alpine/.devcontainer/base.Dockerf

# [Choice] Alpine version: 3.14, 3.13, 3.12, 3.11
ARG VARIANT="3.14"
FROM mcr.microsoft.com/vscode/devcontainers/base:0-alpine-${VARIANT}

RUN apk update && \
    apk --no-cache add \
        ca-certificates \
        git \
        bash \
        openssh-client \
        sshpass \
        openssl \
        rsync && \
    rm -rf /var/cache/apk/*

# Install terraform

ARG TF_VERSION=0.12.31
RUN wget https://releases.hashicorp.com/terraform/${TF_VERSION}/terraform_${TF_VERSION}_linux_amd64.zip -O tmp.zip \
    && unzip tmp.zip \
    && chmod +x terraform \
    && mv terraform /usr/local/bin \
    && rm tmp.zip
