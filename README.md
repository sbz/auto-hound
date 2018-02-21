# auto-hound

Auto-configure etsy's [hound](https://github.com/etsy/hound) project to search into multiple repositories.

This contains the script [gh-list-repo](/gh-list-repo.go) which use the [go-github](https://github.com/google/go-github) project to query GitHub.

You can pass the following environment variables to the script:

* `ORG_NAME`: github organization name
* `ORG_TYPE`: github organization privacy type

# Build

```
docker build -t sbz/auto-hound .
```

# Run

```
docker run -d -p 6080:6080 sbz/auto-hound
```

# Run with environment and random ports

```
docker run -d -e ORG_TYPE=public -e ORG_NAME=freebsd -P sbz/auto-hound
```
