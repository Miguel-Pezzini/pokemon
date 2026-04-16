# AGENTS.md

## Objetivo
Este arquivo define o padrao obrigatorio para qualquer contribuicao humana ou automatizada neste repositorio.

O foco e simples:
- zero magic numbers sem significado explicito
- codigo reaproveitavel
- baixo acoplamento
- estrutura preparada para crescer

Se houver conflito entre velocidade e padrao, o padrao vence.

## Regras inegociaveis

### 1. Nao use magic numbers
Todo numero que represente regra, limite, chance, delay, largura, altura, coordenada importante, identificador especial ou comportamento de gameplay precisa ter nome.

Regras:
- Nao deixe numeros "soltos" dentro de funcoes de dominio, UI, mapa ou combate.
- O nome da constante precisa explicar o significado do valor.
- Se o valor varia por contexto, use struct de configuracao em vez de espalhar varias constantes.
- Se o numero representa um identificador de dominio, prefira tipo nomeado em vez de `int` cru.

Excecoes aceitaveis:
- `0`, `1` e `-1` apenas em usos idiomaticos e locais, como incremento de loop, comparacoes triviais, indexacao basica ou inicializacao sem significado de negocio.
- Fora disso, nomeie.

Exemplos deste projeto:
- IDs de tiles como `0`, `1`, `3`, `4`, `5`, `6` e `7` nao devem ficar espalhados no codigo.
- Coordenadas fixas de spawn e transicao entre mapas devem ser nomeadas.
- Larguras de caixa como `50`, `52` e `60` devem virar constantes de UI.
- Chances de encontro e multiplicadores de dano/XP devem ser nomeados ou movidos para configuracao.
- Faixas de nivel por rota devem viver nos dados da rota, nao em condicionais espalhadas.

Bom:

```go
const (
    battleLogWidth = 60
    wildEncounterChanceDenominator = 10
    wildEncounterTriggerValue = 7
)
```

Ruim:

```go
drawTitledBox("BATTLE LOG", messages, 60, colorBlue, colorCyan)
if rand.Intn(10) == 7 { ... }
```

### 2. Escreva codigo para reaproveitamento
O codigo deve ser desenhado para ser usado em mais de um lugar sem depender do contexto atual.

Regras:
- Nenhuma funcao deve depender do "mapa atual", "rota atual" ou "tela atual" se essa informacao puder ser recebida por parametro.
- Nenhuma regra de negocio deve depender diretamente da UI de terminal.
- Evite acesso direto a posicoes especiais como `character.Pokemons[0]` quando a regra correta e "pokemon ativo".
- Evite `switch` com IDs fixos para selecionar comportamento quando um registro, tabela ou configuracao resolve melhor.
- Se uma regra vai se repetir para novas rotas, NPCs, mapas, menus ou batalhas, transforme em estrutura reutilizavel agora.

Bom:

```go
type Position struct {
    Row int
    Col int
}

type MapTransition struct {
    FromTile TileID
    NextState GameState
    Spawn Position
}
```

Ruim:

```go
if optionPath == 3 {
    currentState = MAP_ONE
    x = 16
    y = 15
}
```

### 3. Prefira design orientado a dados
Sempre que o crescimento esperado do jogo indicar repeticao futura, prefira modelar dados em vez de duplicar fluxo.

Use dados/configuracao para:
- definicao de mapas
- tiles e seus significados
- pontos de spawn
- transicoes entre mapas
- encontros por rota
- parametros de batalha
- menus e opcoes
- dialogos reutilizaveis

Se um novo mapa ou nova rota exigir copiar e colar uma funcao inteira, o design ainda esta acoplado demais.

### 4. Separe responsabilidades
Cada bloco do codigo deve ter um motivo claro para mudar.

Diretriz:
- dominio: regras de batalha, progressao, status, itens, rotas, encontros
- orquestracao: loop principal, transicoes de estado, fluxo do jogo
- UI/infra: ANSI, leitura de teclado, desenho, caixas, renderizacao

Evite funcoes que:
- decidem regra de negocio
- alteram estado
- fazem renderizacao
- leem input

tudo ao mesmo tempo.

Quando necessario, quebre em helpers menores com responsabilidade unica.

### 5. Use tipos e nomes que expressem o dominio
Prefira tipos nomeados a valores numericos crus.

Exemplos recomendados:

```go
type TileID int
type RouteID int
type MenuOption int
type EncounterChance struct {
    Denominator int
    TriggerValue int
}
```

Regras:
- O nome deve explicar o papel do valor no jogo.
- Nao crie constantes genericas como `defaultValue`, `tempNumber`, `specialCase`.
- Nome ruim com constante ruim continua sendo magic number disfarcado.

### 6. Parametros de configuracao devem ficar perto do dominio
Cada conjunto de regras deve ter sua propria configuracao declarada perto do modulo que representa esse dominio.

Exemplos:
- combate: delays, largura de barras, formulas e multiplicadores
- mapa: tamanho, tiles, spawn, saidas
- rota: especies possiveis, faixa de nivel, chance de encontro
- UI: larguras padrao, textos fixos, atalhos

Evite colocar configuracoes importantes escondidas dentro de funcoes grandes.

### 7. Preparar para escala tambem significa facilitar teste
Sempre que fizer sentido, injete dependencias em vez de fixar implementacoes.

Exemplos:
- RNG em combate e encontros
- relogio para delays e animacoes
- input do teclado
- renderer de terminal

Isso reduz acoplamento e facilita testar regras sem depender da tela.

## Convencoes praticas para este repositorio

### Mapas
- Nao usar inteiros crus para representar tile.
- Nao embutir coordenadas de transicao em varios pontos do codigo.
- Centralizar pontos de entrada/saida em estruturas nomeadas.
- Se um mapa puder ser descrito por dados, prefira isso a fluxo manual.

### Combate
- Nao deixar multiplicadores, divisores e chances espalhados na formula.
- Nomear cada parte importante da regra.
- Separar calculo de dano, aplicacao de dano, animacao e renderizacao.

### Pokemon e equipe
- Nao assumir que o pokemon ativo sempre esta no indice `0`.
- Tornar o conceito de pokemon ativo explicito.
- Regras de equipe devem operar em colecoes e estado, nao em posicoes implicitas.

### UI de terminal
- Larguras, delays, textos especiais e teclas devem ser constantes nomeadas ou configuracao.
- A UI nao deve definir regra de negocio.

## Checklist antes de concluir uma alteracao
- Existe algum numero com significado de dominio ainda solto no codigo?
- Esse codigo consegue ser usado em outro mapa, rota, tela ou fluxo sem copiar e colar?
- A regra esta orientada a dados ou presa ao caso atual?
- A funcao esta acumulando responsabilidade demais?
- O nome das constantes e tipos explica o motivo do valor existir?
- A regra de negocio esta separada da renderizacao e da leitura de input?

## Regra final
Se um valor ou comportamento parece especifico demais para o lugar onde esta, extraia.
Se um trecho so funciona porque "hoje o jogo usa assim", desacople.
Se um numero precisa ser explicado em comentario, ele provavelmente deveria ser uma constante nomeada ou um campo de configuracao.
