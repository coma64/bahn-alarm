<p align="center">
    <img style="margin: auto" alt="Bahn Alarm Logo" height="100" src="frontend/src/assets/logo.svg">
</p>

Bahn alarm is a pet project of mine that sends real-time notifications about train delays from the deutsche bahn.
It does this by regularly polling the API of the next "Next DB Navigator" app.

**I am by no means affiliated with the DB. This project was built for educational purposes only.**

## Tech stack

- Angular frontend
  - pnpm as a package manager
  - HSL color scheme based on [web.dev - Color schemes with HSL](https://web.dev/patterns/theming/)
- Go backend
  - Cobra for managing the different management commands
  - Echo as a server
  - sqlx and squirrel for interacting with postgres
- Postgres
- Kubernetes with traefik
- Prometheus for metrics

## Developing locally

### Frontend

```bash
# Install pnpm if you don't have it yet
npm i -g pnpm

cd frontend
pnpm install

pnpm start
```

### Backend

```bash
npm i -g nodemon
make be

# or if you don't need automatic reloading
cd backend
CONFIGOR_ENV=dev go run . serve
```