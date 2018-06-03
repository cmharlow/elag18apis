.PHONY: test

PROJECT_UNIT_TESTS     =$(shell go list ./... | grep -v test | grep -v db)
PROJECT_INT_TESTS      =$(shell go list ./db)
PROJECT_E2E_TESTS      =$(shell go list ./test)
LOCALSTACK_SERVICES    =dynamodb,s3
LOCAL_ENDPOINT_HOST    ?=localhost
LOCAL_ENDPOINT         =--endpoint-url=http://${LOCAL_ENDPOINT_HOST}
DYNAMO_ENDPOINT        =${LOCAL_ENDPOINT}:4569
PROJ_TABLE_NAME        =resources
PROJ_BUCKET_NAME       =taco-deposited-files
PROJ_AWS_REGION        =us-east-1
PROJ_AWS_ACCESS_KEY_ID =999999
PROJ_AWS_SECRET_KEY    =999999
PROJ_ENV_VARS          =AWS_REGION=${PROJ_AWS_REGION} AWS_ACCESS_KEY_ID=${PROJ_AWS_ACCESS_KEY_ID} AWS_SECRET_ACCESS_KEY=${PROJ_AWS_SECRET_KEY}

default: dependencies resources

dependencies:
	go get github.com/golang/dep/cmd/dep
	dep ensure

resources: table

table:
	$(eval TABLE=$(shell aws $(DYNAMO_ENDPOINT) dynamodb list-tables | jq '.TableNames[0] // ""'))
	@if [[ $(TABLE) != "" ]]; \
	  then echo "$(TABLE) table found"; \
		else aws $(DYNAMO_ENDPOINT) dynamodb create-table --table-name $(PROJ_TABLE_NAME) \
				--attribute-definitions AttributeName=tacoIdentifier,AttributeType=S \
				AttributeName=externalIdentifier,AttributeType=S \
				AttributeName=version,AttributeType=N \
				--key-schema "AttributeName=tacoIdentifier,KeyType=HASH" \
				--provisioned-throughput=ReadCapacityUnits=10,WriteCapacityUnits=10 \
				--global-secondary-indexes "IndexName=ResourceByExternalIDAndVersion, \
					KeySchema=[{AttributeName=externalIdentifier,KeyType=HASH}, \
					           {AttributeName=version,KeyType=RANGE}], \
					Projection={ProjectionType=ALL}, \
					ProvisionedThroughput={ReadCapacityUnits=10,WriteCapacityUnits=10}" ; \
	fi;
