#!/bin/bash
go build

for d in consumer_*/ ; do
    cd $d
    go build -o consumer
    cd ..
done