#!/bin/bash
go build .
sleep 5
git add .
git commit -m " ++ "
git push -u origin main 
