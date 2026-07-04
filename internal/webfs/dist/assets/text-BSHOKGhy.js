import{m as C,n as i,d as $,h as l,p as B,q as h,v as T,x as c,ca as z,aA as R}from"./index-B5SBiCPF.js";import{u as S}from"./Space-C-IUYuEv.js";const V=C("text",`
 transition: color .3s var(--n-bezier);
 color: var(--n-text-color);
`,[i("strong",`
 font-weight: var(--n-font-weight-strong);
 `),i("italic",{fontStyle:"italic"}),i("underline",{textDecoration:"underline"}),i("code",`
 line-height: 1.4;
 display: inline-block;
 font-family: var(--n-font-famliy-mono);
 transition: 
 color .3s var(--n-bezier),
 border-color .3s var(--n-bezier),
 background-color .3s var(--n-bezier);
 box-sizing: border-box;
 padding: .05em .35em 0 .35em;
 border-radius: var(--n-code-border-radius);
 font-size: .9em;
 color: var(--n-code-text-color);
 background-color: var(--n-code-color);
 border: var(--n-code-border);
 `)]),P=Object.assign(Object.assign({},h.props),{code:Boolean,type:{type:String,default:"default"},delete:Boolean,strong:Boolean,italic:Boolean,underline:Boolean,depth:[String,Number],tag:String,as:{type:String,validator:()=>!0,default:void 0}}),D=$({name:"Text",props:P,setup(e){const{mergedClsPrefixRef:s,inlineThemeDisabled:r}=B(e),o=h("Typography","-text",V,z,e,s),n=c(()=>{const{depth:a,type:d}=e,u=d==="default"?a===void 0?"textColor":`textColor${a}Depth`:R("textColor",d),{common:{fontWeightStrong:m,fontFamilyMono:g,cubicBezierEaseInOut:x},self:{codeTextColor:f,codeBorderRadius:b,codeColor:v,codeBorder:y,[u]:p}}=o.value;return{"--n-bezier":x,"--n-text-color":p,"--n-font-weight-strong":m,"--n-font-famliy-mono":g,"--n-code-border-radius":b,"--n-code-text-color":f,"--n-code-color":v,"--n-code-border":y}}),t=r?T("text",c(()=>`${e.type[0]}${e.depth||""}`),n,e):void 0;return{mergedClsPrefix:s,compitableTag:S(e,["as","tag"]),cssVars:r?void 0:n,themeClass:t==null?void 0:t.themeClass,onRender:t==null?void 0:t.onRender}},render(){var e,s,r;const{mergedClsPrefix:o}=this;(e=this.onRender)===null||e===void 0||e.call(this);const n=[`${o}-text`,this.themeClass,{[`${o}-text--code`]:this.code,[`${o}-text--delete`]:this.delete,[`${o}-text--strong`]:this.strong,[`${o}-text--italic`]:this.italic,[`${o}-text--underline`]:this.underline}],t=(r=(s=this.$slots).default)===null||r===void 0?void 0:r.call(s);return this.code?l("code",{class:n,style:this.cssVars},this.delete?l("del",null,t):t):this.delete?l("del",{class:n,style:this.cssVars},t):l(this.compitableTag||"span",{class:n,style:this.cssVars},t)}});export{D as N};
