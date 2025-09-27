#!/usr/bin/env bash

(
  WORKDIR=/data
  docker run --rm -it -v "$(pwd)":"${WORKDIR}" mando9/seru -o ${WORKDIR} -i ${WORKDIR}/"$1" -t ${WORKDIR}/"$2"
)