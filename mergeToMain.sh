#!/bin/bash
echo -n "Git push on branch develop "
echo ""
git add .
git commit -m " 💠💠💠 Add feature for storing user input and search it in Net - this is under developing, not complete yet.  "
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
git commit -m " 💠💠💠 Add feature for storing user input and search it in Net - this is under developing, not complete yet.  "
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
git commit -m " 💠💠💠 Add feature for storing user input and search it in Net - this is under developing, not complete yet.  "
git push -u origin main
echo "====================================================================="
echo -n "git checkout develop "
echo ""
git checkout develop
echo "====================================================================="


