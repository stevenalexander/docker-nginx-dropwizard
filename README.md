# Docker nginx with Dropwizard application [![Build Status](https://travis-ci.org/stevenalexander/docker-nginx-dropwizard.svg?branch=master)](https://travis-ci.org/stevenalexander/docker-nginx-dropwizard)

Example using Docker to create an nginx image and dropwizard image, then link them together so nginx acts as a reverse proxy for Dropwizard. Can be extended to link together multiple Dropwizard applications. Uses [Docker Compose](http://docs.docker.com/compose/) to create and configure the images.

Requires:
* [Docker](https://www.docker.com/)
* [Boot2Docker](http://boot2docker.io/)
* [Docker-compose](http://docs.docker.com/compose/)
* JDK (to compile java file locally)
* [Gradle](https://gradle.org/) (for build automation)

To run locally:

```
gradle run
# ./go
```

To run containers:
```
gradle buildJar
docker-compose up -d

# retrieve your docker host IP from boot2docker
boot2docker ip

# curl dropwizard/nginx containers using docker host IP
curl http://192.168.59.103:8080/hello
curl http://192.168.59.103:8090/hello
```

# Details

The `docker-compose.yml` file configures the two images, creating a dropwizard container and linking it to an nginx container. With the link in place, docker creates a hosts entry for the dropwizard container which can be used in the nginx config `volumes-nginx-conf.d/default.conf` when setting up the reverse proxy.

Based Travis build on [moul/travis-docker](https://github.com/moul/travis-docker).
