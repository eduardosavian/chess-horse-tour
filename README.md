# Knight's tour algorithms

## Description

Fazer um programa que resolva o problema da volta do cavalo no tabuleiro de xadrez, utilizando métodos de busca. A equipe deve fundamentar o método ou heurística utilizada e explicar como foi implementado.

O usuário deverá informar qual a casa de inicio do passeio do cavalo e o programa indicar todo o restante do percurso, numerando cada casa visitada de 1 (primeira, informada pelo usuário) a 64 (ultima casa visitada).

Gravar um video com a apresentação da solução. Deve ser explicado a heurística utilizada e como ela foi implementada, além da execução do programa.O video deve ser colocado no youtube (ou similar) e o link para visualização entregue no material didático junto com os fontes e executavel da aplicação.

Trabalho pode ser feito em equipes de até 3 integrantes e todos devem participar efetivamente da apresentação.

OBS.: É obrigatório a entrega de uma versão executavel (que possa ser executada em qualquer maquina sem necessidade de compilação ou qualquer outra dependencia).

Exemplo de saida do programa:

![Horse moved](images/horse_moved.jpeg)

Exemplo de movimentação do cavalo no Xadrez:

![Horse movement](images/horse_movement.jpeg)

## Run

### Linux

```bash
go build -o libs/knight_tour src/main.go src/methods.go
```

```bash
pip install -r requirements.txt
```

```bash
pyinstaller --onefile --name knights_tour --distpath exec --add-data "libs:libs" src/menu.py
```

```bash
./knights_tour.exe
```

## Bibliographic references

[An efficient algorithm for the Knight’s tour problem - Ian Parberry](https://core.ac.uk/download/pdf/81964499.pdf)

[Knight's tour - Wikipedua](https://en.wikipedia.org/wiki/Knight's_tour)
