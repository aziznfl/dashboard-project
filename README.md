# Dashboard Project

A full-stack dashboard application built with Go (Backend) and Vue.js (Frontend).

## 🚀 Prerequisites

Before you begin, ensure you have the following installed:

- **Backend**:
  - [Go](https://golang.org/doc/install) (1.21+)
  - GCC (for SQLite/CGO support)
- **Frontend**:
  - [Node.js](https://nodejs.org/) (v18+)
  - [npm](https://www.npmjs.com/)
- **Optional**:
  - [Docker](https://www.docker.com/) & [Docker Compose](https://docs.docker.com/compose/)

---

## 🛠️ Environment Setup

### 1. Backend Setup
1. Navigate to the backend directory:
   ```bash
   cd backend
   ```
2. Copy the sample environment file:
   ```bash
   cp env.sample .env
   ```
3. (Optional) Customize `.env` if needed (e.g., ports, JWT secret).

### 2. Frontend Setup
1. Navigate to the frontend directory:
   ```bash
   cd frontend
   ```
2. Copy the sample environment file:
   ```bash
   cp .env.sample .env
   ```

---

## 🏃 Running the Project

### Local Development

#### Start Backend
```bash
cd backend
make dep    # Install dependencies
make run    # Run API server
```
*The server will start at `http://localhost:8080` (default).*

#### Start Frontend
```bash
cd frontend
make dep    # Install dependencies
make dev    # Start Vite dev server
```
*The application will be available at `http://localhost:5173`.*

### Docker (Recommended)
You can run both services using Docker Compose:

**Run Backend:**
```bash
cd backend
docker-compose up -d
```

**Run Frontend:**
```bash
cd frontend
docker-compose up -d
```

---

## 📊 Data Seeding

This project features **automatic data seeding**. On the first run, the backend will:
1. Initialize the SQLite database (`dashboard.db`).
2. Create necessary tables (`users`, `payments`).
3. Seed default users and 1,000 random payment records if the database is empty.

**Default Credentials:**
- Email: `cs@test.com` / Password: `password` (Role: CS)
- Email: `operation@test.com` / Password: `password` (Role: Operation)

---

## 🧪 Testing & Building

### Backend
- **Test**: `go test ./...`
- **Build**: `make build`

### Frontend
- **Build**: `make build`
- **Preview Build**: `make preview`

---

## 📖 API Documentation

The API is documented using the **OpenAPI 3.0** specification.

- **File**: `backend/openapi.yaml`
- **View**: You can view the documentation by pasting the content into [Swagger Editor](https://editor.swagger.io/) or using a browser extension like Redocly.

Key Endpoints:
- `POST /dashboard/v1/auth/login`: Authentication
- `GET /dashboard/v1/payments`: List payments (with filters & pagination)
