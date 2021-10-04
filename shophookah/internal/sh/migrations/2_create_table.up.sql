DROP TABLE IF EXISTS hookah.user_role;

DROP TABLE IF EXISTS hookah.roles;

ALTER TABLE hookah.users ADD permission INT NOT NULL default 0;

COMMENT ON COLUMN hookah.users.permission IS '0-Гость, 1-Пользователь, 777-admin';
