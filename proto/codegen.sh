#!/bin/sh
####### Variables #######
PROTO_DIR=./proto
GO_DIR=./api
JS_DIR=./client/src
DESCRIPTOR_DIR=./proxy/pb

####### Generate Go and JS code #######
CMD=
for proto in $(find $PROTO_DIR -name '*.proto' | sed 's!^.*/!!'); do
  PROTO_NAME=$(echo $proto | sed 's!\.proto!!')
  mkdir -p $GO_DIR/$PROTO_NAME
  mkdir -p $JS_DIR/$PROTO_NAME
  mkdir -p $DESCRIPTOR_DIR

  /usr/local/bin/entrypoint.sh -f $PROTO_DIR/$proto -l go -o $GO_DIR/$PROTO_NAME
  /usr/local/bin/entrypoint.sh -f $PROTO_DIR/$proto -l descriptor_set -o $DESCRIPTOR_DIR/$PROTO_NAME
  /usr/local/bin/entrypoint.sh -f $PROTO_DIR/$proto -l web -o $JS_DIR/$PROTO_NAME --js_out=import_style=commonjs --grpc-web_out=import_style=commonjs+dts
  # protoc \
  # --js_out=import_style=commonjs:$JS_DIR/$PROTO_NAME \
  # --grpc-web_out=import_style=commonjs+dts,mode=grpcwebtext:$JS_DIR/$PROTO_NAME \

  #   $PROTO_DIR/$proto
done
