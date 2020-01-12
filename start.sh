#!/bin/bash

docker run -d \
	-p "127.0.0.1:43156:8080" \
	--name "tiona" \
	tiona
