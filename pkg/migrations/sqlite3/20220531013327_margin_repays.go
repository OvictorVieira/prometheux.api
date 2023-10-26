package sqlite3

import (
	"context"

	"github.com/c9s/rockhopper"
)

func init() {
	AddMigration(upMarginRepays, downMarginRepays)

}

func upMarginRepays(ctx context.Context, tx rockhopper.SQLExecutor) (err error) {
	// This code is executed when the migration is applied.

	_, err = tx.ExecContext(ctx, "CREATE TABLE `margin_repays`\n(\n    `gid`                   BIGINT UNSIGNED         NOT NULL AUTO_INCREMENT,\n    `transaction_id`        BIGINT UNSIGNED         NOT NULL,\n    `user_exchanges_id`     BIGINT                  NOT NULL,\n    `asset`                 VARCHAR(24)             NOT NULL DEFAULT '',\n    `isolated_symbol`       VARCHAR(24)             NOT NULL DEFAULT '',\n    -- quantity is the quantity of the trade that makes profit\n    `principle`             DECIMAL(16, 8) UNSIGNED NOT NULL,\n    `time`                  DATETIME(3)             NOT NULL,\n    PRIMARY KEY (`gid`),\n    UNIQUE KEY (`transaction_id`),\n    FOREIGN KEY (`user_exchanges_id`) REFERENCES user_exchanges(id)\n);")
	if err != nil {
		return err
	}

	return err
}

func downMarginRepays(ctx context.Context, tx rockhopper.SQLExecutor) (err error) {
	// This code is executed when the migration is rolled back.

	_, err = tx.ExecContext(ctx, "DROP TABLE IF EXISTS `margin_repays`;")
	if err != nil {
		return err
	}

	return err
}