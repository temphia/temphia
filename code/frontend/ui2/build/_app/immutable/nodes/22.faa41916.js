import{s as C,a as N,c as S,i as h,d as D,B as F,P as L,A as P}from"../chunks/scheduler.e2ee220a.js";import{S as $,i as q,b as g,d,m as b,t as _,c as B,a as f,e as k,g as E}from"../chunks/index.4aee2103.js";import"../chunks/paths.2eaeb908.js";import{F as j}from"../chunks/floating_add.1495f8dc.js";import"../chunks/loading_spinner.svelte_svelte_type_style_lang.c28228c7.js";import{A as z}from"../chunks/autotable.0057460b.js";import{L as G}from"../chunks/loading_spinner.4ef87ddf.js";import{p as H}from"../chunks/index.5458542a.js";function I(s){let t,n;return t=new z({props:{action_key:"slug",actions:[{Name:"Edit",Action:s[2],icon:"pencil-alt"},{Name:"Delete",Class:"bg-red-400",Action:s[3],icon:"trash"}],key_names:[["name","Name"],["slug","Slug"],["ctype","Column type"],["description","Description"]],color:[],datas:s[0]}}),{c(){g(t.$$.fragment)},l(e){d(t.$$.fragment,e)},m(e,o){b(t,e,o),n=!0},p(e,o){const i={};o&1&&(i.datas=e[0]),t.$set(i)},i(e){n||(f(t.$$.fragment,e),n=!0)},o(e){_(t.$$.fragment,e),n=!1},d(e){k(t,e)}}}function J(s){let t,n;return t=new G({}),{c(){g(t.$$.fragment)},l(e){d(t.$$.fragment,e)},m(e,o){b(t,e,o),n=!0},p:P,i(e){n||(f(t.$$.fragment,e),n=!0)},o(e){_(t.$$.fragment,e),n=!1},d(e){k(t,e)}}}function K(s){let t,n,e,o,i;const m=[J,I],c=[];function u(a,l){return a[1]?0:1}return t=u(s),n=c[t]=m[t](s),o=new j({props:{onClick:s[4]}}),{c(){n.c(),e=N(),g(o.$$.fragment)},l(a){n.l(a),e=S(a),d(o.$$.fragment,a)},m(a,l){c[t].m(a,l),h(a,e,l),b(o,a,l),i=!0},p(a,[l]){let p=t;t=u(a),t===p?c[t].p(a,l):(E(),_(c[p],1,1,()=>{c[p]=null}),B(),n=c[t],n?n.p(a,l):(n=c[t]=m[t](a),n.c()),f(n,1),n.m(e.parentNode,e))},i(a){i||(f(n),f(o.$$.fragment,a),i=!0)},o(a){_(n),_(o.$$.fragment,a),i=!1},d(a){a&&D(e),c[t].d(a),k(o,a)}}}function M(s,t,n){let e;F(s,H,r=>n(8,e=r));let{source:o=e.source}=t,{group:i=e.group}=t,{table:m=e.table}=t;const c=L("__app__");let u=[],a=!0;const l=c.api_manager.get_admin_data_api();(async()=>{const r=await l.list_column(o,i,m);r.ok&&(n(0,u=r.data),n(1,a=!1))})();const y=r=>c.nav.admin_data_column(o,i,m,r),w=async r=>{},A=()=>{};return s.$$set=r=>{"source"in r&&n(5,o=r.source),"group"in r&&n(6,i=r.group),"table"in r&&n(7,m=r.table)},[u,a,y,w,A,o,i,m]}class Y extends ${constructor(t){super(),q(this,t,M,K,C,{source:5,group:6,table:7})}}export{Y as component};
