# Design Your API Section

## Bootcamp steps

1. Group: Walk through the TAQUITO API Contract
2. Group: ReST refresher (pros / cons / overview)
3. Group: Understand TAQUITO Data Model & Profiles (or data shapes)
4. Group: Define our TAQUITO API Routes
5. Group: Define our TAQUITO API Data Definitions (what does it need to know or check)
6. Group: Swagger refresher (overview / pros / cons)
7. Group: Walk through TAQUITO Swagger starter specification at [swagger.json](swagger.json)
8. Individual / Pairs: Add GET route to the TAQUITO Swagger starter specification at [swagger.json](swagger.json)
9. Individual / Pairs: Add DELETE route to the TAQUITO Swagger starter specification at [swagger.json](swagger.json)

Check your final specification against the proposed final TAQUITO Swagger specification at [swagger-done.json](swagger-done.json)

## TAQUITO Swagger Specification

See our starter TAQUITO Swagger specification at [swagger.json](swagger.json). It will be missing the `DELETE` and `GET` routes, as those are to be added through your work.

## Swagger & Go

This presumes you [already have `go` installed on your machine](https://github.com/cmh2166/elag18apis/tree/master#technical-prep) and are running the commands in this repository within your Go working directory.

### Install Go-Swagger

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
