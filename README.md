# Pokemon

Jogo de Pokemon feito em Go para rodar no terminal, com mapas em ASCII, cores ANSI, exploracao por teclado e batalhas contra Pokemon selvagens.

## Funcionalidades atuais

- Menu principal com as opcoes `PLAY`, `ABOUT` e `LEAVE`.
- Exploracao em terminal com personagem representado por `@`.
- Tres areas jogaveis no estado atual:
  casa inicial, mapa externo e laboratorio.
- Interacao com o Professor Silva para escolher o Pokemon inicial.
- Escolha entre `Bulbasaur`, `Charmander` e `Squirtle`.
- Encontros aleatorios em mato alto.
- Batalhas por turnos com:
  exibicao de HP, nivel, tipo e log de combate.
- Sistema de vantagem e desvantagem de tipos no dano.
- Ganho de XP ao derrotar Pokemon selvagens.
- Menu interno durante a exploracao para ver equipe e mochila.
- Retorno para a casa inicial com cura total quando toda a equipe desmaia.

## Botoes, teclas e opcoes

### Menu principal

- `1` -> `PLAY`
- `2` -> `ABOUT`
- `3` -> `LEAVE`

### Exploracao

- `W` -> mover para cima
- `A` -> mover para a esquerda
- `S` -> mover para baixo
- `D` -> mover para a direita
- `/` -> abrir o menu do jogo
- `Enter` -> avancar dialogos e mensagens

### Menu dentro do jogo

- `1` -> `POKEMON`
- `2` -> `BAG`
- `3` -> `CANCEL`
- `4` -> `LEAVE GAME`

### Escolha do inicial no laboratorio

- `1` -> Bulbasaur
- `2` -> Charmander
- `3` -> Squirtle

### Menu de combate

- `1` -> `FIGHT`
- `2` -> `BAG`
- `3` -> `RUN`
- `4` -> `POKEMON`

## Como jogar

1. Execute o projeto.
2. No menu principal, escolha `1` para iniciar.
3. Use `W`, `A`, `S` e `D` para sair da casa e explorar o mapa.
4. Entre no laboratorio e fale com o Professor Silva para pegar seu Pokemon inicial.
5. Va para o mato alto no mapa externo para encontrar Pokemon selvagens.
6. No combate, escolha `FIGHT` para atacar ou `RUN` para tentar fugir.
7. Use `/` durante a exploracao para abrir o menu e consultar sua equipe.
8. Continue explorando e vencendo batalhas para acumular XP.

## Pokemon selvagens atuais

Na rota atual, os encontros aleatorios podem gerar Pokemon entre os niveis `2` e `5`, incluindo:

- Caterpie
- Metapod
- Weedle
- Kakuna
- Pidgey
- Rattata

## Como executar

### Com Go

```bash
go run .
```

### No Windows com executavel gerado

```powershell
.\pokemon.exe
```

## Requisitos

- Go `1.21`
- Terminal com suporte a ANSI colors

## Limitacoes atuais

- A mochila aparece no jogo, mas ainda nao possui uso completo nas batalhas.
- A opcao `POKEMON` dentro do combate ainda nao realiza a troca de Pokemon.
- O jogo ainda esta concentrado em uma area inicial curta, com uma rota e um laboratorio.
- O sistema mostra ganho de XP, mas o ciclo completo de subir de nivel, evoluir e capturar Pokemon ainda nao faz parte da experiencia jogavel atual.
