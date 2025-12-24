# 🎮 Quiz Swipe PWA - Documentation Complète du Projet

## 📋 Vue d'ensemble

**Application de quiz interactive type Tinder** avec système de swipe (gauche/droite) pour répondre aux questions.

### Concept

- Affichage d'une carte avec une question et une image
- Swipe **droite** = Oui/Vrai
- Swipe **gauche** = Non/Faux
- Si mauvaise réponse → Game Over
- Si bonne réponse → Question suivante

**Exemple :**

- Question : "Poule ou Coq ?"
- Image : Photo de poule
- Swipe droite (Poule) → ✅ Correct, question suivante
- Swipe gauche (Coq) → ❌ Perdu, fin du jeu

---

## 🏗️ Architecture Technique

### Stack Choisie

| --- | --- | --- | --- |

---

## 📁 Structure du Projet

```
quiz-swipe/
├── frontend/                      # Application Vue.js PWA
│   ├── public/
│   │   ├── manifest.json         # PWA manifest (plus tard)
│   │   ├── icons/                # Icônes PWA (plus tard)
│   │   └── sw.js                 # Service Worker (plus tard)
│   ├── src/
│   │   ├── assets/
│   │   ├── components/
│   │   │   ├── SwipeCard.vue    # Carte swipable principale
│   │   │   └── GameOver.vue     # Écran game over
│   │   ├── views/
│   │   │   └── Home.vue         # Vue principale
│   │   ├── services/
│   │   │   └── api.js           # Client HTTP (Fetch natif)
│   │   ├── App.vue
│   │   ├── main.js
│   │   └── style.css
│   ├── package.json
│   ├── vite.config.js
│   ├── tailwind.config.js
│   └── index.html
│
├── backend/                       # API Go
│   ├── cmd/
│   │   └── server/
│   │       └── main.go           # Point d'entrée serveur
│   ├── internal/
│   │   ├── handlers/
│   │   │   └── quiz.go           # Handlers HTTP
│   │   ├── models/
│   │   │   └── question.go       # Structures de données
│   │   └── database/
│   │       └── postgres.go       # Connexion DB
│   ├── migrations/
│   │   └── 001_init.sql          # Schema initial + données test
│   ├── go.mod                     # Dépendances Go
│   └── go.sum
│
├── docker-compose.yml             # Config Docker (optionnel)
├── .gitignore
└── README.md

```

---

## 🎨 Design & UX

### Principes Mobile-First

1. **Interface verticale** optimisée pour smartphone
2. **Gestes tactiles** : Swipe gauche/droite
3. **Animations fluides** : Transitions smooth entre cartes
4. **PWA** : Installable comme app native
5. **DaisyUI** : Composants pré-stylés, responsive

### Composants UI Principaux

### SwipeCard

- Carte plein écran avec image
- Question en overlay
- Indicateurs visuels (gauche/droite)
- Gestion touch/mouse events
- Animations de swipe

### GameOver

- Écran de fin
- Score affiché
- Bouton "Rejouer"

---

## 🔌 API Backend

### Endpoints

### `GET /api/question`

Récupère une question aléatoire

**Response:**

```json
{
  "id": 1,
  "question": "Poule ou Coq ?",
  "image_url": "<https://images.unsplash.com/photo->...",
  "category": "animaux"
}
```

**Note:** La bonne réponse (`correct_swipe`) n'est PAS envoyée au frontend pour éviter la triche.

### `POST /api/swipe`

Vérifie si le swipe est correct

**Request:**

```json
{
  "question_id": 1,
  "direction": "right"
}
```

**Response:**

```json
{
  "correct": true,
  "message": "Bravo! 🎉"
}
```

### `GET /health`

Health check du serveur

**Response:**

```json
{
  "status": "ok"
}
```

---

## 🗄️ Base de Données

### Schema PostgreSQL

```sql
CREATE TABLE questions (
    id SERIAL PRIMARY KEY,
    question TEXT NOT NULL,
    image_url TEXT NOT NULL,
    correct_swipe VARCHAR(10) NOT NULL CHECK (correct_swipe IN ('left', 'right')),
    category VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_category ON questions(category);

```

### Données Exemple

```sql
INSERT INTO questions (question, image_url, correct_swipe, category) VALUES
('Poule ou Coq ?', 'https://...', 'right', 'animaux'),
('Chat ou Chien ?', 'https://...', 'left', 'animaux'),
('Soleil ou Lune ?', 'https://...', 'right', 'nature');

```

