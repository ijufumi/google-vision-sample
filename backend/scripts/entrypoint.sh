#!/bin/bash

declare exit
exit=0
while [ $exit -eq 0 ]
do
  echo "Waiting 10 seconds...."
  sleep 10s
  /app/db version
  if [ $? -eq 0 ]; then
    exit=1
  else
    echo "App didn't connect to database."
  fi
done

# execute migration
/app/db up
# launch app
/app/app