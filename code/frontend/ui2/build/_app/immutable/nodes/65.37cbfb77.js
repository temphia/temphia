import{s as b,a as k,c as y,i as v,d as w,P as A,A as C}from"../chunks/scheduler.e2ee220a.js";import{S as N,i as P,b as u,d,m as $,t as p,c as D,a as m,e as g,g as E}from"../chunks/index.4aee2103.js";import"../chunks/paths.3df37c61.js";import{F}from"../chunks/floating_add.1495f8dc.js";import"../chunks/loading_spinner.svelte_svelte_type_style_lang.0368a52d.js";import{A as L}from"../chunks/autotable.0057460b.js";import{L as S}from"../chunks/loading_spinner.ede256fc.js";function h(i){let e,n;return e=new L({props:{action_key:"id",actions:[{Name:"Edit",Action:i[5]},{Name:"Delete",Class:"bg-red-400",Action:i[6]}],key_names:[["id","ID"],["name","Name"],["provider","Provider"]],color:["provider"],datas:i[0]}}),{c(){u(e.$$.fragment)},l(t){d(e.$$.fragment,t)},m(t,r){$(e,t,r),n=!0},p(t,r){const c={};r&1&&(c.datas=t[0]),e.$set(c)},i(t){n||(m(e.$$.fragment,t),n=!0)},o(t){p(e.$$.fragment,t),n=!1},d(t){g(e,t)}}}function q(i){let e,n;return e=new S({}),{c(){u(e.$$.fragment)},l(t){d(e.$$.fragment,t)},m(t,r){$(e,t,r),n=!0},p:C,i(t){n||(m(e.$$.fragment,t),n=!0)},o(t){p(e.$$.fragment,t),n=!1},d(t){g(e,t)}}}function I(i){let e,n,t,r,c;const l=[q,h],s=[];function f(a,o){return a[1]?0:1}return e=f(i),n=s[e]=l[e](i),r=new F({props:{onClick:i[2].nav.admin_repo_new}}),{c(){n.c(),t=k(),u(r.$$.fragment)},l(a){n.l(a),t=y(a),d(r.$$.fragment,a)},m(a,o){s[e].m(a,o),v(a,t,o),$(r,a,o),c=!0},p(a,[o]){let _=e;e=f(a),e===_?s[e].p(a,o):(E(),p(s[_],1,1,()=>{s[_]=null}),D(),n=s[e],n?n.p(a,o):(n=s[e]=l[e](a),n.c()),m(n,1),n.m(t.parentNode,t))},i(a){c||(m(n),m(r.$$.fragment,a),c=!0)},o(a){p(n),p(r.$$.fragment,a),c=!1},d(a){a&&w(t),s[e].d(a),g(r,a)}}}function j(i,e,n){const t=A("__app__"),r=t.api_manager.get_admin_repo_api();let c=[],l=!0;const s=async()=>{const o=await r.list();if(o.status!==200){console.log("Err",o);return}n(0,c=o.data),n(1,l=!1)};return s(),[c,l,t,r,s,o=>t.nav.admin_repo_edit(o),async o=>{await r.delete(o),s()}]}class O extends N{constructor(e){super(),P(this,e,j,I,b,{})}}export{O as component};