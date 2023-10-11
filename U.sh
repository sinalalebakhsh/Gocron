#!/bin/bash
go build .
for i in {1..4}
do
    echo -n "* "
    sleep 1
done
git add .
git commit -m " ++ "
git push -u origin main 
