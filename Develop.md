# Develop Your API Section

## Steps in the workshop

TBD

## Go-Swagger Code Generation

Validate using go-swagger, then generate using go-swagger

```shell
$ swagger validate swagger.json
```

If all good, we can now generate the start of our API code from our Swagger spec by running:

```shell
$ git rm -rf generated
	$ mkdir generated
	$ swagger generate server -t generated --exclude-main --principal authorization.Agent
(there appears to be no best way to handle specification-based re-generation of the generated/ API code)
```

## Generated Code Deep dive
