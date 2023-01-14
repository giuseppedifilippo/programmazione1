#!/bin/bash

testa() {
    ORACOLO=$1
    STUDENTE=$2
    DIFFW=$3

    echo
    echo === SINISTRA oracolo === DESTRA studente ===
    if
        diff -y -W $DIFFW <($1) <($2)
    then
        echo
        echo "=== NON ci sono differenze ==="
    else
        echo
        echo "=== ATTENZIONE: ci sono differenze con l'oracolo! ==="
    fi
}

testa "./oracolo.sh 3 uno.input" "./tail 3 uno.input" 100

