BEGIN ;

CREATE TABLE IF NOT EXISTS users(
                                    iid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                                    email varchar(255) NOT NULL ,
                                    password varchar(255) NOT NULL ,
                                    role varchar(50),
                                    address text,
                                    phone varchar(20)
);

COMMIT ;