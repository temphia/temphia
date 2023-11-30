import{s as P,e as E,i as w,A as H,d as _,f as p,g as m,C as L,j as c,h as j,M as R,w as h,a as T,c as I,D as B,_ as G,J,B as K,P as Q}from"../chunks/scheduler.e2ee220a.js";import{S as q,i as z,b as W,d as X,m as Y,a as Z,t as $,e as ee}from"../chunks/index.4aee2103.js";import{p as te}from"../chunks/index.008d0d8b.js";function se(a){let e,s='<div class="bg-text-white rounded text-gray-400 p-1 border"><svg xmlns="http://www.w3.org/2000/svg" class="h-48 w-auto" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 21h10a2 2 0 002-2V9.414a1 1 0 00-.293-.707l-5.414-5.414A1 1 0 0012.586 3H7a2 2 0 00-2 2v14a2 2 0 002 2z"></path></svg></div>';return{c(){e=p("div"),e.innerHTML=s,this.h()},l(t){e=m(t,"DIV",{class:!0,"data-svelte-h":!0}),L(e)!=="svelte-ntqjv1"&&(e.innerHTML=s),this.h()},h(){c(e,"class","flex justify-center")},m(t,l){w(t,e,l)},p:H,d(t){t&&_(e)}}}function le(a){let e,s,t;return{c(){e=p("div"),s=p("img"),this.h()},l(l){e=m(l,"DIV",{class:!0});var n=j(e);s=m(n,"IMG",{class:!0,src:!0,alt:!0}),n.forEach(_),this.h()},h(){c(s,"class","h-48 w-auto p-2 rounded border"),R(s.src,t=URL.createObjectURL(a[0]))||c(s,"src",t),c(s,"alt",""),c(e,"class","flex justify-center")},m(l,n){w(l,e,n),h(e,s)},p(l,n){n&1&&!R(s.src,t=URL.createObjectURL(l[0]))&&c(s,"src",t)},d(l){l&&_(e)}}}function ne(a){let e,s=`<p class="pointer-none text-gray-500"><span class="text-sm">Drag and drop</span>
                files here <br/>
                or
                <span class="text-blue-600 hover:underline">select a file</span>
                from your computer</p>`;return{c(){e=p("div"),e.innerHTML=s,this.h()},l(t){e=m(t,"DIV",{class:!0,"data-svelte-h":!0}),L(e)!=="svelte-1amu24y"&&(e.innerHTML=s),this.h()},h(){c(e,"class","h-full w-full text-center flex flex-col items-center justify-center cursor-pointer")},m(t,l){w(t,e,l)},p:H,d(t){t&&_(e)}}}function O(a){let e,s;function t(r,i){return r[0]?(e==null&&(e=!!r[1]()),e?le:se):ne}let l=t(a),n=l(a);return{c(){n.c(),s=E()},l(r){n.l(r),s=E()},m(r,i){n.m(r,i),w(r,s,i)},p(r,i){l===(l=t(r))&&n?n.p(r,i):(n.d(1),n=l(r),n&&(n.c(),n.m(s.parentNode,s)))},d(r){r&&_(s),n.d(r)}}}function re(a){let e=a[0],s,t=O(a);return{c(){t.c(),s=E()},l(l){t.l(l),s=E()},m(l,n){t.m(l,n),w(l,s,n)},p(l,[n]){n&1&&P(e,e=l[0])?(t.d(1),t=O(l),t.c(),t.m(s.parentNode,s)):t.p(l,n)},i:H,o:H,d(l){l&&_(s),t.d(l)}}}function ae(a,e,s){let{file:t}=e,{filename:l}=e;const n=["jpg","png","jpeg"],r=()=>{const i=l.split(".");return i.length<=1?!1:!!n.includes(i.pop())};return a.$$set=i=>{"file"in i&&s(0,t=i.file),"filename"in i&&s(2,l=i.filename)},[t,r,l]}class ie extends q{constructor(e){super(),z(this,e,ae,re,P,{file:0,filename:2})}}function S(a){let e,s="Upload";return{c(){e=p("button"),e.textContent=s,this.h()},l(t){e=m(t,"BUTTON",{class:!0,"data-svelte-h":!0}),L(e)!=="svelte-1xteojl"&&(e.textContent=s),this.h()},h(){c(e,"class","btn variant-filled-primary")},m(t,l){w(t,e,l)},d(t){t&&_(e)}}}function ce(a){let e,s,t,l="Attach Document",n,r,i,b,x,f,C,g,u="<span>File type: any</span>",U,k,D,V,A;f=new ie({props:{file:a[1],filename:a[0]}});let o=a[1]&&S();return{c(){e=p("form"),s=p("div"),t=p("span"),t.textContent=l,n=T(),r=p("div"),i=p("label"),b=p("input"),x=T(),W(f.$$.fragment),C=T(),g=p("p"),g.innerHTML=u,U=T(),k=p("div"),o&&o.c(),this.h()},l(d){e=m(d,"FORM",{class:!0});var v=j(e);s=m(v,"DIV",{class:!0});var y=j(s);t=m(y,"SPAN",{class:!0,"data-svelte-h":!0}),L(t)!=="svelte-1jgu5f8"&&(t.textContent=l),n=I(y),r=m(y,"DIV",{class:!0});var F=j(r);i=m(F,"LABEL",{class:!0});var M=j(i);b=m(M,"INPUT",{type:!0,class:!0}),x=I(M),X(f.$$.fragment,M),M.forEach(_),F.forEach(_),y.forEach(_),C=I(v),g=m(v,"P",{class:!0,"data-svelte-h":!0}),L(g)!=="svelte-m441pr"&&(g.innerHTML=u),U=I(v),k=m(v,"DIV",{class:!0});var N=j(k);o&&o.l(N),N.forEach(_),v.forEach(_),this.h()},h(){c(t,"class","text-sm font-bold text-gray-500 tracking-wide"),c(b,"type","file"),c(b,"class","hidden"),c(i,"class","flex flex-col rounded-lg border-4 border-dashed w-full h-60 p-5 group text-center"),c(r,"class","flex items-center justify-center w-full"),c(s,"class","grid grid-cols-1 space-y-2"),c(g,"class","text-sm text-gray-300"),c(k,"class","flex justify-end"),c(e,"class","mt-4 space-y-3")},m(d,v){w(d,e,v),h(e,s),h(s,t),h(s,n),h(s,r),h(r,i),h(i,b),h(i,x),Y(f,i,null),h(e,C),h(e,g),h(e,U),h(e,k),o&&o.m(k,null),D=!0,V||(A=[B(b,"change",a[2]),B(e,"submit",G(a[3]))],V=!0)},p(d,[v]){const y={};v&2&&(y.file=d[1]),v&1&&(y.filename=d[0]),f.$set(y),d[1]?o||(o=S(),o.c(),o.m(k,null)):o&&(o.d(1),o=null)},i(d){D||(Z(f.$$.fragment,d),D=!0)},o(d){$(f.$$.fragment,d),D=!1},d(d){d&&_(e),ee(f),o&&o.d(),V=!1,J(A)}}}function oe(a,e,s){let t;K(a,te,u=>s(5,t=u));let{source:l=t.source||"default"}=e,n=t.folder;const b=Q("__app__").get_cabinet_service().get_source_api(l);let x="",f;const C=u=>{console.log(u),s(1,f=u.target.files[0]),s(0,x=f.name),console.log(f)},g=async()=>{const u=new FormData;u.append("file",f),b.uploadFile(n,x,u)};return a.$$set=u=>{"source"in u&&s(4,l=u.source)},[x,f,C,g,l]}class pe extends q{constructor(e){super(),z(this,e,oe,ce,P,{source:4})}}export{pe as component};