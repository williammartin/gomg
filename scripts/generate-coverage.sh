#!/usr/bin/env bash

ginkgo -p -r --randomizeAllSpecs --failOnPending --randomizeSuites --race -cover -coverprofile=coverage.txt -outputdir=.

