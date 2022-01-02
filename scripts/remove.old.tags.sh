#!/usr/bin/env bash

for tag in $(git tag --sort=v:refname | head -n 28); do
  git push --delete origin "$tag"
  git tag -d "$tag"
done
