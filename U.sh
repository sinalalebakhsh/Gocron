#!/bin/bash
go build .
git add .
git commit -m " ++ "
git push -u origin main 
