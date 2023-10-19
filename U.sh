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
