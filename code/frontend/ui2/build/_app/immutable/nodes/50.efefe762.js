import{s as k,a as y,c as w,i as A,d as C,B as I,P as N,A as P}from"../chunks/scheduler.e2ee220a.js";import{S,i as h,b as f,d as g,m as $,t as m,c as v,a as u,e as b,g as F}from"../chunks/index.4aee2103.js";import"../chunks/paths.555c70f8.js";import{F as L}from"../chunks/floating_add.1495f8dc.js";import"../chunks/loading_spinner.svelte_svelte_type_style_lang.9f9ec163.js";import{A as q}from"../chunks/autotable.0057460b.js";import{L as B}from"../chunks/loading_spinner.4ceee987.js";import{p as D}from"../chunks/index.37ddda1f.js";function E(r){let t,n;return t=new q({props:{action_key:"slug",actions:[{Name:"Edit",Action:r[2],icon:"pencil-alt"},{Name:"Delete",Class:"bg-red-400",Action:r[3],icon:"trash"}],key_names:[["slug","Slug"],["resource_id","Resource Id"],["plug_id","Plug Id"],["agent_id","Agent Id"]],color:["type"],datas:r[0]}}),{c(){f(t.$$.fragment)},l(e){g(t.$$.fragment,e)},m(e,o){$(t,e,o),n=!0},p(e,o){const i={};o&1&&(i.datas=e[0]),t.$set(i)},i(e){n||(u(t.$$.fragment,e),n=!0)},o(e){m(t.$$.fragment,e),n=!1},d(e){b(t,e)}}}function R(r){let t,n;return t=new B({}),{c(){f(t.$$.fragment)},l(e){g(t.$$.fragment,e)},m(e,o){$(t,e,o),n=!0},p:P,i(e){n||(u(t.$$.fragment,e),n=!0)},o(e){m(t.$$.fragment,e),n=!1},d(e){b(t,e)}}}function j(r){let t,n,e,o,i;const _=[R,E],s=[];function p(a,c){return a[1]?0:1}return t=p(r),n=s[t]=_[t](r),o=new L({props:{onClick:r[4]}}),{c(){n.c(),e=y(),f(o.$$.fragment)},l(a){n.l(a),e=w(a),g(o.$$.fragment,a)},m(a,c){s[t].m(a,c),A(a,e,c),$(o,a,c),i=!0},p(a,[c]){let d=t;t=p(a),t===d?s[t].p(a,c):(F(),m(s[d],1,1,()=>{s[d]=null}),v(),n=s[t],n?n.p(a,c):(n=s[t]=_[t](a),n.c()),u(n,1),n.m(e.parentNode,e))},i(a){i||(u(n),u(o.$$.fragment,a),i=!0)},o(a){m(n),m(o.$$.fragment,a),i=!1},d(a){a&&C(e),s[t].d(a),b(o,a)}}}function z(r,t,n){let e;I(r,D,l=>n(5,e=l));const o=N("__app__");let i=[],_=!0;const s=o.api_manager.get_admin_plug_api(),p=async()=>{const l=await s.list_agent_resource(e.pid,e.aid);l.ok&&(n(0,i=l.data),n(1,_=!1))};return p(),[i,_,l=>o.nav.admin_agent_res_edit(e.pid,e.aid,l),async l=>{(await s.delete_agent_resource(e.pid,e.aid,l)).ok&&p()},()=>o.nav.admin_agent_res_new(e.pid,e.aid)]}class V extends S{constructor(t){super(),h(this,t,z,j,k,{})}}export{V as component};