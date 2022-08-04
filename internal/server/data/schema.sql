
CREATE TABLE access_keys (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text,
    issued_for bigint,
    provider_id bigint,
    expires_at timestamp with time zone,
    extension bigint,
    extension_deadline timestamp with time zone,
    key_id text,
    secret_checksum bytea,
    scopes text
);

CREATE SEQUENCE access_keys_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE access_keys_id_seq OWNED BY access_keys.id;

CREATE TABLE credentials (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    identity_id bigint,
    password_hash bytea,
    one_time_password boolean
);

CREATE SEQUENCE credentials_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE credentials_id_seq OWNED BY credentials.id;

CREATE TABLE destinations (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text,
    unique_id text,
    connection_url text,
    connection_ca text,
    last_seen_at timestamp with time zone,
    version text,
    resources text,
    roles text
);

CREATE SEQUENCE destinations_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE destinations_id_seq OWNED BY destinations.id;

CREATE TABLE encryption_keys (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    key_id integer,
    name text,
    encrypted bytea,
    algorithm text,
    root_key_id text
);

CREATE SEQUENCE encryption_keys_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE encryption_keys_id_seq OWNED BY encryption_keys.id;

CREATE TABLE grants (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    subject text,
    privilege text,
    resource text,
    created_by bigint
);

CREATE SEQUENCE grants_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE grants_id_seq OWNED BY grants.id;

CREATE TABLE groups (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text,
    created_by bigint,
    created_by_provider bigint
);

CREATE SEQUENCE groups_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE groups_id_seq OWNED BY groups.id;

CREATE TABLE identities (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text,
    last_seen_at timestamp with time zone,
    created_by bigint
);

CREATE TABLE identities_groups (
    identity_id bigint NOT NULL,
    group_id bigint NOT NULL
);

CREATE SEQUENCE identities_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE identities_id_seq OWNED BY identities.id;

CREATE TABLE organizations (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text,
    created_by bigint
);

CREATE SEQUENCE organizations_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE organizations_id_seq OWNED BY organizations.id;

CREATE TABLE provider_users (
    identity_id bigint NOT NULL,
    provider_id bigint NOT NULL,
    email text,
    groups text,
    last_update timestamp with time zone,
    redirect_url text,
    access_token text,
    refresh_token text,
    expires_at timestamp with time zone
);

CREATE TABLE providers (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text,
    url text,
    client_id text,
    client_secret text,
    created_by bigint,
    kind text,
    auth_url text,
    scopes text,
    private_key text,
    client_email text,
    domain_admin_email text
);

CREATE SEQUENCE providers_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE providers_id_seq OWNED BY providers.id;

CREATE TABLE settings (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    private_jwk bytea,
    public_jwk bytea,
    lowercase_min bigint DEFAULT 0,
    uppercase_min bigint DEFAULT 0,
    number_min bigint DEFAULT 0,
    symbol_min bigint DEFAULT 0,
    length_min bigint DEFAULT 8
);

CREATE SEQUENCE settings_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE settings_id_seq OWNED BY settings.id;

ALTER TABLE ONLY access_keys ALTER COLUMN id SET DEFAULT nextval('access_keys_id_seq'::regclass);

ALTER TABLE ONLY credentials ALTER COLUMN id SET DEFAULT nextval('credentials_id_seq'::regclass);

ALTER TABLE ONLY destinations ALTER COLUMN id SET DEFAULT nextval('destinations_id_seq'::regclass);

ALTER TABLE ONLY encryption_keys ALTER COLUMN id SET DEFAULT nextval('encryption_keys_id_seq'::regclass);

ALTER TABLE ONLY grants ALTER COLUMN id SET DEFAULT nextval('grants_id_seq'::regclass);

ALTER TABLE ONLY groups ALTER COLUMN id SET DEFAULT nextval('groups_id_seq'::regclass);

ALTER TABLE ONLY identities ALTER COLUMN id SET DEFAULT nextval('identities_id_seq'::regclass);

