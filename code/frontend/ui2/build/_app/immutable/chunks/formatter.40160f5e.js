const p="shorttext",f="phonenumber",g="select",l="file",y="multifile",d="checkbox",h="number",C="location",m="datetime",b="singleuser",_="multiuser",x="__id",k="less_than",u=[p,f,h,g],a=["name","title"],j=s=>{const c={},t=[];return u.forEach(e=>{Object.values(s).forEach(i=>{c[i.slug]||i.ref_type||i.ctype===e&&(t.push(i.slug),c[i.slug]=!0)})}),Object.values(s).forEach(e=>{u.includes(e.ctype)||c[e.slug]||e.ref_type||(t.push(e.slug),c[e.slug]=!0)}),Object.values(s).forEach(e=>{c[e.slug]||e.ref_type&&(t.push(e.slug),c[e.slug]=!0)}),t},O=(s,c)=>{const t={other:[]},e={};for(let r=0;r<a.length;r++){const n=a[r],o=s[n];o&&o.ctype===p&&(t.name=n,e[n]=!0);break}if(s.image&&(t.image="image",e.image=!0),s.description&&(t.description="description",e.description=!0),s.user&&(t.user="user",e.user=!0),s.tag&&(t.tag="tag",e.tag=!0),!t.image){const r=Object.values(s).filter(n=>n.ctype===y||n.ctype===l);r.length&&(t.image=r[0].slug,e[t.image]=!0)}const i=Object.keys(s);for(let r=0;r<i.length;r++){const n=i[r];e[n]||t.other.push(n)}return t};export{y as C,k as F,x as K,l as a,m as b,O as c,d,C as e,b as f,j as g,_ as h};
