[![Build Status](https://travis-ci.com/williammartin/omg.svg?branch=master)](https://travis-ci.com/williammartin/omg)

# omg
A collection of Open Microservice Guide structures

---

# Schema generation

Under `cmd/generate-omg-schema` you will find a utility tool that will generate a schema for microservice validation.

You can fetch this tool by running:

```
go get github.com/williammartin/cmd/generate-omg-schema
```

and use it as follows:

```
generate-omg-schema
```

*TIP*: Piping `generate-omg-schema` into a tool like `jq` makes it much more readable!
