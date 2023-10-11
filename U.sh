#!/bin/bash
go build .
echo "Go build is run"
for i in {1..5}
do
    echo -n "* "
    sleep 1
done
echo ""
echo "Git is Pushing"
for i in {1..2}
do
    echo -n "* "
    sleep 1
done
echo ""
git add .
git commit -m " ++ "
git push -u origin main 
