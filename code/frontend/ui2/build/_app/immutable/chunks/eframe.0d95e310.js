import{s as w,f as y,g as b,h as x,d as u,M as f,j as i,i as E,D as M,A as d,E as k,p as S}from"./scheduler.e2ee220a.js";import{S as j,i as D}from"./index.4aee2103.js";function F(t){let e,s,l,o;return{c(){e=y("iframe"),this.h()},l(n){e=b(n,"IFRAME",{src:!0,title:!0,class:!0,allow:!0,sandbox:!0}),x(e).forEach(u),this.h()},h(){f(e.src,s=t[3])||i(e,"src",s),i(e,"title",t[2]),i(e,"class","border-green-200 w-full h-full transition-all"),i(e,"allow","accelerometer; ambient-light-sensor; autoplay; battery; camera; clipboard-write; document-domain; encrypted-media; fullscreen; geolocation; gyroscope; layout-animations; legacy-image-formats; magnetometer; microphone; midi; oversized-images; payment; picture-in-picture; publickey-credentials-get; sync-xhr; usb; vr ; wake-lock; xr-spatial-tracking"),i(e,"sandbox","allow-forms allow-modals allow-popups allow-popups-to-escape-sandbox allow-same-origin allow-scripts allow-downloads allow-storage-access-by-user-activation")},m(n,r){E(n,e,r),t[5](e),l||(o=M(e,"load",t[6]),l=!0)},p(n,[r]){r&8&&!f(e.src,s=n[3])&&i(e,"src",s),r&4&&i(e,"title",n[2])},i:d,o:d,d(n){n&&u(e),t[5](null),l=!1,o()}}}function q(t,e,s){let{name:l}=e,{exec_data:o}=e,{iframe:n=null}=e,{url:r}=e,{chan:c=new MessageChannel}=e;const g=k(),_=a=>{console.log("@onFramwData",a);const m=JSON.parse(a.data);if(m.mtype!=="get_exec_data"){g("eframe_message",m);return}console.log("sending_exec_data"),c.port1.postMessage(JSON.stringify({mtype:"exec_data",data:o}),[])};c.port1.onmessage=_;function h(a){S[a?"unshift":"push"](()=>{n=a,s(0,n)})}const p=a=>{n.contentWindow.postMessage("transfer_port","*",[c.port2])};return t.$$set=a=>{"name"in a&&s(2,l=a.name),"exec_data"in a&&s(4,o=a.exec_data),"iframe"in a&&s(0,n=a.iframe),"url"in a&&s(3,r=a.url),"chan"in a&&s(1,c=a.chan)},[n,c,l,r,o,h,p]}class B extends j{constructor(e){super(),D(this,e,q,F,w,{name:2,exec_data:4,iframe:0,url:3,chan:1})}}const C=t=>{let e="";return t.auth_type==="auto_inject"&&(e="/z/pages/agent/inject"),t.start_page&&(e===""?e=`/${t.start_page}`:e=`${e}?&start_page=${t.start_page}`),`http://${t.domain}:${location.port||80}${e}`};export{C as B,B as E};