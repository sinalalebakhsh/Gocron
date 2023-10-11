#!/bin/bash
go build .
sleep 1
echo "*"
sleep 1
echo "*"
sleep 1
echo "*"
sleep 1
echo "*"
git add .
git commit -m " ++ "
git push -u origin main 
