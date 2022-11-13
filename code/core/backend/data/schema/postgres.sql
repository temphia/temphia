create EXTENSION if not exists hstore;

create table system_events(
    id serial primary key,
    type text not null,
    data text not null,
    extra_meta json not null default '{}',
    tenant_id text not null default ''
);

create table system_kv(
    id serial primary key,
    key text not null,
    type text not null default '',
    value text not null default '',
    tenant_id text not null default '',
    unique(key, type, tenant_id)
);

create table tenants(
    slug text not null,
    name text not null default '',
    org_bio text not null default '',
    
    default_ugroup text not null default '',
    default_domain text not null default '',

    smtp_user text not null default '',
    smtp_pass text not null default '',
    master_secret text not null default '',
    disable_p2p boolean not null default FALSE,
    allow_raw_query boolean not null default FALSE,
    extra_meta json not null default '{}'
);


create table tenant_domains(
    id serial primary key,
    name text not null default '',
    about text not null default '',
    default_ugroup text not null default '',
    cors_policy text not null default '',
    adapter_policy text not null default '',
    adapter_type text not null default '',
    adapter_opts json not null default '{}',
    adapter_cab_source text not null default '',
    adapter_cab_folder text not null default '',
    adapter_template_bprints text not null default '',
    tenant_id text not null,
    extra_meta json not null default '{}',
    unique(name, tenant_id)
);


create table target_apps(
    id serial primary key,
    name text not null default '', 
    target_type text not null,
    target text not null, 
    icon text not null default '',
    policy text not null default '',
    plug_id text not null,
    agent_id text not null,
    extra_meta json not null default '{}',
    exec_meta json not null default '{}',
    tenant_id text not null
);

create table target_hooks(
    id serial primary key,
    name text not null default '', 
    target_type text not null,
    target text not null,
    policy text not null default '',
    plug_id text not null default '',
    agent_id text not null default '',
    extra_meta json not null default '{}',
    exec_meta json not null default '{}',
    tenant_id text not null
);


-- create table target_kv(
--     key text not null,
--     value text not null,
--     target_type text not null,
--     target text not null,
--     tenant_id text not null,
--     PRIMARY KEY(tenant_id, key, plug_id)
-- );

-- create table target_event(
--     id serial primary key,
--     target_type text not null,
--     target text not null,
--     data text not null,
--     extra_meta json not null default '{}',
--     tenant_id text not null,
--     PRIMARY KEY(tenant_id, key, plug_id)
-- );


create table tenant_repos(
    id serial primary key,
    name text not null default '',
    provider text not null default '',
    url text not null default '',
    extra_meta json not null default '{}',
    tenant_id text not null
);

create table user_groups(
    name text not null default '',
    slug text not null,
    icon text not null default '',
    scopes text not null default '',
    tenant_id text not null,
    enable_pass_auth boolean not null default TRUE,
    open_sign_up boolean not null default FALSE,
    extra_meta json not null default '{}',
    default_domain text not null default '',

    -- mod_version integer not null default 0,
    -- apn_device_push boolean not null default FALSE,

    primary KEY(slug, tenant_id)
);

create table users(
    user_id text not null,
    full_name text not null default '',
    bio text not null default '',
    pub_key text not null default '',
    password text not null default '',
    email text not null default '',
    tenant_id text not null,
    group_id text not null,
    created_at timestamp not null default now(),
    active boolean not null default false,
    extra_meta json not null default '{}',
    foreign KEY(group_id, tenant_id) references user_groups(slug, tenant_id),
    unique(tenant_id, email),
    primary KEY(user_id, tenant_id)
);

create table user_datas(
    user_id text not null,
    pending_email_verify boolean not null default TRUE,
    pending_pass_change boolean not null default TRUE,

    mfa_enabled boolean not null default FALSE,
    mfa_type text not null default '',
    mfa_data text not null default '',
    
    extra_meta json not null default '{}',
    tenant_id text not null,
    foreign KEY(user_id, tenant_id) references users(user_id, tenant_id),
    primary KEY(user_id, tenant_id)
);

