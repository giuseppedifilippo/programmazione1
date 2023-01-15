#!/bin/bash

LARGH=$1
FILE=$2

#cat $FILE|tr "\n" " " | fold -w $1
cat $FILE|tr -s "\n" " " | fold -s -w $1
echo
