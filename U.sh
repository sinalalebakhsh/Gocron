#!/bin/bash
go build .
for i in {1..4}
do
    echo -n "* "
    sleep 1
done
echo ""
git add .
git commit -m " ++ "
git push -u origin main 
