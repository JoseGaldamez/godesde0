#!/bin/bash

commit=$1

if [ -z "$commit" ]; then
    echo "Please enter a commit message"
    exit 1
fi

branch=$2
if [ -z "$branch" ]; then
    branch="main"
fi

echo $branch

git add .
git commit -m $1
git push origin $branch