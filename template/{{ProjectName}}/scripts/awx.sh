#!/usr/bin/env bash

VERSION=`${PWD}/scripts/version.sh`

curl -X POST \
  https://1001.bale.ai/api/v2/job_templates/46/launch/ \
  -H 'Authorization: Basic amVua2luczpqc0BscyZzcyMx' \
  -H 'Content-Type: application/json' \
  -H 'Postman-Token: 6d5c2995-d8c2-4c9e-886b-f8a11487aaa2' \
  -H 'cache-control: no-cache' \
  -d '{
	"extra_vars":{
 "_config_repo_url": "ssh://git@prd-gitlab.c002.obale.ir:2222/bale/prd/{{ProjectName}}.git",
 "_service": "{{ProjectName}}",
 "_hosts": "nav1152",
 "_version": "docker.bale.ai/goft/{{ProjectName}}:'$VERSION'"
	}
}'