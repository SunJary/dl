#!/bin/bash
git fetch godl &&

git merge godl/master -m "merge goland/dl master" && 

sed -i "s#github.com/SunJary/dl#github.com/SunJary/dl#g"  `grep go -rl .`  && 

git commit -a -m"new version" &&

git push origin master