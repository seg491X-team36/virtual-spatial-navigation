CREATE TYPE USER_ACCOUNT_STATE AS ENUM ('REGISTERED', 'REJECTED', 'ACCEPTED');

CREATE TABLE users (
    id UUID NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    email TEXT NOT NULL UNIQUE,
    state USER_ACCOUNT_STATE NOT NULL,
    source TEXT NOT NULL
);

CREATE TABLE experiments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    config JSON NOT NULL
);

CREATE TABLE experiment_results (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    user_id UUID NOT NULL,
    experiment_id UUID NOT NULL,
    CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id),
    CONSTRAINT fk_experiment FOREIGN KEY(experiment_id) REFERENCES experiments(id)
);

CREATE TABLE invites (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    user_id UUID NOT NULL,
    experiment_id UUID NOT NULL,
    CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id),
    CONSTRAINT fk_experiment FOREIGN KEY(experiment_id) REFERENCES experiments(id)
);