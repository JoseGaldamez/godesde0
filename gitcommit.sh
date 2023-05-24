#!/bin/bash

branch=$2
if [ -z "$branch" ]; then
    branch="main"
fi

echo $branch

git add .
git commit -m $1
git push origin $branch