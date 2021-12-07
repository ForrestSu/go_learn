#!/bin/bash

# gofmt -w -r 'interface{} -> any' ./...
set -x

gofmt -w -r 'TestTask -> testTask' ./
