import{s as L,e as C,i as E,d as b,B as q,P as B,r as N,f as h,a as V,l as z,g as v,h as y,C as F,c as D,m as G,j as w,w as g,n as H,A as J}from"../chunks/scheduler.e2ee220a.js";import{S as K,i as M,t as k,c as O,a as x,g as Q,b as I,d as P,m as S,e as j}from"../chunks/index.4aee2103.js";import"../chunks/paths.2eaeb908.js";import"../chunks/loading_spinner.svelte_svelte_type_style_lang.c28228c7.js";import{L as R}from"../chunks/loading_spinner.4ef87ddf.js";import{p as T}from"../chunks/index.5458542a.js";import{V as U}from"../chunks/_value_inner.d4db59dc.js";function W(m){let t,e,a,r="Add View",_,o,i=(m[0]||"")+"",s,c,l,f;return l=new U({props:{columns:m[2],data:{},onSave:m[3]}}),{c(){t=h("div"),e=h("div"),a=h("div"),a.textContent=r,_=V(),o=h("p"),s=z(i),c=V(),I(l.$$.fragment),this.h()},l(u){t=v(u,"DIV",{class:!0});var d=y(t);e=v(d,"DIV",{class:!0});var n=y(e);a=v(n,"DIV",{class:!0,"data-svelte-h":!0}),F(a)!=="svelte-1wjomul"&&(a.textContent=r),_=D(n),o=v(n,"P",{class:!0});var p=y(o);s=G(p,i),p.forEach(b),c=D(n),P(l.$$.fragment,n),n.forEach(b),d.forEach(b),this.h()},h(){w(a,"class","text-2xl text-indigo-900"),w(o,"class","text-red-500"),w(e,"class","p-5 bg-white w-full "),w(t,"class","h-full w-full bg-indigo-100 p-10 overflow-auto")},m(u,d){E(u,t,d),g(t,e),g(e,a),g(e,_),g(e,o),g(o,s),g(e,c),S(l,e,null),f=!0},p(u,d){(!f||d&1)&&i!==(i=(u[0]||"")+"")&&H(s,i);const n={};d&4&&(n.columns=u[2]),l.$set(n)},i(u){f||(x(l.$$.fragment,u),f=!0)},o(u){k(l.$$.fragment,u),f=!1},d(u){u&&b(t),j(l)}}}function X(m){let t,e;return t=new R({}),{c(){I(t.$$.fragment)},l(a){P(t.$$.fragment,a)},m(a,r){S(t,a,r),e=!0},p:J,i(a){e||(x(t.$$.fragment,a),e=!0)},o(a){k(t.$$.fragment,a),e=!1},d(a){j(t,a)}}}function Y(m){let t,e,a,r;const _=[X,W],o=[];function i(s,c){return s[1]?0:1}return t=i(m),e=o[t]=_[t](m),{c(){e.c(),a=C()},l(s){e.l(s),a=C()},m(s,c){o[t].m(s,c),E(s,a,c),r=!0},p(s,[c]){let l=t;t=i(s),t===l?o[t].p(s,c):(Q(),k(o[l],1,1,()=>{o[l]=null}),O(),e=o[t],e?e.p(s,c):(e=o[t]=_[t](s),e.c()),x(e,1),e.m(a.parentNode,a))},i(s){r||(x(e),r=!0)},o(s){k(e),r=!1},d(s){s&&b(a),o[t].d(s)}}}function Z(m,t,e){let a;q(m,T,n=>e(7,a=n));let{source:r=a.source}=t,{group:_=a.group}=t,{table:o=a.table}=t;const i=B("__app__"),s=i.api_manager.get_admin_data_api();let c="",l=!0,f=[];const u=async n=>{n.selects||(n.selects=f.map(A=>A.slug));const p=await s.add_view(r,_,o,n);if(!p.ok){e(0,c=p.data);return}i.nav.admin_data_views(r,_,o)};return(async()=>{const p=await i.api_manager.get_admin_data_api().list_column(r,_,o);if(!p.ok){e(0,c=p.data),e(1,l=!1);return}e(2,f=p.data),e(1,l=!1)})(),N("__data_context__",{get_modal:()=>({open:i.utils.small_modal_open,close:i.utils.small_modal_close}),table_service:null}),m.$$set=n=>{"source"in n&&e(4,r=n.source),"group"in n&&e(5,_=n.group),"table"in n&&e(6,o=n.table)},[c,l,f,u,r,_,o]}class le extends K{constructor(t){super(),M(this,t,Z,Y,L,{source:4,group:5,table:6})}}export{le as component};
