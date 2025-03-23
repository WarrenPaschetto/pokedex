# ⚡💥 Pokedex CLI

A terminal-based Pokémon explorer written in Go!  
Search regions, explore Pokémon locations, catch and inspect Pokémon, and manage your own Pokedex — all from your command line.

## 📦 Project Structure

This project consists of a modular architecture using Go packages:

- **`main/`** – Handles REPL loop, user input, and commands.
- **`internal/pokeapi/`** – Interfaces with the public [PokeAPI](https://pokeapi.co) to retrieve Pokémon data.
- **`internal/pokecache/`** – Provides a caching mechanism to reduce API calls and improve performance.

## 🔧 Installation

1. **Clone the repository**
```
git clone https://github.com/WarrenPaschetto/pokedex.git
cd pokedex
```
2. **✏️ Edit go.mod file**

  Be sure to enter in your Github username for module link and enter the proper Go version you are using
```
module github.com/{enter your Github username here}/pokedexcli
go 1.24.1
```
   
4. **🚀 Run the App**

```
go build && ./pokedexcli
```

⚠️ Requires Go 1.18+

## 🕹️ CLI Commands

Enter commands in the REPL interface (prompt: Pokedex >) to interact with the Pokémon world.

![pokedexChart](https://github.com/user-attachments/assets/d2e49915-5b0b-4c7f-9493-f6aa7f9ceae6)

## 🧪 Example
```
Pokedex > map
Location areas:
canalave-city-area
eterna-city-area
pastoria-city-area
sunyshore-city-area
sinnoh-pokemon-league-area
oreburgh-mine-1f
oreburgh-mine-b1f
valley-windworks-area
eterna-forest-area
fuego-ironworks-area
mt-coronet-1f-route-207
mt-coronet-2f
mt-coronet-3f
mt-coronet-exterior-snowfall
mt-coronet-exterior-blizzard
mt-coronet-4f
mt-coronet-4f-small-room
mt-coronet-5f
mt-coronet-6f
mt-coronet-1f-from-exterior
Pokedex > explore eterna-city-area
Pokemon in eterna-city-area:
psyduck
golduck
magikarp
gyarados
barboach
whiscash
Pokedex > catch golduck
....Adding to cache
Throwing a Pokeball at golduck...:
golduck escaped!
Pokedex > catch golduck
Cache hit!!
Throwing a Pokeball at golduck...:
golduck escaped!
Pokedex > catch golduck
Cache hit!!
Throwing a Pokeball at golduck...:
golduck was caught!
You may now inspect it with the inspect command.
Pokedex > inspect golduck
Name: golduck
Weight: 766
Stats:
 -hp: 80
 -attack: 82
 -defense: 78
 -special-attack: 95
 -special-defense: 80
 -speed: 85
Types:
 -water
Pokedex > pokedex
Your Pokedex:
 -golduck
Pokedex > exit
Closing the Pokedex... Goodbye!
```

## ✨ Features
- 🌍 Browse Pokémon location areas with pagination
- 🐾 Explore wild Pokémon by location
- 🎣 Catch Pokémon using randomized mechanics based on their difficulty
- 🧾 Inspect Pokémon stats and types
- 🧠 Caching system to avoid redundant API calls and improve speed

## 🧪 Tests
Run tests using:
```
go test ./...
```

