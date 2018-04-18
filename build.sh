#!/usr/bin/bash

### dir structure
# /bfe
#     /bfe-common
#         /go
#             /output
#         /golang-lib
#         /mini_spider
#             build.sh

### restore working dir
WORKROOT=$(pwd) 

cd ${WORKROOT}

# prepare PATH, GOROOT and GOPATH
export GOPATH=$(pwd)

# export golang-lib to GOPATH
cd ${WORKROOT}
export GOPATH=$(pwd)/../bfe-common/golang-lib:$GOPATH

# run go test for all subdirectory
cd ${WORKROOT}/src/mini_spider
go test -c -o ./testRun
if [ $? -ne 0 ];
then
    echo "go compile test failed"
    exit 1
fi

go test -run testRun
if [ $? -ne 0 ];
then
    echo "go run test failed"
    exit 1
fi
rm -rf ./testRun
echo "OK for go test"

### build 
cd ${WORKROOT}/src/main
go build -o mini_spider
if [ $? -ne 0 ];
then
    echo "fail to go build mini_spider.go"
    exit 1
fi
echo "OK for go build mini_spider.go"

### create directory for output
cd ../../
if [ -d "./output" ]
then
    rm -rf output
fi
mkdir output

# copy config
mkdir output/conf
cp conf/spider.conf output/conf

# copy data
cp -r data output/

# copy file to bin
mkdir output/bin
mv src/main/mini_spider output/bin

# create dir for log
mkdir output/log

# change mode of files in /bin
chmod +x output/bin/mini_spider


echo "OK for build mini_spider"
