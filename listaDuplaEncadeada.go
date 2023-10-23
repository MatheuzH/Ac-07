package main

import "fmt"


type No struct {
    Valor    int
    Prox     *No
    Ant *No
}


type ListaDuplaEncadeada struct {
    Cabeca *No
}


func CriarNo(valor int) *No {
    return &No{Valor: valor, Prox: nil, Ant: nil}
}


func ExibirValoresOrdenados(l *ListaDuplaEncadeada, v int) {
    if l.Cabeca == nil {
        fmt.Println("Lista vazia, nada para exibir.")
    } else {
        cabeca := l.Cabeca
        for cabeca != nil {
            if cabeca.Valor >= v {
                fmt.Printf("Valor encontrado: %d\n", cabeca.Valor)
                return
            }
            fmt.Printf("%d -> ", cabeca.Valor)
            cabeca = cabeca.Prox
        }
    }
}


func EncontrarNoDoMeio(inicio, fim *No) *No {
    for inicio != fim && inicio.Prox != fim {
        inicio = inicio.Prox
        fim = fim.Ant
    }
    return inicio
}


func BuscaBinaria(l *ListaDuplaEncadeada, v int) string {
    if l.Cabeca == nil {
        return "Lista vazia, valor não encontrado."
    } else {
        inicio := l.Cabeca
        final := EncontrarUltimoNo(l)

        for inicio != nil && inicio != final {
            meio := EncontrarNoDoMeio(inicio, final)

            if meio.Valor == v {
                return "Valor encontrado."
            } else if meio.Valor > v {
                final = meio.Ant
            } else {
                inicio = meio.Prox
            }
        }

        return "Valor não encontrado."
    }
}


func EncontrarUltimoNo(l *ListaDuplaEncadeada) *No {
    cabeca := l.Cabeca
    for cabeca.Prox != nil {
        cabeca = cabeca.Prox
    }
    return cabeca
}


func InserirNoOrdenado(l *ListaDuplaEncadeada, v int) {
    novoNo := CriarNo(v)

    if l.Cabeca == nil {
        l.Cabeca = novoNo
    } else if v <= l.Cabeca.Valor {
        novoNo.Prox = l.Cabeca
        l.Cabeca.Ant = novoNo
        l.Cabeca = novoNo
    } else {
        cabeca := l.Cabeca
        for cabeca.Prox != nil && cabeca.Prox.Valor < v {
            cabeca = cabeca.Prox
        }

        if cabeca.Prox == nil {
            novoNo.Ant = cabeca
            cabeca.Prox = novoNo
        } else {
            novoNo.Ant = cabeca
            novoNo.Prox = cabeca.Prox
            cabeca.Prox.Ant = novoNo
            cabeca.Prox = novoNo
        }
    }
}


func RemoverValor(l *ListaDuplaEncadeada, v int) {
    if l.Cabeca == nil {
        return
    }

    cabeca := l.Cabeca
    for cabeca != nil && cabeca.Valor != v {
        cabeca = cabeca.Prox
    }

    if cabeca == nil {
        return 
    } else if cabeca.Ant == nil {
        l.Cabeca = cabeca.Prox
        if cabeca.Prox != nil {
            cabeca.Prox.Ant = nil
        }
    } else {
        cabeca.Ant.Prox = cabeca.Prox
        if cabeca.Prox != nil {
            cabeca.Prox.Ant = cabeca.Ant
        }
    }
}

func main() {
    l := &ListaDuplaEncadeada{}
    valorOrdenado := 4

    InserirNoOrdenado(l, 3)
    InserirNoOrdenado(l, 6)
    InserirNoOrdenado(l, 2)
    InserirNoOrdenado(l, 5)
    InserirNoOrdenado(l, 1)

    fmt.Printf("Valores ordenados maiores ou iguais a %d: \n", valorOrdenado)
    ExibirValoresOrdenados(l, valorOrdenado)
    
    fmt.Println(BuscaBinaria(l, 5)) 

    RemoverValor(l, 3)

    fmt.Println("Lista após a remoção:")
    cabeca := l.Cabeca
    for cabeca != nil {
        fmt.Printf("%d -> ", cabeca.Valor)
        cabeca = cabeca.Prox
    }
}
