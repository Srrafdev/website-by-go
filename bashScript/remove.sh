#!/bin/bash

# stop all the container
docker stop $(docker ps -q)
echo "______________________________________________________________________________________________________________________________"
# remove all the containers          or this
docker rm $(docker ps -aq) #&& docker rm $(docker ps -aq)
echo "______________________________________________________________________________________________________________________________"
# remove all images
docker rmi -f $(docker images -q)
