# Pokédex CLI
An interactive Pokédex you can use right from your terminal.
This project is part of the [Boot.dev guided project series](https://www.boot.dev/courses/build-pokedex-cli-golang).

## Features

- Accesses [PokeAPI](https://pokeapi.co/) for up-to-date Pokémon info
- Caching of regularly accessed data
- Dynamic Pokémon capture rates

## Usage

1. Clone the repo:
   ```sh
   git clone https://github.com/charliej2005/pokedex-cli.git
   cd pokedex-cli
   ```

2. Build the project:
   ```sh
   go build
   ```

3. Run the CLI:
   ```sh
   ./pokedex-cli
   ```

4. **Available commands:**
   - `help` — Show available commands
   - `map` — List next page of Pokémon locations
   - `mapb` — List previous page of Pokémon locations
   - `explore <location>` — List Pokémon in a specific location
   - `catch <pokemon>` — Attempt to catch a Pokémon
   - `inspect <pokemon>` — Show details about a caught Pokémon
   - `pokedex` — Show your caught Pokémon
   - `exit` — Quit the application


## Requirements

- Go 1.18+
- Internet connection (for PokeAPI access)