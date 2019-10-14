#!/bin/bash

curl -s https://api.github.com/users/adilSmile | jq '.id'
