# Pokedex Go

A command-line interface (CLI) application written in Go that allows you to explore the Pokémon world, catch Pokémon, and build your own Pokedex using the [PokeAPI](https://pokeapi.co/).

## Features

- **Explore Location Areas**: Browse through different location areas in the Pokémon world
- **Catch Pokémon**: Attempt to catch Pokémon with a probability-based catching system
- **Inspect Pokémon**: View detailed information about caught Pokémon including stats, types, height, and weight
- **Pokedex**: Keep track of all the Pokémon you've caught
- **Caching**: Built-in caching system to reduce API calls and improve performance
- **Pagination**: Navigate through location areas with forward and backward pagination

## Prerequisites

- Go 1.25.1 or higher
- Internet connection (to access the PokeAPI)

## Installation

1. Clone the repository:
```bash
git clone https://github.com/angelchiav/pokedex-go.git
cd pokedex-go
```

2. Build the application:
```bash
go build
```

3. Run the application:
```bash
./pokedex-go
```

Or run directly with Go:
```bash
go run .
```

## Usage

Once the application is running, you'll see a prompt: `Pokedex >`

### Available Commands

- `help` - Displays a help message with all available commands
- `exit` - Exit the Pokedex application
- `map` - Show the next 20 location areas
- `mapb` - Show the previous 20 location areas
- `explore <area_name>` - Show all Pokémon found in a specific location area
- `catch <pokemon_name>` - Attempt to catch a Pokémon (success depends on base experience)
- `inspect <pokemon_name>` - Show detailed information about a caught Pokémon
- `pokedex` - Display all Pokémon you've caught

### Example Usage

```
Pokedex > map
canalave-city-area
eterna-city-area
pastoria-city-area
...

Pokedex > explore canalave-city-area
Pokémon in canalave-city-area (15):
 - tentacool
 - tentacruel
 - magikarp
 ...

Pokedex > catch pikachu
Throwing a Pokeball at pikachu...
pikachu was caught!

Pokedex > inspect pikachu
Name: pikachu
Height: 4 decimeters
Weight: 60 hectograms
Stats:
  -hp: 35
  -attack: 55
  ...
Types:
 - electric

Pokedex > pokedex
Your Pokedex:
 - pikachu
 - charizard
 ...
```

## How Catching Works

The catch probability is calculated based on the Pokémon's base experience:
- Higher base experience = Lower catch chance
- Catch probability ranges from 10% to 90%
- Formula: `0.6 - (base_experience / 800.0)`, clamped between 0.1 and 0.9

## Project Structure

```
pokedex-go/
├── main.go              # Main entry point and REPL loop
├── commands.go          # Command implementations and API interactions
├── go.mod               # Go module definition
├── internal/
│   └── pokecache/
│       ├── cache.go     # Caching implementation with TTL
│       └── cache_test.go # Cache tests
└── README.md            # This file
```

## Technical Details

- **Caching**: The application uses a time-based cache (TTL: 5 seconds) to store API responses and reduce redundant network calls
- **API**: All data is fetched from the [PokeAPI](https://pokeapi.co/) REST API
- **Concurrency**: The cache uses goroutines for automatic cleanup of expired entries

## Testing

Run the tests with:
```bash
go test ./...
```

## License

This project is open source and available for educational purposes.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.