create table user_devices(
    xid text not null,
    user_id text not null,
    name text not null default '',
    last_addr text not null,
    apn_token text not null,
    tenant_id text not null,
    using_auth integer not null default 0,
    manual_generated boolean not null default FALSE,
    blocked boolean not null default FALSE,
    foreign KEY(user_id, tenant_id) references users(user_id, tenant_id),
    primary KEY(xid, tenant_id)
);

create table user_messages(
    id serial primary key,
    title text not null default '',
    read boolean not null default false,
    type text not null,
    contents text not null,
    user_id text not null,
    from_user text not null default '',
    from_plug text not null default '',
    from_agent text not null default '',
    plug_callback text not null default '',
    warn_level integer not null default 0,

    -- tags text not null default '',

    encrypted boolean not null default false,
    created_at timestamptz not null default now(),
    tenant_id text null,
    foreign KEY(user_id, tenant_id) references users(user_id, tenant_id)
);

create table roles(
    slug text not null,
    tenant_id text not null,
    PRIMARY KEY(slug, tenant_id)
);

create table user_roles(
    role_id text not null,
    user_id text not null,
    tenant_id text not null,
    FOREIGN KEY(role_id, tenant_id) references roles(slug, tenant_id),
    FOREIGN KEY(user_id, tenant_id) references users(user_id, tenant_id),
    PRIMARY KEY(role_id, tenant_id, user_id)
);

create table permissions(
    id serial primary key,
    object_type text not null,
    object_id text not null default '',
    role_id text not null,
    extra_meta json not null default '{}',
    tenant_id text not null,
    FOREIGN KEY(role_id, tenant_id) references roles(slug, tenant_id)
);


create table user_group_auths(
    id serial primary key,
    name text not null default '',
    type text not null,
    provider text not null default '',
    provider_opts json default '{}',
    scopes text not null default '',
    policy text not null default '',
    user_group text not null,
    tenant_id text not null,
    extra_meta json default '{}',
    foreign KEY(user_group, tenant_id) references user_groups(slug, tenant_id)
);

create table user_group_datas(
    id serial primary key,
    data_source text not null,
    data_group text not null,
    policy text not null default '',
    user_group text not null,
    tenant_id text not null,
    extra_meta json default '{}',
    foreign KEY(user_group, tenant_id) references user_groups(slug, tenant_id)
);

create table saved_tokens(
    id text not null,
    type text not null,
    user_id text not null default '',
    user_group text not null default '',
    target text not null default '',
    payload text not null,
    tenant_id text not null,
    expires_on timestamptz not null,
    extra_meta json default '{}',
    primary KEY(id, tenant_id)
);



--- DYNDB

create table data_table_groups (
    slug TEXT not null,
    name TEXT not null,
    description TEXT not null,
    source_db TEXT not null,
    tenant_id TEXT not null,
    cabinet_source TEXT not null default '',
    cabinet_folder TEXT not null default '',
    extra_meta json not null default '{}',
    active BOOLEAN not null default false,
    primary KEY(slug, tenant_id)
);
create table data_tables (
    slug TEXT not null,
    name TEXT not null,
    description TEXT not null,
    icon TEXT not null,
    extra_meta json not null default '{}',
    main_column TEXT not null default '',
    main_view TEXT not null default '',

    activity_type TEXT not null default '',
    sync_type TEXT not null default '',

    tenant_id text not null,
    group_id TEXT not null,
    foreign KEY(group_id, tenant_id) references data_table_groups(slug, tenant_id),
    primary KEY(slug, group_id, tenant_id)
);
create table data_table_columns (
    slug TEXT not null,
    name TEXT not null,
    order_id SERIAL,
    description TEXT not null,
    ctype TEXT not null,
    icon TEXT,
    pattern TEXT not null,
    strict_pattern BOOLEAN not null default false,
    table_id TEXT not null,
    group_id TEXT not null,
    options TEXT [],
    ref_id TEXT default '',
    ref_type TEXT default '',
    ref_target TEXT default '',
    ref_object TEXT default '',
    ref_copy TEXT default '',
    extra_meta json not null default '{}',
    tenant_id text not null,
    foreign KEY(group_id, tenant_id) references data_table_groups (slug, tenant_id),
    foreign KEY(table_id, group_id, tenant_id) references data_tables (slug, group_id, tenant_id),
    primary KEY(slug, table_id, group_id, tenant_id)
);
create table data_views (
    id serial primary key,
    name text not null default '',
    count integer not null default 0,
    filter_conds json not null default '[]',
    selects text not null default '',
    main_column text not null default '',
    search_term text not null default '',
    ascending boolean not null default true,
    table_id TEXT not null,
    group_id TEXT not null,
    tenant_id TEXT not null,
    extra_meta json not null default '{}',
    foreign KEY(group_id, tenant_id) references data_table_groups (slug, tenant_id),
    foreign KEY(table_id, group_id, tenant_id) references data_tables (slug, group_id, tenant_id)
);

