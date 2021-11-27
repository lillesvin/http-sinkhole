# http-sinkhole

Simple tool to just dump the headers and bodies of incoming HTTP requests on stdout.

## Installation

```
go get github.com/lillesvin/http-sinkhole@v0.1.1
```

## Usage

```
http-sinkhole [-p INT]
```

Will listen for incoming requests on whatever port you specified with `-p` (default: 1234).

