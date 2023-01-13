#include <stdio.h>
#include <stdlib.h>
#include <time.h>
void stampa_array(int *a, int n ) {
    printf("[");
    for (int i = 0; i<n; i++) printf("%d",a[i]);
    printf("]")
    
}
void riempi_array(int *a, int n, int q) {
    for (int i = 0 ; i<n; i++) a[i] = rand() %q;
}
int main (int argc, char **argv) {
    //argv[i] e os.aargs di golang 
    srand(time(NULL))
    int n ,q ;
    n = atoi(argv[1]);
    q = atoi(argv[2]);

    int *a ;
    riempi_array(a,n,q);
    stampa_array(a, n);
    a = (int *)malloc(n * sizeof(int));

    //printf("n=%d, q= %d", n ,q );
    return 0;
}