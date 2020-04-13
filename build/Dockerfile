FROM debian:buster
ENV DEBIAN_FRONTEND noninteractive
RUN apt-get update && apt-get install -y \
    build-essential \
    curl \
    devscripts \
    equivs \
    git-buildpackage \
    git \
    lsb-release \
    make \
    openssh-client \
    pristine-tar \
    rake \
    rsync \
    ruby \
    ruby-dev \
    rubygems \
    wget

RUN curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | sudo apt-key add - ; \
    echo "deb https://dl.yarnpkg.com/debian/ stable main" | sudo tee /etc/apt/sources.list.d/yarn.list ; \
    apt-get update && apt-get install  -y yarn

WORKDIR /tmp
RUN wget https://dl.google.com/go/go1.14.1.linux-$(dpkg --print-architecture).tar.gz; tar -C /usr/local -xzf go1.14.1.linux-$(dpkg --print-architecture).tar.gz
ENV PATH="/usr/local/go/bin:${PATH}"
ENV BUNDLE_BIN "/data/docker/.bundle/ruby/2.5.0/bin"
ENV BUNDLE_PATH "/data/docker/.bundle/ruby/2.5.0"
ENV BUNDLE_APP_CONFIG: "/data/.bundle"


RUN echo "gem: --no-ri --no-rdoc" >/etc/gemrc
RUN gem install bundler

RUN mkdir /data
WORKDIR /data
