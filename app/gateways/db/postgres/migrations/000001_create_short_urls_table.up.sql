CREATE TABLE IF NOT EXISTS short_urls (
         id text PRIMARY KEY NOT NULL,
         url text NOT NULL,
         active boolean NOT NULL  DEFAULT true,
         created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP);
