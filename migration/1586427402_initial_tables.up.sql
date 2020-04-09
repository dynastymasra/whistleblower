CREATE TABLE IF NOT EXISTS articles (
    id UUID PRIMARY KEY NOT NULL,
    created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS viewers (
    id UUID PRIMARY KEY NOT NULL,
    article_id UUID NOT NULL,
    created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (article_id) REFERENCES articles(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE INDEX IF NOT EXISTS viewers_created_at_idx ON viewers (created_at);
CREATE INDEX IF NOT EXISTS viewers_article_id_idx ON viewers (article_id);