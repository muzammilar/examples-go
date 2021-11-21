ARG UBUNTU_VERSION=20.04

FROM ubuntu:${UBUNTU_VERSION} AS ppbuilder
# ARG are reset after every FROM statement, so need redeclaration
ARG UBUNTU_VERSION
ARG GO_VERSION=1.17.3
# Install basic tools
RUN apt update
# Install container build requirements
RUN DEBIAN_FRONTEND="noninteractive" apt install -y apt-utils wget tar
# Install your favourite editor (we don't discriminate)
RUN DEBIAN_FRONTEND="noninteractive" apt install -y vim emacs nano
# Install Makefile requirements
RUN DEBIAN_FRONTEND="noninteractive" apt install -y build-essential make git lsb-release
# Install build tools (equivs installs fakeroot and debhelper)
RUN DEBIAN_FRONTEND="noninteractive" apt install -y automake equivs devscripts dh-sysuser

# Install golang
ARG GO_BINARY="go${GO_VERSION}.linux-amd64.tar.gz"
RUN wget "https://dl.google.com/go/${GO_BINARY}"
RUN tar -C /usr/local -xzf ${GO_BINARY}
RUN rm ${GO_BINARY}
# Update path variable (use $PATH and not ${PATH} due to some behavior difference)
ENV PATH="$PATH:/usr/local/go/bin:/root/go/bin"
