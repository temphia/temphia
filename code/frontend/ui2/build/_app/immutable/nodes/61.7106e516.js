import{s as b,e as d,i as E,d as T,B as h,P as v,A as L}from"../chunks/scheduler.e2ee220a.js";import{S as X,i as w,t as _,c as A,a as f,g as S,b as g,d as k,m as y,e as $}from"../chunks/index.4aee2103.js";import{A as P}from"../chunks/auto_form.ee101f61.js";import"../chunks/paths.2eaeb908.js";import"../chunks/loading_spinner.svelte_svelte_type_style_lang.c28228c7.js";import{L as U}from"../chunks/loading_spinner.4ef87ddf.js";import{p as q}from"../chunks/index.5458542a.js";function C(o){let e,a;return e=new P({props:{message:o[1],schema:{fields:[{name:"Id",ftype:"TEXT_SLUG",key_name:"id"},{name:"Name",ftype:"TEXT",key_name:"name"},{name:"Provider",ftype:"TEXT",key_name:"provider"},{name:"URL",ftype:"TEXT",key_name:"url"},{name:"Extra Meta",ftype:"KEY_VALUE_TEXT",key_name:"extra_meta"}],name:"Edit Repo",required_fields:[]},onSave:o[3],data:o[0]}}),{c(){g(e.$$.fragment)},l(t){k(e.$$.fragment,t)},m(t,n){y(e,t,n),a=!0},p(t,n){const s={};n&2&&(s.message=t[1]),n&1&&(s.data=t[0]),e.$set(s)},i(t){a||(f(e.$$.fragment,t),a=!0)},o(t){_(e.$$.fragment,t),a=!1},d(t){$(e,t)}}}function H(o){let e,a;return e=new U({}),{c(){g(e.$$.fragment)},l(t){k(e.$$.fragment,t)},m(t,n){y(e,t,n),a=!0},p:L,i(t){a||(f(e.$$.fragment,t),a=!0)},o(t){_(e.$$.fragment,t),a=!1},d(t){$(e,t)}}}function N(o){let e,a,t,n;const s=[H,C],m=[];function c(r,i){return r[2]?0:1}return e=c(o),a=m[e]=s[e](o),{c(){a.c(),t=d()},l(r){a.l(r),t=d()},m(r,i){m[e].m(r,i),E(r,t,i),n=!0},p(r,[i]){let p=e;e=c(r),e===p?m[e].p(r,i):(S(),_(m[p],1,1,()=>{m[p]=null}),A(),a=m[e],a?a.p(r,i):(a=m[e]=s[e](r),a.c()),f(a,1),a.m(t.parentNode,t))},i(r){n||(f(a),n=!0)},o(r){_(a),n=!1},d(r){r&&T(t),m[e].d(r)}}}function R(o,e,a){let t;h(o,q,l=>a(4,t=l));const n=v("__app__"),s=n.api_manager.get_admin_repo_api();let c=t.slug,r={},i="",p=!0;return(async()=>{const l=await s.get(c);if(l.status!==200){console.log("Err",l);return}a(0,r=l.data),a(2,p=!1)})(),[r,i,p,async l=>{const u=await s.update(c,l);if(!u.ok){console.log("Err",u),a(1,i=u.data);return}n.nav.admin_repos()}]}class z extends X{constructor(e){super(),w(this,e,R,N,b,{})}}export{z as component};
