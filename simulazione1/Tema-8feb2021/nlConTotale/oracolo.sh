#!/bin/bash

FILE=$1

TOT=$(wc -l $1|cut -f1 -d" ")

#echo $TOT

LINEA=1
cat $1 | while
 read riga
do
 echo -e "$LINEA/$TOT:\t"$riga
 let LINEA=LINEA+1
done
