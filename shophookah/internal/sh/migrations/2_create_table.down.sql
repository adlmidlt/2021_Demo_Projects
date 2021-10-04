CREATE TABLE IF NOT EXISTS hookah.roles
(
    id    UUID PRIMARY KEY NOT NULL,
    title VARCHAR(15)        NOT NULL
);

CREATE TABLE IF NOT EXISTS hookah.user_role
(
    role_id UUID NOT NULL,
    user_id UUID NOT NULL,
    PRIMARY KEY (role_id, user_id),
    CONSTRAINT fk_user_role_users foreign key (user_id)
        REFERENCES users (id)
        ON DELETE CASCADE
        ON UPDATE CASCADE,
    CONSTRAINT fk_user_role_roles foreign key (role_id)
        REFERENCES roles (id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

ALTER TABLE hookah.users DROP COLUMN permission;