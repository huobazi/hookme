#!/bin/bash
#----------------------------------------------------
# This is a example for jq read the json from stdin
#----------------------------------------------------
STDIN=$(cat - | jq .)
echo "STDIN is : $STDIN"

TAG=$(echo $STDIN | jq .push_data.tag)
PUSHED_AT=$(echo $STDIN | jq .push_data.pushed_at)
PUSHER=$(echo $STDIN | jq .push_data.pusher)

echo "TAG is $TAG"
echo "PUSHED_AT is $PUSHED_AT"
echo "PUSHER is $PUSHER"

