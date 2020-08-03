# go-jaeger-tracing

## Before running this app, make sure you already have jaeger installed. run below command to install jaeger locally using docker

```$ docker run -d -p 6831:6831/udp -p 16686:16686 jaegertracing/all-in-one:latest```

### Run command below to execute go app

```$ go run main.go```

### Then you can view tracing result by access jaeger dashboard on http://localhost:16686
