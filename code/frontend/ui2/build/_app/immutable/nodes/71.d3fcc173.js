import{s as E,e as d,i as $,d as h,B as A,P as L,A as X}from"../chunks/scheduler.e2ee220a.js";import{S as w,i as C,t as f,c as P,a as l,g as S,b as g,d as y,m as k,e as b}from"../chunks/index.4aee2103.js";import{A as N}from"../chunks/auto_form.ee101f61.js";import"../chunks/paths.9c1b57c4.js";import"../chunks/loading_spinner.svelte_svelte_type_style_lang.0368a52d.js";import{L as O}from"../chunks/loading_spinner.ede256fc.js";import{p as U}from"../chunks/index.5ed445fc.js";function q(r){let e,a;return e=new N({props:{message:r[0],schema:{fields:[{name:"Name",ftype:"TEXT",key_name:"name"},{name:"About",ftype:"LONG_TEXT",key_name:"about"},{name:"Default User Group",ftype:"TEXT",key_name:"default_ugroup"},{name:"CORS Policy",ftype:"TEXT_POLICY",key_name:"cors_policy"},{name:"Extra Meta",ftype:"KEY_VALUE_TEXT",key_name:"extra_meta"}],name:"Update Domain",required_fields:[]},onSave:r[3],data:r[1]}}),{c(){g(e.$$.fragment)},l(t){y(e.$$.fragment,t)},m(t,o){k(e,t,o),a=!0},p(t,o){const m={};o&1&&(m.message=t[0]),o&2&&(m.data=t[1]),e.$set(m)},i(t){a||(l(e.$$.fragment,t),a=!0)},o(t){f(e.$$.fragment,t),a=!1},d(t){b(e,t)}}}function v(r){let e,a;return e=new O({}),{c(){g(e.$$.fragment)},l(t){y(e.$$.fragment,t)},m(t,o){k(e,t,o),a=!0},p:X,i(t){a||(l(e.$$.fragment,t),a=!0)},o(t){f(e.$$.fragment,t),a=!1},d(t){b(e,t)}}}function D(r){let e,a,t,o;const m=[v,q],s=[];function c(n,_){return n[2]?0:1}return e=c(r),a=s[e]=m[e](r),{c(){a.c(),t=d()},l(n){a.l(n),t=d()},m(n,_){s[e].m(n,_),$(n,t,_),o=!0},p(n,[_]){let p=e;e=c(n),e===p?s[e].p(n,_):(S(),f(s[p],1,1,()=>{s[p]=null}),P(),a=s[e],a?a.p(n,_):(a=s[e]=m[e](n),a.c()),l(a,1),a.m(t.parentNode,t))},i(n){o||(l(a),o=!0)},o(n){f(a),o=!1},d(n){n&&h(t),s[e].d(n)}}}function G(r,e,a){let t;A(r,U,i=>a(6,t=i));let{did:o=t.did}=e;const m=L("__app__"),s=m.api_manager.get_admin_tenant_api();let c="",n={},_=!0;m.api_manager.self_api.list_adapter_providers().then(i=>{i.data}).catch(()=>{}),(async()=>{const i=await s.get_domain(o);i.ok&&(a(1,n=i.data),a(2,_=!1))})();const T=async i=>{const u=await s.edit_domain(o,i);if(!u.ok){a(0,c=u.data);return}m.nav.admin_tenant()};return r.$$set=i=>{"did"in i&&a(4,o=i.did)},[c,n,_,T,o]}class V extends w{constructor(e){super(),C(this,e,G,D,E,{did:4})}}export{V as component};
