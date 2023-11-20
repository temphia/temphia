import{s as p,P as _}from"../chunks/scheduler.e2ee220a.js";import{S as f,i as u,b as c,d as l,m as y,a as d,t as g,e as T}from"../chunks/index.4aee2103.js";import{v as k,a as E}from"../chunks/index.3b48e8d3.js";import{A as v}from"../chunks/auto_form.ee101f61.js";import"../chunks/paths.9c1b57c4.js";import"../chunks/loading_spinner.svelte_svelte_type_style_lang.0368a52d.js";function X(r){let a,n;return a=new v({props:{message:r[0],schema:{fields:[{name:"Full Name",ftype:"TEXT",key_name:"full_name"},{name:"User Id",ftype:"TEXT",key_name:"user_id"},{name:"Email",ftype:"TEXT",key_name:"email"},{name:"Group",ftype:"TEXT",key_name:"group_id"},{name:"Bio",ftype:"LONG_TEXT",key_name:"bio"},{name:"Password",ftype:"TEXT",key_name:"password"},{name:"Public Key",ftype:"LONG_TEXT",key_name:"pub_key"},{name:"Active",ftype:"BOOL",key_name:"active"}],name:"New User",required_fields:[]},onSave:r[1],data:{}}}),{c(){c(a.$$.fragment)},l(e){l(a.$$.fragment,e)},m(e,s){y(a,e,s),n=!0},p(e,[s]){const t={};s&1&&(t.message=e[0]),a.$set(t)},i(e){n||(d(a.$$.fragment,e),n=!0)},o(e){g(a.$$.fragment,e),n=!1},d(e){T(a,e)}}}function $(r,a,n){const e=_("__app__"),s=e.api_manager.get_admin_user_api();let t="";return[t,async m=>{if(!k(m.email)){n(0,t="Invalid Email");return}const o=m.user_id||"";if(!o||!E(o)){n(0,t="Invalid user id");return}n(0,t="");const i=await s.new(m);if(!i.ok){n(0,t=i.data);return}e.nav.admin_users()}]}class A extends f{constructor(a){super(),u(this,a,$,X,p,{})}}export{A as component};