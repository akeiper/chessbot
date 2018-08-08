package game

import (
	"fmt"
	"math/rand"

	"github.com/notnil/chess"
)

type Color string

const (
	White Color = "White"
	Black Color = "Black"
)

type Player struct {
	ID string
}

type Game struct {
	game    *chess.Game
	Players map[Color]Player
}

func NewGame(players ...Player) *Game {
	gm := Game{
		game: chess.NewGame(chess.UseNotation(chess.LongAlgebraicNotation{})),
	}
	randomOrder := make([]Player, len(players))
	perm := rand.Perm(len(players))
	for i, v := range perm {
		randomOrder[v] = players[i]
	}
	gm.Players = make(map[Color]Player)
	gm.Players[White] = randomOrder[0]
	gm.Players[Black] = randomOrder[1]
	return &gm
}

func NewGameFromFEN(fen string) (Game, error) {
	gameState, err := chess.FEN(fen)
	if err != nil {
		return Game{}, err
	}
	return Game{
		game: chess.NewGame(gameState, chess.UseNotation(chess.LongAlgebraicNotation{})),
	}, nil
}

func (g Game) Turn() Color {
	switch g.game.Position().Turn() {
	case chess.White:
		return White
	case chess.Black:
		return Black
	default:
		return White
	}
}

func (g Game) FEN() string {
	return g.game.FEN()
}

func (g Game) PGN() string {
	return g.game.String()
}

func (g Game) Outcome() chess.Outcome {
	return g.game.Outcome()
}

func (g Game) ResultText() string {
	return fmt.Sprintf("Game completed. %s by %s.", g.Outcome(), g.game.Method())
}

func (g Game) Move(gridMove string) error {
	return g.game.MoveStr(gridMove)
}

func (g Game) ValidMoves() []*chess.Move {
	return g.game.ValidMoves()
}

func (g Game) String() string {
	return g.game.Position().Board().Draw()
}
