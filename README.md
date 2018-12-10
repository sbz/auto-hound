# auto-hound

Auto-configure etsy's [hound](https://github.com/etsy/hound) project to search into multiple repositories.

This contains the script [gh-list-repo](/gh-list-repo.go) which use the [go-github](https://github.com/google/go-github) project to query GitHub.

You can pass the following environment variables to the script:

* `ORG_NAME`: github organization name
* `ORG_TYPE`: github organization privacy type
* `ORG_TOPIC`: github organization topics (repos topics filter)

# Build

```
docker build --tag sbz/auto-hound .
```

# Run

```
docker run --name auto-hound --rm --detach --publish 6080:6080 sbz/auto-hound
```

# Run with environment and random ports

```
docker run --name auto-hound --rm --detach -e ORG_TYPE=public -e ORG_NAME=freebsd -P sbz/auto-hound
```

To figure out the port bound on the host, you can use `docker port auto-hound` or the
following command [docker inspect](https://docs.docker.com/engine/reference/commandline/inspect):

```
docker inspect --format '{{ (index (index .NetworkSettings.Ports "6080/tcp") 0).HostPort }}' auto-hound
```
