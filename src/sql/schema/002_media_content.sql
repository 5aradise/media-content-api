-- +goose Up
CREATE TABLE media_content (
    id SERIAL NOT NULL,
    title VARCHAR(45) NOT NULL,
    description VARCHAR(255),
    body TEXT NOT NULL,
    content_type VARCHAR(36) NOT NULL,
    created_at DATE NOT NULL,
    user_id INT NOT NULL 
        REFERENCES users(id) 
        ON DELETE NO ACTION 
        ON UPDATE NO ACTION
);

-- +goose Down
DROP TABLE media_content;