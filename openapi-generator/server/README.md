# Go API Server for openapi

This is the public Twilio REST API.

## Overview
This server was generated by the [openapi-generator]
(https://openapi-generator.tech) project.
By using the [OpenAPI-Spec](https://github.com/OAI/OpenAPI-Specification) from a remote server, you can easily generate a server stub.
-

To see how to make this your own, look here:

[README](https://openapi-generator.tech)

- API version: 1.29.1
- Build date: 2022-06-10T11:57:13.852018Z[Etc/UTC]
For more information, please visit [https://support.twilio.com](https://support.twilio.com)


### Running the server
To run the server, follow these simple steps:

```
go run main.go
```

To run the server in a docker container
```
docker build --network=host -t openapi .
```

Once image is built use
```
docker run --rm -it openapi
```