---

## 🚀 Installation & Setup

### Prérequis

**Obligatoire:**

- Go 1.21+ → `brew install go`
- Node.js 20+ → `brew install node` (ou utilise Bun)
- PostgreSQL 16-17 → `brew install postgresql@17`

**Recommandé:**

- Bun → `curl -fsSL <https://bun.sh/install> | bash`
- Docker Desktop → `brew install --cask docker`

**Optionnel:**

- Git → Déjà installé sur macOS

### Installation

### Backend Go

```bash
# 1. Créer le dossier backend
mkdir -p backend/cmd/server
cd backend

# 2. Initialiser Go modules
go mod init github.com/tonusername/quiz-swipe

# 3. Créer main.go (copier le code fourni)

# 4. Les dépendances s'installent automatiquement
go mod tidy

# 5. Lancer le serveur
go run cmd/server/main.go

```

### Frontend Vue.js

```bash
# 1. Créer le projet Vue
npm create vite@latest frontend -- --template vue
# OU avec Bun
bun create vite frontend --template vue

cd frontend

# 2. Installer les dépendances
npm install
# OU
bun install

# 3. Installer DaisyUI + Tailwind
npm install -D tailwindcss daisyui autoprefixer
npx tailwindcss init

# 4. Configuration Tailwind (voir section Tailwind + DaisyUI)

# 5. Lancer le dev server
npm run dev -- --host
# OU
bun run dev -- --host

```

### PostgreSQL

**Option 1 : Homebrew (local)**

```bash
# Démarrer PostgreSQL
brew services start postgresql@17

# Créer le rôle postgres
/opt/homebrew/opt/postgresql@17/bin/createuser -s postgres

# Créer la base de données
/opt/homebrew/opt/postgresql@17/bin/createdb quizdb

# Exécuter les migrations
/opt/homebrew/opt/postgresql@17/bin/psql -U postgres -d quizdb -f backend/migrations/001_init.sql

```

**Option 2 : Docker (recommandé)**

```bash
# Lancer PostgreSQL
docker run --name quiz-postgres \\
  -e POSTGRES_PASSWORD=password \\
  -e POSTGRES_DB=quizdb \\
  -p 5432:5432 \\
  -d postgres:17-alpine

# Exécuter les migrations
docker exec -i quiz-postgres \\
  psql -U postgres -d quizdb < backend/migrations/001_init.sql

```

---

## 🎨 Configuration DaisyUI + Tailwind

### Installation

```bash
cd frontend
npm install -D tailwindcss daisyui autoprefixer postcss
npx tailwindcss init -p

```

### Configuration

**tailwind.config.js:**

```jsx
/** @type {import('tailwindcss').Config} */
export default {
  content: ["./index.html", "./src/**/*.{vue,js,ts,jsx,tsx}"],
  theme: {
    extend: {},
  },
  plugins: [require("daisyui")],
  daisyui: {
    themes: ["light", "dark", "cupcake"], // Thèmes disponibles
    darkTheme: "dark",
    base: true,
    styled: true,
    utils: true,
  },
};
```

**src/style.css:**

```css
@tailwind base;
@tailwind components;
@tailwind utilities;

/* Styles globaux pour mobile-first */
body {
  @apply bg-base-100 text-base-content;
  margin: 0;
  padding: 0;
  overflow: hidden;
}

#app {
  width: 100vw;
  height: 100vh;
  overflow: hidden;
}
```

**main.js:**

```jsx
import { createApp } from "vue";
import "./style.css"; // Import Tailwind + DaisyUI
import App from "./App.vue";

createApp(App).mount("#app");
```

### Composants DaisyUI Utiles

```
<!-- Boutons -->
<button class="btn btn-primary">Rejouer</button>
<button class="btn btn-ghost">Annuler</button>

<!-- Cards -->
<div class="card bg-base-100 shadow-xl">
  <div class="card-body">
    <h2 class="card-title">Question</h2>
  </div>
</div>

<!-- Loading -->
<span class="loading loading-spinner loading-lg"></span>

<!-- Alert -->
<div class="alert alert-success">
  <span>Bravo! 🎉</span>
</div>

<!-- Modal -->
<dialog class="modal">
  <div class="modal-box">
    <h3 class="font-bold text-lg">Game Over</h3>
  </div>
</dialog>

```

---

## 📱 Test sur Mobile (Local)

### Configuration

**vite.config.js:**

