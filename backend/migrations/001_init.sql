CREATE TABLE IF NOT EXISTS urls (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    address VARCHAR(255) NOT NULL,
    title VARCHAR(255) DEFAULT '',
    html_version VARCHAR(20) DEFAULT '',
    h1_count INT DEFAULT 0,
    h2_count INT DEFAULT 0,
    h3_count INT DEFAULT 0,
    internal_links INT DEFAULT 0,
    external_links INT DEFAULT 0,
    broken_links INT DEFAULT 0,
    has_login_form BOOLEAN DEFAULT FALSE,
    status VARCHAR(20) DEFAULT 'queued',
    created_at DATETIME,
    updated_at DATETIME
);
