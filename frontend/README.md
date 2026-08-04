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