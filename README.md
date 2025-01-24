
# Gokedex

Gokedex is a fully functional REPL (Read-Eval-Print Loop) and caching system for the PokeAPI, built as part of the Go courses at [Boot.dev](https://boot.dev). It showcases the use of Go to interact with external APIs, implement caching, and create an interactive command-line application.

## Features

- Interactive REPL for querying Pokémon and location data.
- Integration with the PokeAPI to fetch details about Pokémon.
- Local caching to minimize API calls and improve performance.
- Interactive Pokémon catching with random chance based on EXP.
- Using Go concepts such as:
  - HTTP requests and JSON parsing.
  - In-memory caching.
  - Structs, slices, and maps.
  - Handling user input in a REPL environment.

## Installation

1. Clone the repository:
  ```bash
  git clone https://github.com/abtsousa/gokedex.git
  cd gokedex
  ```

2. Run the application:

  ```bash
  go run .
  ```

## Usage

### General Commands
- `help` – Display a help message with available commands.
- `exit` – Exit the Gokedex.

### Exploring the World
- `map` – Show the next 20 locations.
- `mapb` – Show the previous 20 locations.
- `explore` – Show all the Pokémon in the current area.

### Catching and Managing Pokémon
- `catch <pokemon-name>` – Catch a Pokémon by name (e.g., `catch pikachu`).
- `pokedex` – List all caught Pokémon.
- `inspect <location>` – Inspect a specific area for detailed Pokémon information (e.g., `inspect eterna-city-area`).


## Acknowledgments

This project was developed as part of Boot.dev's Go courses.
Data is sourced from the PokeAPI.

## License

MIT
