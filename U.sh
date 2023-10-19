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


git checkout beforeMergeToMain
echo -n "git checkout beforeMergeToMain "
for i in {1..3}
do
    echo -n ". "
    sleep 1
done
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
<<<<<<< HEAD
<<<<<<< HEAD
=======
echo ""
=======
>>>>>>> develop
echo "====================================================================="


git checkout develop
echo ""
echo "====================================================================="



>>>>>>> develop
