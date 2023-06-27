FROM node:20.3.0-alpine3.17 as base

WORKDIR /app

RUN npm install -g pnpm

COPY ./package.json ./pnpm-lock.yaml ./
RUN pnpm install

COPY . .

CMD ["pnpm", "run", "start", "--host", "0.0.0.0"]
