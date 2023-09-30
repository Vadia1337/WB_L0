CREATE TABLE orders (
    id SERIAL NOT NULL,
    order_uid VARCHAR NOT NULL,
    track_number VARCHAR UNIQUE NOT NULL,
    entry VARCHAR NOT NULL,
    delivery_id SERIAL NOT NULL
        REFERENCES deliveries (id)  ON DELETE CASCADE,
    payment_id SERIAL NOT NULL
        REFERENCES payments (id)  ON DELETE CASCADE,
    locale VARCHAR NOT NULL,
    internal_signature VARCHAR NOT NULL,
    customer_id VARCHAR NOT NULL,
    delivery_service VARCHAR NOT NULL,
    shardkey VARCHAR NOT NULL,
    sm_id INTEGER NOT NULL,
    date_created TIMESTAMP NOT NULL,
    oof_shard VARCHAR NOT NULL,
    PRIMARY KEY(id)
);