#!/bin/bash

# build images
cd app/ && sudo docker image build -t mahjongfront .
cd ../core/ && sudo docker image build -t mahjongback

# run docker-compose on detach mode
cd .. && sudo docker-compose up -d