```jsx
import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";

export default defineConfig({
  plugins: [vue()],
  server: {
    host: "0.0.0.0", // Accessible sur le réseau local
    port: 5173,
  },
});
```

### Procédure

```bash
# 1. Trouver ton IP locale (macOS)
ipconfig getifaddr en0
# Exemple : 192.168.1.45

# 2. Lancer le serveur
npm run dev

# 3. Sur ton iPhone/Android
# - Connecte-toi au même WiFi
# - Ouvre Safari/Chrome
# - Va sur : <http://192.168.1.45:5173>

# ✅ Ton app s'affiche !

```

---

## 🐳 Docker (Optionnel - Pour Plus Tard)

### docker-compose.yml

```yaml
version: "3.8"

services:
  postgres:
    image: postgres:17-alpine
    container_name: quiz-postgres
    environment:
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: password
      POSTGRES_DB: quizdb
    ports:
      - "5432:5432"
    volumes:
      - postgres_data:/var/lib/postgresql/data
      - ./backend/migrations:/docker-entrypoint-initdb.d

  backend:
    build: ./backend
    container_name: quiz-backend
    environment:
      DATABASE_URL: postgres://postgres:password@postgres:5432/quizdb?sslmode=disable
    ports:
      - "8080:8080"
    depends_on:
      - postgres

  frontend:
    build: ./frontend
    container_name: quiz-frontend
    ports:
      - "80:80"
    depends_on:
      - backend

volumes:
  postgres_data:
```

### Lancer avec Docker

```bash
# Tout démarrer
docker-compose up -d

# Voir les logs
docker-compose logs -f

# Arrêter
docker-compose down

# Rebuild après modifications
docker-compose up --build

```

---

## 🔄 Workflow de Développement

### Démarrage Quotidien

```bash
# Terminal 1 : PostgreSQL
brew services start postgresql@17
# OU
docker start quiz-postgres

# Terminal 2 : Backend Go
cd backend
go run cmd/server/main.go
# Serveur sur <http://localhost:8080>

# Terminal 3 : Frontend Vue
cd frontend
npm run dev -- --host
# App sur <http://localhost:5173>

```

### Hot Reload

- **Frontend** : Rechargement instantané (Vite HMR)
- **Backend** : Redémarrage manuel du serveur après modifications
- **Base de données** : Modifications SQL nécessitent redémarrage

---

## 📊 Équivalences Commandes

| --- | --- | --- | --- |

---

## 🎯 Roadmap du Projet

### Phase 1 : MVP (Semaines 1-2) ✅ En cours

- [x] Setup Go + Vue.js + PostgreSQL
- [x] Choisir DaisyUI comme framework CSS
- [ ] Créer composant SwipeCard
- [ ] Implémenter logique de swipe (touch + mouse)
- [ ] Connecter frontend ↔ backend
- [ ] Tester sur mobile (local WiFi)
- [ ] Styling mobile-first avec DaisyUI

### Phase 2 : Features (Semaine 3)

- [ ] Système de score
- [ ] Animations de transition
- [ ] Sons (optionnel)
- [ ] Catégories de questions
- [ ] Difficulté progressive

### Phase 3 : PWA (Semaine 4)

- [ ] Installer vite-plugin-pwa
- [ ] Créer manifest.json
- [ ] Générer icônes PWA
- [ ] Service Worker pour cache offline
- [ ] Testable comme app installée

### Phase 4 : Production (Semaine 5+)

