[![Build Status](https://travis-ci.com/williammartin/gomg.svg?branch=master)](https://travis-ci.com/williammartin/gomg) [![codecov](https://codecov.io/gh/williammartin/gomg/branch/master/graph/badge.svg)](https://codecov.io/gh/williammartin/gomg)

# gomg

`gomg` provides a set of utilities for interacting with the [Open Microservice Guide](https://microservice.guide).

## Getting gomg

Getting `gomg` is easy, simply:

```
go get github.com/williammartin/gomg
```

and after that you can explore the `gomg` CLI using:

```
gomg --help
```

### Validating your microservice

You can validate your microservice against the `OMG` by navigating to the directory containing your `microservice.yml` and running:

```
gomg validate
```

### Building your microservice

You can build a docker image for your microservice by navigating to the directory containing your `microservice.yml` and `Dockerfile` and running:

```
gomg build
```

You can then run:

```
docker images
```

and see your microservice tagged in the format `omg-<MICROSERVICE>:latest`

**Note**: You will need a local Docker daemon running for this workflow to succeed.

## Running the tests

You can use the provided makefile to run tests:

```
make test
```

Please run this before providing any contributions.
