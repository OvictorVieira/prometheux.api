-- +up
CREATE TABLE `klines`
(
    `gid`           BIGINT UNSIGNED         NOT NULL AUTO_INCREMENT,
    `exchange_id`   BIGINT                  NOT NULL,
    `start_time`    DATETIME(3)             NOT NULL,
    `end_time`      DATETIME(3)             NOT NULL,
    `interval`      VARCHAR(3)              NOT NULL,
    `symbol`        VARCHAR(20)             NOT NULL,
    `open`          DECIMAL(20, 8) UNSIGNED NOT NULL,
    `high`          DECIMAL(20, 8) UNSIGNED NOT NULL,
    `low`           DECIMAL(20, 8) UNSIGNED NOT NULL,
    `close`         DECIMAL(20, 8) UNSIGNED NOT NULL DEFAULT 0.0,
    `volume`        DECIMAL(20, 8) UNSIGNED NOT NULL DEFAULT 0.0,
    `closed`        BOOL                    NOT NULL DEFAULT TRUE,
    `last_trade_id` INT UNSIGNED            NOT NULL DEFAULT 0,
    `num_trades`    INT UNSIGNED            NOT NULL DEFAULT 0,

    PRIMARY KEY (`gid`)
);

CREATE INDEX `klines_end_time_symbol_interval` ON klines (`end_time`, `symbol`, `interval`);

-- +down
DROP INDEX `klines_end_time_symbol_interval` ON `klines`;

