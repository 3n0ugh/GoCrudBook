CREATE TABLE book (
    isbn INT PRIMARY KEY,
    book_name VARCHAR(50) NOT NULL,
    author VARCHAR(100) NOT NULL,
    page_count INT NOT NULL,
    book_count INT NOT NULL,
    borrow_times INT NOT NULL,
    barrow_date DATETIME, 
    last_recieved_day DATETIME 
);

INSERT INTO book VALUES (
    1933988673,
    'Unlocking Android: A Developer Guide',
    'Charlie Collins', 
    416,
    1,
    0,
    NULL,
    NULL
);

INSERT INTO book VALUES (
    1933988746,
    'Flex 3 in Action',
    'Tariq Ahmed with Jon Hirschi',
    576,
    1,
    0,
    NULL,
    NULL
);
