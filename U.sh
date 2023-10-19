#!/bin/bash
echo -n "Git push on branch develop "
for i in {1..3}
do
    echo -n ". "
    sleep 1
done
echo ""
git add .
git commit -m " ++ "
git push -u origin develop
echo "====================================================================="
echo -n "git checkout beforeMergeToMain "
for i in {1..3}
do
    echo -n ". "
    sleep 1
done
echo ""
git checkout beforeMergeToMain
echo "====================================================================="
echo -n "git merge develop "
for i in {1..3}
do
    echo -n ". "
    sleep 1
done
echo ""
git merge develop
echo "====================================================================="
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
echo "====================================================================="
echo -n "git checkout develop "
for i in {1..3}
do
    echo -n ". "
    sleep 1
done
echo ""
git checkout develop
echo "====================================================================="



