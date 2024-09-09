FROM node:20.16-alpine3.20 AS base
ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
RUN corepack enable

FROM base AS builder

RUN mkdir /app
COPY pnpm-lock.yaml /app
# COPY . /app
WORKDIR /app

RUN pnpm fetch --prod

COPY . /app
RUN pnpm build

# RUN --mount=type=cache,id=pnpm,target=/pnpm/store pnpm install --prod --frozen-lockfile
# RUN pnpm build


FROM lipanski/docker-static-website:latest AS final

COPY --from=builder /app/dist .
