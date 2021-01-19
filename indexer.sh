#!/bin/sh

bin=$1
os=$2
platform=$3

version=v$(./bin/$1-$2-$3 -v | awk '{ printf $3 }')

cp bin/$1-$2-$3 bin/$1-$2-$3-$version

url=https://s3.iamstudent.dev/public/grgd/$1-$2-$3-$version
size=$(wc -c bin/$1-$2-$3-$version | awk '{ printf $1}')
md5=$(md5 -q bin/$1-$2-$3-$version )

echo "          $version:"
echo "            url: \"$url\""
echo "            description: \"macos blababla\""
echo "            released: \"$(date '+%Y-%m-%dT%H:%M:%SZ')\""
echo "            author: \"gregod\""
echo "            size: $size"
echo "            md5: \"$md5\""
