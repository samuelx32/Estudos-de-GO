package main
import "fmt"

func main() {
  var notas = []int{5,3,9,3,4,3,6,3,7,8}
  //var moda int
  //var mediana int
  //var desvio_medio float32
  //var variancia float32
  //var desvio_padrao float32
  
  ordenando(notas)
  fmt.Println(notas)
  fmt.Println("Média: ", media(notas))
  fmt.Println("Moda: ",moda(notas))
  mediana(notas) 
    
}

func media(notas [] int) float32{
    var soma int
    for i:=0;i<len(notas);i++ {
      soma += notas[i]
    }
    return float32(soma) / float32(len(notas))
}

func mediana(notas[]int) float32{
    meio := len(notas)/2
    var aux float32
    if len(notas) % 2 == 0{
        return (float32(notas[meio]) + float32(notas[meio+1])) / 2
    }else{
        aux = int((float32(meio) + 0.5))
        
        return float32(notas[aux])
    }
    
}

func moda(notas [] int) int{
    aux := 0
    aux2 := 0
    var moda int
    for i:=0;i<len(notas);i++{
        for x:=1;x<len(notas);x++{
            if notas[i] == notas[x]{
                aux2 += 1
            }
        }
        if aux2 > aux{
            aux = aux2
            moda = notas[i]
        }
        aux2 = 0
    }
    return moda
}

func ordenando(notas []int) int {
    var aux int
    for i:=0;i<len(notas);i++ {
      for x:=0;x<len(notas);x++{
          if notas[i] < notas[x]{
              aux = notas[x]
              notas[x] = notas[i]
              notas[i] = aux
          }
      }
    }
    return 0
}












