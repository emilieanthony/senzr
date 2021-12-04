# Senzr API

## Prerequisites

Before running the server, install the [google cloud CLI](https://cloud.google.com/sdk/docs/install)
for your system. Then run the following commands:

1. `gcloud auth login`
2. `gcloud beta auth application-default login`

You need to Have Go 1.16 installed. Install for your system [here](https://go.dev/doc/install)

## How to start

_Make sure you have done the prerequisites_

1. Run `make install` to install required dependencies
2. Run `make develop` to start the server
