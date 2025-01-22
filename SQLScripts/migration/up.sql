CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE public.accounts (
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

CREATE TABLE public.profiles (
	id uuid DEFAULT public.uuid_generate_v4(),
	account_id uuid NOT NULL, 
	"name" varchar NOT NULL,
	address varchar NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT profiles_pkey PRIMARY KEY (id)
);

CREATE TABLE public.farms (
	id uuid DEFAULT public.uuid_generate_v4(),
	profile_id uuid NOT NULL, 
	"name" varchar NOT NULL,
	address varchar NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT farms_pkey PRIMARY KEY (id)
);

CREATE TABLE public.system_units (
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

CREATE TABLE public.growth_plans (
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

CREATE TABLE public.growth_hist(
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

CREATE TABLE public.tank_trans (
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

ALTER TABLE ONLY public.profiles ADD CONSTRAINT fk_profiles_accounts FOREIGN KEY (account_id) REFERENCES public.accounts(id) ON UPDATE CASCADE ON DELETE SET NULL;
ALTER TABLE ONLY public.farms ADD CONSTRAINT fk_farms_profiles FOREIGN KEY (profile_id) REFERENCES public.profiles(id) ON UPDATE CASCADE ON DELETE SET NULL;
ALTER TABLE ONLY public.system_units ADD CONSTRAINT fk_system_units_farms FOREIGN KEY (farm_id) REFERENCES public.farms(id) ON UPDATE CASCADE ON DELETE SET NULL;
ALTER TABLE ONLY public.growth_plans ADD CONSTRAINT fk_growth_plans_farms FOREIGN KEY (farm_id) REFERENCES public.farms(id) ON UPDATE CASCADE ON DELETE SET NULL;
ALTER TABLE ONLY public.growth_plans ADD CONSTRAINT fk_growth_plans_system FOREIGN KEY (system_id) REFERENCES public.system_units(id) ON UPDATE CASCADE ON DELETE SET NULL;
ALTER TABLE ONLY public.growth_hist ADD CONSTRAINT fk_growth_hist_system FOREIGN KEY (system_id) REFERENCES public.system_units(id) ON UPDATE CASCADE ON DELETE SET NULL;
ALTER TABLE ONLY public.growth_hist ADD CONSTRAINT fk_growth_hist_farms FOREIGN KEY (farm_id) REFERENCES public.farms(id) ON UPDATE CASCADE ON DELETE SET NULL;
ALTER TABLE ONLY public.tank_trans ADD CONSTRAINT fk_tank_trans_system FOREIGN KEY (system_id) REFERENCES public.system_units(id) ON UPDATE CASCADE ON DELETE SET NULL;
ALTER TABLE ONLY public.tank_trans ADD CONSTRAINT fk_tank_trans_farms FOREIGN KEY (farm_id) REFERENCES public.farms(id) ON UPDATE CASCADE ON DELETE SET NULL;






