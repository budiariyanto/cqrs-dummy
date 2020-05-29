-- +migrate Up
CREATE TABLE campaign (
    id BIGINT AUTO_INCREMENT,
    name varchar(50),
    is_closed TINYINT(1),
    target_donation BIGINT,
    reason varchar(100),
    PRIMARY KEY(id)
);

-- +migrate Down
DROP TABLE campaign;