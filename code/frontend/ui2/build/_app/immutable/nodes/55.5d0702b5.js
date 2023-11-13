import{s as nt,e as bt,i as I,d as b,f as x,g as v,C as H,j as k,A as j,a as U,h as B,c as N,w as p,D as q,J as $,l as F,m as z,B as wt,P as Et,k as Lt,Q as Dt,n as Q}from"../chunks/scheduler.e2ee220a.js";import{S as st,i as at,g as ot,t as V,c as rt,a as A,b as Y,d as G,m as W,e as X}from"../chunks/index.4aee2103.js";import{e as xt}from"../chunks/each.e59479a4.js";import{p as Ut}from"../chunks/index.5ed445fc.js";import"../chunks/paths.9c1b57c4.js";import{F as Nt}from"../chunks/floating_add.1495f8dc.js";import"../chunks/loading_spinner.svelte_svelte_type_style_lang.0368a52d.js";import{L as it}from"../chunks/loading_spinner.ede256fc.js";import{T as Ot}from"../chunks/top_actions.3bf62959.js";function St(o){let t,e='<h3 class="mb-2 text-2xl font-bold text-gray-800">Done</h3>';return{c(){t=x("div"),t.innerHTML=e,this.h()},l(l){t=v(l,"DIV",{class:!0,"data-svelte-h":!0}),H(t)!=="svelte-1r1jfrs"&&(t.innerHTML=e),this.h()},h(){k(t,"class","p-2 text-center overflow-y-auto")},m(l,n){I(l,t,n)},p:j,i:j,o:j,d(l){l&&b(t)}}}function Bt(o){let t,e,l="Export State",n,i,f="Do you want to export plug states ?",u,s,a,r="Ok",_,d,C="Cancel",m,O;return{c(){t=x("div"),e=x("h3"),e.textContent=l,n=U(),i=x("p"),i.textContent=f,u=U(),s=x("div"),a=x("button"),a.textContent=r,_=U(),d=x("button"),d.textContent=C,this.h()},l(E){t=v(E,"DIV",{class:!0});var L=B(t);e=v(L,"H3",{class:!0,"data-svelte-h":!0}),H(e)!=="svelte-1ss3mvb"&&(e.textContent=l),n=N(L),i=v(L,"P",{class:!0,"data-svelte-h":!0}),H(i)!=="svelte-ie6lge"&&(i.textContent=f),u=N(L),s=v(L,"DIV",{class:!0});var y=B(s);a=v(y,"BUTTON",{class:!0,"data-svelte-h":!0}),H(a)!=="svelte-vn5fsz"&&(a.textContent=r),_=N(y),d=v(y,"BUTTON",{class:!0,"data-svelte-h":!0}),H(d)!=="svelte-13q4hqb"&&(d.textContent=C),y.forEach(b),L.forEach(b),this.h()},h(){k(e,"class","mb-2 text-2xl font-bold text-gray-800"),k(i,"class","text-gray-500"),k(a,"class","btn variant-filled-primary"),k(d,"class","btn variant-filled-secondary"),k(s,"class","mt-6 flex justify-end gap-2"),k(t,"class","p-2 text-center overflow-y-auto")},m(E,L){I(E,t,L),p(t,e),p(t,n),p(t,i),p(t,u),p(t,s),p(s,a),p(s,_),p(s,d),m||(O=[q(a,"click",o[5]),q(d,"click",o[6])],m=!0)},p:j,i:j,o:j,d(E){E&&b(t),m=!1,$(O)}}}function It(o){let t,e;return t=new it({props:{classes:""}}),{c(){Y(t.$$.fragment)},l(l){G(t.$$.fragment,l)},m(l,n){W(t,l,n),e=!0},p:j,i(l){e||(A(t.$$.fragment,l),e=!0)},o(l){V(t.$$.fragment,l),e=!1},d(l){X(t,l)}}}function jt(o){let t,e,l,n;const i=[It,Bt,St],f=[];function u(s,a){return s[1]?0:s[2]==="SELECT"?1:2}return t=u(o),e=f[t]=i[t](o),{c(){e.c(),l=bt()},l(s){e.l(s),l=bt()},m(s,a){f[t].m(s,a),I(s,l,a),n=!0},p(s,[a]){let r=t;t=u(s),t===r?f[t].p(s,a):(ot(),V(f[r],1,1,()=>{f[r]=null}),rt(),e=f[t],e?e.p(s,a):(e=f[t]=i[t](s),e.c()),A(e,1),e.m(l.parentNode,l))},i(s){n||(A(e),n=!0)},o(s){V(e),n=!1},d(s){s&&b(l),f[t].d(s)}}}function Ht(o,t,e){let{id:l}=t,{app:n}=t,i=!1,f="SELECT";const u=async()=>{e(1,i=!0);const _=await n.api_manager.get_admin_plug_api().export_plug_state(l);if(!_.ok)return;let d=document.createElement("a");document.body.appendChild(d);let C=_.data;typeof C=="object"&&(C=JSON.stringify(_.data,null,4));let m=window.URL.createObjectURL(new Blob([C]));d.href=m,d.download="plug_state.json",d.click(),window.URL.revokeObjectURL(m),e(2,f="RESULT"),e(1,i=!1)},s=()=>u(),a=()=>n.utils.small_modal_close();return o.$$set=r=>{"id"in r&&e(4,l=r.id),"app"in r&&e(0,n=r.app)},[n,i,f,u,l,s,a]}class Pt extends st{constructor(t){super(),at(this,t,Ht,jt,nt,{id:4,app:0})}}function Rt(o){let t,e="Done";return{c(){t=x("h3"),t.textContent=e,this.h()},l(l){t=v(l,"H3",{class:!0,"data-svelte-h":!0}),H(t)!=="svelte-ulkybn"&&(t.textContent=e),this.h()},h(){k(t,"class","mb-2 text-2xl font-bold text-gray-800 text-center")},m(l,n){I(l,t,n)},p:j,i:j,o:j,d(l){l&&b(t)}}}function Vt(o){let t,e="Import State",l,n,i,f,u,s,a,r,_,d,C,m,O="Cancel",E,L,y=o[3]&&vt(o);return{c(){t=x("h3"),t.textContent=e,l=U(),n=x("label"),i=F(`JSON file
      `),f=x("input"),u=U(),s=x("label"),a=F(`Clean Previous States
      `),r=x("input"),_=U(),d=x("div"),y&&y.c(),C=U(),m=x("button"),m.textContent=O,this.h()},l(g){t=v(g,"H3",{class:!0,"data-svelte-h":!0}),H(t)!=="svelte-19hxypx"&&(t.textContent=e),l=N(g),n=v(g,"LABEL",{});var h=B(n);i=z(h,`JSON file
      `),f=v(h,"INPUT",{type:!0}),h.forEach(b),u=N(g),s=v(g,"LABEL",{});var c=B(s);a=z(c,`Clean Previous States
      `),r=v(c,"INPUT",{type:!0}),c.forEach(b),_=N(g),d=v(g,"DIV",{class:!0});var T=B(d);y&&y.l(T),C=N(T),m=v(T,"BUTTON",{class:!0,"data-svelte-h":!0}),H(m)!=="svelte-13q4hqb"&&(m.textContent=O),T.forEach(b),this.h()},h(){k(t,"class","mb-2 text-center text-2xl font-bold text-gray-800"),k(f,"type","file"),k(r,"type","checkbox"),k(m,"class","btn variant-filled-secondary"),k(d,"class","mt-6 flex justify-end gap-2")},m(g,h){I(g,t,h),I(g,l,h),I(g,n,h),p(n,i),p(n,f),I(g,u,h),I(g,s,h),p(s,a),p(s,r),r.checked=o[2],I(g,_,h),I(g,d,h),y&&y.m(d,null),p(d,C),p(d,m),E||(L=[q(f,"change",o[7]),q(r,"change",o[8]),q(m,"click",o[9])],E=!0)},p(g,h){h&4&&(r.checked=g[2]),g[3]?y?y.p(g,h):(y=vt(g),y.c(),y.m(d,C)):y&&(y.d(1),y=null)},i:j,o:j,d(g){g&&(b(t),b(l),b(n),b(u),b(s),b(_),b(d)),y&&y.d(),E=!1,$(L)}}}function At(o){let t,e;return t=new it({props:{classes:""}}),{c(){Y(t.$$.fragment)},l(l){G(t.$$.fragment,l)},m(l,n){W(t,l,n),e=!0},p:j,i(l){e||(A(t.$$.fragment,l),e=!0)},o(l){V(t.$$.fragment,l),e=!1},d(l){X(t,l)}}}function vt(o){let t,e="Ok",l,n;return{c(){t=x("button"),t.textContent=e,this.h()},l(i){t=v(i,"BUTTON",{class:!0,"data-svelte-h":!0}),H(t)!=="svelte-1a1iugl"&&(t.textContent=e),this.h()},h(){k(t,"class","btn variant-filled-primary")},m(i,f){I(i,t,f),l||(n=q(t,"click",o[5]),l=!0)},p:j,d(i){i&&b(t),l=!1,n()}}}function qt(o){let t,e,l,n;const i=[At,Vt,Rt],f=[];function u(s,a){return s[1]?0:s[4]==="SELECT"?1:2}return e=u(o),l=f[e]=i[e](o),{c(){t=x("div"),l.c(),this.h()},l(s){t=v(s,"DIV",{class:!0});var a=B(t);l.l(a),a.forEach(b),this.h()},h(){k(t,"class","p-4 text-center overflow-y-auto flex flex-col gap-2")},m(s,a){I(s,t,a),f[e].m(t,null),n=!0},p(s,[a]){let r=e;e=u(s),e===r?f[e].p(s,a):(ot(),V(f[r],1,1,()=>{f[r]=null}),rt(),l=f[e],l?l.p(s,a):(l=f[e]=i[e](s),l.c()),A(l,1),l.m(t,null))},i(s){n||(A(l),n=!0)},o(s){V(l),n=!1},d(s){s&&b(t),f[e].d()}}}function Jt(o,t,e){let{id:l}=t,{app:n}=t,i=!1,f=!1,u,s="SELECT";const a=async()=>{e(1,i=!0),!(!u||!(await n.api_manager.get_admin_plug_api().import_plug_state(l,f,u)).ok)&&(e(4,s="RESULT"),e(1,i=!1))},r=C=>{const m=C.target.files[0];if(m){const O=new FileReader;O.onload=E=>{e(3,u=E.target.result),console.log(u)},O.readAsText(m)}};function _(){f=this.checked,e(2,f)}const d=()=>n.utils.small_modal_close();return o.$$set=C=>{"id"in C&&e(6,l=C.id),"app"in C&&e(0,n=C.app)},[n,i,f,u,s,a,l,r,_,d]}class Mt extends st{constructor(t){super(),at(this,t,Jt,qt,nt,{id:6,app:0})}}function yt(o,t,e){const l=o.slice();return l[18]=t[e],l}function Ft(o){let t,e,l,n,i,f='<tr><th scope="col" class="text-sm font-medium text-gray-900 px-2 py-2 text-left">Key</th> <th scope="col" class="text-sm font-medium text-gray-900 px-2 py-2 text-left">Version</th> <th scope="col" class="text-sm font-medium text-gray-900 px-2 py-2 text-left">Tag1</th> <th scope="col" class="text-sm font-medium text-gray-900 px-2 py-2 text-left">Tag2</th> <th scope="col" class="text-sm font-medium text-gray-900 px-2 py-2 text-left">Tag3</th> <th scope="col" class="text-sm font-medium text-gray-900 px-2 py-2 text-left">TTL</th> <th scope="col" class="text-sm font-medium text-gray-900 px-2 py-2 text-left">Actions</th></tr>',u,s,a,r,_,d="Previous",C,m,O="Next",E,L,y;t=new Ot({props:{actions:{Import:o[10],Export:o[11]}}});let g=xt(o[1]),h=[];for(let c=0;c<g.length;c+=1)h[c]=kt(yt(o,g,c));return{c(){Y(t.$$.fragment),e=U(),l=x("div"),n=x("table"),i=x("thead"),i.innerHTML=f,u=U(),s=x("tbody");for(let c=0;c<h.length;c+=1)h[c].c();a=U(),r=x("div"),_=x("button"),_.textContent=d,C=U(),m=x("button"),m.textContent=O,this.h()},l(c){G(t.$$.fragment,c),e=N(c),l=v(c,"DIV",{class:!0});var T=B(l);n=v(T,"TABLE",{class:!0});var S=B(n);i=v(S,"THEAD",{class:!0,"data-svelte-h":!0}),H(i)!=="svelte-l2hper"&&(i.innerHTML=f),u=N(S),s=v(S,"TBODY",{});var D=B(s);for(let P=0;P<h.length;P+=1)h[P].l(D);D.forEach(b),S.forEach(b),a=N(T),r=v(T,"DIV",{class:!0});var R=B(r);_=v(R,"BUTTON",{class:!0,"data-svelte-h":!0}),H(_)!=="svelte-azbayw"&&(_.textContent=d),C=N(R),m=v(R,"BUTTON",{class:!0,style:!0,"data-svelte-h":!0}),H(m)!=="svelte-1ag2hpy"&&(m.textContent=O),R.forEach(b),T.forEach(b),this.h()},h(){k(i,"class","bg-gray-50 border-b"),k(n,"class","min-w-full shadow rounded-lg"),k(_,"class","flex items-center p-1 text-gray-500 bg-gray-300 rounded-md hover:bg-teal-400 hover:text-white font-bold"),k(m,"class","p-1 text-gray-500 bg-gray-300 rounded-md hover:bg-teal-400 hover:text-white font-bold"),Lt(m,"transition","all 0.2s ease 0s"),k(r,"class","flex justify-between p-1"),k(l,"class","overflow-auto p-4")},m(c,T){W(t,c,T),I(c,e,T),I(c,l,T),p(l,n),p(n,i),p(n,u),p(n,s);for(let S=0;S<h.length;S+=1)h[S]&&h[S].m(s,null);p(l,a),p(l,r),p(r,_),p(r,C),p(r,m),E=!0,L||(y=[q(_,"click",o[14]),q(m,"click",o[15])],L=!0)},p(c,T){const S={};if(T&1&&(S.actions={Import:c[10],Export:c[11]}),t.$set(S),T&227){g=xt(c[1]);let D;for(D=0;D<g.length;D+=1){const R=yt(c,g,D);h[D]?h[D].p(R,T):(h[D]=kt(R),h[D].c(),h[D].m(s,null))}for(;D<h.length;D+=1)h[D].d(1);h.length=g.length}},i(c){E||(A(t.$$.fragment,c),E=!0)},o(c){V(t.$$.fragment,c),E=!1},d(c){c&&(b(e),b(l)),X(t,c),Dt(h,c),L=!1,$(y)}}}function zt(o){let t,e;return t=new it({}),{c(){Y(t.$$.fragment)},l(l){G(t.$$.fragment,l)},m(l,n){W(t,l,n),e=!0},p:j,i(l){e||(A(t.$$.fragment,l),e=!0)},o(l){V(t.$$.fragment,l),e=!1},d(l){X(t,l)}}}function kt(o){let t,e,l=(o[18].key||"")+"",n,i,f,u=(o[18].version||"")+"",s,a,r,_=(o[18].tag1||"")+"",d,C,m,O=(o[18].tag2||"")+"",E,L,y,g=(o[18].tag3||"")+"",h,c,T,S=(o[18].ttl||"")+"",D,R,P,J,ct="Edit",tt,M,ft="Delete",et,lt,ut;function Ct(){return o[12](o[18])}function Tt(){return o[13](o[18])}return{c(){t=x("tr"),e=x("td"),n=F(l),i=U(),f=x("td"),s=F(u),a=U(),r=x("td"),d=F(_),C=U(),m=x("td"),E=F(O),L=U(),y=x("td"),h=F(g),c=U(),T=x("td"),D=F(S),R=U(),P=x("td"),J=x("button"),J.textContent=ct,tt=U(),M=x("button"),M.textContent=ft,et=U(),this.h()},l(K){t=v(K,"TR",{class:!0});var w=B(t);e=v(w,"TD",{class:!0});var pt=B(e);n=z(pt,l),pt.forEach(b),i=N(w),f=v(w,"TD",{class:!0});var dt=B(f);s=z(dt,u),dt.forEach(b),a=N(w),r=v(w,"TD",{class:!0});var _t=B(r);d=z(_t,_),_t.forEach(b),C=N(w),m=v(w,"TD",{class:!0});var ht=B(m);E=z(ht,O),ht.forEach(b),L=N(w),y=v(w,"TD",{class:!0});var mt=B(y);h=z(mt,g),mt.forEach(b),c=N(w),T=v(w,"TD",{class:!0});var gt=B(T);D=z(gt,S),gt.forEach(b),R=N(w),P=v(w,"TD",{class:!0});var Z=B(P);J=v(Z,"BUTTON",{class:!0,"data-svelte-h":!0}),H(J)!=="svelte-u5a084"&&(J.textContent=ct),tt=N(Z),M=v(Z,"BUTTON",{class:!0,"data-svelte-h":!0}),H(M)!=="svelte-14nc668"&&(M.textContent=ft),Z.forEach(b),et=N(w),w.forEach(b),this.h()},h(){k(e,"class","px-2 py-2 whitespace-nowrap text-sm font-medium text-gray-900"),k(f,"class","text-sm text-gray-900 font-light px-2 py-2 whitespace-nowrap"),k(r,"class","text-sm text-gray-900 font-light px-2 py-2 whitespace-nowrap"),k(m,"class","text-sm text-gray-900 font-light px-2 py-2 whitespace-nowrap"),k(y,"class","text-sm text-gray-900 font-light px-2 py-2 whitespace-nowrap"),k(T,"class","text-sm text-gray-900 font-light px-2 py-2 whitespace-nowrap"),k(J,"class","btn variant-filled-primary"),k(M,"class","btn btn-sm variant-filled-secondary"),k(P,"class","text-sm text-gray-900 font-light px-2 py-2 whitespace-nowrap flex gap-2"),k(t,"class","bg-white border-b transition duration-300 ease-in-out hover:bg-gray-100")},m(K,w){I(K,t,w),p(t,e),p(e,n),p(t,i),p(t,f),p(f,s),p(t,a),p(t,r),p(r,d),p(t,C),p(t,m),p(m,E),p(t,L),p(t,y),p(y,h),p(t,c),p(t,T),p(T,D),p(t,R),p(t,P),p(P,J),p(P,tt),p(P,M),p(t,et),lt||(ut=[q(J,"click",Ct),q(M,"click",Tt)],lt=!0)},p(K,w){o=K,w&2&&l!==(l=(o[18].key||"")+"")&&Q(n,l),w&2&&u!==(u=(o[18].version||"")+"")&&Q(s,u),w&2&&_!==(_=(o[18].tag1||"")+"")&&Q(d,_),w&2&&O!==(O=(o[18].tag2||"")+"")&&Q(E,O),w&2&&g!==(g=(o[18].tag3||"")+"")&&Q(h,g),w&2&&S!==(S=(o[18].ttl||"")+"")&&Q(D,S)},d(K){K&&b(t),lt=!1,$(ut)}}}function Kt(o){let t,e,l,n,i;const f=[zt,Ft],u=[];function s(a,r){return a[2]?0:1}return t=s(o),e=u[t]=f[t](o),n=new Nt({props:{onClick:o[16]}}),{c(){e.c(),l=U(),Y(n.$$.fragment)},l(a){e.l(a),l=N(a),G(n.$$.fragment,a)},m(a,r){u[t].m(a,r),I(a,l,r),W(n,a,r),i=!0},p(a,[r]){let _=t;t=s(a),t===_?u[t].p(a,r):(ot(),V(u[_],1,1,()=>{u[_]=null}),rt(),e=u[t],e?e.p(a,r):(e=u[t]=f[t](a),e.c()),A(e,1),e.m(l.parentNode,l));const d={};r&1&&(d.onClick=a[16]),n.$set(d)},i(a){i||(A(e),A(n.$$.fragment,a),i=!0)},o(a){V(e),V(n.$$.fragment,a),i=!1},d(a){a&&b(l),u[t].d(a),X(n,a)}}}let Qt=0;function Yt(o,t,e){let l;wt(o,Ut,c=>e(17,l=c));let{pid:n=l.pid}=t,i=[],f=!0,u="",s=[];const a=Et("__app__"),r=a.api_manager.get_admin_plug_api(),_=async()=>{e(2,f=!0);const c=await r.list_plug_state(n,{key_cursor:u,page:Qt,no_value:!0});if(!c.ok){console.log("Err",c);return}e(1,i=c.data),e(2,f=!1)},d=c=>{a.utils.small_modal_open(Pt,{app:a,id:c})},C=c=>{a.utils.small_modal_open(Mt,{app:a,id:c})};_();const m=()=>d(n),O=()=>C(n),E=c=>a.nav.admin_plug_state_edit(n,c.key),L=async c=>{await r.delete_plug_state(n,c.key),_()},y=()=>{s.length>0?e(3,u=s.pop()):e(3,u=""),_()},g=()=>{i.length>0?(s.push(u),e(3,u=i[i.length-1].key||"")):(e(3,u=""),e(4,s=[])),_()},h=()=>a.nav.admin_plug_state_new(n);return o.$$set=c=>{"pid"in c&&e(0,n=c.pid)},[n,i,f,u,s,a,r,_,d,C,m,O,E,L,y,g,h]}class se extends st{constructor(t){super(),at(this,t,Yt,Kt,nt,{pid:0})}}export{se as component};