-- data table trigger function

CREATE OR REPLACE FUNCTION data_activity_tg() RETURNS TRIGGER AS $func$
BEGIN
    RAISE INFO 'OUTER BEGIN';

    IF NOT(NEW IS NULL) THEN
        DECLARE
			_id integer = NEW.__id;
        	_version integer = 0;
        	_raw_mod_ctx json= COALESCE(NEW.__mod_sig, '{}');
        	_user_id text =  COALESCE((_raw_mod_ctx->>'user_id')::text, '');
        	_user_sign text = COALESCE((_raw_mod_ctx->>'user_sign')::text, '');
        	_init_sign text =  COALESCE((_raw_mod_ctx->>'init_sign')::text, '');
        	_table text = (_raw_mod_ctx->>'table_name')::text;
        	_type text = LOWER(TG_OP);
        	_hpayload hstore = hstore(NEW) - ARRAY['__id','__mod_sig', '__version' ];
        	_payload text = hstore_to_json(_hpayload)::text;
			BEGIN
                RAISE INFO 'NEW IS %s', _payload;
                RAISE INFO 'INNER BEGIN [table=>%s, user_id=>%s, user_sign=>%s init_sign=>%s]', _table, _user_id, _user_sign, _init_sign;
                IF NOT(OLD IS NULL) THEN
                    _version = OLD.__version;
                    NEW.__version = OLD.__version + 1;
                END IF;

                RAISE INFO '@@version => %d', _version;
                
				EXECUTE format(
					'INSERT INTO dact_%s( type,row_id,row_version,user_id,user_sign,init_sign,payload) VALUES ($1,$2,$3,$4,$5,$6,$7);',
					_table
				)
				USING _type,_id,_version,_user_id,_user_sign,_init_sign,_payload;
			END;
    END IF;
RETURN NEW;
END;
$func$ LANGUAGE plpgsql;


-- data table delete helper function
CREATE OR REPLACE FUNCTION data_row_delete(_table text, _id integer, mod_ctx text) RETURNS void AS $func$
DECLARE
    _raw_mod_ctx json = mod_ctx::json;
    _user_id text =  COALESCE((_raw_mod_ctx->>'user_id')::text, '');
    _user_sign text = COALESCE((_raw_mod_ctx->>'user_sign')::text, '');
    _init_sign text =  COALESCE((_raw_mod_ctx->>'init_sign')::text, '');
	_old record = null;
