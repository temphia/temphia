import{s as E,e as d,i as L,d as v,B as h,P as w,A as O}from"../chunks/scheduler.e2ee220a.js";import{S as P,i as X,t as c,c as A,a as f,g as B,b as g,d as y,m as k,e as b}from"../chunks/index.4aee2103.js";import{A as I}from"../chunks/auto_form.ee101f61.js";import"../chunks/paths.2eaeb908.js";import"../chunks/loading_spinner.svelte_svelte_type_style_lang.c28228c7.js";import{L as S}from"../chunks/loading_spinner.4ef87ddf.js";import{p as C}from"../chunks/index.5458542a.js";function N(s){let e,a;return e=new I({props:{message:s[0],schema:{fields:[{name:"Id",ftype:"TEXT_SLUG",key_name:"id",disabled:!0},{name:"Name",ftype:"TEXT",key_name:"name"},{name:"Live",ftype:"BOOL",key_name:"live"},{name:"Dev",ftype:"BOOL",key_name:"dev"},{name:"Bprint Id",ftype:"TEXT",key_name:"bprint_id"},{name:"Invoke Policy",ftype:"TEXT_POLICY",key_name:"invoke_policy"},{name:"Extra Meta",ftype:"KEY_VALUE_TEXT",key_name:"extra_meta"}],name:"New Plug",required_fields:["bprint_id"]},onSave:s[3],data:s[1]}}),{c(){g(e.$$.fragment)},l(t){y(e.$$.fragment,t)},m(t,r){k(e,t,r),a=!0},p(t,r){const i={};r&1&&(i.message=t[0]),r&2&&(i.data=t[1]),e.$set(i)},i(t){a||(f(e.$$.fragment,t),a=!0)},o(t){c(e.$$.fragment,t),a=!1},d(t){b(e,t)}}}function q(s){let e,a;return e=new S({}),{c(){g(e.$$.fragment)},l(t){y(e.$$.fragment,t)},m(t,r){k(e,t,r),a=!0},p:O,i(t){a||(f(e.$$.fragment,t),a=!0)},o(t){c(e.$$.fragment,t),a=!1},d(t){b(e,t)}}}function U(s){let e,a,t,r;const i=[q,N],o=[];function _(n,m){return n[2]?0:1}return e=_(s),a=o[e]=i[e](s),{c(){a.c(),t=d()},l(n){a.l(n),t=d()},m(n,m){o[e].m(n,m),L(n,t,m),r=!0},p(n,[m]){let l=e;e=_(n),e===l?o[e].p(n,m):(B(),c(o[l],1,1,()=>{o[l]=null}),A(),a=o[e],a?a.p(n,m):(a=o[e]=i[e](n),a.c()),f(a,1),a.m(t.parentNode,t))},i(n){r||(f(a),r=!0)},o(n){c(a),r=!1},d(n){n&&v(t),o[e].d(n)}}}function Y(s,e,a){let t;h(s,C,p=>a(4,t=p));let i=t.slug;const o=w("__app__"),_=o.api_manager.get_admin_plug_api();let n="",m={},l=!0;const $=async()=>{const p=await _.get_plug(i);if(!p.ok){a(0,n=p.data);return}a(1,m=p.data),a(2,l=!1)},T=async p=>{console.log("@@data",p);const u=await _.update_plug(i,p);if(!u.ok){a(0,n=u.data);return}o.nav.admin_plugs()};return $(),[n,m,l,T]}class z extends P{constructor(e){super(),X(this,e,Y,U,E,{})}}export{z as component};
