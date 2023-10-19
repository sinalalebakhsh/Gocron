#!/bin/bash
git checkout beforeMergeToMain
echo -n "git checkout beforeMergeToMain "
for i in {1..3}
do
    echo -n ". "
    sleep 1
done
echo ""
# -------------------------------------------------------------------
git merge develop
echo -n "git merge develop "
for i in {1..3}
do
    echo -n ". "
    sleep 1
done
echo ""
# -------------------------------------------------------------------
echo -n "Git is Push for beforeMergeToMain "
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
