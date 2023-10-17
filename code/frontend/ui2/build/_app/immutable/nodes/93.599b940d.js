import{s as ye,A as Ve,a as D,c as T,i as ee,d as p,P as Le,f as g,g as v,h as w,C as de,j as $,w as i,$ as be,D as se,Q as $e,l as M,m as U,W as ke,n as Z,M as we,J as Oe}from"../chunks/scheduler.e2ee220a.js";import{S as Ne,i as Se,b as R,d as W,m as z,a as P,t as B,e as K,c as He,g as Pe}from"../chunks/index.4aee2103.js";import{e as ne}from"../chunks/each.e59479a4.js";import{I as fe}from"../chunks/Icon.9e22c3e4.js";import{s as Ee}from"../chunks/index.3b48e8d3.js";import"../chunks/paths.2eaeb908.js";import{F as qe}from"../chunks/floating_add.1495f8dc.js";import"../chunks/loading_spinner.svelte_svelte_type_style_lang.c28228c7.js";import{L as Fe}from"../chunks/loading_spinner.4ef87ddf.js";import{A as Me}from"../chunks/action_picker.7081fe98.js";function Ue(l){let e,a;return e=new Me({props:{actions:l[0],title:"Create new data tables."}}),{c(){R(e.$$.fragment)},l(t){W(e.$$.fragment,t)},m(t,r){z(e,t,r),a=!0},p:Ve,i(t){a||(P(e.$$.fragment,t),a=!0)},o(t){B(e.$$.fragment,t),a=!1},d(t){K(e,t)}}}function Ge(l,e,a){let{app:t}=e;const r=[{name:"Quick Sheets",icon:"table",info:"New Data sheets for storing data.",action:()=>{t.utils.small_modal_close(),t.nav.data_sheets_new()}},{name:"From Store",icon:"shopping-bag",info:"You can import required data tables from store, if you have proper scope.",action:()=>{t.utils.small_modal_close(),t.nav.repo_loader()}},{name:"From Builder",icon:"document-add",info:"You can design schema for data table but it needs knowledge about database system.",action:()=>{t.utils.small_modal_close()}}];return l.$$set=u=>{"app"in u&&a(1,t=u.app)},[r,t]}class Qe extends Ne{constructor(e){super(),Se(this,e,Ge,Ue,ye,{app:1})}}function xe(l,e,a){const t=l.slice();return t[12]=e[a],t}function Ie(l,e,a){const t=l.slice();return t[15]=e[a],t}function Ye(l){let e,a,t,r,u='<h2 class="text-gray-700 text-3xl font-medium">Data Table Groups</h2> <div class="h-1 w-64 bg-indigo-500 rounded"></div>',c,m,b,f,h,k,N,L,y=ne(l[2]),o=[];for(let s=0;s<y.length;s+=1)o[s]=Ce(Ie(l,y,s));let E=ne(l[3]),d=[];for(let s=0;s<E.length;s+=1)d[s]=Te(xe(l,E,s));const X=s=>B(d[s],1,1,()=>{d[s]=null});return{c(){e=g("div"),a=g("div"),t=g("div"),r=g("div"),r.innerHTML=u,c=D(),m=g("div"),b=g("select");for(let s=0;s<o.length;s+=1)o[s].c();f=D(),h=g("div");for(let s=0;s<d.length;s+=1)d[s].c();this.h()},l(s){e=v(s,"DIV",{class:!0});var _=w(e);a=v(_,"DIV",{class:!0});var n=w(a);t=v(n,"DIV",{class:!0});var C=w(t);r=v(C,"DIV",{class:!0,"data-svelte-h":!0}),de(r)!=="svelte-ij96ck"&&(r.innerHTML=u),c=T(C),m=v(C,"DIV",{});var G=w(m);b=v(G,"SELECT",{class:!0});var V=w(b);for(let S=0;S<o.length;S+=1)o[S].l(V);V.forEach(p),G.forEach(p),C.forEach(p),f=T(n),h=v(n,"DIV",{class:!0});var j=w(h);for(let S=0;S<d.length;S+=1)d[S].l(j);j.forEach(p),n.forEach(p),_.forEach(p),this.h()},h(){$(r,"class","w-full mb-6 lg:mb-0"),$(b,"class","rounded p-2 bg-white"),$(t,"class","flex justify-around w-full mb-4 p-4"),$(h,"class","flex flex-wrap -m-4"),$(a,"class","container px-2 py-8 mx-auto max-w-7x1"),$(e,"class","w-full h-full overflow-auto")},m(s,_){ee(s,e,_),i(e,a),i(a,t),i(t,r),i(t,c),i(t,m),i(m,b);for(let n=0;n<o.length;n+=1)o[n]&&o[n].m(b,null);be(b,l[0]),i(a,f),i(a,h);for(let n=0;n<d.length;n+=1)d[n]&&d[n].m(h,null);k=!0,N||(L=se(b,"change",We),N=!0)},p(s,_){if(_&4){y=ne(s[2]);let n;for(n=0;n<y.length;n+=1){const C=Ie(s,y,n);o[n]?o[n].p(C,_):(o[n]=Ce(C),o[n].c(),o[n].m(b,null))}for(;n<o.length;n+=1)o[n].d(1);o.length=y.length}if((!k||_&5)&&be(b,s[0]),_&57){E=ne(s[3]);let n;for(n=0;n<E.length;n+=1){const C=xe(s,E,n);d[n]?(d[n].p(C,_),P(d[n],1)):(d[n]=Te(C),d[n].c(),P(d[n],1),d[n].m(h,null))}for(Pe(),n=E.length;n<d.length;n+=1)X(n);He()}},i(s){if(!k){for(let _=0;_<E.length;_+=1)P(d[_]);k=!0}},o(s){d=d.filter(Boolean);for(let _=0;_<d.length;_+=1)B(d[_]);k=!1},d(s){s&&p(e),$e(o,s),$e(d,s),N=!1,L()}}}function Je(l){let e,a;return e=new Fe({}),{c(){R(e.$$.fragment)},l(t){W(e.$$.fragment,t)},m(t,r){z(e,t,r),a=!0},p:Ve,i(t){a||(P(e.$$.fragment,t),a=!0)},o(t){B(e.$$.fragment,t),a=!1},d(t){K(e,t)}}}function Ce(l){let e,a=l[15]+"",t,r;return{c(){e=g("option"),t=M(a),this.h()},l(u){e=v(u,"OPTION",{});var c=w(e);t=U(c,a),c.forEach(p),this.h()},h(){e.__value=r=l[15],ke(e,e.__value)},m(u,c){ee(u,e,c),i(e,t)},p(u,c){c&4&&a!==(a=u[15]+"")&&Z(t,a),c&4&&r!==(r=u[15])&&(e.__value=r,ke(e,e.__value))},d(u){u&&p(e)}}}function De(l){let e,a,t=(l[12].renderer||"")+"",r;return{c(){e=g("h3"),a=M("#"),r=M(t),this.h()},l(u){e=v(u,"H3",{class:!0});var c=w(e);a=U(c,"#"),r=U(c,t),c.forEach(p),this.h()},h(){$(e,"class","tracking-widest text-blue-500 text-xs uppercase font-medium title-font bg-yellow-400 rounded")},m(u,c){ee(u,e,c),i(e,a),i(e,r)},p(u,c){c&8&&t!==(t=(u[12].renderer||"")+"")&&Z(r,t)},d(u){u&&p(e)}}}function Te(l){let e,a,t,r,u,c,m,b,f,h,k,N=l[12].name+"",L,y,o,E=l[12].description+"",d,X,s,_,n,C,G,V,j,S,Q,_e="Raw",re,O,F,oe,Y,me="Setting",ie,q,ce,pe,x=l[12].renderer&&De(l);n=new fe({props:{name:"view-list",class:"h-5 w-5"}});function je(){return l[6](l[12])}j=new fe({props:{name:"hashtag",class:"h-5 w-5"}});function Ae(){return l[7](l[12])}F=new fe({props:{name:"cog",class:"h-5 w-5"}});function Be(){return l[8](l[12])}return{c(){e=g("div"),a=g("div"),t=g("img"),u=D(),c=g("div"),m=g("h3"),b=M(l[0]),f=D(),x&&x.c(),h=D(),k=g("h2"),L=M(N),y=D(),o=g("p"),d=M(E),X=D(),s=g("div"),_=g("button"),R(n.$$.fragment),C=M(`
                  Explore`),G=D(),V=g("button"),R(j.$$.fragment),S=D(),Q=g("span"),Q.textContent=_e,re=D(),O=g("button"),R(F.$$.fragment),oe=D(),Y=g("span"),Y.textContent=me,ie=D(),this.h()},l(I){e=v(I,"DIV",{class:!0});var H=w(e);a=v(H,"DIV",{class:!0});var A=w(a);t=v(A,"IMG",{class:!0,src:!0,alt:!0}),u=T(A),c=v(A,"DIV",{class:!0});var te=w(c);m=v(te,"H3",{class:!0});var he=w(m);b=U(he,l[0]),he.forEach(p),f=T(te),x&&x.l(te),te.forEach(p),h=T(A),k=v(A,"H2",{class:!0});var ge=w(k);L=U(ge,N),ge.forEach(p),y=T(A),o=v(A,"P",{class:!0});var ve=w(o);d=U(ve,E),ve.forEach(p),X=T(A),s=v(A,"DIV",{class:!0});var J=w(s);_=v(J,"BUTTON",{class:!0});var ue=w(_);W(n.$$.fragment,ue),C=U(ue,`
                  Explore`),ue.forEach(p),G=T(J),V=v(J,"BUTTON",{class:!0});var ae=w(V);W(j.$$.fragment,ae),S=T(ae),Q=v(ae,"SPAN",{"data-svelte-h":!0}),de(Q)!=="svelte-1fwv1ds"&&(Q.textContent=_e),ae.forEach(p),re=T(J),O=v(J,"BUTTON",{class:!0});var le=w(O);W(F.$$.fragment,le),oe=T(le),Y=v(le,"SPAN",{"data-svelte-h":!0}),de(Y)!=="svelte-14iox2q"&&(Y.textContent=me),le.forEach(p),J.forEach(p),A.forEach(p),ie=T(H),H.forEach(p),this.h()},h(){$(t,"class","lg:h-60 xl:h-56 md:h-64 sm:h-72 xs:h-72 h-72 rounded w-full object-cover object-center mb-6"),we(t.src,r="https://picsum.photos/seed/"+Ee(l[12].name+l[12].slug)+"d/800/400")||$(t,"src",r),$(t,"alt",""),$(m,"class","tracking-widest text-indigo-500 text-xs uppercase font-medium title-font h3"),$(c,"class","flex gap-1"),$(k,"class","text-lg text-gray-900 font-medium title-font mb-4 h2"),$(o,"class","leading-relaxed selection:bg-red-200 text-base"),$(_,"class","btn btn-md variant-filled-primary"),$(V,"class","btn btn-sm variant-filled-secondary"),$(O,"class","btn btn-sm variant-filled-secondary"),$(s,"class","flex p-5 gap-2"),$(a,"class","card p-6"),$(e,"class","xl:w-1/3 md:w-1/2 p-4")},m(I,H){ee(I,e,H),i(e,a),i(a,t),i(a,u),i(a,c),i(c,m),i(m,b),i(c,f),x&&x.m(c,null),i(a,h),i(a,k),i(k,L),i(a,y),i(a,o),i(o,d),i(a,X),i(a,s),i(s,_),z(n,_,null),i(_,C),i(s,G),i(s,V),z(j,V,null),i(V,S),i(V,Q),i(s,re),i(s,O),z(F,O,null),i(O,oe),i(O,Y),i(e,ie),q=!0,ce||(pe=[se(_,"click",je),se(V,"click",Ae),se(O,"click",Be)],ce=!0)},p(I,H){l=I,(!q||H&8&&!we(t.src,r="https://picsum.photos/seed/"+Ee(l[12].name+l[12].slug)+"d/800/400"))&&$(t,"src",r),(!q||H&1)&&Z(b,l[0]),l[12].renderer?x?x.p(l,H):(x=De(l),x.c(),x.m(c,null)):x&&(x.d(1),x=null),(!q||H&8)&&N!==(N=l[12].name+"")&&Z(L,N),(!q||H&8)&&E!==(E=l[12].description+"")&&Z(d,E)},i(I){q||(P(n.$$.fragment,I),P(j.$$.fragment,I),P(F.$$.fragment,I),q=!0)},o(I){B(n.$$.fragment,I),B(j.$$.fragment,I),B(F.$$.fragment,I),q=!1},d(I){I&&p(e),x&&x.d(),K(n),K(j),K(F),ce=!1,Oe(pe)}}}function Re(l){let e,a,t,r,u;const c=[Je,Ye],m=[];function b(f,h){return f[1]?0:1}return e=b(l),a=m[e]=c[e](l),r=new qe({props:{onClick:l[9]}}),{c(){a.c(),t=D(),R(r.$$.fragment)},l(f){a.l(f),t=T(f),W(r.$$.fragment,f)},m(f,h){m[e].m(f,h),ee(f,t,h),z(r,f,h),u=!0},p(f,[h]){let k=e;e=b(f),e===k?m[e].p(f,h):(Pe(),B(m[k],1,1,()=>{m[k]=null}),He(),a=m[e],a?a.p(f,h):(a=m[e]=c[e](f),a.c()),P(a,1),a.m(t.parentNode,t))},i(f){u||(P(a),P(r.$$.fragment,f),u=!0)},o(f){B(a),B(r.$$.fragment,f),u=!1},d(f){f&&p(t),m[e].d(f),K(r,f)}}}const We=l=>{};function ze(l,e,a){let{source:t="default"}=e;const r=Le("__app__"),u=r.api_manager.get_admin_data_api();let c=!1,m=[],b=[];r.api_manager.self_data.get_data_sources().then(o=>{a(2,m=o)}),(async()=>{const o=await u.list_group(t);o.ok&&(a(3,b=o.data),a(1,c=!1))})();const h=o=>{console.log("@group =>",o);const E=o.slug;switch(o.renderer){case"sheet":r.nav.data_sheets_page(t,E);break;default:r.nav.data_group_page(t,E);break}},k=o=>h(o),N=o=>r.nav.data_group_page(t,o.slug),L=o=>r.nav.admin_data_group(t,o.slug),y=()=>{r.utils.small_modal_open(Qe,{app:r})};return l.$$set=o=>{"source"in o&&a(0,t=o.source)},[t,c,m,b,r,h,k,N,L,y]}class ot extends Ne{constructor(e){super(),Se(this,e,ze,Re,ye,{source:0})}}export{ot as component};
