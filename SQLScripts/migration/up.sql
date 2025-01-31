CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create schema hydroponic_system;

CREATE TABLE hydroponic_system.accounts (
	id uuid DEFAULT public.uuid_generate_v4(),
	username varchar NOT NULL,
	email varchar NOT NULL,
	"password" varchar NOT NULL,
	"role" varchar NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT accounts_pkey PRIMARY KEY (id)
);

CREATE TABLE hydroponic_system.profiles (
	id uuid DEFAULT public.uuid_generate_v4(),
	account_id uuid NOT NULL, 
	"name" varchar NOT NULL,
	address varchar NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT profiles_pkey PRIMARY KEY (id)
);

CREATE TABLE hydroponic_system.farms (
	id uuid DEFAULT public.uuid_generate_v4(),
	profile_id uuid NOT NULL, 
	"name" varchar NOT NULL,
	address varchar NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT farms_pkey PRIMARY KEY (id)
);

CREATE TABLE hydroponic_system.system_units (
	id uuid DEFAULT public.uuid_generate_v4(),
	farm_id uuid NOT NULL, 
	tank_volume int NOT NULL,
	tank_a_volume int NOT NULL,
	tank_b_volume int NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT system_units_pkey PRIMARY KEY (id)
);

CREATE TABLE hydroponic_system.growth_plans (
	id uuid DEFAULT public.uuid_generate_v4(),
	farm_id uuid NOT NULL,
	system_id uuid NOT NULL,
	seed_name varchar NOT NULL,
	seed_source varchar NOT NULL,
	seed_qty int NOT NULL,
	start_plan timestamptz NOT NULL,
	harvest_plan timestamptz NOT NULL,
	start_act timestamptz NOT NULL,
	harvest_act timestamptz NOT NULL,
	harvest_qty int NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT growth_plans_pkey PRIMARY KEY (id)
);

CREATE TABLE hydroponic_system.growth_hist(
	id uuid DEFAULT public.uuid_generate_v4(),
	farm_id uuid NOT NULL, 
	system_id uuid NOT NULL, 
	ppm int NOT NULL,
	ph int NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT growth_hist_pkey PRIMARY KEY (id)
);

CREATE TABLE hydroponic_system.tank_trans (
	id uuid DEFAULT public.uuid_generate_v4(),
	farm_id uuid NOT NULL, 
	system_id uuid NOT NULL, 
	water_volume int NOT NULL,
	a_volume int NOT NULL,
	b_volume int NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT tank_trans_pkey PRIMARY KEY (id)
);

ALTER TABLE ONLY hydroponic_system.profiles ADD CONSTRAINT fk_profiles_accounts FOREIGN KEY (account_id) REFERENCES hydroponic_system.accounts(id) ON UPDATE CASCADE ON DELETE SET NULL;
ALTER TABLE ONLY hydroponic_system.farms ADD CONSTRAINT fk_farms_profiles FOREIGN KEY (profile_id) REFERENCES hydroponic_system.profiles(id) ON UPDATE CASCADE ON DELETE SET NULL;
ALTER TABLE ONLY hydroponic_system.system_units ADD CONSTRAINT fk_system_units_farms FOREIGN KEY (farm_id) REFERENCES hydroponic_system.farms(id) ON UPDATE CASCADE ON DELETE SET NULL;
ALTER TABLE ONLY hydroponic_system.growth_plans ADD CONSTRAINT fk_growth_plans_farms FOREIGN KEY (farm_id) REFERENCES hydroponic_system.farms(id) ON UPDATE CASCADE ON DELETE SET NULL;
ALTER TABLE ONLY hydroponic_system.growth_plans ADD CONSTRAINT fk_growth_plans_system FOREIGN KEY (system_id) REFERENCES hydroponic_system.system_units(id) ON UPDATE CASCADE ON DELETE SET NULL;
ALTER TABLE ONLY hydroponic_system.growth_hist ADD CONSTRAINT fk_growth_hist_system FOREIGN KEY (system_id) REFERENCES hydroponic_system.system_units(id) ON UPDATE CASCADE ON DELETE SET NULL;
ALTER TABLE ONLY hydroponic_system.growth_hist ADD CONSTRAINT fk_growth_hist_farms FOREIGN KEY (farm_id) REFERENCES hydroponic_system.farms(id) ON UPDATE CASCADE ON DELETE SET NULL;
ALTER TABLE ONLY hydroponic_system.tank_trans ADD CONSTRAINT fk_tank_trans_system FOREIGN KEY (system_id) REFERENCES hydroponic_system.system_units(id) ON UPDATE CASCADE ON DELETE SET NULL;
ALTER TABLE ONLY hydroponic_system.tank_trans ADD CONSTRAINT fk_tank_trans_farms FOREIGN KEY (farm_id) REFERENCES hydroponic_system.farms(id) ON UPDATE CASCADE ON DELETE SET NULL;




create schema super_admin;

CREATE TABLE super_admin.accounts (
	id uuid DEFAULT public.uuid_generate_v4(),
	username varchar NOT NULL,
	"password" varchar NOT NULL,
	created_at timestamptz NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT accounts_pkey PRIMARY KEY (id)
);

CREATE TABLE super_admin.unit_ids (
	id uuid DEFAULT public.uuid_generate_v4(),
	created_at timestamptz NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT unit_ids_pkey PRIMARY KEY (id)
);

CREATE TABLE super_admin.system_logs(
	id uuid DEFAULT public.uuid_generate_v4(),
	message varchar NOT NULL,
	created_at timestamptz NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT system_logs_pkey PRIMARY KEY (id)
);