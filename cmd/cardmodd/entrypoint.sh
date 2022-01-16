#!/bin/bash

# migrate database
# TODO: seed database with known cards, magics, and rations (new entity/edge)
migrate -database "${CARDMOD_DATABASE_ENDPOINT}" -path ./migrations up

# run service
exec ./cardmodd