import{a4 as x,ah as l}from"./scheduler.e2ee220a.js";function $(o){const t=o-1;return t*t*t+1}function S(o,{delay:t=0,duration:s=400,easing:a=x}={}){const c=+getComputedStyle(o).opacity;return{delay:t,duration:s,easing:a,css:r=>`opacity: ${r*c}`}}function U(o,{delay:t=0,duration:s=400,easing:a=$,x:c=0,y:r=0,opacity:e=0}={}){const n=getComputedStyle(o),i=+n.opacity,y=n.transform==="none"?"":n.transform,u=i*(1-e),[p,f]=l(c),[d,g]=l(r);return{delay:t,duration:s,easing:a,css:(m,_)=>`
			transform: ${y} translate(${(1-m)*p}${f}, ${(1-m)*d}${g});
			opacity: ${i-u*_}`}}function V(o,{delay:t=0,duration:s=400,easing:a=$,start:c=0,opacity:r=0}={}){const e=getComputedStyle(o),n=+e.opacity,i=e.transform==="none"?"":e.transform,y=1-c,u=n*(1-r);return{delay:t,duration:s,easing:a,css:(p,f)=>`
			transform: ${i} scale(${1-y*f});
			opacity: ${n-u*f}
		`}}export{S as a,U as f,V as s};
