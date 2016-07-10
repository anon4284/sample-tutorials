#!/bin/bash
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix st -o app projects/sample-tutorials
