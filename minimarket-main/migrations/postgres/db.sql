CREATE TABLE branch (
  id UUID PRIMARY KEY,
  name VARCHAR(30),
  address VARCHAR(30),
  create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TYPE tarif_type_enum AS ENUM ('percent', 'fixed');
CREATE TABLE staff_tarif (
  id UUID PRIMARY KEY,
  name VARCHAR(30),
  tarif_type tarif_type_enum,
  amount_for_cash INT,
  amount_for_card INT,
  create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


 CREATE TYPE staff_type_enum AS ENUM ('shopassistant', 'cashier');
CREATE TABLE staff (
 
  id UUID PRIMARY KEY,
  branch_id UUID REFERENCES branch(id),
  tarif_id UUID REFERENCES staff_tarif(id),
  staff_type staff_type_enum,
  name VARCHAR(30),
  balance VARCHAR(30),
  age INT,
  birthdate INT,
  login VARCHAR(30),
  password VARCHAR(30),
  create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TYPE payment_type_enum AS ENUM ('card', 'cash');
CREATE TYPE status_enum AS ENUM ('success', 'cancel', 'inprocses');
CREATE TABLE sale (
  id UUID PRIMARY KEY,
  branch_id UUID REFERENCES branch(id),
  shopassistant_id UUID REFERENCES staff(id),
  cashier_id UUID REFERENCES staff(id),
  payment_type payment_type_enum,
  price INT,
  status_type status_enum,
  clientname VARCHAR(30),
  create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);



CREATE TYPE transaction_type_enum AS ENUM ('withdraw', 'topup');

CREATE TYPE source_type_enum AS ENUM ('bonus', 'sales');
CREATE TABLE transaction (
  id UUID PRIMARY KEY,
  sale_id UUID REFERENCES sale(id),
  staff_id UUID REFERENCES staff(id),
  transaction_type transaction_type_enum,
  source_type source_type_enum,
  amount INT,
  description VARCHAR(30),
  create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);









