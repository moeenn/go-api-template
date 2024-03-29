CREATE TABLE
  users (
    user_id UUID NOT NULL,
    email VARCHAR(255) UNIQUE,
    password VARCHAR(255),
    role user_role NOT NULL,
    is_active BOOL DEFAULT(true),
    PRIMARY KEY (user_id)
  );
