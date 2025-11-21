FROM node:20 AS builder

WORKDIR /app
COPY package*.json ./
RUN npm install

# Copy source
COPY . .

# Build production bundle
RUN npm run build

# Stage 2: Serve with Nginx
FROM nginx:alpine

WORKDIR /usr/share/nginx/html

# Copy build output
COPY --from=builder /app/build /usr/share/nginx/html/

# Replace default Nginx config
COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 3000
CMD ["nginx", "-g", "daemon off;"]
