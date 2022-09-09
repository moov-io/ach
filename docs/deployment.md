---
layout: page
title: Kubernetes
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Kubernetes

The following snippet runs the ACH Server on [Kubernetes](https://kubernetes.io/docs/tutorials/kubernetes-basics/) in the `apps` namespace. You could reach the ach instance at the following URL from inside the cluster.

```
# Needs to be ran from inside the cluster
$ curl http://ach.apps.svc.cluster.local:8080/ping
PONG

$ curl http://localhost:8080/files
{"files":[],"error":null}
```

Kubernetes manifest - save in a file (`ach.yaml`) and apply with `kubectl apply -f ach.yaml`.

```yaml
apiVersion: v1
kind: Namespace
metadata:
  name: apps
---
apiVersion: v1
kind: Service
metadata:
  name: ach
  namespace: apps
spec:
  type: ClusterIP
  selector:
    app: ach
  ports:
    - name: http
      protocol: TCP
      port: 8080
      targetPort: 8080
    - name: metrics
      protocol: TCP
      port: 9090
      targetPort: 9090
---
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: ach
  namespace: apps
  labels:
    app: ach
spec:
  replicas: 1
  selector:
    matchLabels:
      app: ach
  template:
    metadata:
      labels:
        app: ach
    spec:
      containers:
      - image: moov/ach:v1.0.0
        imagePullPolicy: Always
        name: ach
        args:
          - -http.addr=:8080
          - -admin.addr=:9090
        env:
          - name: ACH_FILE_TTL
            value: 30m # 30 minutes
        ports:
          - containerPort: 8080
            name: http
            protocol: TCP
          - containerPort: 9090
            name: metrics
            protocol: TCP
        resources:
          limits:
            cpu: 0.1
            memory: 50Mi
          requests:
            cpu: 25m
            memory: 10Mi
        readinessProbe:
          httpGet:
            path: /ping
            port: 8080
          initialDelaySeconds: 5
          periodSeconds: 10
        livenessProbe:
          httpGet:
            path: /ping
            port: 8080
          initialDelaySeconds: 5
          periodSeconds: 10
      restartPolicy: Always
```
