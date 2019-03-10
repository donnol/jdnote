#!/bin/bash

go clean -cache && go install && env UNIT_TEST_ENV=$@ /mnt/d/go/bin/server
