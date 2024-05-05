BEGIN;

CREATE TABLE IF NOT EXISTS shippings(
                                           id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
                                           transaction_id uuid NOT NULL,
                                           is_shipped BOOLEAN NOT NULL DEFAULT false,
                                           shipping_cost numeric(10,2),
                                           created_at timestamptz NOT NULL default now(),
                                           updated_at timestamptz NOT NULL default now(),
                                           deleted_at timestamptz
);

COMMIT;