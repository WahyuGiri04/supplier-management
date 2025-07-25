CREATE SCHEMA IF NOT EXISTS supplier_management;

CREATE EXTENSION IF NOT EXISTS "pgcrypto";

DROP TABLE IF EXISTS "supplier_management".supplier_group CASCADE;
DROP TABLE IF EXISTS "supplier_management".supplier CASCADE;
DROP TABLE IF EXISTS "supplier_management".address CASCADE;
DROP TABLE IF EXISTS "supplier_management".contacts CASCADE;


CREATE TABLE IF NOT EXISTS "supplier_management".supplier (
    id SERIAL PRIMARY KEY,
    uuid UUID DEFAULT gen_random_uuid(),
    code VARCHAR(255) UNIQUE,
    name VARCHAR(255),
    nickname VARCHAR(255),
    logo VARCHAR(500),
    created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by INTEGER,
    last_updated_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    last_updated_by INTEGER,
    is_active BOOLEAN DEFAULT TRUE,
    is_deleted BOOLEAN DEFAULT FALSE
);


CREATE TABLE IF NOT EXISTS "supplier_management".supplier_group (
    id SERIAL PRIMARY KEY,
    uuid UUID DEFAULT gen_random_uuid(),
    supplier_id INT REFERENCES "supplier_management".supplier(id) ON DELETE CASCADE,
    group_name VARCHAR(255),
    value VARCHAR(255),
    logo VARCHAR(500),
    created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by INTEGER,
    last_updated_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    last_updated_by INTEGER,
    is_active BOOLEAN DEFAULT TRUE,
    is_deleted BOOLEAN DEFAULT FALSE
);


CREATE TABLE IF NOT EXISTS "supplier_management".address (
    id SERIAL PRIMARY KEY,
    uuid UUID DEFAULT gen_random_uuid(),
    supplier_id INT REFERENCES "supplier_management".supplier(id) ON DELETE CASCADE,
    name VARCHAR(255),
    full_address VARCHAR(255),
    is_main BOOLEAN,
    created_by INTEGER,
    created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    last_updated_by INTEGER,
    last_updated_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    is_active BOOLEAN DEFAULT TRUE,
    is_deleted BOOLEAN DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS "supplier_management".contacts (
    id SERIAL PRIMARY KEY,
    uuid UUID DEFAULT gen_random_uuid(),
    supplier_id INT REFERENCES "supplier_management".supplier(id) ON DELETE CASCADE,
    name VARCHAR(255),
    job_position VARCHAR(255),
    email VARCHAR(255),
    phone_number VARCHAR(255),
    mobile_phone_number VARCHAR(255),
    is_main BOOLEAN,
    created_by INTEGER,
    created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    last_updated_by INTEGER,
    last_updated_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    is_active BOOLEAN DEFAULT TRUE,
    is_deleted BOOLEAN DEFAULT FALSE
);

select * from "supplier_management".supplier;
select * from "supplier_management".supplier_group;
select * from "supplier_management".address;
