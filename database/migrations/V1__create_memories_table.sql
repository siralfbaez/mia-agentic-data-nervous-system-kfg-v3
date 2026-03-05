-- Enable Google's AlloyDB AI and pgvector extensions
CREATE EXTENSION IF NOT EXISTS vector;
CREATE EXTENSION IF NOT EXISTS google_ml_integration;

CREATE TABLE memories (
    memory_id UUID PRIMARY KEY,
    correlation_id VARCHAR(255),
    context_summary TEXT,
    embedding_vector vector(768), -- Optimized for Vertex AI 'textembedding-gecko'
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
