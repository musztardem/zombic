# Zombic

A top-down zombie shooter game built with Go and Ebiten.

## About

Zombic is a 2D survival shooter where you fight against waves of zombies. The game features:

- **Multiple zombie types**: BigZombie, KidZombie, SkinnyZombie, and TurretZombie
- **Automatic shooting system** that targets the nearest enemy
- **Wave-based spawning** with enemies spawning from the edges of the screen
- **Player movement** and collision detection
- **Missile/projectile system** for combat

## Engine

This game is built using [Ebiten](https://ebiten.org/), a dead simple 2D game engine for Go. Ebiten provides:
- Cross-platform support (Windows, macOS, Linux, and more)
- Easy-to-use 2D graphics rendering
- Input handling
- Audio support

## Requirements

- Go 1.24.4 or higher

## Installation & Running

1. **Clone the repository**:
```bash
git clone https://github.com/musztardem/zombic.git
cd zombic
```

2. **Install dependencies**:
```bash
go mod download
```

3. **Run the game**:
```bash
go run main.go
```

Alternatively, you can build an executable:
```bash
go build -o zombic
./zombic
```

## Game Controls

(Add your control scheme here)

## Project Structure

- `main.go` - Main game loop and initialization
- `components/` - ECS components (position, velocity, collider, etc.)
- `entities/` - Game entities (player, enemies, weapons, missiles)
- `systems/` - Game systems (shooting, collision detection, etc.)
- `assets/` - Game assets (sprites, maps, tilesets)
- `images/` - Image loading utilities
- `config/` - Game configuration
- `vectors/` - Vector math utilities
