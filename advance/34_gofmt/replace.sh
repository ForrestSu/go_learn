#!/bin/bash

# gofmt -w -r 'interface{} -> any' ./...

demos=(
  'TestTask -> testTask'
  'SecondTask -> secondTask'
)

for rule in "${demos[@]}"; do
   echo ${rule}
   gofmt -w -r "${rule}" ./
done