ALTER TABLE ONLY organizations ALTER COLUMN id SET DEFAULT nextval('organizations_id_seq'::regclass);

ALTER TABLE ONLY providers ALTER COLUMN id SET DEFAULT nextval('providers_id_seq'::regclass);

ALTER TABLE ONLY settings ALTER COLUMN id SET DEFAULT nextval('settings_id_seq'::regclass);

ALTER TABLE ONLY access_keys
    ADD CONSTRAINT access_keys_pkey PRIMARY KEY (id);

ALTER TABLE ONLY credentials
    ADD CONSTRAINT credentials_pkey PRIMARY KEY (id);

ALTER TABLE ONLY destinations
    ADD CONSTRAINT destinations_pkey PRIMARY KEY (id);

ALTER TABLE ONLY encryption_keys
    ADD CONSTRAINT encryption_keys_pkey PRIMARY KEY (id);

ALTER TABLE ONLY grants
    ADD CONSTRAINT grants_pkey PRIMARY KEY (id);

ALTER TABLE ONLY groups
    ADD CONSTRAINT groups_pkey PRIMARY KEY (id);

ALTER TABLE ONLY identities_groups
    ADD CONSTRAINT identities_groups_pkey PRIMARY KEY (identity_id, group_id);

ALTER TABLE ONLY identities
    ADD CONSTRAINT identities_pkey PRIMARY KEY (id);

ALTER TABLE ONLY organizations
    ADD CONSTRAINT organizations_pkey PRIMARY KEY (id);

ALTER TABLE ONLY provider_users
    ADD CONSTRAINT provider_users_pkey PRIMARY KEY (identity_id, provider_id);

ALTER TABLE ONLY providers
    ADD CONSTRAINT providers_pkey PRIMARY KEY (id);

ALTER TABLE ONLY settings
    ADD CONSTRAINT settings_pkey PRIMARY KEY (id);

CREATE UNIQUE INDEX idx_access_keys_key_id ON access_keys USING btree (key_id) WHERE (deleted_at IS NULL);

CREATE UNIQUE INDEX idx_access_keys_name ON access_keys USING btree (name) WHERE (deleted_at IS NULL);

CREATE UNIQUE INDEX idx_credentials_identity_id ON credentials USING btree (identity_id) WHERE (deleted_at IS NULL);

CREATE UNIQUE INDEX idx_destinations_unique_id ON destinations USING btree (unique_id) WHERE (deleted_at IS NULL);

CREATE UNIQUE INDEX idx_encryption_keys_key_id ON encryption_keys USING btree (key_id);

CREATE UNIQUE INDEX idx_grant_srp ON grants USING btree (subject, privilege, resource) WHERE (deleted_at IS NULL);

CREATE UNIQUE INDEX idx_groups_name ON groups USING btree (name) WHERE (deleted_at IS NULL);

CREATE UNIQUE INDEX idx_identities_name ON identities USING btree (name) WHERE (deleted_at IS NULL);

CREATE UNIQUE INDEX idx_organizations_name ON organizations USING btree (name) WHERE (deleted_at IS NULL);

CREATE UNIQUE INDEX idx_providers_name ON providers USING btree (name) WHERE (deleted_at IS NULL);

ALTER TABLE ONLY access_keys
    ADD CONSTRAINT fk_access_keys_issued_for_identity FOREIGN KEY (issued_for) REFERENCES identities(id);

ALTER TABLE ONLY identities_groups
    ADD CONSTRAINT fk_identities_groups_group FOREIGN KEY (group_id) REFERENCES groups(id);

ALTER TABLE ONLY identities_groups
    ADD CONSTRAINT fk_identities_groups_identity FOREIGN KEY (identity_id) REFERENCES identities(id);

ALTER TABLE ONLY provider_users
    ADD CONSTRAINT fk_provider_users_identity FOREIGN KEY (identity_id) REFERENCES identities(id);

ALTER TABLE ONLY provider_users
    ADD CONSTRAINT fk_provider_users_provider FOREIGN KEY (provider_id) REFERENCES providers(id);
