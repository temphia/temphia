import{s as F,e as d,i as C,d as L,B as g,P,A as S}from"../chunks/scheduler.e2ee220a.js";import{S as q,i as A,t as f,c as B,a as u,g as N,b,d as k,m as $,e as h}from"../chunks/index.4aee2103.js";import{F as V}from"../chunks/FolderView.e91969cd.js";import"../chunks/loading_spinner.svelte_svelte_type_style_lang.9f9ec163.js";import{L as j}from"../chunks/loading_spinner.4ceee987.js";import{p as v}from"../chunks/index.37ddda1f.js";import{c as z,s as D}from"../chunks/select.297e1f2d.js";function G(s){let e,r;return e=new V({props:{selected:s[4].item,files:s[2]}}),e.$on("open_item",s[7]),e.$on("select_item",s[8]),{c(){b(e.$$.fragment)},l(t){k(e.$$.fragment,t)},m(t,a){$(e,t,a),r=!0},p(t,a){const i={};a&16&&(i.selected=t[4].item),a&4&&(i.files=t[2]),e.$set(i)},i(t){r||(u(e.$$.fragment,t),r=!0)},o(t){f(e.$$.fragment,t),r=!1},d(t){h(e,t)}}}function H(s){let e,r;return e=new j({}),{c(){b(e.$$.fragment)},l(t){k(e.$$.fragment,t)},m(t,a){$(e,t,a),r=!0},p:S,i(t){r||(u(e.$$.fragment,t),r=!0)},o(t){f(e.$$.fragment,t),r=!1},d(t){h(e,t)}}}function I(s){let e,r,t,a;const i=[H,G],c=[];function m(o,l){return o[3]?0:1}return e=m(s),r=c[e]=i[e](s),{c(){r.c(),t=d()},l(o){r.l(o),t=d()},m(o,l){c[e].m(o,l),C(o,t,l),a=!0},p(o,[l]){let _=e;e=m(o),e===_?c[e].p(o,l):(N(),f(c[_],1,1,()=>{c[_]=null}),B(),r=c[e],r?r.p(o,l):(r=c[e]=i[e](o),r.c()),u(r,1),r.m(t.parentNode,t))},i(o){a||(u(r),a=!0)},o(o){f(r),a=!1},d(o){o&&L(t),c[e].d(o)}}}function J(s,e,r){let t,a;g(s,v,n=>r(6,t=n)),g(s,z,n=>r(4,a=n));let{source:i=t.source||"default"}=e,c;const m=P("__app__"),o=m.get_cabinet_service();let l=[],_=!0;const w=async n=>{r(1,c=n);const p=await o.get_source_api(i).listFolder(n);p.ok&&(r(2,l=p.data),r(3,_=!1))},y=n=>{console.log("@EEEE",n.detail),n.detail.is_dir?m.nav.cab_folder(i,`${c}/${n.detail.name}`):m.nav.cab_file(i,c,n.detail.name)},E=n=>{D(c,n.detail.name)};return s.$$set=n=>{"source"in n&&r(0,i=n.source)},s.$$.update=()=>{s.$$.dirty&64&&w(t.folder)},[i,c,l,_,a,m,t,y,E]}class X extends q{constructor(e){super(),A(this,e,J,I,F,{source:0})}}export{X as component};