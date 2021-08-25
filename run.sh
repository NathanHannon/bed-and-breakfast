#!/bin/bash

go build -o bed-and-breakfast cmd/web/*.go && ./bed-and-breakfast -dbname=bed-and-breakfast -dbuser=nathan -cache=false -production=false