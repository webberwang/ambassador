if [ -z "${DOCKER_REGISTRY}" ]; then
    DOCKER_REGISTRY=ark3
fi
cat <<EOF
---
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  creationTimestamp: null
  name: statsd-sink
spec:
  replicas: 1
  strategy: {}
  template:
    metadata:
      creationTimestamp: null
      labels:
        service: statsd-sink
    spec:
      containers:
      - name: statsd-sink
        image: ${DOCKER_REGISTRY}/prom-statsd-exporter:0.5.0
        resources: {}
      restartPolicy: Always
status: {}
---
apiVersion: v1
kind: Service
metadata:
  creationTimestamp: null
  labels:
    service: statsd-sink
  name: statsd-sink
spec:
  ports:
  - protocol: UDP
    port: 8125
    name: statsd-metrics
  - protocol: TCP
    port: 9102
    name: prometheus-metrics
  selector:
    service: statsd-sink
EOF
