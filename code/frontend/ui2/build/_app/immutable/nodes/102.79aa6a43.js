import{s as v,a as A,c as D,i as k,d as b,P as C,f as L,g as N,h as x,j as E,A as h}from"../chunks/scheduler.e2ee220a.js";import{S as y,i as w,b as m,d as _,m as p,g as F,t as f,c as I,a as u,e as d}from"../chunks/index.4aee2103.js";import"../chunks/paths.2eaeb908.js";import{F as P}from"../chunks/floating_add.1495f8dc.js";import"../chunks/loading_spinner.svelte_svelte_type_style_lang.c28228c7.js";import{A as S}from"../chunks/autotable.0057460b.js";import{L as j}from"../chunks/loading_spinner.4ef87ddf.js";function q(c){let e,t,n;return t=new S({props:{action_key:"id",actions:[{Name:"Delete",Class:"bg-red-400",Action:c[5]}],key_names:[["id","ID"],["name","Name"],["device_type","Device Type"],["last_addr","Last Address"]],color:["device_type"],datas:c[0]}}),{c(){e=L("div"),m(t.$$.fragment),this.h()},l(a){e=N(a,"DIV",{class:!0});var s=x(e);_(t.$$.fragment,s),s.forEach(b),this.h()},h(){E(e,"class","p-4 w-full h-full bg-indigo-100")},m(a,s){k(a,e,s),p(t,e,null),n=!0},p(a,s){const l={};s&1&&(l.datas=a[0]),t.$set(l)},i(a){n||(u(t.$$.fragment,a),n=!0)},o(a){f(t.$$.fragment,a),n=!1},d(a){a&&b(e),d(t)}}}function T(c){let e,t;return e=new j({}),{c(){m(e.$$.fragment)},l(n){_(e.$$.fragment,n)},m(n,a){p(e,n,a),t=!0},p:h,i(n){t||(u(e.$$.fragment,n),t=!0)},o(n){f(e.$$.fragment,n),t=!1},d(n){d(e,n)}}}function V(c){let e,t,n,a,s;const l=[T,q],i=[];function $(r,o){return r[1]?0:1}return e=$(c),t=i[e]=l[e](c),a=new P({props:{onClick:c[6]}}),{c(){t.c(),n=A(),m(a.$$.fragment)},l(r){t.l(r),n=D(r),_(a.$$.fragment,r)},m(r,o){i[e].m(r,o),k(r,n,o),p(a,r,o),s=!0},p(r,[o]){let g=e;e=$(r),e===g?i[e].p(r,o):(F(),f(i[g],1,1,()=>{i[g]=null}),I(),t=i[e],t?t.p(r,o):(t=i[e]=l[e](r),t.c()),u(t,1),t.m(n.parentNode,n))},i(r){s||(u(t),u(a.$$.fragment,r),s=!0)},o(r){f(t),f(a.$$.fragment,r),s=!1},d(r){r&&b(n),i[e].d(r),d(a,r)}}}function z(c,e,t){const n=C("__app__"),a=n.api_manager.get_self_api();let s=[],l=!0;const i=async()=>{t(1,l=!0);const o=await a.list_devices();if(o.status!==200){console.log("Err",o);return}t(0,s=o.data),t(1,l=!1)};return i(),[s,l,n,a,i,async o=>{await a.delete_device(o),i()},()=>n.nav.self_device_new()]}class B extends y{constructor(e){super(),w(this,e,z,V,v,{})}}function G(c){let e,t;return e=new B({}),{c(){m(e.$$.fragment)},l(n){_(e.$$.fragment,n)},m(n,a){p(e,n,a),t=!0},p:h,i(n){t||(u(e.$$.fragment,n),t=!0)},o(n){f(e.$$.fragment,n),t=!1},d(n){d(e,n)}}}class U extends y{constructor(e){super(),w(this,e,null,G,v,{})}}export{U as component};
