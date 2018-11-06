# Example About CERNBox API implementation

This repository is for illustration purposes.
This repository contains a gRPC server implementation
of the [About CERNBox API](https://github.com/cernbox/cernboxapis).
It also contains a gRPC gateway for this API.


# Quickstart

Start the gRPC server that will listen on port 9901.

```
cd grpcserver
go get ./...
go build
./grpcserver &
```

Start the HTTP-to-gRPC gateway server that will listen of port 8801.

```
cd ../gateway
go get ./...
go build
./gateway &
```


# Usage

You can send a request to the gateway with curl.

```
$ curl -s localhost:8081/v1/about/members| python -m json.tool
{
    "members": [
        {
            "display_name": "Hugo Gonzalez Labrador",
            "email": "hugo.gonzalez.labrador@cern.ch"
        }
    ],
    "status": {
        "code": "CODE_OK"
    }
}
```

You can send a request to the gRPC server using [prototool](https://github.com/uber/prototool):

```
$ prototool grpc --address localhost:9901 --method cernbox.aboutv1.About/GetDocumentation --data '{}'
{
  "status": {
    "code": "CODE_OK"
  },
  "documentation": {
    "serviceUrl": "https://cernbox.cern.ch",
    "githubUrl": "https://github.com/cernbox"
  }
}
```
