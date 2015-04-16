gradle clean buildJar
java -jar -Done-jar.silent=true volume-dropwizard/dropwizard-application.jar server volume-dropwizard/config.yml
