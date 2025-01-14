#!/bin/bash
git push origin master &&

git fetch upstream &&

git merge upstream/master -m "merge goland/dl master" && 

sed -i "s#golang.org/dl#github.com/SunJary/dl#g"  `grep go -rl --include '*.go' .` && 

git commit -a -m"new version" &&

git push origin master
