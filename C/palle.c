#include <stdio.h>

int main() {
    int n ;
    scanf("%d", &n);
    int x =4;
    if (n == x ) {printf("hai indovinato\n");}
    else if (n < 1 || n > 10 ){printf("valore non valido\n");}
    else {printf("non hai indovinato\n");}
    return 0;
} 