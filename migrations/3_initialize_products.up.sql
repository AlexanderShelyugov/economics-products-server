WITH INS (NAME, PRODUCT_TYPE, WEIGHT) AS (
VALUES
    ('Notebook', 'Office supplies', 1),
    ('Notebook', 'Office supplies', 2),
    ('Pencil', 'Office supplies', 2),
    ('Pencil', 'Office supplies', 1),
    ('Eraser', 'Office supplies', 1),
    ('Coffee', 'Groceries', 1),
    ('Cookies', 'Groceries', 2),
    ('Puzzles', 'Toys', 5)
)
INSERT INTO PRODUCTS (NAME, PRODUCT_TYPE_ID, WEIGHT)
SELECT ins.NAME, PRODUCT_TYPES.ID, ins.WEIGHT
FROM
    PRODUCT_TYPES JOIN INS
    ON PRODUCT_TYPES.NAME = INS.PRODUCT_TYPE
;