BEGIN
	EXECUTE format('SELECT __version FROM %s WHERE __id = $1', _table) INTO _old USING _id;
	EXECUTE format(
		'INSERT INTO 
			dact_%s(
				type,row_id,row_version,user_id,user_sign,init_sign
				) VALUES ($1,$2,$3,$4,$5);',_table) 
		USING 'delete',_id,_old.__version,_user_id,_user_sign,_init_sign;
    EXECUTE format('DELETE FROM %s WHERE __id=$1', _table) USING _id;
END
$func$  LANGUAGE plpgsql;


-- DROP TABLE IF EXISTS data_table_groups CASCADE;
-- DROP TABLE IF EXISTS  data_tables CASCADE;
-- DROP TABLE IF EXISTS  data_table_columns CASCADE;
-- DROP TABLE IF EXISTS  data_views CASCADE;
-- DROP TABLE IF EXISTS  data_hooks CASCADE;
-- DROP FUNCTION IF EXISTS  data_activity_tg CASCADE;
-- DROP FUNCTION IF EXISTS  data_row_delete CASCADE;



create table plug_states(
    key text not null,
    value text not null,
    version integer not null default 0,
	
    tag1 text not null default '',
    tag2 text not null default '',
	tag3 text not null default '',
	
    ttl timestamp,
    plug_id text not null,
    tenant_id text not null,
    PRIMARY KEY(tenant_id, key, plug_id)
);


create table bprints(
    id text not null,
    slug text not null default '',
    name text not null default '',
    type text not null,
    sub_type text not null default '',
    inline_schema text not null default '',
    description text not null default '',
    icon text not null default '',
    source_id text not null default '',
    files text not null default '',
    tags text not null default '',
    extra_meta json not null default '{}',
    tenant_id text not null,
    primary KEY(id, tenant_id)
);
create table plugs(
    id text not null,
    name text not null default '',
    live boolean not null default false,
    dev boolean not null default false,
    owner text not null default '',
    bprint_id text not null default '',
    invoke_policy text not null default '',
    extra_meta json not null default '{}',
    tenant_id text not null,
    primary KEY(id, tenant_id)
);
create table agents(
    id text not null,
    name text not null default '',
    type text not null,
    executor text not null,
    iface_file text not null default '',

    -- mod_version integer not null default 0,

    web_entry text not null default '',
    web_script text not null default '',
    web_style text not null default '',
    web_loader text not null default '',
    web_files json not null default '{}',
    env_vars json not null default '{}',
    extra_meta json not null default '{}',
    tenant_id text not null,
    plug_id text not null,
    foreign KEY(plug_id, tenant_id) references plugs(id, tenant_id),
    primary KEY(id, plug_id, tenant_id)
);

create table resources(
    id text not null,
    name text not null default '',
    type text not null,
    sub_type text not null default '',
    payload text not null default '',
    target text not null default '',
    policy text not null default '',
    plug_id text not null default '',
    extra_meta json not null default '{}',
    tenant_id text not null,
    primary KEY(id, tenant_id)
);

create table agent_links(
    id serial primary key,
    name text not null default '',
    from_plug_id text not null,
    from_agent_id text not null,
    to_plug_id text not null,
    to_agent_id text not null,
    to_handler text not null default '',
    tenant_id text not null,
    foreign KEY(to_plug_id, tenant_id) references plugs(id, tenant_id),
    foreign KEY(to_plug_id, to_agent_id, tenant_id) references agents(id,plug_id, tenant_id),

    foreign KEY(from_plug_id, tenant_id) references plugs(id, tenant_id),
    foreign KEY( from_agent_id, from_plug_id, tenant_id) references agents(id,plug_id, tenant_id),

    extra_meta json not null default '{}'
);

create table agent_resources(
    slug text not null,
    plug_id text not null,
    agent_id text not null,
    resource_id text not null,
    tenant_id text not null,
    actions text not null default '', 
    policy text not null default '',

    foreign KEY(resource_id, tenant_id) references resources(id, tenant_id),
    foreign KEY(plug_id, tenant_id) references plugs(id, tenant_id),
    foreign KEY(agent_id, plug_id, tenant_id) references agents(id,plug_id, tenant_id),
    primary KEY(slug, plug_id, agent_id, tenant_id)
);

create table agent_extensions(
    id serial primary key,
    name text not null default '',
    plug_id text not null,
    agent_id text not null,
    brpint_id text not null,
    ref_file text not null,
    tenant_id text not null,
    extra_meta json not null default '{}',
    foreign KEY(plug_id, tenant_id) references plugs(id, tenant_id),
    foreign KEY(agent_id, plug_id, tenant_id) references agents(id,plug_id, tenant_id)
);

