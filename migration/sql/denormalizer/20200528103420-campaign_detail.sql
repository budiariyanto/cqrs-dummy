-- +migrate Up
CREATE TABLE campaign_detail (
    id BIGINT AUTO_INCREMENT,
    name varchar(50),
    total_donation BIGINT,
    target_donation BIGINT,
    is_closed TINYINT(1),
    reason varchar(100),
    PRIMARY KEY(id)
);

-- +migrate Down
DROP TABLE campaign_detail;
