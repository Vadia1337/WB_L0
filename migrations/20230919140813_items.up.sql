CREATE TABLE items (
    id SERIAL NOT NULL,
    chrt_id INTEGER NOT NULL,
    track_number VARCHAR NOT NULL
        REFERENCES orders (track_number) ON DELETE CASCADE,
    price INTEGER NOT NULL,
    rid VARCHAR NOT NULL,
    name VARCHAR NOT NULL,
    sale INTEGER NOT NULL,
    size VARCHAR NOT NULL,
    total_price INTEGER NOT NULL,
    nm_id INTEGER NOT NULL,
    brand VARCHAR NOT NULL,
    status INTEGER NOT NULL,
    PRIMARY KEY(id)
);