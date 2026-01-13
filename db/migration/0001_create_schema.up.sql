CREATE TABLE IF NOT EXISTS coupon (
    coupon_id UUID NOT NULL,
    coupon_name VARCHAR(20),
    ammount INT NOT NULL DEFAULT COMMENT 'initial default stock',
    remaining_amount INT NOT NULL DEFAULT 0 CHECK (remaining_amount > 0) COMMENT 'remaining stock ammount',
    created_at TIMESTAMPTZ DEFAULT NOW,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    PRIMARY KEY (coupon_id)
);

CREATE UNIQUE INDEX idx_coupon_name_uniq
ON coupon(coupon_name) WHERE deleted_at IS NULL;

CREATE TABLE IF NOT EXISTS coupon_claim_history (
    coupon_claim_history_id UUID NOT NULL,
    coupon_name VARCHAR(20),
    user_id VARCHAR(20) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    PRIMARY KEY (coupon_claim_history_id)
);

CREATE UNIQUE INDEX idx_coupon_claim_history_uniq
ON coupon(coupon_name, user_id) WHERE deleted_at IS NULL;