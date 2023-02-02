CREATE TYPE USER_ACCOUNT_ROLE AS ENUM ('REGISTERED', 'REJECTED', 'ACCEPTED');

CREATE TABLE users (
    id UUID NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    email TEXT NOT NULL UNIQUE,
    state USER_ACCOUNT_ROLE NOT NULL,
    source TEXT NOT NULL
);

CREATE TABLE experiments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    description TEXT NOT NULL
    -- TODO: not arenas aren't implemented
    -- arena_id UUID NOT NULL,
    -- CONSTRAINT fk_arena FOREIGN KEY(arena_id) REFERENCES arenas(id)
);

CREATE TABLE experiment_results (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    experiment_id UUID NOT NULL,
    completed TIMESTAMP NOT NULL,
    CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id),
    CONSTRAINT fk_experiment FOREIGN KEY(experiment_id) REFERENCES experiments(id)
);

CREATE TABLE invites (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    experiment_id UUID NOT NULL,
    supervised BOOLEAN NOT NULL,
    CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id),
    CONSTRAINT fk_experiment FOREIGN KEY(experiment_id) REFERENCES experiments(id)
);