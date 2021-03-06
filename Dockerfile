FROM ubuntu:14.04
RUN apt-get update
RUN apt-get install -y build-essential mercurial git subversion wget curl

# env vars

ENV GOPATH /goprojects
ENV PATH /goprojects/bin:/usr/local/go/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/games
ENV MONGOHQ_URL 'mongodb://172.17.0.36:27017'

# go 1.4.1 tarball
RUN wget -qO- https://storage.googleapis.com/golang/go1.4.1.linux-amd64.tar.gz | tar -C /usr/local -xzf -


# GOPATH
RUN mkdir -p /goprojects
RUN mkdir -p /goprojects/bin
RUN mkdir -p /goprojects/pkg
RUN mkdir -p /goprojects/src
RUN mkdir -p /goprojects/src/gogo

#Install Revel framework
RUN go get github.com/revel/revel
RUN go get github.com/revel/cmd/revel
RUN go get gopkg.in/mgo.v2
RUN go get gopkg.in/mgo.v2/bson
RUN go get gopkg.in/pipe.v2
RUN go get code.google.com/p/go.net/websocket
RUN go get github.com/disintegration/imaging

RUN chmod 775 -R /goprojects

WORKDIR /goprojects/src/gogo

EXPOSE 27017