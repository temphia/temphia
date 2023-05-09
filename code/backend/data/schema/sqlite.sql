create table system_events(
        id integer primary key autoincrement not null,
    type text not null,
    data text not null,
    extra_meta json not null default '{}',
    tenant_id text not null default ''
);

create table system_kv(
        id integer primary key autoincrement not null,
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
        id integer primary key autoincrement not null,
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
        id integer primary key autoincrement not null,
    name text not null default '', 
    target_type text not null,
    target text not null, 
    context_type text not null default '',
    icon text not null default '',
    policy text not null default '',
    plug_id text not null,
    agent_id text not null,
    extra_meta json not null default '{}',
    exec_meta json not null default '{}',
    tenant_id text not null,

    unique(target_type, target, context_type, tenant_id)
);

create table target_hooks(
        id integer primary key autoincrement not null,
    name text not null default '', 
    target_type text not null,
    target text not null,
    event_type text not null default '',
    policy text not null default '',
    plug_id text not null default '',
    agent_id text not null default '',
    extra_meta json not null default '{}',
    exec_meta json not null default '{}',
    tenant_id text not null,
    
    unique(target_type, target, event_type, tenant_id)
);




create table tenant_repos(
        id integer primary key autoincrement not null,
    name text not null default '',
    provider text not null default '',
    url text not null default '',
    extra_meta json not null default '{}',
    tenant_id text not null
);

create table user_groups(
    name text not null default '',
    slug text not null,
    scopes text not null default '',
    tenant_id text not null,
    features text not null default '',
    feature_opts json not null default '{}',
    extra_meta json not null default '{}',
    mod_version integer not null default 0,
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
    created_at timestamp not null default current_timestamp,
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
    id bigint not null,
    name text not null default '',
    user_id text not null,
    device_type text not null default 'device', --- device/login
    apn_token text not null default '',
    scopes text not null default '',
    last_data json not null default '{}', -- browser/ip/timestamp
    expires_on timestamptz not null,
    pair_options json not null default '{}',
    extra_meta json not null default '{}',
    tenant_id text not null,
    foreign KEY(user_id, tenant_id) references users(user_id, tenant_id),
    primary key(id)
);


create table user_messages(
        id integer primary key autoincrement not null,
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
    created_at timestamptz not null default current_timestamp,
    tenant_id text null,
    foreign KEY(user_id, tenant_id) references users(user_id, tenant_id)
);


create table user_group_auths(
        id integer primary key autoincrement not null,
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
        id integer primary key autoincrement not null,
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
    bprint_id text not null default '',
    renderer text not null default '',
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
        id integer primary key autoincrement not null,
    name text not null default '',
    count integer not null default 0,
    filter_conds json not null default '[]',
    selects text not null default '',
    main_column text not null default '',
    search_term text not null default '',
    ascending boolean not null default true,
    tags json not null default '{}',
    table_id TEXT not null,
    group_id TEXT not null,
    tenant_id TEXT not null,
    extra_meta json not null default '{}',
    foreign KEY(group_id, tenant_id) references data_table_groups (slug, tenant_id),
    foreign KEY(table_id, group_id, tenant_id) references data_tables (slug, group_id, tenant_id)
);



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
    entry_file text not null default '',
    iface_file text not null default '',
    mod_version integer not null default 0,
    web_entry text not null default '',
    web_script text not null default '',
    web_style text not null default '',
    web_loader text not null default '',
    web_files json not null default '{}',
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
    bprint_id text not null default '',
    extra_meta json not null default '{}',
    tenant_id text not null,
    primary KEY(id, tenant_id)
);

create table agent_links(
        id integer primary key autoincrement not null,
    name text not null default '',
    from_plug_id text not null,
    from_agent_id text not null,
    to_plug_id text not null,
    to_agent_id text not null,
    to_handler text not null default '',
    tenant_id text not null,
	 extra_meta json not null default '{}',
    foreign KEY(to_plug_id, tenant_id) references plugs(id, tenant_id),
    foreign KEY(to_agent_id, to_plug_id, tenant_id) references agents(id,plug_id, tenant_id),

    foreign KEY(from_plug_id, tenant_id) references plugs(id, tenant_id),
    foreign KEY( from_agent_id, from_plug_id, tenant_id) references agents(id,plug_id, tenant_id)

   
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
        id integer primary key autoincrement not null,
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

