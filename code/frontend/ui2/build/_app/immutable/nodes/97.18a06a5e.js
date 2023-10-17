import{s as J,f as g,g as $,h as y,d as u,j as d,i as b,P as Q,e as A,A as C,Q as Y,a as I,l as V,C as P,c as x,m as H,w as v,D as F,n as M,k as Z,W as N,J as ee,p as te,U as le}from"../chunks/scheduler.e2ee220a.js";import{S as R,i as W,g as X,t as k,c as z,a as w,b as S,d as D,m as L,e as j,j as ne}from"../chunks/index.4aee2103.js";import{S as se,a as B}from"../chunks/stepper.ce9fe1e4.js";import{e as U}from"../chunks/each.e59479a4.js";import"../chunks/paths.2eaeb908.js";import"../chunks/loading_spinner.svelte_svelte_type_style_lang.c28228c7.js";import{L as K}from"../chunks/loading_spinner.4ef87ddf.js";function q(c,e,l){const t=c.slice();return t[7]=e[l][0],t[8]=e[l][1],t}function ae(c){let e,l=U(Object.entries(c[2])),t=[];for(let a=0;a<l.length;a+=1)t[a]=G(q(c,l,a));return{c(){for(let a=0;a<t.length;a+=1)t[a].c();e=A()},l(a){for(let o=0;o<t.length;o+=1)t[o].l(a);e=A()},m(a,o){for(let n=0;n<t.length;n+=1)t[n]&&t[n].m(a,o);b(a,e,o)},p(a,o){if(o&5){l=U(Object.entries(a[2]));let n;for(n=0;n<l.length;n+=1){const s=q(a,l,n);t[n]?t[n].p(s,o):(t[n]=G(s),t[n].c(),t[n].m(e.parentNode,e))}for(;n<t.length;n+=1)t[n].d(1);t.length=l.length}},i:C,o:C,d(a){a&&u(e),Y(t,a)}}}function re(c){let e,l;return e=new K({props:{classes:""}}),{c(){S(e.$$.fragment)},l(t){D(e.$$.fragment,t)},m(t,a){L(e,t,a),l=!0},p:C,i(t){l||(w(e.$$.fragment,t),l=!0)},o(t){k(e.$$.fragment,t),l=!1},d(t){j(e,t)}}}function G(c){let e,l,t='<img alt="" class="block w-full h-auto object-cover object-center cursor-pointer" src="https://dummyimage.com/400x200"/>',a,o,n,s=c[8].name+"",r,i,f,m,E;function T(){return c[3](c[7])}return{c(){e=g("div"),l=g("div"),l.innerHTML=t,a=I(),o=g("div"),n=g("h2"),r=V(s),i=I(),this.h()},l(_){e=$(_,"DIV",{class:!0});var p=y(e);l=$(p,"DIV",{class:!0,"data-svelte-h":!0}),P(l)!=="svelte-h24tje"&&(l.innerHTML=t),a=x(p),o=$(p,"DIV",{class:!0});var h=y(o);n=$(h,"H2",{class:!0});var O=y(n);r=H(O,s),O.forEach(u),h.forEach(u),i=x(p),p.forEach(u),this.h()},h(){d(l,"class","block overflow-hidden rounded"),d(n,"class","title-font text-lg font-medium text-gray-900"),d(o,"class","mt-4"),d(e,"class",f="p-4 w-52 m-1 shadow rounded hover:bg-blue-100 "+(c[7]===c[0]?"border border-blue-400":""))},m(_,p){b(_,e,p),v(e,l),v(e,a),v(e,o),v(o,n),v(n,r),v(e,i),m||(E=F(e,"click",T),m=!0)},p(_,p){c=_,p&4&&s!==(s=c[8].name+"")&&M(r,s),p&5&&f!==(f="p-4 w-52 m-1 shadow rounded hover:bg-blue-100 "+(c[7]===c[0]?"border border-blue-400":""))&&d(e,"class",f)},d(_){_&&u(e),m=!1,E()}}}function oe(c){let e,l,t,a;const o=[re,ae],n=[];function s(r,i){return r[1]?0:1}return l=s(c),t=n[l]=o[l](c),{c(){e=g("div"),t.c(),this.h()},l(r){e=$(r,"DIV",{class:!0});var i=y(e);t.l(i),i.forEach(u),this.h()},h(){d(e,"class","flex flex-wrap p-2 gap-2 overflow-auto justify-center")},m(r,i){b(r,e,i),n[l].m(e,null),a=!0},p(r,[i]){let f=l;l=s(r),l===f?n[l].p(r,i):(X(),k(n[f],1,1,()=>{n[f]=null}),z(),t=n[l],t?t.p(r,i):(t=n[l]=o[l](r),t.c()),w(t,1),t.m(e,null))},i(r){a||(w(t),a=!0)},o(r){k(t),a=!1},d(r){r&&u(e),n[l].d()}}}function ce(c,e,l){let{template:t=""}=e;const a=Q("__app__");let o=!0;const n=a.api_manager.get_self_api();let s=[];(async()=>{const f=await n.list_sheet_templates();f.ok&&(l(2,s=f.data),l(1,o=!1),console.log("@templates",s))})();const i=f=>{l(0,t=f)};return c.$$set=f=>{"template"in f&&l(0,t=f.template)},[t,o,s,i]}class ie extends R{constructor(e){super(),W(this,e,ce,oe,J,{template:0})}}function fe(c){let e;return{c(){e=V("New Sheet")},l(l){e=H(l,"New Sheet")},m(l,t){b(l,e,t)},d(l){l&&u(e)}}}function ue(c){let e,l,t="Name",a,o,n,s,r,i="Info",f,m,E,T;return{c(){e=g("div"),l=g("label"),l.textContent=t,a=I(),o=g("input"),n=I(),s=g("div"),r=g("label"),r.textContent=i,f=I(),m=g("textarea"),this.h()},l(_){e=$(_,"DIV",{class:!0});var p=y(e);l=$(p,"LABEL",{for:!0,class:!0,"data-svelte-h":!0}),P(l)!=="svelte-1ckt5v7"&&(l.textContent=t),a=x(p),o=$(p,"INPUT",{type:!0,class:!0}),p.forEach(u),n=x(_),s=$(_,"DIV",{class:!0});var h=y(s);r=$(h,"LABEL",{for:!0,class:!0,"data-svelte-h":!0}),P(r)!=="svelte-1npi386"&&(r.textContent=i),f=x(h),m=$(h,"TEXTAREA",{class:!0}),y(m).forEach(u),h.forEach(u),this.h()},h(){d(l,"for",""),d(l,"class","pb-2 text-gray-700 font-semibold"),d(o,"type","text"),d(o,"class","p-2 rounded-lg bg-gray-100 outline-none focus:bg-gray-200"),d(e,"class","flex-col flex py-3"),d(r,"for",""),d(r,"class","pb-2 text-gray-700 font-semibold"),d(m,"class","p-2 rounded-lg bg-gray-100 outline-none focus:bg-gray-200"),d(s,"class","flex-col flex py-3")},m(_,p){b(_,e,p),v(e,l),v(e,a),v(e,o),N(o,c[1]),b(_,n,p),b(_,s,p),v(s,r),v(s,f),v(s,m),N(m,c[2]),E||(T=[F(o,"input",c[7]),F(m,"input",c[8])],E=!0)},p(_,p){p&2&&o.value!==_[1]&&N(o,_[1]),p&4&&N(m,_[2])},d(_){_&&(u(e),u(n),u(s)),E=!1,ee(T)}}}function pe(c){let e;return{c(){e=V("Select Template")},l(l){e=H(l,"Select Template")},m(l,t){b(l,e,t)},d(l){l&&u(e)}}}function _e(c){let e,l,t;function a(n){c[9](n)}let o={};return c[3]!==void 0&&(o.template=c[3]),e=new ie({props:o}),te.push(()=>ne(e,"template",a)),{c(){S(e.$$.fragment)},l(n){D(e.$$.fragment,n)},m(n,s){L(e,n,s),t=!0},p(n,s){const r={};!l&&s&8&&(l=!0,r.template=n[3],le(()=>l=!1)),e.$set(r)},i(n){t||(w(e.$$.fragment,n),t=!0)},o(n){k(e.$$.fragment,n),t=!1},d(n){j(e,n)}}}function me(c){let e,l="Sheet is ready. Go explore.";return{c(){e=g("p"),e.textContent=l},l(t){e=$(t,"P",{"data-svelte-h":!0}),P(e)!=="svelte-i5foot"&&(e.textContent=l)},m(t,a){b(t,e,a)},p:C,i:C,o:C,d(t){t&&u(e)}}}function de(c){let e,l;return{c(){e=g("p"),l=V(c[5]),this.h()},l(t){e=$(t,"P",{class:!0});var a=y(e);l=H(a,c[5]),a.forEach(u),this.h()},h(){d(e,"class","text-red-500")},m(t,a){b(t,e,a),v(e,l)},p(t,a){a&32&&M(l,t[5])},i:C,o:C,d(t){t&&u(e)}}}function he(c){let e,l;return e=new K({props:{classes:""}}),{c(){S(e.$$.fragment)},l(t){D(e.$$.fragment,t)},m(t,a){L(e,t,a),l=!0},p:C,i(t){l||(w(e.$$.fragment,t),l=!0)},o(t){k(e.$$.fragment,t),l=!1},d(t){j(e,t)}}}function ge(c){let e,l,t,a;const o=[he,de,me],n=[];function s(r,i){return r[4]?0:r[5]?1:2}return e=s(c),l=n[e]=o[e](c),{c(){l.c(),t=A()},l(r){l.l(r),t=A()},m(r,i){n[e].m(r,i),b(r,t,i),a=!0},p(r,i){let f=e;e=s(r),e===f?n[e].p(r,i):(X(),k(n[f],1,1,()=>{n[f]=null}),z(),l=n[e],l?l.p(r,i):(l=n[e]=o[e](r),l.c()),w(l,1),l.m(t.parentNode,t))},i(r){a||(w(l),a=!0)},o(r){k(l),a=!1},d(r){r&&u(t),n[e].d(r)}}}function $e(c){let e=c[0]?"Finished":"Instancing",l;return{c(){l=V(e)},l(t){l=H(t,e)},m(t,a){b(t,l,a)},p(t,a){a&1&&e!==(e=t[0]?"Finished":"Instancing")&&M(l,e)},d(t){t&&u(l)}}}function be(c){let e,l,t,a,o,n;return e=new B({props:{locked:!c[1]||!c[2],$$slots:{default:[ue],header:[fe]},$$scope:{ctx:c}}}),t=new B({props:{back_locked:c[0],locked:!c[3],$$slots:{default:[_e],header:[pe]},$$scope:{ctx:c}}}),o=new B({props:{back_locked:c[0],$$slots:{header:[$e],default:[ge]},$$scope:{ctx:c}}}),{c(){S(e.$$.fragment),l=I(),S(t.$$.fragment),a=I(),S(o.$$.fragment)},l(s){D(e.$$.fragment,s),l=x(s),D(t.$$.fragment,s),a=x(s),D(o.$$.fragment,s)},m(s,r){L(e,s,r),b(s,l,r),L(t,s,r),b(s,a,r),L(o,s,r),n=!0},p(s,r){const i={};r&6&&(i.locked=!s[1]||!s[2]),r&8198&&(i.$$scope={dirty:r,ctx:s}),e.$set(i);const f={};r&1&&(f.back_locked=s[0]),r&8&&(f.locked=!s[3]),r&8200&&(f.$$scope={dirty:r,ctx:s}),t.$set(f);const m={};r&1&&(m.back_locked=s[0]),r&8241&&(m.$$scope={dirty:r,ctx:s}),o.$set(m)},i(s){n||(w(e.$$.fragment,s),w(t.$$.fragment,s),w(o.$$.fragment,s),n=!0)},o(s){k(e.$$.fragment,s),k(t.$$.fragment,s),k(o.$$.fragment,s),n=!1},d(s){s&&(u(l),u(a)),j(e,s),j(t,s),j(o,s)}}}function ve(c){let e,l,t,a;return t=new se({props:{buttonCompleteLabel:"",$$slots:{default:[be]},$$scope:{ctx:c}}}),t.$on("next",c[6]),t.$on("back",ke),t.$on("step",we),t.$on("complete",ye),{c(){e=g("div"),l=g("div"),S(t.$$.fragment),this.h()},l(o){e=$(o,"DIV",{class:!0});var n=y(e);l=$(n,"DIV",{class:!0,style:!0});var s=y(l);D(t.$$.fragment,s),s.forEach(u),n.forEach(u),this.h()},h(){d(l,"class","card p-4 text-token border shadow mx-auto my-4 bg-white"),Z(l,"max-width","750px"),d(e,"class","w-full bg-gray-50 h-full py-4 px-1")},m(o,n){b(o,e,n),v(e,l),L(t,l,null),a=!0},p(o,[n]){const s={};n&8255&&(s.$$scope={dirty:n,ctx:o}),t.$set(s)},i(o){a||(w(t.$$.fragment,o),a=!0)},o(o){k(t.$$.fragment,o),a=!1},d(o){o&&u(e),j(t)}}}function ke(c){console.log("event:prev",c.detail)}function we(c){console.log("event:step",c.detail)}function ye(c){console.log("event:complete",c.detail)}function Ce(c,e,l){const t=Q("__app__");let a=!1,o="",n="",s="";const r=t.api_manager.get_self_api();let i=!0,f="";const m=async()=>{const h=await r.instance_sheet_template(o,n,s);if(!h.ok){l(5,f=h.data),l(4,i=!1);return}l(4,i=!1),l(0,a=!0)};function E(h){console.log("event:next",h.detail.step),h.detail.step===1&&m()}function T(){o=this.value,l(1,o)}function _(){n=this.value,l(2,n)}function p(h){s=h,l(3,s)}return[a,o,n,s,i,f,E,T,_,p]}class Te extends R{constructor(e){super(),W(this,e,Ce,ve,J,{})}}export{Te as component};
