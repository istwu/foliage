DROP TABLE IF EXISTS posts CASCADE;
DROP TABLE IF EXISTS portfolios CASCADE;
DROP TABLE IF EXISTS users CASCADE;

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(100) UNIQUE NOT NULL,
    -- email VARCHAR(255) UNIQUE NOT NULL,
    -- password_hash VARCHAR(255) NOT NULL,
    role VARCHAR(20) DEFAULT 'user',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP DEFAULT NULL
);
INSERT INTO users VALUES (
    1,
    'admin',
    'admin'
),
(
    2,
    'user',
    'user'
);

CREATE TABLE portfolios (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP DEFAULT NULL
);
INSERT INTO portfolios VALUES (
    1,
    2,
    'My First Portfolio'
),
(
    2,
    2,
    'My Second Portfolio'
);

CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    portfolio_id INT NOT NULL REFERENCES portfolios(id) ON DELETE CASCADE,
    post_type TEXT CHECK (post_type IN ('image', 'text')),
    image_url TEXT,
    text_body TEXT,
    position INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP DEFAULT NULL
);
INSERT INTO posts VALUES (
    1,
    2,
    1,
    'image',
    'https://mbevia.com/wp-content/uploads/2021/09/easy-tabby-cat-painting-tutorial-miguel-bevia.jpg',
    NULL,
    1
),
(
    2,
    2,
    1,
    'text',
    NULL,
    'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam finibus ante sit amet ligula iaculis pellentesque. Suspendisse semper, lacus et ultrices venenatis, tortor urna cursus mauris, at porttitor nulla nibh at ante. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Praesent non augue porta turpis rutrum efficitur ac finibus ipsum. Suspendisse potenti. Suspendisse vehicula pharetra consequat. Sed in sodales orci. Praesent elementum ipsum at libero pulvinar tincidunt. Suspendisse faucibus, tortor a iaculis venenatis, risus enim placerat nisi, sit amet varius nulla sem at metus. Duis dignissim nisi sapien, ac pretium nisi facilisis dapibus.',
    2
),
(
    3,
    2,
    2,
    'image',
    'https://i.ebayimg.com/images/g/1qwAAOSwi3VlBaKQ/s-l1200.jpg',
    NULL,
    1
);

ALTER SEQUENCE portfolios_id_seq RESTART WITH 3;
ALTER SEQUENCE users_id_seq RESTART WITH 3;
ALTER SEQUENCE posts_id_seq RESTART WITH 4;
