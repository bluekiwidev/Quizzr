# Quizzr
An open source, fun quizzing/revision platform for putting the fun into learning. 

## Instructions of how to run both parts
# Frontend
The frontend runs on svelte with tailwind css, shadcn and some other packages.

## How to run
You will need the following
- Latest version of [nodejs](https://nodejs.org/en/download)
- Latest version of pnpm

Clone the repo
```bash
git clone https://github.com/bluekiwidev/Quizzr.git
```

Then cd into the repo's frontend
```bash
cd Quizzr/frontend
```

Install the deps with pnpm
```bash
pnpm install
```

Make your .env file
```bash
cp .env.example .env
```
> [!IMPORTANT]
> Remember to fill out the dummy .env vars with your real config.

Run the webserver
```bash
pnpm run dev --open
```

Now it will open a tab in your browser with the website on it.

# Backend
The backend runs on purely on go. It hadles api requests, db requests, users, etc

There are two ways of running the backend
- [Locally](#how-to-run-locally)
- [In docker](#how-to-run-with-docker)

## How to run locally
You will need the following
- Latest version of go
- A MariaDB to connect to

Clone the repo
```bash
git clone https://github.com/bluekiwidev/Quizzr.git
```

Then cd into the backend
```bash
cd Quizzr/backend
```

Copy over the .env example into the src
```bash
cp .env.example src/.env
```
> [!IMPORTANT]
> Remember to fill out the dummy .env vars with your real config.

cd into the src folder
```bash
cd src/
```

Install the deps
```bash
go mod tidy
```

Build the app
```bash
go build -o Quizzr.exe .
```

Run the app
```bash
./Quizzr.exe
```

## How to run with docker
You will need the following
- Latest version of docker

Clone the repo
```bash
git clone https://github.com/bluekiwidev/Quizzr.git
```

Then cd into the backend
```bash
cd Quizzr/backend
```
Edit the docker compose to match your config
```bash
nano docker-compose.yml
```
> [!IMPORTANT]
> Remember to fill out the dummy .env vars with your real config.

Run the docker compose stack
```bash
docker compose up
```
