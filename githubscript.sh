#!/usr/local/bin/bash

# curl -u jdyang11:c285ba9c37f1692271d868f93fe3e7266cb31603 https://api.github.com/user/repos -d "{\"name\":\"$1\"}";
echo $1
# git init;
git remote add origin git@github.com:jdyang11/$1.git;


