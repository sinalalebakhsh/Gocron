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
git checkout beforeMergeToMain
echo ""
echo "====================================================================="


echo -n "git merge develop "
for i in {1..3}
do
    echo -n ". "
    sleep 1
done
git merge develop
echo ""
echo "====================================================================="


echo -n "Git is Push for beforeMergeToMain "
echo ""
for i in {1..3}
do
    echo -n ". "
    sleep 1
done
git add .
git commit -m " ++ "
git push -u origin beforeMergeToMain
echo "====================================================================="


git checkout develop
echo ""
echo "====================================================================="



