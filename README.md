API de criação de times de pokemon utilizando pokeapi.co

Inicie o projeto com:
```sh
docker compose up
```

- Rotas
  - GET /api/teams - Lista todos os times registrados
  - GET /api/teams/{nomeDoDono} - Busca um time registrado pelo nome do dono do time
  - POST /api/teams - Rota para criação de um time
  ```JSON
    {
      "user": "exemplo",
      "team": [
        "blastoise",
        "pikachu",
        "charizard",
        "venusaur",
        "lapras",
        "dragonite"
      ]
    }
    ```

