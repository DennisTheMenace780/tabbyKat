#!/bin/bash

# Create a TestRepo directory 
mkdir TestRepo

# Change directories into TestRepo2
cd TestRepo || exit

# Initialize a Git repository
git init
git config --global user.name "userName"
git config --global user.email "email123@google.com"

# Create a file called file.txt with "hello world" in it
echo "hello world" > file.txt

# Optionally, you can add and commit the file to the Git repository:
git add file.txt && git commit --no-verify -m "Inital Commit"

# Iterate through creating branches
count=1
while [ $count -lt 4 ]; do
    git checkout -b "branch-$count"
    count=$((count + 1))
done

# Reset to the first branch when finished
git checkout "branch-1"

