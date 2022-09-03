create table tenants(
    slug text not null,
    name text not null default '',
    org_bio text not null default '',
    root_plug text not null default '',
    root_agent text not null default '',
    smtp_user text not null default '',
    smtp_password text not null default '',
    master_secret text not null default '',
    disable_p2p boolean not null default FALSE,
    extra_meta json not null default '{}'
);
