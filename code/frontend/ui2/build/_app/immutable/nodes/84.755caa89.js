import{s as E,a as C,c as N,i as y,d as v,P as R,A as P}from"../chunks/scheduler.e2ee220a.js";import{S,i as h,b as d,d as $,m as g,t as p,c as D,a as u,e as A,g as F}from"../chunks/index.4aee2103.js";import"../chunks/paths.2eaeb908.js";import{F as G}from"../chunks/floating_add.1495f8dc.js";import"../chunks/loading_spinner.svelte_svelte_type_style_lang.c28228c7.js";import{A as L}from"../chunks/autotable.0057460b.js";import{L as T}from"../chunks/loading_spinner.4ef87ddf.js";import{T as I}from"../chunks/top_actions.3bf62959.js";import{R as U,C as q}from"../chunks/reset_password.47e4c928.js";function H(r){let o,a;return o=new L({props:{action_key:"user_id",show_drop:!0,actions:[{Name:"Profile",Action:r[4],icon:"user-circle"},{Name:"Edit",Action:r[3],icon:"pencil-alt"},{Name:"Disable",Action:W,icon:"user-remove",drop:!0},{Name:"Reset Password",Action:r[6],icon:"lock-open",drop:!0},{Name:"Change Email",Action:r[7],icon:"at-symbol",drop:!0},{Name:"Roles",Action:r[8],icon:"identification",drop:!0},{Name:"Delete",Action:r[5],icon:"trash",drop:!0}],key_names:[["user_id","User Id"],["full_name","Full Name"],["group_id","Group"],["created_at","Created At"],["active","Active"]],color:["group_id","active"],datas:r[0]}}),{c(){d(o.$$.fragment)},l(e){$(o.$$.fragment,e)},m(e,n){g(o,e,n),a=!0},p(e,n){const s={};n&1&&(s.datas=e[0]),o.$set(s)},i(e){a||(u(o.$$.fragment,e),a=!0)},o(e){p(o.$$.fragment,e),a=!1},d(e){A(o,e)}}}function M(r){let o,a;return o=new T({}),{c(){d(o.$$.fragment)},l(e){$(o.$$.fragment,e)},m(e,n){g(o,e,n),a=!0},p:P,i(e){a||(u(o.$$.fragment,e),a=!0)},o(e){p(o.$$.fragment,e),a=!1},d(e){A(o,e)}}}function O(r){let o,a,e,n,s,c,m;o=new I({props:{actions:{"User Groups":r[10]}}});const b=[M,H],_=[];function k(t,l){return t[1]?0:1}return e=k(r),n=_[e]=b[e](r),c=new G({props:{onClick:r[9]}}),{c(){d(o.$$.fragment),a=C(),n.c(),s=C(),d(c.$$.fragment)},l(t){$(o.$$.fragment,t),a=N(t),n.l(t),s=N(t),$(c.$$.fragment,t)},m(t,l){g(o,t,l),y(t,a,l),_[e].m(t,l),y(t,s,l),g(c,t,l),m=!0},p(t,[l]){let f=e;e=k(t),e===f?_[e].p(t,l):(F(),p(_[f],1,1,()=>{_[f]=null}),D(),n=_[e],n?n.p(t,l):(n=_[e]=b[e](t),n.c()),u(n,1),n.m(s.parentNode,s))},i(t){m||(u(o.$$.fragment,t),u(n),u(c.$$.fragment,t),m=!0)},o(t){p(o.$$.fragment,t),p(n),p(c.$$.fragment,t),m=!1},d(t){t&&(v(a),v(s)),A(o,t),_[e].d(t),A(c,t)}}}const W=r=>{};function j(r,o,a){const e=R("__app__");let n=[],s=!0;const c=e.api_manager.get_admin_user_api(),m=async()=>{a(1,s=!0);const i=await c.list();i.ok&&(a(0,n=i.data),a(1,s=!1))};return m(),[n,s,e,i=>e.nav.admin_user_edit(i),i=>e.nav.user_profile(i),async i=>{await c.delete(i),m()},i=>{e.utils.small_modal_open(U,{uid:i,onComplete:w=>{console.log("RESET PASSWORD",w),e.utils.small_modal_close()}})},i=>{e.utils.small_modal_open(q,{uid:i,onComplete:w=>{console.log("CHANGE EMAIL",w),e.utils.small_modal_close()}})},i=>{},()=>e.nav.admin_user_new(),()=>e.nav.admin_ugroups()]}class te extends S{constructor(o){super(),h(this,o,j,O,E,{})}}export{te as component};
