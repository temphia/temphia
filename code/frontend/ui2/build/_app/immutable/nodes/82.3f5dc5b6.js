import{s as p,P as _}from"../chunks/scheduler.e2ee220a.js";import{S as i,i as u,b as f,d as c,m as g,a as l,t as y,e as T}from"../chunks/index.4aee2103.js";import{A as $}from"../chunks/auto_form.ee101f61.js";import"../chunks/paths.2eaeb908.js";import"../chunks/loading_spinner.svelte_svelte_type_style_lang.c28228c7.js";function d(o){let a,t;return a=new $({props:{message:o[0],schema:{fields:[{name:"Name",ftype:"TEXT",key_name:"name"},{name:"Slug",ftype:"TEXT_SLUG",key_name:"slug",slug_gen:null},{name:"Scopes",ftype:"MULTI_TEXT",key_name:"scopes"},{name:"Features",ftype:"MULTI_TEXT",key_name:"features"},{name:"Feature Options",ftype:"KEY_VALUE_TEXT",key_name:"feature_opts"},{name:"Extra Meta",ftype:"KEY_VALUE_TEXT",key_name:"extra_meta"}],name:"New User Group",required_fields:[]},onSave:o[1],data:{}}}),{c(){f(a.$$.fragment)},l(e){c(a.$$.fragment,e)},m(e,n){g(a,e,n),t=!0},p(e,[n]){const s={};n&1&&(s.message=e[0]),a.$set(s)},i(e){t||(l(a.$$.fragment,e),t=!0)},o(e){y(a.$$.fragment,e),t=!1},d(e){T(a,e)}}}function E(o,a,t){const e=_("__app__"),n=e.api_manager.get_admin_ugroup_api();let s="";return[s,async m=>{const r=await n.new(m);if(!r.ok){t(0,s=r.data);return}e.nav.admin_ugroups()}]}class w extends i{constructor(a){super(),u(this,a,E,d,p,{})}}export{w as component};
