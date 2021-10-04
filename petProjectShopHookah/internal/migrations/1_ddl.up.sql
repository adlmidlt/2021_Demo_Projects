CREATE TABLE IF NOT EXISTS hookah.roles
(
    id    SERIAL PRIMARY KEY NOT NULL,
    title VARCHAR(15)        NOT NULL
);

CREATE TABLE IF NOT EXISTS hookah.users
(
    id            SERIAL PRIMARY KEY NOT NULL,
    first_name    VARCHAR(35)        NOT NULL,
    second_name   VARCHAR(35)        NOT NULL,
    data_of_birth DATE               NOT NULL,
    gender        CHAR(1)            NOT NULL,
    number_phone  VARCHAR(12)        NOT NULL,
    login         VARCHAR(15)        NOT NULL,
    password      VARCHAR(10)        NOT NULL
);

CREATE TABLE IF NOT EXISTS hookah.manufactures
(
    id    SERIAL PRIMARY KEY NOT NULL,
    title VARCHAR(50)        NOT NULL
);

CREATE TABLE IF NOT EXISTS hookah.categories
(
    id    SERIAL PRIMARY KEY NOT NULL,
    title VARCHAR(15)        NOT NULL
);

CREATE TABLE IF NOT EXISTS hookah.products
(
    id             SERIAL PRIMARY KEY NOT NULL,
    title          VARCHAR(25)        NOT NULL,
    description    VARCHAR(150)       NOT NULL,
    price          NUMERIC(18, 2)     NULL,
    manufacture_id SERIAL             NOT NULL,
    date_delivery  DATE               NULL,
    CONSTRAINT fk_product_manufacture
        FOREIGN KEY (manufacture_id)
            REFERENCES manufactures (id)
            ON DELETE CASCADE
            ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS hookah.user_role
(
    role_id SERIAL NOT NULL,
    user_id SERIAL NOT NULL,
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

CREATE TABLE IF NOT EXISTS hookah.basket
(
    id          SERIAL PRIMARY KEY NOT NULL,
    description VARCHAR(150)       NOT NULL,
    price       NUMERIC(18, 2)     NOT NULL,
    user_id     SERIAL             NOT NULL,
    product_id  SERIAL             NOT NULL,
    CONSTRAINT fk_user_basket
        FOREIGN KEY (user_id)
            REFERENCES users (id)
            ON DELETE CASCADE
            ON UPDATE CASCADE,
    CONSTRAINT fk_products_basket
        FOREIGN KEY (user_id)
            REFERENCES products (id)
            ON DELETE CASCADE
            ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS hookah.categories_products
(
    category_id SERIAL NOT NULL,
    product_id  SERIAL NOT NULL,
    CONSTRAINT fk_categories_products_category
        FOREIGN KEY (category_id)
            REFERENCES categories (id)
            ON DELETE CASCADE
            ON UPDATE CASCADE,
    CONSTRAINT fk_categories_products_product
        FOREIGN KEY (product_id)
            REFERENCES products (id)
            ON DELETE CASCADE
            ON UPDATE CASCADE
);


COMMENT ON TABLE hookah.roles is 'Роли';
COMMENT ON COLUMN hookah.roles.id is 'Идентификатор роли';
COMMENT ON COLUMN hookah.roles.title is 'Наименование';

COMMENT ON TABLE hookah.users is 'Пользователи';
COMMENT ON COLUMN hookah.users.id is 'Идентификатор пользователя';
COMMENT ON COLUMN hookah.users.first_name is 'Имя';
COMMENT ON COLUMN hookah.users.second_name is 'Фамилия';
COMMENT ON COLUMN hookah.users.data_of_birth is 'Дата рождения';
COMMENT ON COLUMN hookah.users.gender is 'Пол (Мужской/Женский)';
COMMENT ON COLUMN hookah.users.number_phone is 'Номер телефона';
COMMENT ON COLUMN hookah.users.login is 'Логин';
COMMENT ON COLUMN hookah.users.password is 'Пароль';

COMMENT ON TABLE hookah.manufactures is 'Производители продукций';
COMMENT ON COLUMN hookah.manufactures.id is 'Идентификатор производителя';
COMMENT ON COLUMN hookah.manufactures.title is 'Наименование';

COMMENT ON TABLE hookah.categories is 'Категории';
COMMENT ON COLUMN hookah.categories.id is 'Идентификатор категории';
COMMENT ON COLUMN hookah.categories.title is 'Наименование';

COMMENT ON TABLE hookah.products is 'Продукты';
COMMENT ON COLUMN hookah.products.id is 'Идентификатор продукта';
COMMENT ON COLUMN hookah.products.title is 'Наименование';
COMMENT ON COLUMN hookah.products.description is 'Описание';
COMMENT ON COLUMN hookah.products.price is 'Цена';
COMMENT ON COLUMN hookah.products.manufacture_id is 'Идентификатор производителя';
COMMENT ON COLUMN hookah.products.date_delivery is 'Дата доставки';

COMMENT ON TABLE hookah.user_role is 'Роль пользователя';
COMMENT ON COLUMN hookah.user_role.user_id is 'Идентификатор пользователя';
COMMENT ON COLUMN hookah.user_role.role_id is 'Идентификатор роли';

COMMENT ON TABLE hookah.basket is 'Корзина';
COMMENT ON COLUMN hookah.basket.id is 'Идентификатор корзины';
COMMENT ON COLUMN hookah.basket.description is 'Описание';
COMMENT ON COLUMN hookah.basket.user_id is 'Идентификатор пользователя';
COMMENT ON COLUMN hookah.basket.product_id is 'Идентификатор продукта';

COMMENT ON TABLE hookah.categories_products is 'Категории продуктов';
COMMENT ON COLUMN hookah.categories_products.product_id is 'Идентификатор продукта';
COMMENT ON COLUMN hookah.categories_products.category_id is 'Идентификатор категории';

