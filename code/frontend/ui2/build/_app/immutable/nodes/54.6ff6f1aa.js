import{s as i,P as p}from"../chunks/scheduler.e2ee220a.js";import{S as _,i as f,b as c,d as u,m as l,a as g,t as y,e as d}from"../chunks/index.4aee2103.js";import{A as $}from"../chunks/auto_form.ee101f61.js";import"../chunks/paths.2eaeb908.js";import"../chunks/loading_spinner.svelte_svelte_type_style_lang.c28228c7.js";function v(m){let a,t;return a=new $({props:{message:m[0],schema:{fields:[{name:"Id",ftype:"TEXT_SLUG",key_name:"id"},{name:"Name",ftype:"TEXT",key_name:"name"},{name:"Bprint Id",ftype:"TEXT",key_name:"bprint_id"},{name:"Invoke Policy",ftype:"TEXT_POLICY",key_name:"invoke_policy"},{name:"Live",ftype:"BOOL",key_name:"live"},{name:"Dev",ftype:"BOOL",key_name:"dev"},{name:"Extra Meta",ftype:"KEY_VALUE_TEXT",key_name:"extra_meta"}],name:"New Plug",required_fields:["bprint_id"]},onSave:m[1],data:{}}}),{c(){c(a.$$.fragment)},l(e){u(a.$$.fragment,e)},m(e,n){l(a,e,n),t=!0},p(e,[n]){const o={};n&1&&(o.message=e[0]),a.$set(o)},i(e){t||(g(a.$$.fragment,e),t=!0)},o(e){y(a.$$.fragment,e),t=!1},d(e){d(a,e)}}}function k(m,a,t){const e=p("__app__"),n=e.api_manager.get_admin_plug_api();let o="";return[o,async s=>{console.log("@@data",s);const r=await n.new_plug(s);if(!r.ok){t(0,o=r.data);return}e.nav.admin_plugs()}]}class w extends _{constructor(a){super(),f(this,a,k,v,i,{})}}export{w as component};
