# Develop Your API Section

## Bootcamp steps

1. Group: Go refresher (overview / pros / cons / setup)
2. SimpleDev Work
    3. Group: Go dependencies (overview / pros / cons / setup)
    2. Individual / Pairs: Check your Go Environment, that this repository is within it, & add dependencies
    6. Group: `Go-Swagger` review (overview / usage)
    6. Group: `Go-Swagger` generated server code review
    7. Individual / Pairs: Install go-swagger, validate your Swagger spec, check out the generated docs, then generate your swagger code.
    6. Group: Review our stub handler.
    7. Group: Review how to run our code locally.
    7. Individual / Pairs: Test run your SimpleDev server and cURL against our stub Deposit route.
6. TAQUITO Work
    7. Group: Review Localstack (how to run, why, interacting with DynamoDB there)
    7. Individual / Pairs: Run localstack and make your DynamoDB resources.
    8. Group: Review TAQUITO main.go / server.go
    7. Group: Review TAQUITO full Handlers.
    7. Group: Review TAQUITO internal services.
    7. Group: Review TAQUITO validators & dataUtils / hard data problems.
    7. Individual / Pairs: Spin up LocalStack with DynamoDB created
 7. Group: Some notes on TACO development parts missing in this session.

## Installing Dep

On a Mac / OSX, use homebrew:

```bash
$ brew install dep
$ brew upgrade dep
```

For other installs, follow the instructions here (cURL a shell script & run it, basically): https://github.com/golang/dep#installation

## Swagger & Go

This presumes you [already have `go` installed on your machine](https://github.com/cmh2166/elag18apis/tree/master#technical-prep) and are running the commands in this repository within your Go working directory.

### Install Go-Swagger

https://github.com/go-swagger/go-swagger/blob/master/docs/install.md <= has link to binary retrieval for installing swagger.

We want to validate and eventually generate some code based off of `swagger.json` using the `go-swagger` library. You'll need to install the `go-swagger` library before running these commands (commands for those using Mac OSX):

```shell
$ brew tap go-swagger/go-swagger
$ brew install go-swagger
$ brew upgrade go-swagger
```

This should give you the `swagger` binary command in your $GOPATH and allow you to manage versions better. The version of your go-swagger binary is **0.13.0** (run `swagger version` to check this).

### Nota Bene on go-swagger install from source

If instead of brew, you decide to install go-swagger from source (i.e. `go get -u github.com/go-swagger/go-swagger/cmd/swagger`), you will be running the library at its Github `dev` branch. You will need to change into that library (`cd $GOPATH/src/github.com/go-swagger/go-swagger`) and checkout from Github the latest release (`git checkout tags/0.13.0`). Then you will need to run `go install github.com/go-swagger/go-swagger/cmd/swagger` to generate the go-swagger binary in your `$GOPATH/bin`.

## Validate Your Swagger Specification

Once you are done with your TAQUITO `swagger.json` specification, you can use `go-swagger` to validate the file. Run:

```shell
$ swagger validate swagger.json
The swagger spec at "swagger.json" is valid against swagger specification 2.0
```

## Generate Some Swagger Documentation

You can use the `go-swagger` library to do a number of other things, like create a simple webpage & server for Swagger-based API documentation. Just run:

```shell
$ swagger serve swagger.json
2018/05/15 15:23:37 serving docs at http://localhost:50035/docs
```

Then go see your documentation at that URL. Once you're done, you can just hit `cntl+c` to stop the web server. You can see the TACO API specification documentation here: https://sul-dlss-labs.github.io/taco/.
