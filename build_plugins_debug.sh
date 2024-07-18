#!/bin/bash

go build -buildmode plugin -gcflags all="-N -l" ./languages/**