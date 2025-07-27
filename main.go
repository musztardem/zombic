package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/musztardem/zombic/components"
	"github.com/musztardem/zombic/entities"
)

const (
	SCREEN_WIDTH  = 320
	SCREEN_HEIGHT = 180
)

type Game struct {
	player  *entities.Player
	spawner *entities.Spawner
	enemies *[]entities.EnemyBehaviour
}

func (g *Game) Update() error {
	g.player.Update()
	g.spawner.Update()

	enemyColliders := make([]*components.Collider, 0)
	for _, enemy := range *g.enemies {
		enemyColliders = append(enemyColliders, enemy.GetCollider())
	}

	for _, enemy := range *g.enemies {
		if err := enemy.Update(enemyColliders); err != nil {
			log.Fatalf("failed to update enemy: %v", err)
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw the background
	screen.Fill(color.RGBA{130, 206, 235, 255})

	// Draw the player
	g.player.Draw(screen)

	// Draw enemies
	for _, enemy := range *g.enemies {
		enemy.Draw(screen)
	}
}

func (g *Game) Layout(_, _ int) (int, int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}

func main() {
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Zombic")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	player := entities.NewPlayer(
		&components.Position{
			X: 150.0,
			Y: 80.0,
		},
		&components.Velocity{
			Val: 1.0,
		},
	)

	enemies := make([]entities.EnemyBehaviour, 0)

	spawnerPath := &components.Path{
		Points: []*components.Position{
			{X: -5.0, Y: -5.0},
			{X: SCREEN_WIDTH + 5.0, Y: -5.0},
			{X: SCREEN_WIDTH + 5.0, Y: SCREEN_HEIGHT + 5.0},
			{X: -5.0, Y: SCREEN_HEIGHT + 5.0},
		},
	}

	spawner := entities.NewSpawner(
		spawnerPath,
		&components.Velocity{Val: 1.0},
		player.Position,
		&enemies,
		3,
	)

	game := &Game{
		player:  player,
		spawner: spawner,
		enemies: &enemies,
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
