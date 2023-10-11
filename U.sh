#!/bin/bash
go build .
sleep 3
git add .
git commit -m " ++ "
git push -u origin main 
