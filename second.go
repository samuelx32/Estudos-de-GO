package main
import "fmt"

type Produto struct{
    nome string
    preco float32
    qtn int
}

func main() {
    var produtos []Produto
    var n string
    var p float32
    var q int
        
    for i:=0; i<5; i++ {
        
        fmt.Scan(&n)
        fmt.Scan(&p)
        fmt.Scan(&q)
        produtos = append(produtos, Produto{n,p,q})
    }

  fmt.Println(produtos)
}