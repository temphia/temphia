import{s as b,f as d,g as _,h as m,d as r,j as p,i as g,A as c,Q as v,l as y,m as j,w as k,D as x,R as T,n as w}from"./scheduler.e2ee220a.js";import{e as h}from"./each.e59479a4.js";import{S as O,i as D}from"./index.4aee2103.js";function f(i,n,a){const e=i.slice();return e[1]=n[a][0],e[2]=n[a][1],e}function u(i){let n,a=i[1]+"",e,s,l;return{c(){n=d("button"),e=y(a),this.h()},l(t){n=_(t,"BUTTON",{class:!0});var o=m(n);e=j(o,a),o.forEach(r),this.h()},h(){p(n,"class","px-2 py-1 rounded-full bg-white hover:text-white hover:bg-slate-500 border border-slate-600")},m(t,o){g(t,n,o),k(n,e),s||(l=x(n,"click",function(){T(i[2])&&i[2].apply(this,arguments)}),s=!0)},p(t,o){i=t,o&1&&a!==(a=i[1]+"")&&w(e,a)},d(t){t&&r(n),s=!1,l()}}}function E(i){let n,a=h(Object.entries(i[0])),e=[];for(let s=0;s<a.length;s+=1)e[s]=u(f(i,a,s));return{c(){n=d("div");for(let s=0;s<e.length;s+=1)e[s].c();this.h()},l(s){n=_(s,"DIV",{class:!0});var l=m(n);for(let t=0;t<e.length;t+=1)e[t].l(l);l.forEach(r),this.h()},h(){p(n,"class","flex justify-end p-2 gap-1")},m(s,l){g(s,n,l);for(let t=0;t<e.length;t+=1)e[t]&&e[t].m(n,null)},p(s,[l]){if(l&1){a=h(Object.entries(s[0]));let t;for(t=0;t<a.length;t+=1){const o=f(s,a,t);e[t]?e[t].p(o,l):(e[t]=u(o),e[t].c(),e[t].m(n,null))}for(;t<e.length;t+=1)e[t].d(1);e.length=a.length}},i:c,o:c,d(s){s&&r(n),v(e,s)}}}function S(i,n,a){let{actions:e={}}=n;return i.$$set=s=>{"actions"in s&&a(0,e=s.actions)},[e]}class C extends O{constructor(n){super(),D(this,n,S,E,b,{actions:0})}}export{C as T};