#!/bin/bash
echo -n "Git push on branch develop "
for i in {1..3}
do
    echo -n ". "
    sleep 0.09
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
    sleep 0.09
done
echo ""
git checkout beforeMergeToMain
echo "====================================================================="
echo -n "git merge develop "
for i in {1..5}
do
    echo -n ". "
    sleep 0.09
done
echo ""
git merge develop
echo "====================================================================="
echo -n "Git is Push for beforeMergeToMain "
for i in {1..3}
do
    echo -n ". "
    sleep 0.09
done
echo ""
git add .
git commit -m " ++ "
git push -u origin beforeMergeToMain
echo "====================================================================="
echo -n "git checkout main "
for i in {1..3}
do
    echo -n ". "
    sleep 0.09
done
echo ""
git checkout main
echo "====================================================================="
echo -n "git merge beforeMergeToMain "
for i in {1..5}
do
    echo -n ". "
    sleep 0.09
done
echo ""
git merge beforeMergeToMain
echo "====================================================================="
go build .
echo -n "Go Build "
for i in {1..6}
do
    echo -n "*"
    sleep 0.09
done
echo ""
echo "====================================================================="
echo -n "Git is Push for main "
for i in {1..3}
do
    echo -n ". "
    sleep 0.09
done
echo ""
git add .
git commit -m " ++ "
git push -u origin main
echo "====================================================================="
echo -n "git checkout develop "
for i in {1..3}
do
    echo -n ". "
    sleep 0.09
done
echo ""
git checkout develop
echo "====================================================================="



