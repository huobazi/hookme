#!/bin/bash
STDIN=$(cat - | jq .)
echo "STDIN is : $STDIN"

TAG=$(echo $STDIN | jq .push_data.tag)
PUSHED_AT=$(echo $STDIN | jq .push_data.pushed_at)
DIGEST=$(echo $STDIN | jq .push_data.digest)

echo "TAG is $TAG"
echo "PUSHED_AT is $PUSHED_AT"
echo "DIGEST is $DIGEST"

