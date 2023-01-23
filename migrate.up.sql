CREATE TABLE IF NOT EXISTS items(
                                    id SERIAL PRIMARY KEY,
                                    name varchar(100),
                                    file_url varchar(100),
                                    main_rack varchar(100)
);

CREATE TABLE IF NOT EXISTS client_order(
    id SERIAL PRIMARY KEY,
    order_number integer,
    item_id integer,
    amount integer,
    FOREIGN KEY (item_id) REFERENCES items (id)
);

CREATE TABLE IF NOT EXISTS storage(
    id SERIAL PRIMARY KEY,
    item_id integer,
    rack_name varchar(100),
    FOREIGN KEY (item_id) REFERENCES items (id)
);

INSERT INTO items (name, file_url,main_rack) VALUES ('Ноутбук','Картинка ноутбука','А');
INSERT INTO items (name, file_url,main_rack) VALUES ('Монитор','Картинка монитора','А');
INSERT INTO items (name, file_url,main_rack) VALUES ('Телефон','Картинка телефона','Б');
INSERT INTO items (name, file_url,main_rack) VALUES ('Компьютер','Картинка компьютера','Ж');
INSERT INTO items (name, file_url,main_rack) VALUES ('Часы','Картинка часов','Ж');
INSERT INTO items (name, file_url,main_rack) VALUES ('Микрофон','Картинка микрофона','Ж');

INSERT INTO storage (item_id, rack_name) VALUES (1,'А');
INSERT INTO storage (item_id, rack_name) VALUES (2,'А');
INSERT INTO storage (item_id, rack_name) VALUES (3,'Б');
INSERT INTO storage (item_id, rack_name) VALUES (3,'З');
INSERT INTO storage (item_id, rack_name) VALUES (3,'В');
INSERT INTO storage (item_id, rack_name) VALUES (4,'Ж');
INSERT INTO storage (item_id, rack_name) VALUES (5,'Ж');
INSERT INTO storage (item_id, rack_name) VALUES (5,'А');
INSERT INTO storage (item_id, rack_name) VALUES (6,'Ж');


INSERT INTO client_order (order_number, item_id, amount) VALUES (10,1,2);
INSERT INTO client_order (order_number, item_id, amount) VALUES (11,2,3);
INSERT INTO client_order (order_number, item_id, amount) VALUES (10,3,1);
INSERT INTO client_order (order_number, item_id, amount) VALUES (14,4,4);
INSERT INTO client_order (order_number, item_id, amount) VALUES (15,5,1);
INSERT INTO client_order (order_number, item_id, amount) VALUES (10,6,1);
