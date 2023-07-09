# crab-stash-protofiles
Repository containing protofiles and protoc generated golang code

## How To Use
To generate golang code create <microservice_name>.proto file in <microservice_name>/proto directory
Then add in Makefile under the project variable and .phony <microservice_name>
Then call make <microservice_name>
