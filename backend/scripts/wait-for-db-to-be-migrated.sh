#!/usr/bin/env sh

until ! bahn-alarm migrate status status 2>&1 | tail -n +3 | grep Pending > /dev/null; do
  bahn-alarm migrate status
  echo "Waiting for DB to be migrated"
  sleep 1
done

bahn-alarm migrate status

echo "All migrations are applied"