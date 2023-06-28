#!/usr/bin/env sh

until ! migrate status 2>&1 | tail -n +3 | grep Pending > /dev/null; do echo "Waiting for DB to be migrated"; sleep 1; done
