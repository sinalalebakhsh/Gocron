#!/bin/bash
git checkout main
git merge beforeMergeToMain
go build .
echo "Go build is run"
for i in {1..5}
do
    echo -n "* "
    sleep 1
done
echo ""
echo "Git is Push"
for i in {1..5}
do
    echo -n "* "
    sleep 1
done
echo ""
git add .
git commit -m " ++ "
git push -u origin main 
