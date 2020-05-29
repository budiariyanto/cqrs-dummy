-- +migrate Up
CREATE TABLE donation (
    id BIGINT AUTO_INCREMENT,
    campaign_id BIGINT NOT NULL,
    donor_name varchar(50) NOT NULL,
    amount INT NOT NULL,
    PRIMARY KEY(id)
);

-- +migrate Down
DROP TABLE donation;
