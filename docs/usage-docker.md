---
layout: page
title: Docker
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Docker

We publish a [public Docker image `moov/ach`](https://hub.docker.com/r/moov/ach/) from Docker Hub or use this repository. No configuration is required to serve on `:8080` and metrics at `:9090/metrics` in Prometheus format. We also have Docker images for [OpenShift](https://quay.io/repository/moov/ach?tab=tags) published as `quay.io/moov/ach`.

Pull & start the Docker image:
```
docker pull moov/ach:latest
docker run -p 8080:8080 -p 9090:9090 moov/ach:latest
```

List files stored in-memory:
```
curl localhost:8080/files
```
```
{"files":[],"error":null}
```

Create a file on the HTTP server:
```
curl -X POST --data-binary "@./test/testdata/ppd-debit.ach" http://localhost:8080/files/create
```
```
{"id":"<YOUR-UNIQUE-FILE-ID>","error":null}
```

Read the ACH file (in JSON form):
```
curl http://localhost:8080/files/<YOUR-UNIQUE-FILE-ID>
```
```
{"file":{"id":"<YOUR-UNIQUE-FILE-ID>","fileHeader":{"id":"","immediateDestination":"231380104","immediateOrigin":"121042882", ...
```