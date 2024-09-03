#!/bin/zsh
CURRENT_DIR=$1
rm -rf ${CURRENT_DIR}/genproto
for x in $(find ${CURRENT_DIR}/Google_Docs_proto/* -type d); do
  protoc -I=${x} -I=${CURRENT_DIR}/Google_Docs_proto -I /usr/local/go --go_out=${CURRENT_DIR} \
   --go-grpc_out=${CURRENT_DIR} ${x}/*.proto
done