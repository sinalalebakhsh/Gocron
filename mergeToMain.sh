#!/bin/bash
echo -n "Git push on branch develop "
echo ""
git add .
git commit -m " 💠💠💠 Adding There are complete => Numbers for Keys Map until 480  "
git push -u origin develop
echo "====================================================================="
echo -n "git checkout beforeMergeToMain "
echo ""
git checkout beforeMergeToMain
echo "====================================================================="
echo -n "git merge develop "
echo ""
git merge develop
echo "====================================================================="
echo -n "Git is Push for beforeMergeToMain "
echo ""
git add .
git commit -m " 💠💠💠 Adding There are complete => Numbers for Keys Map until 480  "
git push -u origin beforeMergeToMain
echo "====================================================================="
echo -n "git checkout main "
echo ""
git checkout main
echo "====================================================================="
echo -n "git merge beforeMergeToMain "
echo ""
git merge beforeMergeToMain
echo "====================================================================="
go build .
echo -n "Go Build "
echo ""
echo "====================================================================="
echo -n "Git is Push for main "
echo ""
git add .
git commit -m " 💠💠💠 Adding There are complete => Numbers for Keys Map until 480  "
git push -u origin main
echo "====================================================================="
echo -n "git checkout develop "
echo ""
git checkout develop
echo "====================================================================="


