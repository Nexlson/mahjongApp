#!/bin/bash

# build images
cd app/ && docker image build -t mahjongfront .
cd ../core/ && docker image build -t mahjongback

# run docker-compose
cd .. && docker-compose up