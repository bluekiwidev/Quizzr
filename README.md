# Voxlet

## Star History

<a href="https://www.star-history.com/?repos=bluekiwidev%2Fvoxlet&type=date&legend=top-left">
 <picture>
   <source media="(prefers-color-scheme: dark)" srcset="https://api.star-history.com/chart?repos=bluekiwidev/voxlet&type=date&theme=dark&legend=top-left&sealed_token=YCK2sk-5yRTJeId8g_2Hy-K0XrPT4Cai3xn71em-fmAV6Xgxrdc_Er0Z6VSiaL9pWw2c8IZ8wtPM4SdZ8_XTBy6q-R_lNkooIidaQQtrQTDwCNDCH5p5ZoWIHPo3roceHAhk3k_nA9anmUEpTq3I4Z55lNYssdvkBT9NmTJ-VBheptJb2mPkQZJNvIsA" />
   <source media="(prefers-color-scheme: light)" srcset="https://api.star-history.com/chart?repos=bluekiwidev/voxlet&type=date&legend=top-left&sealed_token=YCK2sk-5yRTJeId8g_2Hy-K0XrPT4Cai3xn71em-fmAV6Xgxrdc_Er0Z6VSiaL9pWw2c8IZ8wtPM4SdZ8_XTBy6q-R_lNkooIidaQQtrQTDwCNDCH5p5ZoWIHPo3roceHAhk3k_nA9anmUEpTq3I4Z55lNYssdvkBT9NmTJ-VBheptJb2mPkQZJNvIsA" />
   <img alt="Star History Chart" src="https://api.star-history.com/chart?repos=bluekiwidev/voxlet&type=date&legend=top-left&sealed_token=YCK2sk-5yRTJeId8g_2Hy-K0XrPT4Cai3xn71em-fmAV6Xgxrdc_Er0Z6VSiaL9pWw2c8IZ8wtPM4SdZ8_XTBy6q-R_lNkooIidaQQtrQTDwCNDCH5p5ZoWIHPo3roceHAhk3k_nA9anmUEpTq3I4Z55lNYssdvkBT9NmTJ-VBheptJb2mPkQZJNvIsA" />
 </picture>
</a>

YES, I KNOW THE NAME IS A WIP.

## Instructions of how to run both parts
# Frontend
The frontend runs on svelte with tailwind css, shadcn and some other packages.

## How to run
You will need the following
- Latest version of [nodejs](https://nodejs.org/en/download)
- Latest version of pnpm

Clone the repo
```bash
git clone https://github.com/bluekiwidev/voxlet.git
```

Then cd into the repo's frontend
```bash
cd voxlet/frontend
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
git clone https://github.com/bluekiwidev/voxlet.git
```

Then cd into the backend
```bash
cd voxlet/backend
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
go build -o Voxlet.exe .
```

Run the app
```bash
./Voxlet.exe
```

## How to run with docker
You will need the following
- Latest version of docker

Clone the repo
```bash
git clone https://github.com/bluekiwidev/voxlet.git
```

Then cd into the backend
```bash
cd voxlet/backend
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
