package main

import "fmt"
// import "math"

func main() {
  fmt.Println("50 million primes")
  prime := make([]int, 1, 1000000)
  multiple := make([]int, 1, 1000000)
  prime[0] = 2
  multiple[0] = 2 * 2

  for x := 3; ; x++ {
	  if x < multiple[0] {
	    prime = append(prime, x)
	    multiple = append(multiple, x * x)
		  // fmt.Println(x, prime, multiple)
	    // if len(prime) % 100000 == 0 {
	    //   fmt.Println( 
	    //   	"prime:", x, 
	    //   	"#primes:", len(prime), 
	    //   	"P(n) / (n/ln(n))", float64(len(prime)) / (float64(x) / math.Log(float64(x))),
	    //   )
	    // }
	    if len(prime) == 50000000 { return }
	  } else {
	    for x == multiple[0] {
	      multiple[0] = multiple[0] + prime[0]
	      i := 0;
	      for {
	        test := multiple[i]
	        index := 2 * i + 1;
	        if index >= len(multiple) { break }
	        left := multiple[index];
	        if index + 1 < len(multiple) {
	          right := multiple[index + 1];
	          if test > right || test > left {
	            if right < left { index = index + 1 }
	          } else {
	          	break
	          }
	        } else {
	         	if test < left { break }
          }         
          prime[i], prime[index] = prime[index], prime[i]  
          multiple[i], multiple[index] = multiple[index], multiple[i]  
          i = index
	      }
	    }
	  }
	}
}

/*
// binary tree for faster re-sorting
var prime = [], multiple = [];
prime[0] = 2;
multiple[0] = 2 * 2;

function swap(i,j){
  var temp = prime[i];
  prime[i] = prime[j];
  prime[j] = temp;
  temp = multiple[i];
  multiple[i] = multiple[j];
  multiple[j] = temp;
}

for ( var x = 3; ; x++ ) {
  // console.log(x, prime, multiple);
  if (x < multiple[0]) {
    var firstMultiple = x * x;
    if (firstMultiple > Number.MAX_SAFE_INTEGER) {
      multiplier = Math.floor(Number.MAX_SAFE_INTEGER / x);
      if ( multiplier > 1 ){
        firstMultiple = multiplier * x;
      } else {
        console.log( 'prime:', x, '#primes:', prime.length, 'P(n) / (n/ln(n))', prime.length / (x / Math.log(x)) );
        process.exit(0);
      }
    }
    prime.push(x)
    multiple.push(firstMultiple) // this will always be the largest, right?
    if (prime.length % 100000 === 0){
      console.log( 'prime:', x, '#primes:', prime.length, 'P(n) / (n/ln(n))', prime.length / (x / Math.log(x)) );
    }
  } else {
    while (x === multiple[0]){
      // bump the multiple for the lowest
      multiple[0] += prime[0]
      // swap with smallest child down tree
      var i = 0;
      while ( true ){
        var test = multiple[i]
        var index = 2 * i + 1;
        var left = multiple[index];
        var right = multiple[index + 1];
        if ( test > right || test > left ){
          if ( left < right) {
            swap(i, index);
            i = index;
          } else {
            swap(i, index + 1);
            i = index + 1
          }
        } else {
          break;
        }
      }
    }
  }
}
*/