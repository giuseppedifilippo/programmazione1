#!/bin/bash
# OLD ./reflow 100 lorem.input

echo
echo
echo
echo === WIDTH:40 === SX oracolo === DX studente ===
if
 diff -y -W 100 <(./oracolo.sh 40 lorem.input) <(echo NIL | ./reflow  40 lorem.input)
then
 echo "=== NON ci sono differenze ==="
else
 echo "=== ATTENZIONE: ci sono differenze con l'oracolo! ==="
fi



echo
echo
echo
echo === WIDTH:20 === SX oracolo === DX studente ===
if
 diff -y -W 80 <(./oracolo.sh 20 lorem.input) <(echo NIL | ./reflow  20 lorem.input)
then
 echo "=== NON ci sono differenze ==="
else
 echo "=== ATTENZIONE: ci sono differenze con l'oracolo! ==="
fi



echo
echo
echo
echo === WIDTH:70 === SX oracolo === DX studente ===
if
 diff -y -W 150 <(./oracolo.sh 70 lorem.input) <(echo NIL | ./reflow  70 lorem.input)
then
 echo "=== NON ci sono differenze ==="
else
 echo "=== ATTENZIONE: ci sono differenze con l'oracolo! ==="
fi