- [ ] Dockerisation complète
- [ ] CI/CD (GitHub Actions)
- [ ] Déploiement (Vercel, Railway, [Fly.io](http://fly.io/))
- [ ] Domaine personnalisé
- [ ] Analytics (optionnel)

---

## 🛠️ Commandes Utiles

### Go

```bash
# Initialiser module
go mod init nom-projet

# Télécharger dépendances
go mod tidy

# Lancer serveur
go run cmd/server/main.go

# Build binaire
go build -o server cmd/server/main.go

# Tester
go test ./...

```

### Frontend

```bash
# Installer dépendances
npm install  # ou bun install

# Dev server
npm run dev  # ou bun run dev

# Build production
npm run build  # ou bun run build

# Preview build
npm run preview

```

### PostgreSQL

```bash
# Démarrer service
brew services start postgresql@17

# Stopper service
brew services stop postgresql@17

# Se connecter
psql -U postgres -d quizdb

# Créer database
createdb quizdb

# Dump database
pg_dump quizdb > backup.sql

# Restore database
psql quizdb < backup.sql

```

### Docker

```bash
# Lancer un container
docker run -d --name nom image

# Voir containers actifs
docker ps

# Logs d'un container
docker logs -f nom

# Stopper container
docker stop nom

# Supprimer container
docker rm nom

# Entrer dans un container
docker exec -it nom bash

```

---

## 📚 Ressources & Documentation

### Vue.js

- Documentation : https://vuejs.org/
- Guide PWA : https://vite-pwa-org.netlify.app/

### DaisyUI

- Documentation : https://daisyui.com/
- Composants : https://daisyui.com/components/
- Thèmes : https://daisyui.com/docs/themes/

### Tailwind CSS

- Documentation : https://tailwindcss.com/
- Playground : https://play.tailwindcss.com/

### Go

- Documentation : https://go.dev/doc/
- Gin Framework : https://gin-gonic.com/

### PostgreSQL

- Documentation : https://www.postgresql.org/docs/

---

## 🐛 Troubleshooting

### Backend ne démarre pas

```bash
# Vérifier que PostgreSQL tourne
brew services list | grep postgresql

# Vérifier la connexion DB
psql -U postgres -d quizdb -c "SELECT 1;"

# Vérifier les dépendances Go
go mod tidy

```

### Frontend ne charge pas

```bash
# Vérifier que le backend tourne
curl <http://localhost:8080/health>

# Vérifier les dépendances
npm install  # ou bun install

# Clear cache Vite
rm -rf node_modules/.vite

```

### Erreurs CORS

Vérifie que le backend a bien le middleware CORS configuré pour accepter `http://localhost:5173`.

### PostgreSQL : "role does not exist"

```bash
# Créer le rôle postgres
createuser -s postgres

```

---

## 🎓 Décisions Techniques & Justifications

### Pourquoi Vue.js et pas React ?

- Plus léger (bundle size)
- Courbe d'apprentissage plus douce
- Excellent support PWA natif
- Syntaxe plus intuitive pour débuter

### Pourquoi Go et pas Node.js ?

- Performance 10x supérieure
- Consommation mémoire 5x inférieure
- Binary unique (déploiement simplifié)
- Concurrence native (goroutines)
- Compilé (pas d'erreurs runtime)

### Pourquoi PostgreSQL et pas MongoDB ?

- Structure de données simple (SQL suffit)
- ACID compliance (intégrité garantie)
- Requêtes SQL standard
- Meilleur pour relations entre tables

### Pourquoi DaisyUI et pas autre chose ?

- Composants pré-stylés mobile-first
- Basé sur Tailwind (standard industrie)
- Thèmes changeables facilement
- Pas de JavaScript (pure CSS)
- Léger et performant

### Pourquoi Fetch natif et pas Axios ?

- 0 KB de dépendance
- Standard Web moderne
- Largement supporté
- Suffisant pour ce projet simple

### Pourquoi Bun et pas npm ?

- 10-20x plus rapide
- Compatible à 100% avec npm
- Économise du temps de dev
- Même syntaxe que npm

---

## 💡 Notes Importantes

1. **Mobile-First** : Tous les composants doivent être pensés pour mobile d'abord, desktop ensuite
2. **PWA plus tard** : On peut ajouter les features PWA une fois que l'app fonctionne
3. **Docker optionnel** : Pas nécessaire pour débuter, utile pour déploiement
4. **DaisyUI** : Utiliser les composants DaisyUI plutôt que créer du CSS custom
5. **Performance** : Go + PostgreSQL = Stack très performante, pas besoin d'optimiser prématurément

---

## 🚀 Quick Start (Copy-Paste)

```bash
# Backend
mkdir -p backend/cmd/server backend/migrations
cd backend
go mod init quiz-swipe
# Copier main.go et 001_init.sql
go mod tidy
go run cmd/server/main.go &

# Frontend
cd ..
bun create vite frontend --template vue
cd frontend
bun install
bun add -D tailwindcss daisyui autoprefixer postcss
npx tailwindcss init -p
# Configurer Tailwind + DaisyUI
bun run dev -- --host

# PostgreSQL (Docker)
docker run --name quiz-postgres \\
  -e POSTGRES_PASSWORD=password \\
  -e POSTGRES_DB=quizdb \\
  -p 5432:5432 \\
  -d postgres:17-alpine

# ✅ Tout est prêt !

```

---

**Projet créé le** : 21 novembre 2024

**Dernière mise à jour** : 21 novembre 2024

**Status** : 🟡 En développement (Phase 1 - MVP)
