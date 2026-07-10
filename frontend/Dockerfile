# Stage 1: Build the Vue app
FROM oven/bun:1-alpine AS builder

WORKDIR /app

# Copy dependency files first for caching
COPY package.json bun.lock ./
RUN bun install --frozen-lockfile

# Copy the rest of the source code and build
COPY . .
RUN bun run build

# Stage 2: Serve the built files with nginx
FROM nginx:alpine

# Replace the default nginx config with our custom one
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Copy the built Vue app from the builder stage
COPY --from=builder /app/dist /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
