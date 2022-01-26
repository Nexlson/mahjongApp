#!/bin/bash

# build images
cd app/ && docker image build -t mahjongfront .
cd ../core/ && docker image build -t mahjongback

# run docker-compose on detach mode
cd .. && docker-compose up -d