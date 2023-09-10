-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS time_series_data (
    id int NOT NULL PRIMARY KEY AUTO_INCREMENT,
    task VARCHAR(40) NOT NULL,  
    value VARCHAR(100) NOT NULL,
    date VARCHAR(20) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE key `task_date` (`task`, `date`)
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE time_series_data;
-- +goose StatementEnd
