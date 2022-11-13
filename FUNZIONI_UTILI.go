//funzioni utili per qualsiasi evenienza


//DICE SE UN NUMERO è PRIMO OPPURE NO
func isPrime(x int ) bool {
  for i:=2; i<=x; i++ {
    if x%i==2 {
      return false
    }
  }
  return true
}

//RITORNA LA SOMMA DI TUTTI I DIVISORI PROPRI D IUN NUMERO
func div(x int ) int {
  var c int
  for i:=1; i<x; i++ {
    if x%i==0
    c+=i
  }
  return c
}

// RITORNA IL NUMERO DI CIFRE DI UN NUMERO INTERO POSITIVO
func digit(x int) int {
  var c int
  for x>0 {
    x/=10
    c++
  }
  return c
}

//RITORNA UN SLICE COMPOSTO DALLE CIFRE DI UN NUMERO INTER POSITIVO
func numArr(x int ) []int {
for
}// incompleta

//RITORNA SE UN NUMERO è BINARIO O NO(composto solo da 0 e 1)
func isBin(x int) bool {
for x>0 {
  if !(x%10==0 || x%10==1) {
    return false
  }
  x/=10
}
return true
}

//RITORNA UN RANGETABLE DI UNA STRINGA(COMPOSTA DAI VALORI UNICODE DI OGNI CARATTERE)
