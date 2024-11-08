# golang-grpc-example

## How to use

### pb

```
docker compose run --rm protoc
```

### client

```
cd client/
npm install
npm run dev
```

### envoy

```
cd proxy/
docker buildx build -t golang-grpc-recap-envoy .
docker run -d -p 51051:51051 -p 9901:9901 golang-grpc-recap-envoy
```

### go

```
cd api/
go run ./server.go
```
