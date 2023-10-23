package main

import "fmt"


type No struct {
    Valor  int
    Prox  *No
}


type ListaCircular struct {
    Cabeca *No
}


func CriaNo(valor int) *No {
    return &No{Valor: valor, Prox: nil}
}

func InserirNoInicio(l *ListaCircular, valor int) {
    novoNo := CriaNo(valor)

    if l.Cabeca == nil {
        l.Cabeca = novoNo
        novoNo.Prox = novoNo
    } else {
        novoNo.Prox = l.Cabeca
        antigoNo := EncontraUltimoNo(l)
        l.Cabeca = novoNo
        antigoNo.Prox = novoNo
    }
}

func EncontraUltimoNo(l *ListaCircular) *No {
    cabeca := l.Cabeca

    if cabeca == nil {
        return nil
    }

    for cabeca.Prox != l.Cabeca{
        cabeca = cabeca.Prox
    }

    return cabeca
}


func ExcluirPrimeiroNo(l *ListaCircular) {
    if l.Cabeca == nil {
        return
    } else if l.Cabeca.Prox == l.Cabeca {
        d := l.Cabeca
        l.Cabeca = nil
        fmt.Println("\n Nó removido:", d.Valor)
    } else {
        UltimoNo := EncontraUltimoNo(l)
        PrimeiroNo := l.Cabeca
        l.Cabeca = PrimeiroNo.Prox
        UltimoNo.Prox = l.Cabeca
        fmt.Println("\n Nó removido:", PrimeiroNo.Valor)
    }
}

func main() {
    lista := &ListaCircular{}

    InserirNoInicio(lista, 3)
    InserirNoInicio(lista, 2)
    InserirNoInicio(lista, 1)
    InserirNoInicio(lista, 7)
    InserirNoInicio(lista, 4)
    InserirNoInicio(lista, 9)
    
    l := lista.Cabeca
    for {
        fmt.Printf("%d -> ", l.Valor)
        l = l.Prox
        if l == lista.Cabeca {
            break
        }
    }

   	ExcluirPrimeiroNo(lista)

    l = lista.Cabeca
    for {
        fmt.Printf("%d -> ", l.Valor)
        l = l.Prox
        if l == lista.Cabeca {
            break
        }
    }
}