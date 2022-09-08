create table tenants(
    slug text not null,
    name text not null default '',
    org_bio text not null default '',
    
    root_plug_id text not null default '',
    root_agent_id text not null default '',
    root_handler text not null default '',
    default_ugroup text not null default '',
    default_domain text not null default '',

    smtp_user text not null default '',
    smtp_pass text not null default '',
    master_secret text not null default '',
    disable_p2p boolean not null default FALSE,
    allow_raw_query boolean not null default FALSE,
    extra_meta json not null default '{}'
);

