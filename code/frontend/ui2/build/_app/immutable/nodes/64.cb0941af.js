import{s as u,B as c,P as f}from"../chunks/scheduler.e2ee220a.js";import{S as _,i as l,b as y,d as T,m as k,a as d,t as v,e as E}from"../chunks/index.4aee2103.js";import{A as X}from"../chunks/auto_form.ee101f61.js";import"../chunks/paths.3df37c61.js";import"../chunks/loading_spinner.svelte_svelte_type_style_lang.0368a52d.js";import{p as w}from"../chunks/index.008d0d8b.js";function S(s){let t,n;return t=new X({props:{message:s[0],schema:{fields:[{name:"key",ftype:"TEXT",key_name:"key"},{name:"Value",ftype:"LONG_TEXT",key_name:"value"},{name:"Tag1",ftype:"TEXT",key_name:"tag1"},{name:"Tag2",ftype:"TEXT",key_name:"tag1"},{name:"Tag3",ftype:"TEXT",key_name:"tag1"}],name:"New Plug State",required_fields:["key"]},onSave:s[1],data:{}}}),{c(){y(t.$$.fragment)},l(e){T(t.$$.fragment,e)},m(e,o){k(t,e,o),n=!0},p(e,[o]){const m={};o&1&&(m.message=e[0]),t.$set(m)},i(e){n||(d(t.$$.fragment,e),n=!0)},o(e){v(t.$$.fragment,e),n=!1},d(e){E(t,e)}}}function $(s,t,n){let e;c(s,w,a=>n(3,e=a));let{pid:o=e.pid}=t;const m=f("__app__"),i=m.api_manager.get_admin_plug_api();let r="";const g=async a=>{console.log("@@data",a);const p=await i.new_plug_state(o,{key:a.key,value:a.value||"",options:{tag1:a.tag1||"",tag2:a.tag2||"",tag3:a.tag3||""}});if(!p.ok){n(0,r=p.data);return}m.nav.admin_plug_states(o)};return s.$$set=a=>{"pid"in a&&n(2,o=a.pid)},[r,g,o]}class C extends _{constructor(t){super(),l(this,t,$,S,u,{pid:2})}}export{C as component};