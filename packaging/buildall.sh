#!/usr/bin/env bash

cd "${0%/*}"

for script in ./*/build; do
    bash "$script"
done