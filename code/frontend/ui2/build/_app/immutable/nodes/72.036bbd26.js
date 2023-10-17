import{s as R,f as v,a as A,g as h,h as I,c as k,C as T,d as b,j as m,i as L,w as _,P as G}from"../chunks/scheduler.e2ee220a.js";import{S as W,i as F,b as j,d as B,m as z,a as N,t as S,e as H,c as J,g as M}from"../chunks/index.4aee2103.js";import"../chunks/paths.2eaeb908.js";import"../chunks/loading_spinner.svelte_svelte_type_style_lang.c28228c7.js";import{A as Q}from"../chunks/autotable.0057460b.js";import{A as X}from"../chunks/action_button.2a2ac158.js";import{T as Y}from"../chunks/top_actions.3bf62959.js";function Z(o){let t,a;return t=new X({props:{onClick:o[0],name:o[1],icon_name:"plus"}}),{c(){j(t.$$.fragment)},l(e){B(t.$$.fragment,e)},m(e,l){z(t,e,l),a=!0},p(e,[l]){const n={};l&1&&(n.onClick=e[0]),l&2&&(n.name=e[1]),t.$set(n)},i(e){a||(N(t.$$.fragment,e),a=!0)},o(e){S(t.$$.fragment,e),a=!1},d(e){H(t,e)}}}function tt(o,t,a){let{onClick:e}=t,{name:l="Add"}=t;return o.$$set=n=>{"onClick"in n&&a(0,e=n.onClick),"name"in n&&a(1,l=n.name)},[e,l]}class et extends W{constructor(t){super(),F(this,t,tt,Z,R,{onClick:0,name:1})}}function O(o){let t,a,e="Name",l,n,p,c,u,C="Slug",s,i,d,f,$,w,g,E,q="Domains",U,x,P;return w=new et({props:{onClick:o[8]}}),x=new Q({props:{action_key:"id",show_drop:!0,actions:[{Name:"Adapter Editor",Action:o[9],icon:"lightning-bolt"},{Name:"Edit",Action:o[10],icon:"pencil"},{Name:"Hooks",Action:at,drop:!0,icon:"hashtag"},{Name:"Widgets",Action:st,drop:!0,icon:"hashtag"},{Name:"Reset",drop:!0,icon:"refresh",Action:o[11]},{Name:"Delete",drop:!0,icon:"trash",Action:o[12]}],key_names:[["id","ID"],["name","Name"],["adapter_type","Http Adapter"],["about","About"]],datas:o[0],color:["adapter_type"]}}),{c(){t=v("div"),a=v("label"),a.textContent=e,l=A(),n=v("input"),p=A(),c=v("div"),u=v("label"),u.textContent=C,s=A(),i=v("input"),d=A(),f=v("div"),$=v("div"),j(w.$$.fragment),g=A(),E=v("label"),E.textContent=q,U=A(),j(x.$$.fragment),this.h()},l(r){t=h(r,"DIV",{class:!0});var y=I(t);a=h(y,"LABEL",{class:!0,"data-svelte-h":!0}),T(a)!=="svelte-5pozpj"&&(a.textContent=e),l=k(y),n=h(y,"INPUT",{type:!0,class:!0}),y.forEach(b),p=k(r),c=h(r,"DIV",{class:!0});var D=I(c);u=h(D,"LABEL",{class:!0,"data-svelte-h":!0}),T(u)!=="svelte-bads9j"&&(u.textContent=C),s=k(D),i=h(D,"INPUT",{type:!0,class:!0}),D.forEach(b),d=k(r),f=h(r,"DIV",{class:!0});var V=I(f);$=h(V,"DIV",{class:!0});var K=I($);B(w.$$.fragment,K),K.forEach(b),g=k(V),E=h(V,"LABEL",{class:!0,"data-svelte-h":!0}),T(E)!=="svelte-th1ugr"&&(E.textContent=q),U=k(V),B(x.$$.fragment,V),V.forEach(b),this.h()},h(){m(a,"class","pb-2 text-gray-700 font-semibold"),m(n,"type","text"),n.disabled=!0,n.value=o[4].name||"",m(n,"class","p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"),m(t,"class","flex-col flex py-3"),m(u,"class","pb-2 text-gray-700 font-semibold"),m(i,"type","text"),i.value=o[4].slug,i.disabled=!0,m(i,"class","p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"),m(c,"class","flex-col flex py-3 relative"),m($,"class","absolute right-1"),m(E,"class","pb-2 text-gray-700 font-semibold"),m(f,"class","flex-col flex py-3 relative border rounded p-2")},m(r,y){L(r,t,y),_(t,a),_(t,l),_(t,n),L(r,p,y),L(r,c,y),_(c,u),_(c,s),_(c,i),L(r,d,y),L(r,f,y),_(f,$),z(w,$,null),_(f,g),_(f,E),_(f,U),z(x,f,null),P=!0},p(r,y){const D={};y&1&&(D.datas=r[0]),x.$set(D)},i(r){P||(N(w.$$.fragment,r),N(x.$$.fragment,r),P=!0)},o(r){S(w.$$.fragment,r),S(x.$$.fragment,r),P=!1},d(r){r&&(b(t),b(p),b(c),b(d),b(f)),H(w),H(x)}}}function nt(o){let t,a,e,l,n,p,c="Organization",u,C;a=new Y({props:{actions:{"System KV":o[6],"System Event":o[7]}}});let s=o[1]&&O(o);return{c(){t=v("div"),j(a.$$.fragment),e=A(),l=v("div"),n=v("div"),p=v("div"),p.textContent=c,u=A(),s&&s.c(),this.h()},l(i){t=h(i,"DIV",{class:!0});var d=I(t);B(a.$$.fragment,d),e=k(d),l=h(d,"DIV",{class:!0});var f=I(l);n=h(f,"DIV",{class:!0});var $=I(n);p=h($,"DIV",{class:!0,"data-svelte-h":!0}),T(p)!=="svelte-1omoan6"&&(p.textContent=c),u=k($),s&&s.l($),$.forEach(b),f.forEach(b),d.forEach(b),this.h()},h(){m(p,"class","text-2xl text-indigo-900"),m(n,"class","md:w-1/2-screen m-0 p-5 bg-white w-full tw-h-full shadow md:rounded-lg relative"),m(l,"class","md:p-8 bg-indigo-100 flex flex-row flex-wrap"),m(t,"class","h-full w-full overflow-auto")},m(i,d){L(i,t,d),z(a,t,null),_(t,e),_(t,l),_(l,n),_(n,p),_(n,u),s&&s.m(n,null),C=!0},p(i,[d]){i[1]?s?(s.p(i,d),d&2&&N(s,1)):(s=O(i),s.c(),N(s,1),s.m(n,null)):s&&(M(),S(s,1,1,()=>{s=null}),J())},i(i){C||(N(a.$$.fragment,i),N(s),C=!0)},o(i){S(a.$$.fragment,i),S(s),C=!1},d(i){i&&b(t),H(a),s&&s.d()}}}const at=o=>{},st=o=>{};function ot(o,t,a){const e=G("__app__"),l=e.api_manager.get_admin_tenant_api();let n=[],p={},c=!1;const u=async()=>{const g=await l.get_domains();g.ok&&(a(0,n=g.data),a(1,c=!0))};return u(),[n,c,e,l,p,u,()=>e.nav.admin_tenant_system_kvs(),()=>e.nav.admin_tenant_system_events(),()=>e.nav.admin_tenant_domain_new(),g=>e.nav.admin_tenant_domain_adapter_editor(g),g=>e.nav.admin_tenant_domain_edit(g),g=>{l.domain_adapter_reset(g)},async g=>{await l.delete_domain(g),u()}]}class mt extends W{constructor(t){super(),F(this,t,ot,nt,R,{})}}export{mt as component};
