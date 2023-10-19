#!/bin/bash
git checkout beforeMergeToMain
git merge develop
go build .
echo -n "Go build is run for 'beforeMergeToMain' branch "
for i in {1..3}
do
    echo -n ". "
    sleep 1
done
echo ""
echo -n "Git is Push for 'beforeMergeToMain' branch "
for i in {1..3}
do
    echo -n ". "
    sleep 1
done
echo ""
git add .
git commit -m " ++ "
git push -u origin beforeMergeToMain 
git checkout main
git merge beforeMergeToMain
go build .
echo -n "Go build is run for 'main' branch "
for i in {1..3}
do
    echo -n ". "
    sleep 1
done
echo ""
echo -n "Git is Push for 'main' branch "
for i in {1..3}
do
    echo -n ". "
    sleep 1
done
echo ""
git add .
git commit -m " ++ "
git push -u origin main 
git checkout develop
