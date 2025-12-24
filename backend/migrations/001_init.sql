-- backend/migrations/001_init.sql

CREATE TABLE IF NOT EXISTS questions (
    id SERIAL PRIMARY KEY,
    question TEXT NOT NULL,
    image_url TEXT NOT NULL,
    correct_swipe VARCHAR(10) NOT NULL CHECK (correct_swipe IN ('left', 'right')),
    category VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Index pour améliorer les performances
CREATE INDEX idx_category ON questions(category);

-- dataset test
INSERT INTO questions (question, image_url, correct_swipe, category) VALUES
('Poule ou Coq ?', 'https://images.unsplash.com/photo-1548550023-2bdb3c5beed7', 'right', 'animaux'),
('Chat ou Chien ?', 'https://images.unsplash.com/photo-1574158622682-e40e69881006', 'left', 'animaux'),
('Soleil ou Lune ?', 'https://images.unsplash.com/photo-1473496169904-658ba7c44d8a', 'right', 'nature'),
('Jour ou Nuit ?', 'https://images.unsplash.com/photo-1518173946687-a4c8892bbd9f', 'left', 'nature'),
('Café ou Thé ?', 'https://images.unsplash.com/photo-1495474472287-4d71bcdd2085', 'right', 'boissons');

-- Création de la table scores (pour plus tard > tracking perf)
CREATE TABLE IF NOT EXISTS scores (
    id SERIAL PRIMARY KEY,
    player_name VARCHAR(100),
    score INTEGER DEFAULT 0,
    questions_answered INTEGER DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);