package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/musztardem/zombic/components"
	"github.com/musztardem/zombic/entities"
	"github.com/musztardem/zombic/systems"
)

const (
	SCREEN_WIDTH  = 320
	SCREEN_HEIGHT = 180
)

type Game struct {
	player  *entities.Player
	spawner *entities.Spawner
	enemies *[]entities.EnemyBehaviour
	missles *[]entities.Missle
}

func (g *Game) Update() error {
	g.player.Update(g.enemies, g.missles)
	g.spawner.Update()

	for _, missle := range *g.missles {
		missle.Update()
	}

	enemyColliders := make([]*components.Collider, 0)
	for _, enemy := range *g.enemies {
		enemyColliders = append(enemyColliders, enemy.GetCollider())
	}

	for _, enemy := range *g.enemies {
		if err := enemy.Update(enemyColliders); err != nil {
			log.Fatalf("failed to update enemy: %v", err)
		}
	}

	systems.ShootAtNearestEnemy(g.player, g.enemies, g.missles)
	systems.MissleHit(g.enemies, g.missles)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw the background
	screen.Fill(color.RGBA{130, 206, 235, 255})

	// Draw the player
	g.player.Draw(screen)

	for _, missle := range *g.missles {
		missle.Draw(screen)
	}

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
			{X: -10.0, Y: -10.0},
			{X: SCREEN_WIDTH + 10.0, Y: -10.0},
			{X: SCREEN_WIDTH + 10.0, Y: SCREEN_HEIGHT + 10.0},
			{X: -10.0, Y: SCREEN_HEIGHT + 10.0},
		},
	}

	spawner := entities.NewSpawner(
		spawnerPath,
		&components.Velocity{Val: 1.0},
		player.Position,
		&enemies,
		0.5,
	)

	missles := make([]entities.Missle, 0)

	game := &Game{
		player:  player,
		spawner: spawner,
		enemies: &enemies,
		missles: &missles,
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
