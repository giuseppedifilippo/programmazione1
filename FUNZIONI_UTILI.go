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
    X/=10
    C++
  }
  return c
}
