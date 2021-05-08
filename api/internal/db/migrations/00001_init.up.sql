CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE carbon_dioxide (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    value decimal,
    created_at timestamp
);

CREATE TABLE temperature (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    value decimal,
    created_at timestamp
);

CREATE TABLE humidity (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    value decimal,
    created_at timestamp
);
