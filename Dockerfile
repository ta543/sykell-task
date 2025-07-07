FROM node:18 AS frontend
WORKDIR /app/frontend
COPY frontend ./
RUN npm install && npm run build

FROM golang:1.23-alpine AS backend
WORKDIR /app
COPY backend/go.mod backend/go.sum ./
RUN go mod download || true
COPY backend ./backend
RUN cd backend && go build -o /app/server

FROM alpine:3.18
WORKDIR /app
COPY --from=backend /app/server ./server
COPY --from=frontend /app/frontend/dist ./frontend
EXPOSE 8080
CMD ["./server"]