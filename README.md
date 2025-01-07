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

### go & gRPC proxy (envoy)

```
docker compose up -d api
```

### go test

```
cd api/
go test ./test/service/controller/*
```
