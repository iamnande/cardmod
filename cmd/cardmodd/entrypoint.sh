#!/bin/bash

# database: migrate schemas
migrate -database "${CARDMOD_DATABASE_ENDPOINT}" -path ./migrations up

# gardner: seed the database
# TODO: seed database with known cards and ratios (new entity/edge)
./gardner

# run service
exec ./cardmodd