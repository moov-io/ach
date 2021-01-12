#!/bin/sh

[ $# -eq 0 ] && { echo "Usage: $0 version_tag"; exit 1; }

# releases should be created from master branch
if [ "$(git branch --show-current)" != "master" ]
then
    echo "switching to branch 'master'"
    git checkout master > /dev/null 2>&1
fi

status=$(git status --porcelain)

# only 2 files should be changed
if [ "$(echo "$status" | wc -l)" -ne 2 ]
then
  echo "Only version.go and CHANGELOG.md should be updated!"
  if [ -n "$status" ]
  then
    printf "Pending changes:\n%s\n" "$status"
  fi
  exit
elif ! echo "$status" | grep -q "CHANGELOG.md" && ! echo "$status" | grep -q "version.go"
then
    echo "version.go and changelog.md must be updated to proceed"
    exit
fi

# make sure this is a new tag
if git show-ref --tags | grep -q "$1"; then
  echo "$1 already exists!"
  exit
fi

firstLine=$(head -n 1 CHANGELOG.md)
if ! echo "$firstLine" | grep -q "$1"; then
  echo "new tag ($1) doesn't match CHANGELOG ($firstLine)"
  exit
fi

expectedHeader=$(printf "## $1 (Released %s)" "$(date +"%Y-%m-%d")")
if [ "$firstLine" != "$expectedHeader" ]
then
  echo "Did you update the CHANGELOG's header? Expected \"$expectedHeader\", found \"$firstLine\""
  exit
fi

git add CHANGELOG.md version.go
git commit -m "release $1"
git tag "$1"
git push origin master
git push origin "$1"