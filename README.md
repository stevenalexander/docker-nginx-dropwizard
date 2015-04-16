# Docker nginx with Dropwizard application

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
docker-compose up
```

# TODO need to setup link between nginx container and dropwizard, see [here](https://docs.docker.com/userguide/dockerlinks/).
