#!/bin/sh

/app/main \
  -httpPort "$HTTPPORT" \
  -pgDSN "$PGDSN"
