var G=Object.defineProperty;var K=(n,t,e)=>t in n?G(n,t,{enumerable:!0,configurable:!0,writable:!0,value:e}):n[t]=e;var a=(n,t,e)=>(K(n,typeof t!="symbol"?t+"":t,e),e);import{s as M,f as d,a as w,l as O,g as f,h as L,C as B,c as x,m as V,d as k,j as o,i as Q,w as s,W as S,D,J as Y}from"../chunks/scheduler.e2ee220a.js";import{S as Z,i as tt,a as et,t as st}from"../chunks/index.4aee2103.js";import"../chunks/loading_spinner.svelte_svelte_type_style_lang.c28228c7.js";import{H as X,a as at,S as nt}from"../chunks/routes_v2.65e45c81.js";class it{constructor(t){a(this,"http");a(this,"list_methods",async t=>this.http.get(`/auth?ugroup=${t}`));a(this,"login_next",async t=>this.http.post("/login/next",t));a(this,"login_submit",async t=>this.http.post("/login/submit",t));a(this,"altauth_generate",async(t,e)=>this.http.post(`/alt/${t}/generate`,e));a(this,"altauth_next",async(t,e,r)=>this.http.post(`/alt/${t}/next/${e}`,r));a(this,"altauth_submit",async(t,e)=>this.http.post(`/alt/${t}/submit`,e));a(this,"finish",async t=>this.http.post("/finish",t));a(this,"signup_next",async t=>this.http.post("/signup/next",t));a(this,"signup_submit",async t=>this.http.post("/signup/submit",t));a(this,"reset_submit",async t=>this.http.post("/reset/submit",t));a(this,"reset_finish",async t=>this.http.post("/reset/finish",t));a(this,"about",async t=>new X(this.http.baseURL,{"Content-Type":"application/json",Authorization:t}).get("/about"));this.http=new X(t,{"Content-Type":"application/json"})}}class lt{constructor(){a(this,"api");a(this,"site_utils");this.api=new it(at()),this.site_utils=new nt}async init(){if(this.site_utils.isLogged()){this.site_utils.gotoPortalPage();return}}async loginWithPassword(t,e){const r=await this.api.login_next({user_ident:t,password:e});if(r.data.ok){console.log("@NEXT1");const c=r.data.next_token,i=await this.api.login_submit({next_token:c});if(!i.data.ok)return i.data;const p=await this.api.finish({preauthed_token:i.data.preauthed_token});if(console.log("@FINISH",p.data),!p.ok)return p.data;this.site_utils.setAuthedData({tenant_id:p.data.tenant_id,user_token:p.data.user_token}),this.site_utils.gotoPortalPage()}else return r.data}}function ot(n){let t,e,r="User Login",c,i,p="",P,C,g,m,W="User",N,_,T,y,v,z="Password",I,h,j,b,$,U,H,R,q=rt;return{c(){t=d("div"),e=d("h3"),e.textContent=r,c=w(),i=d("p"),P=O(p),C=w(),g=d("label"),m=d("span"),m.textContent=W,N=w(),_=d("input"),T=w(),y=d("label"),v=d("span"),v.textContent=z,I=w(),h=d("input"),j=w(),b=d("button"),$=O(`

        Login`),this.h()},l(u){t=f(u,"DIV",{});var l=L(t);e=f(l,"H3",{class:!0,"data-svelte-h":!0}),B(e)!=="svelte-22soj8"&&(e.textContent=r),c=x(l),i=f(l,"P",{class:!0});var F=L(i);P=V(F,p),F.forEach(k),C=x(l),g=f(l,"LABEL",{class:!0});var A=L(g);m=f(A,"SPAN",{"data-svelte-h":!0}),B(m)!=="svelte-s6n4r1"&&(m.textContent=W),N=x(A),_=f(A,"INPUT",{class:!0,type:!0,placeholder:!0}),A.forEach(k),T=x(l),y=f(l,"LABEL",{class:!0});var E=L(y);v=f(E,"SPAN",{"data-svelte-h":!0}),B(v)!=="svelte-1kvjhoz"&&(v.textContent=z),I=x(E),h=f(E,"INPUT",{class:!0,title:!0,type:!0,placeholder:!0}),E.forEach(k),j=x(l),b=f(l,"BUTTON",{type:!0,class:!0});var J=L(b);$=V(J,`

        Login`),J.forEach(k),l.forEach(k),this.h()},h(){o(e,"class","h3"),o(i,"class","text-red-500"),o(_,"class","input p-2"),o(_,"type","text"),o(_,"placeholder","User1"),o(g,"class","label my-1"),o(h,"class","input p-2"),o(h,"title","Password"),o(h,"type","password"),o(h,"placeholder","password"),o(y,"class","label my-1"),o(b,"type","button"),o(b,"class","btn variant-filled my-1")},m(u,l){Q(u,t,l),s(t,e),s(t,c),s(t,i),s(i,P),s(t,C),s(t,g),s(g,m),s(g,N),s(g,_),S(_,n[0]),s(t,T),s(t,y),s(y,v),s(y,I),s(y,h),S(h,n[1]),s(t,j),s(t,b),s(b,$),U=!0,H||(R=[D(_,"input",n[4]),D(h,"input",n[5]),D(b,"click",n[6])],H=!0)},p(u,[l]){l&1&&_.value!==u[0]&&S(_,u[0]),l&2&&h.value!==u[1]&&S(h,u[1])},i(u){U||(et(q),U=!0)},o(u){st(q),U=!1},d(u){u&&k(t),H=!1,Y(R)}}}let rt=!1;function ut(n,t,e){const r=new lt;let c="",i="",p="";function P(){c=this.value,e(0,c)}function C(){i=this.value,e(1,i)}return[c,i,p,r,P,C,async()=>{await r.init();const m=await r.loginWithPassword(c,i);m&&e(2,p=m.message)}]}class ft extends Z{constructor(t){super(),tt(this,t,ut,ot,M,{})}}export{ft as component};
