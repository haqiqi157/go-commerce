BEGIN;

CREATE TABLE IF NOT EXISTS transactions(
                                       id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
                                       name
                                       product_id uuid NOT NULL,
                                       qty int not null,
                                       user_id uuid NOT NULL,
                                       discount numeric(10, 2),
                                       is_paid BOOLEAN NOT NULL DEFAULT false,
                                       created_at timestamptz NOT NULL default now(),
                                       updated_at timestamptz NOT NULL default now(),
                                       deleted_at timestamptz
);

COMMIT;