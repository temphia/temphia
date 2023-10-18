import{s as q,e as F,i as L,d as U,A as B,a as R,c as V,P as x}from"../chunks/scheduler.e2ee220a.js";import{S as E,i as H,g as I,t as f,c as j,a as _,b as w,d as y,m as h,e as A}from"../chunks/index.4aee2103.js";import"../chunks/paths.2eaeb908.js";import{F as ee}from"../chunks/floating_add.1495f8dc.js";import"../chunks/loading_spinner.svelte_svelte_type_style_lang.c28228c7.js";import{A as te}from"../chunks/autotable.0057460b.js";import{L as z}from"../chunks/loading_spinner.4ef87ddf.js";import{T as ne}from"../chunks/top_actions.3bf62959.js";import{A as T}from"../chunks/action_picker.7081fe98.js";import{T as Y,a as Z,b as v,c as X}from"../chunks/target.d267d535.js";function ae(c){let t,n;return t=new T({props:{actions:c[5].map(c[17]),title:"Pick Agent"}}),{c(){w(t.$$.fragment)},l(e){y(t.$$.fragment,e)},m(e,a){h(t,e,a),n=!0},p(e,a){const i={};a&32&&(i.actions=e[5].map(e[17])),t.$set(i)},i(e){n||(_(t.$$.fragment,e),n=!0)},o(e){f(t.$$.fragment,e),n=!1},d(e){A(t,e)}}}function re(c){let t,n;return t=new T({props:{actions:c[4].map(c[16]),title:"Pick Plug"}}),{c(){w(t.$$.fragment)},l(e){y(t.$$.fragment,e)},m(e,a){h(t,e,a),n=!0},p(e,a){const i={};a&16&&(i.actions=e[4].map(e[16])),t.$set(i)},i(e){n||(_(t.$$.fragment,e),n=!0)},o(e){f(t.$$.fragment,e),n=!1},d(e){A(t,e)}}}function ie(c){let t,n;return t=new T({props:{actions:c[3].map(c[15]),title:"Pick Data Sheet"}}),{c(){w(t.$$.fragment)},l(e){y(t.$$.fragment,e)},m(e,a){h(t,e,a),n=!0},p(e,a){const i={};a&8&&(i.actions=e[3].map(e[15])),t.$set(i)},i(e){n||(_(t.$$.fragment,e),n=!0)},o(e){f(t.$$.fragment,e),n=!1},d(e){A(t,e)}}}function ce(c){let t,n;return t=new T({props:{actions:c[2].map(c[14]),title:"Pick Data Group"}}),{c(){w(t.$$.fragment)},l(e){y(t.$$.fragment,e)},m(e,a){h(t,e,a),n=!0},p(e,a){const i={};a&4&&(i.actions=e[2].map(e[14])),t.$set(i)},i(e){n||(_(t.$$.fragment,e),n=!0)},o(e){f(t.$$.fragment,e),n=!1},d(e){A(t,e)}}}function oe(c){let t,n;return t=new T({props:{actions:c[1].map(c[13]),title:"Pick Data source"}}),{c(){w(t.$$.fragment)},l(e){y(t.$$.fragment,e)},m(e,a){h(t,e,a),n=!0},p(e,a){const i={};a&2&&(i.actions=e[1].map(e[13])),t.$set(i)},i(e){n||(_(t.$$.fragment,e),n=!0)},o(e){f(t.$$.fragment,e),n=!1},d(e){A(t,e)}}}function se(c){let t,n;return t=new z({props:{classes:""}}),{c(){w(t.$$.fragment)},l(e){y(t.$$.fragment,e)},m(e,a){h(t,e,a),n=!0},p:B,i(e){n||(_(t.$$.fragment,e),n=!0)},o(e){f(t.$$.fragment,e),n=!1},d(e){A(t,e)}}}function pe(c){let t,n,e,a;const i=[se,oe,ce,ie,re,ae],s=[];function k(r,l){return r[0]?0:r[6]==="pick_source"?1:r[6]==="pick_group"?2:r[6]==="pick_sheet"?3:r[6]==="pick_plug"?4:r[6]==="pick_agent"?5:-1}return~(t=k(c))&&(n=s[t]=i[t](c)),{c(){n&&n.c(),e=F()},l(r){n&&n.l(r),e=F()},m(r,l){~t&&s[t].m(r,l),L(r,e,l),a=!0},p(r,[l]){let g=t;t=k(r),t===g?~t&&s[t].p(r,l):(n&&(I(),f(s[g],1,1,()=>{s[g]=null}),j()),~t?(n=s[t],n?n.p(r,l):(n=s[t]=i[t](r),n.c()),_(n,1),n.m(e.parentNode,e)):n=null)},i(r){a||(_(n),a=!0)},o(r){f(n),a=!1},d(r){r&&U(e),~t&&s[t].d(r)}}}function le(c,t,n){let{service:e}=t,a=!0,i="",s="",k="",r="",l="",g=[],p=[],$=[],P=[],u=[],b="pick_source";const C=async()=>{n(1,g=await e.api_manager.self_data.get_data_sources()),n(0,a=!1)},S=async o=>{n(0,a=!0),n(6,b="pick_group"),i=o.name;const d=await(await e.api_manager.get_admin_data_api()).list_group(i);d.ok&&(n(2,p=d.data),n(0,a=!1))},G=async o=>{n(0,a=!0),n(6,b="pick_sheet"),s=o.data.slug;const d=await e.api_manager.get_admin_data_api().list_sheet(i,s);d.ok&&(n(3,$=d.data),n(0,a=!1))},m=async o=>{n(0,a=!0),n(6,b="pick_plug"),k=o.data.__id;const d=await e.api_manager.get_admin_plug_api().list_plug();d.ok&&(n(4,P=d.data),n(0,a=!1))},D=async o=>{n(0,a=!0),n(6,b="pick_agent"),r=o.data.id;const d=await e.api_manager.get_admin_plug_api().list_agent(r);d.ok&&(n(5,u=d.data),n(0,a=!1))},N=async o=>{n(0,a=!0),l=o.data.id,e.nav.admin_target_app_new({target_type:Y,target:`${i}/${s}/${k}`,context_type:"global.1",plug_id:r,agent_id:l}),e.utils.small_modal_close()};C();const J=o=>({action:S,icon:"hashtag",info:"",name:o}),K=o=>({action:G,icon:"hashtag",info:o.description,name:o.slug,data:o}),M=o=>({action:m,icon:"hashtag",info:o.name,name:o.__id,data:o}),O=o=>({action:D,icon:"hashtag",info:o.name,name:o.id,data:o}),Q=o=>({action:N,icon:"hashtag",info:o.name,name:o.id,data:o});return c.$$set=o=>{"service"in o&&n(12,e=o.service)},[a,g,p,$,P,u,b,S,G,m,D,N,e,J,K,M,O,Q]}class fe extends E{constructor(t){super(),H(this,t,le,pe,q,{service:12})}}function _e(c){let t,n;return t=new T({props:{actions:c[5].map(c[17]),title:"Pick Agent"}}),{c(){w(t.$$.fragment)},l(e){y(t.$$.fragment,e)},m(e,a){h(t,e,a),n=!0},p(e,a){const i={};a&32&&(i.actions=e[5].map(e[17])),t.$set(i)},i(e){n||(_(t.$$.fragment,e),n=!0)},o(e){f(t.$$.fragment,e),n=!1},d(e){A(t,e)}}}function ue(c){let t,n;return t=new T({props:{actions:c[4].map(c[16]),title:"Pick Plug"}}),{c(){w(t.$$.fragment)},l(e){y(t.$$.fragment,e)},m(e,a){h(t,e,a),n=!0},p(e,a){const i={};a&16&&(i.actions=e[4].map(e[16])),t.$set(i)},i(e){n||(_(t.$$.fragment,e),n=!0)},o(e){f(t.$$.fragment,e),n=!1},d(e){A(t,e)}}}function me(c){let t,n;return t=new T({props:{actions:c[3].map(c[15]),title:"Pick Data Table"}}),{c(){w(t.$$.fragment)},l(e){y(t.$$.fragment,e)},m(e,a){h(t,e,a),n=!0},p(e,a){const i={};a&8&&(i.actions=e[3].map(e[15])),t.$set(i)},i(e){n||(_(t.$$.fragment,e),n=!0)},o(e){f(t.$$.fragment,e),n=!1},d(e){A(t,e)}}}function ge(c){let t,n;return t=new T({props:{actions:c[2].map(c[14]),title:"Pick Data Group"}}),{c(){w(t.$$.fragment)},l(e){y(t.$$.fragment,e)},m(e,a){h(t,e,a),n=!0},p(e,a){const i={};a&4&&(i.actions=e[2].map(e[14])),t.$set(i)},i(e){n||(_(t.$$.fragment,e),n=!0)},o(e){f(t.$$.fragment,e),n=!1},d(e){A(t,e)}}}function ke(c){let t,n;return t=new T({props:{actions:c[1].map(c[13]),title:"Pick Data source"}}),{c(){w(t.$$.fragment)},l(e){y(t.$$.fragment,e)},m(e,a){h(t,e,a),n=!0},p(e,a){const i={};a&2&&(i.actions=e[1].map(e[13])),t.$set(i)},i(e){n||(_(t.$$.fragment,e),n=!0)},o(e){f(t.$$.fragment,e),n=!1},d(e){A(t,e)}}}function de(c){let t,n;return t=new z({props:{classes:""}}),{c(){w(t.$$.fragment)},l(e){y(t.$$.fragment,e)},m(e,a){h(t,e,a),n=!0},p:B,i(e){n||(_(t.$$.fragment,e),n=!0)},o(e){f(t.$$.fragment,e),n=!1},d(e){A(t,e)}}}function $e(c){let t,n,e,a;const i=[de,ke,ge,me,ue,_e],s=[];function k(r,l){return r[0]?0:r[6]==="pick_source"?1:r[6]==="pick_group"?2:r[6]==="pick_table"?3:r[6]==="pick_plug"?4:r[6]==="pick_agent"?5:-1}return~(t=k(c))&&(n=s[t]=i[t](c)),{c(){n&&n.c(),e=F()},l(r){n&&n.l(r),e=F()},m(r,l){~t&&s[t].m(r,l),L(r,e,l),a=!0},p(r,[l]){let g=t;t=k(r),t===g?~t&&s[t].p(r,l):(n&&(I(),f(s[g],1,1,()=>{s[g]=null}),j()),~t?(n=s[t],n?n.p(r,l):(n=s[t]=i[t](r),n.c()),_(n,1),n.m(e.parentNode,e)):n=null)},i(r){a||(_(n),a=!0)},o(r){f(n),a=!1},d(r){r&&U(e),~t&&s[t].d(r)}}}function be(c,t,n){let{service:e}=t,a=!0,i="",s="",k="",r="",l="",g=[],p=[],$=[],P=[],u=[],b="pick_source";const C=async()=>{n(1,g=await e.api_manager.self_data.get_data_sources()),n(0,a=!1)},S=async o=>{n(0,a=!0),n(6,b="pick_group"),i=o.name;const d=await(await e.api_manager.get_admin_data_api()).list_group(i);d.ok&&(n(2,p=d.data),n(0,a=!1))},G=async o=>{n(0,a=!0),n(6,b="pick_table"),s=o.data.slug;const d=await e.api_manager.get_admin_data_api().list_tables(i,s);d.ok&&(n(3,$=d.data),n(0,a=!1))},m=async o=>{n(0,a=!0),n(6,b="pick_plug"),k=o.data.slug;const d=await e.api_manager.get_admin_plug_api().list_plug();d.ok&&(n(4,P=d.data),n(0,a=!1))},D=async o=>{n(0,a=!0),n(6,b="pick_agent"),r=o.data.id;const d=await e.api_manager.get_admin_plug_api().list_agent(r);d.ok&&(n(5,u=d.data),n(0,a=!1))},N=async o=>{n(0,a=!0),l=o.data.id,e.nav.admin_target_app_new({target_type:Z,target:`${i}/${s}/${k}`,context_type:"global.1",plug_id:r,agent_id:l}),e.utils.small_modal_close()};C();const J=o=>({action:S,icon:"hashtag",info:"",name:o}),K=o=>({action:G,icon:"hashtag",info:o.description,name:o.slug,data:o}),M=o=>({action:m,icon:"hashtag",info:o.name,name:o.slug,data:o}),O=o=>({action:D,icon:"hashtag",info:o.name,name:o.id,data:o}),Q=o=>({action:N,icon:"hashtag",info:o.name,name:o.id,data:o});return c.$$set=o=>{"service"in o&&n(12,e=o.service)},[a,g,p,$,P,u,b,S,G,m,D,N,e,J,K,M,O,Q]}class we extends E{constructor(t){super(),H(this,t,be,$e,q,{service:12})}}function ye(c){let t,n;return t=new T({props:{actions:c[3].map(c[11]),title:"Pick Agent"}}),{c(){w(t.$$.fragment)},l(e){y(t.$$.fragment,e)},m(e,a){h(t,e,a),n=!0},p(e,a){const i={};a&8&&(i.actions=e[3].map(e[11])),t.$set(i)},i(e){n||(_(t.$$.fragment,e),n=!0)},o(e){f(t.$$.fragment,e),n=!1},d(e){A(t,e)}}}function he(c){let t,n;return t=new T({props:{actions:c[2].map(c[10]),title:"Pick Plug"}}),{c(){w(t.$$.fragment)},l(e){y(t.$$.fragment,e)},m(e,a){h(t,e,a),n=!0},p(e,a){const i={};a&4&&(i.actions=e[2].map(e[10])),t.$set(i)},i(e){n||(_(t.$$.fragment,e),n=!0)},o(e){f(t.$$.fragment,e),n=!1},d(e){A(t,e)}}}function Ae(c){let t,n;return t=new T({props:{actions:c[1].map(c[9]),title:"Pick User Group"}}),{c(){w(t.$$.fragment)},l(e){y(t.$$.fragment,e)},m(e,a){h(t,e,a),n=!0},p(e,a){const i={};a&2&&(i.actions=e[1].map(e[9])),t.$set(i)},i(e){n||(_(t.$$.fragment,e),n=!0)},o(e){f(t.$$.fragment,e),n=!1},d(e){A(t,e)}}}function Pe(c){let t,n;return t=new z({props:{classes:""}}),{c(){w(t.$$.fragment)},l(e){y(t.$$.fragment,e)},m(e,a){h(t,e,a),n=!0},p:B,i(e){n||(_(t.$$.fragment,e),n=!0)},o(e){f(t.$$.fragment,e),n=!1},d(e){A(t,e)}}}function Te(c){let t,n,e,a;const i=[Pe,Ae,he,ye],s=[];function k(r,l){return r[0]?0:r[4]==="pick_group"?1:r[4]==="pick_plug"?2:r[4]==="pick_agent"?3:-1}return~(t=k(c))&&(n=s[t]=i[t](c)),{c(){n&&n.c(),e=F()},l(r){n&&n.l(r),e=F()},m(r,l){~t&&s[t].m(r,l),L(r,e,l),a=!0},p(r,[l]){let g=t;t=k(r),t===g?~t&&s[t].p(r,l):(n&&(I(),f(s[g],1,1,()=>{s[g]=null}),j()),~t?(n=s[t],n?n.p(r,l):(n=s[t]=i[t](r),n.c()),_(n,1),n.m(e.parentNode,e)):n=null)},i(r){a||(_(n),a=!0)},o(r){f(n),a=!1},d(r){r&&U(e),~t&&s[t].d(r)}}}function Ne(c,t,n){let{service:e}=t,a=!0,i="",s="",k="",r=[],l=[],g=[],p="pick_group";const $=async()=>{const D=await e.api_manager.get_admin_ugroup_api().list();D.ok&&(n(1,r=D.data),n(0,a=!1))},P=async m=>{n(0,a=!0),n(4,p="pick_plug"),i=m.data.slug;const N=await e.api_manager.get_admin_plug_api().list_plug();N.ok&&(n(2,l=N.data),n(0,a=!1))},u=async m=>{n(0,a=!0),n(4,p="pick_agent"),s=m.data.id;const N=await e.api_manager.get_admin_plug_api().list_agent(s);N.ok&&(n(3,g=N.data),n(0,a=!1))},b=async m=>{n(0,a=!0),k=m.data.id,e.nav.admin_target_app_new({target_type:v,target:i,context_type:"app.1",plug_id:s,agent_id:k}),e.utils.small_modal_close()};$();const C=m=>({action:P,icon:"hashtag",info:m.name,name:m.slug,data:m}),S=m=>({action:u,icon:"hashtag",info:m.name,name:m.id,data:m}),G=m=>({action:b,icon:"hashtag",info:m.name,name:m.id,data:m});return c.$$set=m=>{"service"in m&&n(8,e=m.service)},[a,r,l,g,p,P,u,b,e,C,S,G]}class De extends E{constructor(t){super(),H(this,t,Ne,Te,q,{service:8})}}function Ce(c){let t,n;return t=new T({props:{actions:c[0],title:"New Target app type"}}),{c(){w(t.$$.fragment)},l(e){y(t.$$.fragment,e)},m(e,a){h(t,e,a),n=!0},p:B,i(e){n||(_(t.$$.fragment,e),n=!0)},o(e){f(t.$$.fragment,e),n=!1},d(e){A(t,e)}}}function Se(c,t,n){let{service:e}=t;const a=[{name:v,icon:"user-group",info:"App for group of people",action:()=>e.utils.small_modal_open(De,{service:e})},{name:Z,icon:"table",info:"Datatable widget",action:()=>e.utils.small_modal_open(we,{service:e})},{name:Y,icon:"table",info:"DataSheet widget",action:()=>{e.utils.small_modal_open(fe,{service:e})}},{name:X,icon:"globe-alt",info:"Domain widget",action:()=>{e.nav.admin_target_app_new({target_type:X,context_type:"widget.1"}),e.utils.small_modal_close()}}];return c.$$set=i=>{"service"in i&&n(1,e=i.service)},[a,e]}class Ge extends E{constructor(t){super(),H(this,t,Se,Ce,q,{service:1})}}function We(c){let t,n;return t=new te({props:{action_key:"id",actions:[{Name:"Edit",Action:c[4],icon:"pencil-alt"},{Name:"Delete",Class:"bg-red-400",Action:c[5],icon:"trash"}],key_names:[["name","Name"],["id","Id"],["target_type","Type"],["target","Target"],["context_type","Context Type"]],color:["target_type"],datas:c[1]}}),{c(){w(t.$$.fragment)},l(e){y(t.$$.fragment,e)},m(e,a){h(t,e,a),n=!0},p(e,a){const i={};a&2&&(i.datas=e[1]),t.$set(i)},i(e){n||(_(t.$$.fragment,e),n=!0)},o(e){f(t.$$.fragment,e),n=!1},d(e){A(t,e)}}}function Fe(c){let t,n;return t=new z({}),{c(){w(t.$$.fragment)},l(e){y(t.$$.fragment,e)},m(e,a){h(t,e,a),n=!0},p:B,i(e){n||(_(t.$$.fragment,e),n=!0)},o(e){f(t.$$.fragment,e),n=!1},d(e){A(t,e)}}}function Le(c){let t,n,e,a,i,s,k;t=new ne({props:{actions:c[3]}});const r=[Fe,We],l=[];function g(p,$){return p[2]?0:1}return e=g(c),a=l[e]=r[e](c),s=new ee({props:{onClick:c[0]}}),{c(){w(t.$$.fragment),n=R(),a.c(),i=R(),w(s.$$.fragment)},l(p){y(t.$$.fragment,p),n=V(p),a.l(p),i=V(p),y(s.$$.fragment,p)},m(p,$){h(t,p,$),L(p,n,$),l[e].m(p,$),L(p,i,$),h(s,p,$),k=!0},p(p,[$]){const P={};$&8&&(P.actions=p[3]),t.$set(P);let u=e;e=g(p),e===u?l[e].p(p,$):(I(),f(l[u],1,1,()=>{l[u]=null}),j(),a=l[e],a?a.p(p,$):(a=l[e]=r[e](p),a.c()),_(a,1),a.m(i.parentNode,i));const b={};$&1&&(b.onClick=p[0]),s.$set(b)},i(p){k||(_(t.$$.fragment,p),_(a),_(s.$$.fragment,p),k=!0)},o(p){f(t.$$.fragment,p),f(a),f(s.$$.fragment,p),k=!1},d(p){p&&(U(n),U(i)),A(t,p),l[e].d(p),A(s,p)}}}function Ue(c,t,n){let{ttype:e=void 0}=t,{target:a=void 0}=t;const i=x("__app__");let{action_new:s=()=>{i.utils.small_modal_open(Ge,{service:i})}}=t,k=[],r=!0;const l=i.api_manager.get_admin_target_api(),g=async()=>{let u;n(2,r=!0),e?u=await l.listAppByType(e,a):u=await l.listApp(),u.ok&&(n(1,k=u.data),n(2,r=!1))};g();const p=(u,b)=>i.nav.admin_target_app_edit(b.target_type,Number(u)),$=async(u,b)=>{console.log("@delete"),n(2,r=!0);try{const C=await l.deleteApp(b.target_type,Number(u))}catch{}finally{g()}},P={Hooks:()=>i.nav.admin_target_hooks()};return e&&(P["All Apps"]=()=>i.nav.admin_target_apps()),c.$$set=u=>{"ttype"in u&&n(6,e=u.ttype),"target"in u&&n(7,a=u.target),"action_new"in u&&n(0,s=u.action_new)},[s,k,r,P,p,$,e,a]}class Oe extends E{constructor(t){super(),H(this,t,Ue,Le,q,{ttype:6,target:7,action_new:0})}}export{Oe as component};