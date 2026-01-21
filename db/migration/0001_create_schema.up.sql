BEGIN;

CREATE TABLE IF NOT EXISTS coupon (
    coupon_id UUID NOT NULL,
    coupon_name VARCHAR(20),
    amount INT NOT NULL DEFAULT 0,
    remaining_amount INT NOT NULL DEFAULT 0 CHECK (remaining_amount > 0),
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    PRIMARY KEY (coupon_id)
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_coupon_name_uniq
ON coupon(coupon_name);

CREATE UNIQUE INDEX IF NOT EXISTS idx_coupon_name_del
ON coupon(deleted_at);

COMMENT ON COLUMN coupon.amount IS 'Initial amount';
COMMENT ON COLUMN coupon.remaining_amount IS 'This is amount stock for retrive or add';

CREATE TABLE IF NOT EXISTS coupon_claim_history (
    coupon_claim_history_id UUID NOT NULL,
    coupon_name VARCHAR(20),
    user_id VARCHAR(20) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    PRIMARY KEY (coupon_claim_history_id)
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_coupon_claim_history_uniq
ON coupon_claim_history(coupon_name, user_id);

CREATE UNIQUE INDEX IF NOT EXISTS idx_coupon_claim_history_del
ON coupon_claim_history(deleted_at);

COMMIT;