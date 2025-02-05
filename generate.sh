#!/bin/sh
# This script will codegen a go package for a server stub or client.


docker run \
    --rm \
    --user $(id -u):$(id -g) \
    -v ${PWD}:/local \
    openapitools/openapi-generator-cli generate \
        -i /local/swagger.yml \
        -g go \
        -o /local/openapi \
        -p enumClassPrefix=true \
        -p skipFormModel=true
rm openapi/go.mod \
    openapi/go.sum \
    openapi/git_push.sh
#rm openapi/api_*.go \
#    openapi/.travis.yml \
#    openapi/client.go \
#    openapi/configuration.go \
#    openapi/git_push.sh \
#    openapi/go.mod \
#    openapi/go.sum \
#    openapi/response.go \
#    openapi/api/openapi.yaml \
#    openapi/docs/*Api.md

echo "module github.com/alapidas/tdameritrade-openapi
go 1.13" > go.mod

