import{u as Rn,a as Be,A as _e}from"./client-DbLYzOci.js";import{d as ge,h as T,r as O,ak as fr,al as hr,am as Ve,an as gr,j as br,ao as pr,ap as vr,l as Sn,aq as $t,ar as ue,p as He,C as Et,v as kt,E as Z,x as ne,B as Y,N as Pn,as as mr,I as yr,m as R,z as M,A as q,at as Bn,L as wr,q as ke,au as Cr,D as he,y as Tn,av as _t,n as $,F as St,aw as Je,ax as xr,ay as Rr,az as Sr,aA as K,ad as Ye,aB as de,ac as Pr,Q as Br,K as We,H as Tr,aC as Er,ag as Ge,V as Qe,a5 as je,U as zt,O as kr,aD as zr,a8 as Ar,a3 as Nr,aE as Ir,ah as $r,aF as Oe,aG as Me,X as Ee,u as _,aH as _r,w as D,o as fe,g as J,Y as Lt,b as U,c as Pt,a as qe,Z as Bt,aj as Lr,ai as Mr}from"./index-B5SBiCPF.js";import{A as Ur,N as Mt}from"./InputNumber-DfuGl_yW.js";import{a as Fr,d as Ut,e as Dr,p as jr,o as Or,N as qr}from"./Tag-pmWKzwFY.js";import{u as Ft,N as se}from"./Space-C-IUYuEv.js";import{N as Vr}from"./Select-Brkp7OYw.js";import{a as Dt,u as En,N as Te}from"./use-message-CQekcbBa.js";import{N as Wr}from"./Alert-ssnFm5Cp.js";import{N as jt}from"./text-BSHOKGhy.js";import{N as Hr,a as Kr}from"./DrawerContent-BezLXm6F.js";import{N as Jr}from"./DataTable-CoKW17v_.js";import{N as Yr,a as Ue}from"./FormItem-DAB72hxf.js";import"./Dropdown-CHh3TWO-.js";const Gr=Ut(".v-x-scroll",{overflow:"auto",scrollbarWidth:"none"},[Ut("&::-webkit-scrollbar",{width:0,height:0})]),Qr=ge({name:"XScroll",props:{disabled:Boolean,onScroll:Function},setup(){const e=O(null);function a(t){!(t.currentTarget.offsetWidth<t.currentTarget.scrollWidth)||t.deltaY===0||(t.currentTarget.scrollLeft+=t.deltaY+t.deltaX,t.preventDefault())}const o=fr();return Gr.mount({id:"vueuc/x-scroll",head:!0,anchorMetaName:Fr,ssr:o}),Object.assign({selfRef:e,handleWheel:a},{scrollTo(...t){var n;(n=e.value)===null||n===void 0||n.scrollTo(...t)}})},render(){return T("div",{ref:"selfRef",onScroll:this.onScroll,onWheel:this.disabled?void 0:this.handleWheel,class:"v-x-scroll"},this.$slots)}});var Xr=/\s/;function Zr(e){for(var a=e.length;a--&&Xr.test(e.charAt(a)););return a}var ea=/^\s+/;function ta(e){return e&&e.slice(0,Zr(e)+1).replace(ea,"")}var Ot=NaN,na=/^[-+]0x[0-9a-f]+$/i,ra=/^0b[01]+$/i,aa=/^0o[0-7]+$/i,oa=parseInt;function qt(e){if(typeof e=="number")return e;if(hr(e))return Ot;if(Ve(e)){var a=typeof e.valueOf=="function"?e.valueOf():e;e=Ve(a)?a+"":a}if(typeof e!="string")return e===0?e:+e;e=ta(e);var o=ra.test(e);return o||aa.test(e)?oa(e.slice(2),o?2:8):na.test(e)?Ot:+e}var Xe=function(){return gr.Date.now()},ia="Expected a function",sa=Math.max,la=Math.min;function da(e,a,o){var r,t,n,i,s,l,c=0,h=!1,x=!1,g=!0;if(typeof e!="function")throw new TypeError(ia);a=qt(a)||0,Ve(o)&&(h=!!o.leading,x="maxWait"in o,n=x?sa(qt(o.maxWait)||0,a):n,g="trailing"in o?!!o.trailing:g);function b(d){var u=r,C=t;return r=t=void 0,c=d,i=e.apply(C,u),i}function S(d){return c=d,s=setTimeout(B,a),h?b(d):i}function P(d){var u=d-l,C=d-c,p=a-u;return x?la(p,n-C):p}function F(d){var u=d-l,C=d-c;return l===void 0||u>=a||u<0||x&&C>=n}function B(){var d=Xe();if(F(d))return y(d);s=setTimeout(B,P(d))}function y(d){return s=void 0,g&&r?b(d):(r=t=void 0,i)}function f(){s!==void 0&&clearTimeout(s),c=0,r=l=t=s=void 0}function w(){return s===void 0?i:y(Xe())}function E(){var d=Xe(),u=F(d);if(r=arguments,t=this,l=d,u){if(s===void 0)return S(l);if(x)return clearTimeout(s),s=setTimeout(B,a),b(l)}return s===void 0&&(s=setTimeout(B,a)),i}return E.cancel=f,E.flush=w,E}var ua="Expected a function";function ca(e,a,o){var r=!0,t=!0;if(typeof e!="function")throw new TypeError(ua);return Ve(o)&&(r="leading"in o?!!o.leading:r,t="trailing"in o?!!o.trailing:t),da(e,a,{leading:r,maxWait:a,trailing:t})}function fa(e){const{primaryColor:a,opacityDisabled:o,borderRadius:r,textColor3:t}=e;return Object.assign(Object.assign({},pr),{iconColor:t,textColor:"white",loadingColor:a,opacityDisabled:o,railColor:"rgba(0, 0, 0, .14)",railColorActive:a,buttonBoxShadow:"0 1px 4px 0 rgba(0, 0, 0, 0.3), inset 0 0 1px 0 rgba(0, 0, 0, 0.05)",buttonColor:"#FFF",railBorderRadiusSmall:r,railBorderRadiusMedium:r,railBorderRadiusLarge:r,buttonBorderRadiusSmall:r,buttonBorderRadiusMedium:r,buttonBorderRadiusLarge:r,boxShadowFocus:`0 0 0 2px ${vr(a,{alpha:.2})}`})}const ha={common:br,self:fa},kn=Sn("n-popconfirm"),zn={positiveText:String,negativeText:String,showIcon:{type:Boolean,default:!0},onPositiveClick:{type:Function,required:!0},onNegativeClick:{type:Function,required:!0}},Vt=yr(zn),ga=ge({name:"NPopconfirmPanel",props:zn,setup(e){const{localeRef:a}=Dt("Popconfirm"),{inlineThemeDisabled:o}=He(),{mergedClsPrefixRef:r,mergedThemeRef:t,props:n}=Et(kn),i=ne(()=>{const{common:{cubicBezierEaseInOut:l},self:{fontSize:c,iconSize:h,iconColor:x}}=t.value;return{"--n-bezier":l,"--n-font-size":c,"--n-icon-size":h,"--n-icon-color":x}}),s=o?kt("popconfirm-panel",void 0,i,n):void 0;return Object.assign(Object.assign({},Dt("Popconfirm")),{mergedClsPrefix:r,cssVars:o?void 0:i,localizedPositiveText:ne(()=>e.positiveText||a.value.positiveText),localizedNegativeText:ne(()=>e.negativeText||a.value.negativeText),positiveButtonProps:Z(n,"positiveButtonProps"),negativeButtonProps:Z(n,"negativeButtonProps"),handlePositiveClick(l){e.onPositiveClick(l)},handleNegativeClick(l){e.onNegativeClick(l)},themeClass:s==null?void 0:s.themeClass,onRender:s==null?void 0:s.onRender})},render(){var e;const{mergedClsPrefix:a,showIcon:o,$slots:r}=this,t=$t(r.action,()=>this.negativeText===null&&this.positiveText===null?[]:[this.negativeText!==null&&T(Y,Object.assign({size:"small",onClick:this.handleNegativeClick},this.negativeButtonProps),{default:()=>this.localizedNegativeText}),this.positiveText!==null&&T(Y,Object.assign({size:"small",type:"primary",onClick:this.handlePositiveClick},this.positiveButtonProps),{default:()=>this.localizedPositiveText})]);return(e=this.onRender)===null||e===void 0||e.call(this),T("div",{class:[`${a}-popconfirm__panel`,this.themeClass],style:this.cssVars},ue(r.default,n=>o||n?T("div",{class:`${a}-popconfirm__body`},o?T("div",{class:`${a}-popconfirm__icon`},$t(r.icon,()=>[T(Pn,{clsPrefix:a},{default:()=>T(mr,null)})])):null,n):null),t?T("div",{class:[`${a}-popconfirm__action`]},t):null)}}),ba=R("popconfirm",[M("body",`
 font-size: var(--n-font-size);
 display: flex;
 align-items: center;
 flex-wrap: nowrap;
 position: relative;
 `,[M("icon",`
 display: flex;
 font-size: var(--n-icon-size);
 color: var(--n-icon-color);
 transition: color .3s var(--n-bezier);
 margin: 0 8px 0 0;
 `)]),M("action",`
 display: flex;
 justify-content: flex-end;
 `,[q("&:not(:first-child)","margin-top: 8px"),R("button",[q("&:not(:last-child)","margin-right: 8px;")])])]),pa=Object.assign(Object.assign(Object.assign({},ke.props),jr),{positiveText:String,negativeText:String,showIcon:{type:Boolean,default:!0},trigger:{type:String,default:"click"},positiveButtonProps:Object,negativeButtonProps:Object,onPositiveClick:Function,onNegativeClick:Function}),va=ge({name:"Popconfirm",props:pa,slots:Object,__popover__:!0,setup(e){const{mergedClsPrefixRef:a}=He(),o=ke("Popconfirm","-popconfirm",ba,Cr,e,a),r=O(null);function t(s){var l;if(!(!((l=r.value)===null||l===void 0)&&l.getMergedShow()))return;const{onPositiveClick:c,"onUpdate:show":h}=e;Promise.resolve(c?c(s):!0).then(x=>{var g;x!==!1&&((g=r.value)===null||g===void 0||g.setShow(!1),h&&he(h,!1))})}function n(s){var l;if(!(!((l=r.value)===null||l===void 0)&&l.getMergedShow()))return;const{onNegativeClick:c,"onUpdate:show":h}=e;Promise.resolve(c?c(s):!0).then(x=>{var g;x!==!1&&((g=r.value)===null||g===void 0||g.setShow(!1),h&&he(h,!1))})}return Tn(kn,{mergedThemeRef:o,mergedClsPrefixRef:a,props:e}),{setShow(s){var l;(l=r.value)===null||l===void 0||l.setShow(s)},syncPosition(){var s;(s=r.value)===null||s===void 0||s.syncPosition()},mergedTheme:o,popoverInstRef:r,handlePositiveClick:t,handleNegativeClick:n}},render(){const{$slots:e,$props:a,mergedTheme:o}=this;return T(Dr,Object.assign({},Bn(a,Vt),{theme:o.peers.Popover,themeOverrides:o.peerOverrides.Popover,internalExtraClass:["popconfirm"],ref:"popoverInstRef"}),{trigger:e.trigger,default:()=>{const r=wr(a,Vt);return T(ga,Object.assign({},r,{onPositiveClick:this.handlePositiveClick,onNegativeClick:this.handleNegativeClick}),e)}})}}),ma=R("switch",`
 height: var(--n-height);
 min-width: var(--n-width);
 vertical-align: middle;
 user-select: none;
 -webkit-user-select: none;
 display: inline-flex;
 outline: none;
 justify-content: center;
 align-items: center;
`,[M("children-placeholder",`
 height: var(--n-rail-height);
 display: flex;
 flex-direction: column;
 overflow: hidden;
 pointer-events: none;
 visibility: hidden;
 `),M("rail-placeholder",`
 display: flex;
 flex-wrap: none;
 `),M("button-placeholder",`
 width: calc(1.75 * var(--n-rail-height));
 height: var(--n-rail-height);
 `),R("base-loading",`
 position: absolute;
 top: 50%;
 left: 50%;
 transform: translateX(-50%) translateY(-50%);
 font-size: calc(var(--n-button-width) - 4px);
 color: var(--n-loading-color);
 transition: color .3s var(--n-bezier);
 `,[_t({left:"50%",top:"50%",originalTransform:"translateX(-50%) translateY(-50%)"})]),M("checked, unchecked",`
 transition: color .3s var(--n-bezier);
 color: var(--n-text-color);
 box-sizing: border-box;
 position: absolute;
 white-space: nowrap;
 top: 0;
 bottom: 0;
 display: flex;
 align-items: center;
 line-height: 1;
 `),M("checked",`
 right: 0;
 padding-right: calc(1.25 * var(--n-rail-height) - var(--n-offset));
 `),M("unchecked",`
 left: 0;
 justify-content: flex-end;
 padding-left: calc(1.25 * var(--n-rail-height) - var(--n-offset));
 `),q("&:focus",[M("rail",`
 box-shadow: var(--n-box-shadow-focus);
 `)]),$("round",[M("rail","border-radius: calc(var(--n-rail-height) / 2);",[M("button","border-radius: calc(var(--n-button-height) / 2);")])]),St("disabled",[St("icon",[$("rubber-band",[$("pressed",[M("rail",[M("button","max-width: var(--n-button-width-pressed);")])]),M("rail",[q("&:active",[M("button","max-width: var(--n-button-width-pressed);")])]),$("active",[$("pressed",[M("rail",[M("button","left: calc(100% - var(--n-offset) - var(--n-button-width-pressed));")])]),M("rail",[q("&:active",[M("button","left: calc(100% - var(--n-offset) - var(--n-button-width-pressed));")])])])])])]),$("active",[M("rail",[M("button","left: calc(100% - var(--n-button-width) - var(--n-offset))")])]),M("rail",`
 overflow: hidden;
 height: var(--n-rail-height);
 min-width: var(--n-rail-width);
 border-radius: var(--n-rail-border-radius);
 cursor: pointer;
 position: relative;
 transition:
 opacity .3s var(--n-bezier),
 background .3s var(--n-bezier),
 box-shadow .3s var(--n-bezier);
 background-color: var(--n-rail-color);
 `,[M("button-icon",`
 color: var(--n-icon-color);
 transition: color .3s var(--n-bezier);
 font-size: calc(var(--n-button-height) - 4px);
 position: absolute;
 left: 0;
 right: 0;
 top: 0;
 bottom: 0;
 display: flex;
 justify-content: center;
 align-items: center;
 line-height: 1;
 `,[_t()]),M("button",`
 align-items: center; 
 top: var(--n-offset);
 left: var(--n-offset);
 height: var(--n-button-height);
 width: var(--n-button-width-pressed);
 max-width: var(--n-button-width);
 border-radius: var(--n-button-border-radius);
 background-color: var(--n-button-color);
 box-shadow: var(--n-button-box-shadow);
 box-sizing: border-box;
 cursor: inherit;
 content: "";
 position: absolute;
 transition:
 background-color .3s var(--n-bezier),
 left .3s var(--n-bezier),
 opacity .3s var(--n-bezier),
 max-width .3s var(--n-bezier),
 box-shadow .3s var(--n-bezier);
 `)]),$("active",[M("rail","background-color: var(--n-rail-color-active);")]),$("loading",[M("rail",`
 cursor: wait;
 `)]),$("disabled",[M("rail",`
 cursor: not-allowed;
 opacity: .5;
 `)])]),ya=Object.assign(Object.assign({},ke.props),{size:String,value:{type:[String,Number,Boolean],default:void 0},loading:Boolean,defaultValue:{type:[String,Number,Boolean],default:!1},disabled:{type:Boolean,default:void 0},round:{type:Boolean,default:!0},"onUpdate:value":[Function,Array],onUpdateValue:[Function,Array],checkedValue:{type:[String,Number,Boolean],default:!0},uncheckedValue:{type:[String,Number,Boolean],default:!1},railStyle:Function,rubberBand:{type:Boolean,default:!0},spinProps:Object,onChange:[Function,Array]});let Ie;const wa=ge({name:"Switch",props:ya,slots:Object,setup(e){Ie===void 0&&(typeof CSS<"u"?typeof CSS.supports<"u"?Ie=CSS.supports("width","max(1px)"):Ie=!1:Ie=!0);const{mergedClsPrefixRef:a,inlineThemeDisabled:o,mergedComponentPropsRef:r}=He(e),t=ke("Switch","-switch",ma,ha,e,a),n=xr(e,{mergedSize(p){var N,k;if(e.size!==void 0)return e.size;if(p)return p.mergedSize.value;const z=(k=(N=r==null?void 0:r.value)===null||N===void 0?void 0:N.Switch)===null||k===void 0?void 0:k.size;return z||"medium"}}),{mergedSizeRef:i,mergedDisabledRef:s}=n,l=O(e.defaultValue),c=Z(e,"value"),h=Rn(c,l),x=ne(()=>h.value===e.checkedValue),g=O(!1),b=O(!1),S=ne(()=>{const{railStyle:p}=e;if(p)return p({focused:b.value,checked:x.value})});function P(p){const{"onUpdate:value":N,onChange:k,onUpdateValue:z}=e,{nTriggerFormInput:A,nTriggerFormChange:L}=n;N&&he(N,p),z&&he(z,p),k&&he(k,p),l.value=p,A(),L()}function F(){const{nTriggerFormFocus:p}=n;p()}function B(){const{nTriggerFormBlur:p}=n;p()}function y(){e.loading||s.value||(h.value!==e.checkedValue?P(e.checkedValue):P(e.uncheckedValue))}function f(){b.value=!0,F()}function w(){b.value=!1,B(),g.value=!1}function E(p){e.loading||s.value||p.key===" "&&(h.value!==e.checkedValue?P(e.checkedValue):P(e.uncheckedValue),g.value=!1)}function d(p){e.loading||s.value||p.key===" "&&(p.preventDefault(),g.value=!0)}const u=ne(()=>{const{value:p}=i,{self:{opacityDisabled:N,railColor:k,railColorActive:z,buttonBoxShadow:A,buttonColor:L,boxShadowFocus:G,loadingColor:re,textColor:ee,iconColor:Ce,[K("buttonHeight",p)]:te,[K("buttonWidth",p)]:pe,[K("buttonWidthPressed",p)]:ce,[K("railHeight",p)]:Q,[K("railWidth",p)]:ae,[K("railBorderRadius",p)]:ze,[K("buttonBorderRadius",p)]:ve},common:{cubicBezierEaseInOut:me}}=t.value;let V,H,X;return Ie?(V=`calc((${Q} - ${te}) / 2)`,H=`max(${Q}, ${te})`,X=`max(${ae}, calc(${ae} + ${te} - ${Q}))`):(V=Ye((de(Q)-de(te))/2),H=Ye(Math.max(de(Q),de(te))),X=de(Q)>de(te)?ae:Ye(de(ae)+de(te)-de(Q))),{"--n-bezier":me,"--n-button-border-radius":ve,"--n-button-box-shadow":A,"--n-button-color":L,"--n-button-width":pe,"--n-button-width-pressed":ce,"--n-button-height":te,"--n-height":H,"--n-offset":V,"--n-opacity-disabled":N,"--n-rail-border-radius":ze,"--n-rail-color":k,"--n-rail-color-active":z,"--n-rail-height":Q,"--n-rail-width":ae,"--n-width":X,"--n-box-shadow-focus":G,"--n-loading-color":re,"--n-text-color":ee,"--n-icon-color":Ce}}),C=o?kt("switch",ne(()=>i.value[0]),u,e):void 0;return{handleClick:y,handleBlur:w,handleFocus:f,handleKeyup:E,handleKeydown:d,mergedRailStyle:S,pressed:g,mergedClsPrefix:a,mergedValue:h,checked:x,mergedDisabled:s,cssVars:o?void 0:u,themeClass:C==null?void 0:C.themeClass,onRender:C==null?void 0:C.onRender}},render(){const{mergedClsPrefix:e,mergedDisabled:a,checked:o,mergedRailStyle:r,onRender:t,$slots:n}=this;t==null||t();const{checked:i,unchecked:s,icon:l,"checked-icon":c,"unchecked-icon":h}=n,x=!(Je(l)&&Je(c)&&Je(h));return T("div",{role:"switch","aria-checked":o,class:[`${e}-switch`,this.themeClass,x&&`${e}-switch--icon`,o&&`${e}-switch--active`,a&&`${e}-switch--disabled`,this.round&&`${e}-switch--round`,this.loading&&`${e}-switch--loading`,this.pressed&&`${e}-switch--pressed`,this.rubberBand&&`${e}-switch--rubber-band`],tabindex:this.mergedDisabled?void 0:0,style:this.cssVars,onClick:this.handleClick,onFocus:this.handleFocus,onBlur:this.handleBlur,onKeyup:this.handleKeyup,onKeydown:this.handleKeydown},T("div",{class:`${e}-switch__rail`,"aria-hidden":"true",style:r},ue(i,g=>ue(s,b=>g||b?T("div",{"aria-hidden":!0,class:`${e}-switch__children-placeholder`},T("div",{class:`${e}-switch__rail-placeholder`},T("div",{class:`${e}-switch__button-placeholder`}),g),T("div",{class:`${e}-switch__rail-placeholder`},T("div",{class:`${e}-switch__button-placeholder`}),b)):null)),T("div",{class:`${e}-switch__button`},ue(l,g=>ue(c,b=>ue(h,S=>T(Rr,null,{default:()=>this.loading?T(Sr,Object.assign({key:"loading",clsPrefix:e,strokeWidth:20},this.spinProps)):this.checked&&(b||g)?T("div",{class:`${e}-switch__button-icon`,key:b?"checked-icon":"icon"},b||g):!this.checked&&(S||g)?T("div",{class:`${e}-switch__button-icon`,key:S?"unchecked-icon":"icon"},S||g):null})))),ue(i,g=>g&&T("div",{key:"checked",class:`${e}-switch__checked`},g)),ue(s,g=>g&&T("div",{key:"unchecked",class:`${e}-switch__unchecked`},g)))))}}),At=Sn("n-tabs"),An={tab:[String,Number,Object,Function],name:{type:[String,Number],required:!0},disabled:Boolean,displayDirective:{type:String,default:"if"},closable:{type:Boolean,default:void 0},tabProps:Object,label:[String,Number,Object,Function]},Fe=ge({__TAB_PANE__:!0,name:"TabPane",alias:["TabPanel"],props:An,slots:Object,setup(e){const a=Et(At,null);return a||Pr("tab-pane","`n-tab-pane` must be placed inside `n-tabs`."),{style:a.paneStyleRef,class:a.paneClassRef,mergedClsPrefix:a.mergedClsPrefixRef}},render(){return T("div",{class:[`${this.mergedClsPrefix}-tab-pane`,this.class],style:this.style},this.$slots)}}),Ca=Object.assign({internalLeftPadded:Boolean,internalAddable:Boolean,internalCreatedByPane:Boolean},Bn(An,["displayDirective"])),Tt=ge({__TAB__:!0,inheritAttrs:!1,name:"Tab",props:Ca,setup(e){const{mergedClsPrefixRef:a,valueRef:o,typeRef:r,closableRef:t,tabStyleRef:n,addTabStyleRef:i,tabClassRef:s,addTabClassRef:l,tabChangeIdRef:c,onBeforeLeaveRef:h,triggerRef:x,handleAdd:g,activateTab:b,handleClose:S}=Et(At);return{trigger:x,mergedClosable:ne(()=>{if(e.internalAddable)return!1;const{closable:P}=e;return P===void 0?t.value:P}),style:n,addStyle:i,tabClass:s,addTabClass:l,clsPrefix:a,value:o,type:r,handleClose(P){P.stopPropagation(),!e.disabled&&S(e.name)},activateTab(){if(e.disabled)return;if(e.internalAddable){g();return}const{name:P}=e,F=++c.id;if(P!==o.value){const{value:B}=h;B?Promise.resolve(B(e.name,o.value)).then(y=>{y&&c.id===F&&b(P)}):b(P)}}}},render(){const{internalAddable:e,clsPrefix:a,name:o,disabled:r,label:t,tab:n,value:i,mergedClosable:s,trigger:l,$slots:{default:c}}=this,h=t??n;return T("div",{class:`${a}-tabs-tab-wrapper`},this.internalLeftPadded?T("div",{class:`${a}-tabs-tab-pad`}):null,T("div",Object.assign({key:o,"data-name":o,"data-disabled":r?!0:void 0},Br({class:[`${a}-tabs-tab`,i===o&&`${a}-tabs-tab--active`,r&&`${a}-tabs-tab--disabled`,s&&`${a}-tabs-tab--closable`,e&&`${a}-tabs-tab--addable`,e?this.addTabClass:this.tabClass],onClick:l==="click"?this.activateTab:void 0,onMouseenter:l==="hover"?this.activateTab:void 0,style:e?this.addStyle:this.style},this.internalCreatedByPane?this.tabProps||{}:this.$attrs)),T("span",{class:`${a}-tabs-tab__label`},e?T(We,null,T("div",{class:`${a}-tabs-tab__height-placeholder`}," "),T(Pn,{clsPrefix:a},{default:()=>T(Ur,null)})):c?c():typeof h=="object"?h:Tr(h??o)),s&&this.type==="card"?T(Er,{clsPrefix:a,class:`${a}-tabs-tab__close`,onClick:this.handleClose,disabled:r}):null))}}),xa=R("tabs",`
 box-sizing: border-box;
 width: 100%;
 display: flex;
 flex-direction: column;
 transition:
 background-color .3s var(--n-bezier),
 border-color .3s var(--n-bezier);
`,[$("segment-type",[R("tabs-rail",[q("&.transition-disabled",[R("tabs-capsule",`
 transition: none;
 `)])])]),$("top",[R("tab-pane",`
 padding: var(--n-pane-padding-top) var(--n-pane-padding-right) var(--n-pane-padding-bottom) var(--n-pane-padding-left);
 `)]),$("left",[R("tab-pane",`
 padding: var(--n-pane-padding-right) var(--n-pane-padding-bottom) var(--n-pane-padding-left) var(--n-pane-padding-top);
 `)]),$("left, right",`
 flex-direction: row;
 `,[R("tabs-bar",`
 width: 2px;
 right: 0;
 transition:
 top .2s var(--n-bezier),
 max-height .2s var(--n-bezier),
 background-color .3s var(--n-bezier);
 `),R("tabs-tab",`
 padding: var(--n-tab-padding-vertical); 
 `)]),$("right",`
 flex-direction: row-reverse;
 `,[R("tab-pane",`
 padding: var(--n-pane-padding-left) var(--n-pane-padding-top) var(--n-pane-padding-right) var(--n-pane-padding-bottom);
 `),R("tabs-bar",`
 left: 0;
 `)]),$("bottom",`
 flex-direction: column-reverse;
 justify-content: flex-end;
 `,[R("tab-pane",`
 padding: var(--n-pane-padding-bottom) var(--n-pane-padding-right) var(--n-pane-padding-top) var(--n-pane-padding-left);
 `),R("tabs-bar",`
 top: 0;
 `)]),R("tabs-rail",`
 position: relative;
 padding: 3px;
 border-radius: var(--n-tab-border-radius);
 width: 100%;
 background-color: var(--n-color-segment);
 transition: background-color .3s var(--n-bezier);
 display: flex;
 align-items: center;
 `,[R("tabs-capsule",`
 border-radius: var(--n-tab-border-radius);
 position: absolute;
 pointer-events: none;
 background-color: var(--n-tab-color-segment);
 box-shadow: 0 1px 3px 0 rgba(0, 0, 0, .08);
 transition: transform 0.3s var(--n-bezier);
 `),R("tabs-tab-wrapper",`
 flex-basis: 0;
 flex-grow: 1;
 display: flex;
 align-items: center;
 justify-content: center;
 `,[R("tabs-tab",`
 overflow: hidden;
 border-radius: var(--n-tab-border-radius);
 width: 100%;
 display: flex;
 align-items: center;
 justify-content: center;
 `,[$("active",`
 font-weight: var(--n-font-weight-strong);
 color: var(--n-tab-text-color-active);
 `),q("&:hover",`
 color: var(--n-tab-text-color-hover);
 `)])])]),$("flex",[R("tabs-nav",`
 width: 100%;
 position: relative;
 `,[R("tabs-wrapper",`
 width: 100%;
 `,[R("tabs-tab",`
 margin-right: 0;
 `)])])]),R("tabs-nav",`
 box-sizing: border-box;
 line-height: 1.5;
 display: flex;
 transition: border-color .3s var(--n-bezier);
 `,[M("prefix, suffix",`
 display: flex;
 align-items: center;
 `),M("prefix","padding-right: 16px;"),M("suffix","padding-left: 16px;")]),$("top, bottom",[q(">",[R("tabs-nav",[R("tabs-nav-scroll-wrapper",[q("&::before",`
 top: 0;
 bottom: 0;
 left: 0;
 width: 20px;
 `),q("&::after",`
 top: 0;
 bottom: 0;
 right: 0;
 width: 20px;
 `),$("shadow-start",[q("&::before",`
 box-shadow: inset 10px 0 8px -8px rgba(0, 0, 0, .12);
 `)]),$("shadow-end",[q("&::after",`
 box-shadow: inset -10px 0 8px -8px rgba(0, 0, 0, .12);
 `)])])])])]),$("left, right",[R("tabs-nav-scroll-content",`
 flex-direction: column;
 `),q(">",[R("tabs-nav",[R("tabs-nav-scroll-wrapper",[q("&::before",`
 top: 0;
 left: 0;
 right: 0;
 height: 20px;
 `),q("&::after",`
 bottom: 0;
 left: 0;
 right: 0;
 height: 20px;
 `),$("shadow-start",[q("&::before",`
 box-shadow: inset 0 10px 8px -8px rgba(0, 0, 0, .12);
 `)]),$("shadow-end",[q("&::after",`
 box-shadow: inset 0 -10px 8px -8px rgba(0, 0, 0, .12);
 `)])])])])]),R("tabs-nav-scroll-wrapper",`
 flex: 1;
 position: relative;
 overflow: hidden;
 `,[R("tabs-nav-y-scroll",`
 height: 100%;
 width: 100%;
 overflow-y: auto; 
 scrollbar-width: none;
 `,[q("&::-webkit-scrollbar, &::-webkit-scrollbar-track-piece, &::-webkit-scrollbar-thumb",`
 width: 0;
 height: 0;
 display: none;
 `)]),q("&::before, &::after",`
 transition: box-shadow .3s var(--n-bezier);
 pointer-events: none;
 content: "";
 position: absolute;
 z-index: 1;
 `)]),R("tabs-nav-scroll-content",`
 display: flex;
 position: relative;
 min-width: 100%;
 min-height: 100%;
 width: fit-content;
 box-sizing: border-box;
 `),R("tabs-wrapper",`
 display: inline-flex;
 flex-wrap: nowrap;
 position: relative;
 `),R("tabs-tab-wrapper",`
 display: flex;
 flex-wrap: nowrap;
 flex-shrink: 0;
 flex-grow: 0;
 `),R("tabs-tab",`
 cursor: pointer;
 white-space: nowrap;
 flex-wrap: nowrap;
 display: inline-flex;
 align-items: center;
 color: var(--n-tab-text-color);
 font-size: var(--n-tab-font-size);
 background-clip: padding-box;
 padding: var(--n-tab-padding);
 transition:
 box-shadow .3s var(--n-bezier),
 color .3s var(--n-bezier),
 background-color .3s var(--n-bezier),
 border-color .3s var(--n-bezier);
 `,[$("disabled",{cursor:"not-allowed"}),M("close",`
 margin-left: 6px;
 transition:
 background-color .3s var(--n-bezier),
 color .3s var(--n-bezier);
 `),M("label",`
 display: flex;
 align-items: center;
 z-index: 1;
 `)]),R("tabs-bar",`
 position: absolute;
 bottom: 0;
 height: 2px;
 border-radius: 1px;
 background-color: var(--n-bar-color);
 transition:
 left .2s var(--n-bezier),
 max-width .2s var(--n-bezier),
 opacity .3s var(--n-bezier),
 background-color .3s var(--n-bezier);
 `,[q("&.transition-disabled",`
 transition: none;
 `),$("disabled",`
 background-color: var(--n-tab-text-color-disabled)
 `)]),R("tabs-pane-wrapper",`
 position: relative;
 overflow: hidden;
 transition: max-height .2s var(--n-bezier);
 `),R("tab-pane",`
 color: var(--n-pane-text-color);
 width: 100%;
 transition:
 color .3s var(--n-bezier),
 background-color .3s var(--n-bezier),
 opacity .2s var(--n-bezier);
 left: 0;
 right: 0;
 top: 0;
 `,[q("&.next-transition-leave-active, &.prev-transition-leave-active, &.next-transition-enter-active, &.prev-transition-enter-active",`
 transition:
 color .3s var(--n-bezier),
 background-color .3s var(--n-bezier),
 transform .2s var(--n-bezier),
 opacity .2s var(--n-bezier);
 `),q("&.next-transition-leave-active, &.prev-transition-leave-active",`
 position: absolute;
 `),q("&.next-transition-enter-from, &.prev-transition-leave-to",`
 transform: translateX(32px);
 opacity: 0;
 `),q("&.next-transition-leave-to, &.prev-transition-enter-from",`
 transform: translateX(-32px);
 opacity: 0;
 `),q("&.next-transition-leave-from, &.next-transition-enter-to, &.prev-transition-leave-from, &.prev-transition-enter-to",`
 transform: translateX(0);
 opacity: 1;
 `)]),R("tabs-tab-pad",`
 box-sizing: border-box;
 width: var(--n-tab-gap);
 flex-grow: 0;
 flex-shrink: 0;
 `),$("line-type, bar-type",[R("tabs-tab",`
 font-weight: var(--n-tab-font-weight);
 box-sizing: border-box;
 vertical-align: bottom;
 `,[q("&:hover",{color:"var(--n-tab-text-color-hover)"}),$("active",`
 color: var(--n-tab-text-color-active);
 font-weight: var(--n-tab-font-weight-active);
 `),$("disabled",{color:"var(--n-tab-text-color-disabled)"})])]),R("tabs-nav",[$("line-type",[$("top",[M("prefix, suffix",`
 border-bottom: 1px solid var(--n-tab-border-color);
 `),R("tabs-nav-scroll-content",`
 border-bottom: 1px solid var(--n-tab-border-color);
 `),R("tabs-bar",`
 bottom: -1px;
 `)]),$("left",[M("prefix, suffix",`
 border-right: 1px solid var(--n-tab-border-color);
 `),R("tabs-nav-scroll-content",`
 border-right: 1px solid var(--n-tab-border-color);
 `),R("tabs-bar",`
 right: -1px;
 `)]),$("right",[M("prefix, suffix",`
 border-left: 1px solid var(--n-tab-border-color);
 `),R("tabs-nav-scroll-content",`
 border-left: 1px solid var(--n-tab-border-color);
 `),R("tabs-bar",`
 left: -1px;
 `)]),$("bottom",[M("prefix, suffix",`
 border-top: 1px solid var(--n-tab-border-color);
 `),R("tabs-nav-scroll-content",`
 border-top: 1px solid var(--n-tab-border-color);
 `),R("tabs-bar",`
 top: -1px;
 `)]),M("prefix, suffix",`
 transition: border-color .3s var(--n-bezier);
 `),R("tabs-nav-scroll-content",`
 transition: border-color .3s var(--n-bezier);
 `),R("tabs-bar",`
 border-radius: 0;
 `)]),$("card-type",[M("prefix, suffix",`
 transition: border-color .3s var(--n-bezier);
 `),R("tabs-pad",`
 flex-grow: 1;
 transition: border-color .3s var(--n-bezier);
 `),R("tabs-tab-pad",`
 transition: border-color .3s var(--n-bezier);
 `),R("tabs-tab",`
 font-weight: var(--n-tab-font-weight);
 border: 1px solid var(--n-tab-border-color);
 background-color: var(--n-tab-color);
 box-sizing: border-box;
 position: relative;
 vertical-align: bottom;
 display: flex;
 justify-content: space-between;
 font-size: var(--n-tab-font-size);
 color: var(--n-tab-text-color);
 `,[$("addable",`
 padding-left: 8px;
 padding-right: 8px;
 font-size: 16px;
 justify-content: center;
 `,[M("height-placeholder",`
 width: 0;
 font-size: var(--n-tab-font-size);
 `),St("disabled",[q("&:hover",`
 color: var(--n-tab-text-color-hover);
 `)])]),$("closable","padding-right: 8px;"),$("active",`
 background-color: #0000;
 font-weight: var(--n-tab-font-weight-active);
 color: var(--n-tab-text-color-active);
 `),$("disabled","color: var(--n-tab-text-color-disabled);")])]),$("left, right",`
 flex-direction: column; 
 `,[M("prefix, suffix",`
 padding: var(--n-tab-padding-vertical);
 `),R("tabs-wrapper",`
 flex-direction: column;
 `),R("tabs-tab-wrapper",`
 flex-direction: column;
 `,[R("tabs-tab-pad",`
 height: var(--n-tab-gap-vertical);
 width: 100%;
 `)])]),$("top",[$("card-type",[R("tabs-scroll-padding","border-bottom: 1px solid var(--n-tab-border-color);"),M("prefix, suffix",`
 border-bottom: 1px solid var(--n-tab-border-color);
 `),R("tabs-tab",`
 border-top-left-radius: var(--n-tab-border-radius);
 border-top-right-radius: var(--n-tab-border-radius);
 `,[$("active",`
 border-bottom: 1px solid #0000;
 `)]),R("tabs-tab-pad",`
 border-bottom: 1px solid var(--n-tab-border-color);
 `),R("tabs-pad",`
 border-bottom: 1px solid var(--n-tab-border-color);
 `)])]),$("left",[$("card-type",[R("tabs-scroll-padding","border-right: 1px solid var(--n-tab-border-color);"),M("prefix, suffix",`
 border-right: 1px solid var(--n-tab-border-color);
 `),R("tabs-tab",`
 border-top-left-radius: var(--n-tab-border-radius);
 border-bottom-left-radius: var(--n-tab-border-radius);
 `,[$("active",`
 border-right: 1px solid #0000;
 `)]),R("tabs-tab-pad",`
 border-right: 1px solid var(--n-tab-border-color);
 `),R("tabs-pad",`
 border-right: 1px solid var(--n-tab-border-color);
 `)])]),$("right",[$("card-type",[R("tabs-scroll-padding","border-left: 1px solid var(--n-tab-border-color);"),M("prefix, suffix",`
 border-left: 1px solid var(--n-tab-border-color);
 `),R("tabs-tab",`
 border-top-right-radius: var(--n-tab-border-radius);
 border-bottom-right-radius: var(--n-tab-border-radius);
 `,[$("active",`
 border-left: 1px solid #0000;
 `)]),R("tabs-tab-pad",`
 border-left: 1px solid var(--n-tab-border-color);
 `),R("tabs-pad",`
 border-left: 1px solid var(--n-tab-border-color);
 `)])]),$("bottom",[$("card-type",[R("tabs-scroll-padding","border-top: 1px solid var(--n-tab-border-color);"),M("prefix, suffix",`
 border-top: 1px solid var(--n-tab-border-color);
 `),R("tabs-tab",`
 border-bottom-left-radius: var(--n-tab-border-radius);
 border-bottom-right-radius: var(--n-tab-border-radius);
 `,[$("active",`
 border-top: 1px solid #0000;
 `)]),R("tabs-tab-pad",`
 border-top: 1px solid var(--n-tab-border-color);
 `),R("tabs-pad",`
 border-top: 1px solid var(--n-tab-border-color);
 `)])])])]),Ze=ca,Ra=Object.assign(Object.assign({},ke.props),{value:[String,Number],defaultValue:[String,Number],trigger:{type:String,default:"click"},type:{type:String,default:"bar"},closable:Boolean,justifyContent:String,size:String,placement:{type:String,default:"top"},tabStyle:[String,Object],tabClass:String,addTabStyle:[String,Object],addTabClass:String,barWidth:Number,paneClass:String,paneStyle:[String,Object],paneWrapperClass:String,paneWrapperStyle:[String,Object],addable:[Boolean,Object],tabsPadding:{type:Number,default:0},animated:Boolean,onBeforeLeave:Function,onAdd:Function,"onUpdate:value":[Function,Array],onUpdateValue:[Function,Array],onClose:[Function,Array],labelSize:String,activeName:[String,Number],onActiveNameChange:[Function,Array]}),Sa=ge({name:"Tabs",props:Ra,slots:Object,setup(e,{slots:a}){var o,r,t,n;const{mergedClsPrefixRef:i,inlineThemeDisabled:s,mergedComponentPropsRef:l}=He(e),c=ke("Tabs","-tabs",xa,zr,e,i),h=O(null),x=O(null),g=O(null),b=O(null),S=O(null),P=O(null),F=O(!0),B=O(!0),y=Ft(e,["labelSize","size"]),f=ne(()=>{var v,m;if(y.value)return y.value;const I=(m=(v=l==null?void 0:l.value)===null||v===void 0?void 0:v.Tabs)===null||m===void 0?void 0:m.size;return I||"medium"}),w=Ft(e,["activeName","value"]),E=O((r=(o=w.value)!==null&&o!==void 0?o:e.defaultValue)!==null&&r!==void 0?r:a.default?(n=(t=Ge(a.default())[0])===null||t===void 0?void 0:t.props)===null||n===void 0?void 0:n.name:null),d=Rn(w,E),u={id:0},C=ne(()=>{if(!(!e.justifyContent||e.type==="card"))return{display:"flex",justifyContent:e.justifyContent}});je(d,()=>{u.id=0,A(),L()});function p(){var v;const{value:m}=d;return m===null?null:(v=h.value)===null||v===void 0?void 0:v.querySelector(`[data-name="${m}"]`)}function N(v){if(e.type==="card")return;const{value:m}=x;if(!m)return;const I=m.style.opacity==="0";if(v){const j=`${i.value}-tabs-bar--disabled`,{barWidth:W,placement:oe}=e;if(v.dataset.disabled==="true"?m.classList.add(j):m.classList.remove(j),["top","bottom"].includes(oe)){if(z(["top","maxHeight","height"]),typeof W=="number"&&v.offsetWidth>=W){const ie=Math.floor((v.offsetWidth-W)/2)+v.offsetLeft;m.style.left=`${ie}px`,m.style.maxWidth=`${W}px`}else m.style.left=`${v.offsetLeft}px`,m.style.maxWidth=`${v.offsetWidth}px`;m.style.width="8192px",I&&(m.style.transition="none"),m.offsetWidth,I&&(m.style.transition="",m.style.opacity="1")}else{if(z(["left","maxWidth","width"]),typeof W=="number"&&v.offsetHeight>=W){const ie=Math.floor((v.offsetHeight-W)/2)+v.offsetTop;m.style.top=`${ie}px`,m.style.maxHeight=`${W}px`}else m.style.top=`${v.offsetTop}px`,m.style.maxHeight=`${v.offsetHeight}px`;m.style.height="8192px",I&&(m.style.transition="none"),m.offsetHeight,I&&(m.style.transition="",m.style.opacity="1")}}}function k(){if(e.type==="card")return;const{value:v}=x;v&&(v.style.opacity="0")}function z(v){const{value:m}=x;if(m)for(const I of v)m.style[I]=""}function A(){if(e.type==="card")return;const v=p();v?N(v):k()}function L(){var v;const m=(v=S.value)===null||v===void 0?void 0:v.$el;if(!m)return;const I=p();if(!I)return;const{scrollLeft:j,offsetWidth:W}=m,{offsetLeft:oe,offsetWidth:ie}=I;j>oe?m.scrollTo({top:0,left:oe,behavior:"smooth"}):oe+ie>j+W&&m.scrollTo({top:0,left:oe+ie-W,behavior:"smooth"})}const G=O(null);let re=0,ee=null;function Ce(v){const m=G.value;if(m){re=v.getBoundingClientRect().height;const I=`${re}px`,j=()=>{m.style.height=I,m.style.maxHeight=I};ee?(j(),ee(),ee=null):ee=j}}function te(v){const m=G.value;if(m){const I=v.getBoundingClientRect().height,j=()=>{document.body.offsetHeight,m.style.maxHeight=`${I}px`,m.style.height=`${Math.max(re,I)}px`};ee?(ee(),ee=null,j()):ee=j}}function pe(){const v=G.value;if(v){v.style.maxHeight="",v.style.height="";const{paneWrapperStyle:m}=e;if(typeof m=="string")v.style.cssText=m;else if(m){const{maxHeight:I,height:j}=m;I!==void 0&&(v.style.maxHeight=I),j!==void 0&&(v.style.height=j)}}}const ce={value:[]},Q=O("next");function ae(v){const m=d.value;let I="next";for(const j of ce.value){if(j===m)break;if(j===v){I="prev";break}}Q.value=I,ze(v)}function ze(v){const{onActiveNameChange:m,onUpdateValue:I,"onUpdate:value":j}=e;m&&he(m,v),I&&he(I,v),j&&he(j,v),E.value=v}function ve(v){const{onClose:m}=e;m&&he(m,v)}function me(){const{value:v}=x;if(!v)return;const m="transition-disabled";v.classList.add(m),A(),v.classList.remove(m)}const V=O(null);function H({transitionDisabled:v}){const m=h.value;if(!m)return;v&&m.classList.add("transition-disabled");const I=p();I&&V.value&&(V.value.style.width=`${I.offsetWidth}px`,V.value.style.height=`${I.offsetHeight}px`,V.value.style.transform=`translateX(${I.offsetLeft-de(getComputedStyle(m).paddingLeft)}px)`,v&&V.value.offsetWidth),v&&m.classList.remove("transition-disabled")}je([d],()=>{e.type==="segment"&&Oe(()=>{H({transitionDisabled:!1})})}),zt(()=>{e.type==="segment"&&H({transitionDisabled:!0})});let X=0;function Ae(v){var m;if(v.contentRect.width===0&&v.contentRect.height===0||X===v.contentRect.width)return;X=v.contentRect.width;const{type:I}=e;if((I==="line"||I==="bar")&&me(),I!=="segment"){const{placement:j}=e;Ke((j==="top"||j==="bottom"?(m=S.value)===null||m===void 0?void 0:m.$el:P.value)||null)}}const Ln=Ze(Ae,64);je([()=>e.justifyContent,()=>e.size],()=>{Oe(()=>{const{type:v}=e;(v==="line"||v==="bar")&&me()})});const xe=O(!1);function Mn(v){var m;const{target:I,contentRect:{width:j,height:W}}=v,oe=I.parentElement.parentElement.offsetWidth,ie=I.parentElement.parentElement.offsetHeight,{placement:Se}=e;if(!xe.value)Se==="top"||Se==="bottom"?oe<j&&(xe.value=!0):ie<W&&(xe.value=!0);else{const{value:Ne}=b;if(!Ne)return;Se==="top"||Se==="bottom"?oe-j>Ne.$el.offsetWidth&&(xe.value=!1):ie-W>Ne.$el.offsetHeight&&(xe.value=!1)}Ke(((m=S.value)===null||m===void 0?void 0:m.$el)||null)}const Un=Ze(Mn,64);function Fn(){const{onAdd:v}=e;v&&v(),Oe(()=>{const m=p(),{value:I}=S;!m||!I||I.scrollTo({left:m.offsetLeft,top:0,behavior:"smooth"})})}function Ke(v){if(!v)return;const{placement:m}=e;if(m==="top"||m==="bottom"){const{scrollLeft:I,scrollWidth:j,offsetWidth:W}=v;F.value=I<=0,B.value=I+W>=j}else{const{scrollTop:I,scrollHeight:j,offsetHeight:W}=v;F.value=I<=0,B.value=I+W>=j}}const Dn=Ze(v=>{Ke(v.target)},64);Tn(At,{triggerRef:Z(e,"trigger"),tabStyleRef:Z(e,"tabStyle"),tabClassRef:Z(e,"tabClass"),addTabStyleRef:Z(e,"addTabStyle"),addTabClassRef:Z(e,"addTabClass"),paneClassRef:Z(e,"paneClass"),paneStyleRef:Z(e,"paneStyle"),mergedClsPrefixRef:i,typeRef:Z(e,"type"),closableRef:Z(e,"closable"),valueRef:d,tabChangeIdRef:u,onBeforeLeaveRef:Z(e,"onBeforeLeave"),activateTab:ae,handleClose:ve,handleAdd:Fn}),Or(()=>{A(),L()}),kr(()=>{const{value:v}=g;if(!v)return;const{value:m}=i,I=`${m}-tabs-nav-scroll-wrapper--shadow-start`,j=`${m}-tabs-nav-scroll-wrapper--shadow-end`;F.value?v.classList.remove(I):v.classList.add(I),B.value?v.classList.remove(j):v.classList.add(j)});const jn={syncBarPosition:()=>{A()}},On=()=>{H({transitionDisabled:!0})},It=ne(()=>{const{value:v}=f,{type:m}=e,I={card:"Card",bar:"Bar",line:"Line",segment:"Segment"}[m],j=`${v}${I}`,{self:{barColor:W,closeIconColor:oe,closeIconColorHover:ie,closeIconColorPressed:Se,tabColor:Ne,tabBorderColor:qn,paneTextColor:Vn,tabFontWeight:Wn,tabBorderRadius:Hn,tabFontWeightActive:Kn,colorSegment:Jn,fontWeightStrong:Yn,tabColorSegment:Gn,closeSize:Qn,closeIconSize:Xn,closeColorHover:Zn,closeColorPressed:er,closeBorderRadius:tr,[K("panePadding",v)]:Le,[K("tabPadding",j)]:nr,[K("tabPaddingVertical",j)]:rr,[K("tabGap",j)]:ar,[K("tabGap",`${j}Vertical`)]:or,[K("tabTextColor",m)]:ir,[K("tabTextColorActive",m)]:sr,[K("tabTextColorHover",m)]:lr,[K("tabTextColorDisabled",m)]:dr,[K("tabFontSize",v)]:ur},common:{cubicBezierEaseInOut:cr}}=c.value;return{"--n-bezier":cr,"--n-color-segment":Jn,"--n-bar-color":W,"--n-tab-font-size":ur,"--n-tab-text-color":ir,"--n-tab-text-color-active":sr,"--n-tab-text-color-disabled":dr,"--n-tab-text-color-hover":lr,"--n-pane-text-color":Vn,"--n-tab-border-color":qn,"--n-tab-border-radius":Hn,"--n-close-size":Qn,"--n-close-icon-size":Xn,"--n-close-color-hover":Zn,"--n-close-color-pressed":er,"--n-close-border-radius":tr,"--n-close-icon-color":oe,"--n-close-icon-color-hover":ie,"--n-close-icon-color-pressed":Se,"--n-tab-color":Ne,"--n-tab-font-weight":Wn,"--n-tab-font-weight-active":Kn,"--n-tab-padding":nr,"--n-tab-padding-vertical":rr,"--n-tab-gap":ar,"--n-tab-gap-vertical":or,"--n-pane-padding-left":Me(Le,"left"),"--n-pane-padding-right":Me(Le,"right"),"--n-pane-padding-top":Me(Le,"top"),"--n-pane-padding-bottom":Me(Le,"bottom"),"--n-font-weight-strong":Yn,"--n-tab-color-segment":Gn}}),Re=s?kt("tabs",ne(()=>`${f.value[0]}${e.type[0]}`),It,e):void 0;return Object.assign({mergedClsPrefix:i,mergedValue:d,renderedNames:new Set,segmentCapsuleElRef:V,tabsPaneWrapperRef:G,tabsElRef:h,barElRef:x,addTabInstRef:b,xScrollInstRef:S,scrollWrapperElRef:g,addTabFixed:xe,tabWrapperStyle:C,handleNavResize:Ln,mergedSize:f,handleScroll:Dn,handleTabsResize:Un,cssVars:s?void 0:It,themeClass:Re==null?void 0:Re.themeClass,animationDirection:Q,renderNameListRef:ce,yScrollElRef:P,handleSegmentResize:On,onAnimationBeforeLeave:Ce,onAnimationEnter:te,onAnimationAfterEnter:pe,onRender:Re==null?void 0:Re.onRender},jn)},render(){const{mergedClsPrefix:e,type:a,placement:o,addTabFixed:r,addable:t,mergedSize:n,renderNameListRef:i,onRender:s,paneWrapperClass:l,paneWrapperStyle:c,$slots:{default:h,prefix:x,suffix:g}}=this;s==null||s();const b=h?Ge(h()).filter(E=>E.type.__TAB_PANE__===!0):[],S=h?Ge(h()).filter(E=>E.type.__TAB__===!0):[],P=!S.length,F=a==="card",B=a==="segment",y=!F&&!B&&this.justifyContent;i.value=[];const f=()=>{const E=T("div",{style:this.tabWrapperStyle,class:`${e}-tabs-wrapper`},y?null:T("div",{class:`${e}-tabs-scroll-padding`,style:o==="top"||o==="bottom"?{width:`${this.tabsPadding}px`}:{height:`${this.tabsPadding}px`}}),P?b.map((d,u)=>(i.value.push(d.props.name),et(T(Tt,Object.assign({},d.props,{internalCreatedByPane:!0,internalLeftPadded:u!==0&&(!y||y==="center"||y==="start"||y==="end")}),d.children?{default:d.children.tab}:void 0)))):S.map((d,u)=>(i.value.push(d.props.name),et(u!==0&&!y?Kt(d):d))),!r&&t&&F?Ht(t,(P?b.length:S.length)!==0):null,y?null:T("div",{class:`${e}-tabs-scroll-padding`,style:{width:`${this.tabsPadding}px`}}));return T("div",{ref:"tabsElRef",class:`${e}-tabs-nav-scroll-content`},F&&t?T(Qe,{onResize:this.handleTabsResize},{default:()=>E}):E,F?T("div",{class:`${e}-tabs-pad`}):null,F?null:T("div",{ref:"barElRef",class:`${e}-tabs-bar`}))},w=B?"top":o;return T("div",{class:[`${e}-tabs`,this.themeClass,`${e}-tabs--${a}-type`,`${e}-tabs--${n}-size`,y&&`${e}-tabs--flex`,`${e}-tabs--${w}`],style:this.cssVars},T("div",{class:[`${e}-tabs-nav--${a}-type`,`${e}-tabs-nav--${w}`,`${e}-tabs-nav`]},ue(x,E=>E&&T("div",{class:`${e}-tabs-nav__prefix`},E)),B?T(Qe,{onResize:this.handleSegmentResize},{default:()=>T("div",{class:`${e}-tabs-rail`,ref:"tabsElRef"},T("div",{class:`${e}-tabs-capsule`,ref:"segmentCapsuleElRef"},T("div",{class:`${e}-tabs-wrapper`},T("div",{class:`${e}-tabs-tab`}))),P?b.map((E,d)=>(i.value.push(E.props.name),T(Tt,Object.assign({},E.props,{internalCreatedByPane:!0,internalLeftPadded:d!==0}),E.children?{default:E.children.tab}:void 0))):S.map((E,d)=>(i.value.push(E.props.name),d===0?E:Kt(E))))}):T(Qe,{onResize:this.handleNavResize},{default:()=>T("div",{class:`${e}-tabs-nav-scroll-wrapper`,ref:"scrollWrapperElRef"},["top","bottom"].includes(w)?T(Qr,{ref:"xScrollInstRef",onScroll:this.handleScroll},{default:f}):T("div",{class:`${e}-tabs-nav-y-scroll`,onScroll:this.handleScroll,ref:"yScrollElRef"},f()))}),r&&t&&F?Ht(t,!0):null,ue(g,E=>E&&T("div",{class:`${e}-tabs-nav__suffix`},E))),P&&(this.animated&&(w==="top"||w==="bottom")?T("div",{ref:"tabsPaneWrapperRef",style:c,class:[`${e}-tabs-pane-wrapper`,l]},Wt(b,this.mergedValue,this.renderedNames,this.onAnimationBeforeLeave,this.onAnimationEnter,this.onAnimationAfterEnter,this.animationDirection)):Wt(b,this.mergedValue,this.renderedNames)))}});function Wt(e,a,o,r,t,n,i){const s=[];return e.forEach(l=>{const{name:c,displayDirective:h,"display-directive":x}=l.props,g=S=>h===S||x===S,b=a===c;if(l.key!==void 0&&(l.key=c),b||g("show")||g("show:lazy")&&o.has(c)){o.has(c)||o.add(c);const S=!g("if");s.push(S?Ar(l,[[Nr,b]]):l)}}),i?T(Ir,{name:`${i}-transition`,onBeforeLeave:r,onEnter:t,onAfterEnter:n},{default:()=>s}):s}function Ht(e,a){return T(Tt,{ref:"addTabInstRef",key:"__addable",name:"__addable",internalCreatedByPane:!0,internalAddable:!0,internalLeftPadded:a,disabled:typeof e=="object"&&e.disabled})}function Kt(e){const a=$r(e);return a.props?a.props.internalLeftPadded=!0:a.props={internalLeftPadded:!0},a}function et(e){return Array.isArray(e.dynamicProps)?e.dynamicProps.includes("internalLeftPadded")||e.dynamicProps.push("internalLeftPadded"):e.dynamicProps=["internalLeftPadded"],e}function Pa(e){return e&&e.__esModule&&Object.prototype.hasOwnProperty.call(e,"default")?e.default:e}var Pe={},tt,Jt;function Ba(){return Jt||(Jt=1,tt=function(){return typeof Promise=="function"&&Promise.prototype&&Promise.prototype.then}),tt}var nt={},be={},Yt;function ye(){if(Yt)return be;Yt=1;let e;const a=[0,26,44,70,100,134,172,196,242,292,346,404,466,532,581,655,733,815,901,991,1085,1156,1258,1364,1474,1588,1706,1828,1921,2051,2185,2323,2465,2611,2761,2876,3034,3196,3362,3532,3706];return be.getSymbolSize=function(r){if(!r)throw new Error('"version" cannot be null or undefined');if(r<1||r>40)throw new Error('"version" should be in range from 1 to 40');return r*4+17},be.getSymbolTotalCodewords=function(r){return a[r]},be.getBCHDigit=function(o){let r=0;for(;o!==0;)r++,o>>>=1;return r},be.setToSJISFunction=function(r){if(typeof r!="function")throw new Error('"toSJISFunc" is not a valid function.');e=r},be.isKanjiModeEnabled=function(){return typeof e<"u"},be.toSJIS=function(r){return e(r)},be}var rt={},Gt;function Nt(){return Gt||(Gt=1,(function(e){e.L={bit:1},e.M={bit:0},e.Q={bit:3},e.H={bit:2};function a(o){if(typeof o!="string")throw new Error("Param is not a string");switch(o.toLowerCase()){case"l":case"low":return e.L;case"m":case"medium":return e.M;case"q":case"quartile":return e.Q;case"h":case"high":return e.H;default:throw new Error("Unknown EC Level: "+o)}}e.isValid=function(r){return r&&typeof r.bit<"u"&&r.bit>=0&&r.bit<4},e.from=function(r,t){if(e.isValid(r))return r;try{return a(r)}catch{return t}}})(rt)),rt}var at,Qt;function Ta(){if(Qt)return at;Qt=1;function e(){this.buffer=[],this.length=0}return e.prototype={get:function(a){const o=Math.floor(a/8);return(this.buffer[o]>>>7-a%8&1)===1},put:function(a,o){for(let r=0;r<o;r++)this.putBit((a>>>o-r-1&1)===1)},getLengthInBits:function(){return this.length},putBit:function(a){const o=Math.floor(this.length/8);this.buffer.length<=o&&this.buffer.push(0),a&&(this.buffer[o]|=128>>>this.length%8),this.length++}},at=e,at}var ot,Xt;function Ea(){if(Xt)return ot;Xt=1;function e(a){if(!a||a<1)throw new Error("BitMatrix size must be defined and greater than 0");this.size=a,this.data=new Uint8Array(a*a),this.reservedBit=new Uint8Array(a*a)}return e.prototype.set=function(a,o,r,t){const n=a*this.size+o;this.data[n]=r,t&&(this.reservedBit[n]=!0)},e.prototype.get=function(a,o){return this.data[a*this.size+o]},e.prototype.xor=function(a,o,r){this.data[a*this.size+o]^=r},e.prototype.isReserved=function(a,o){return this.reservedBit[a*this.size+o]},ot=e,ot}var it={},Zt;function ka(){return Zt||(Zt=1,(function(e){const a=ye().getSymbolSize;e.getRowColCoords=function(r){if(r===1)return[];const t=Math.floor(r/7)+2,n=a(r),i=n===145?26:Math.ceil((n-13)/(2*t-2))*2,s=[n-7];for(let l=1;l<t-1;l++)s[l]=s[l-1]-i;return s.push(6),s.reverse()},e.getPositions=function(r){const t=[],n=e.getRowColCoords(r),i=n.length;for(let s=0;s<i;s++)for(let l=0;l<i;l++)s===0&&l===0||s===0&&l===i-1||s===i-1&&l===0||t.push([n[s],n[l]]);return t}})(it)),it}var st={},en;function za(){if(en)return st;en=1;const e=ye().getSymbolSize,a=7;return st.getPositions=function(r){const t=e(r);return[[0,0],[t-a,0],[0,t-a]]},st}var lt={},tn;function Aa(){return tn||(tn=1,(function(e){e.Patterns={PATTERN000:0,PATTERN001:1,PATTERN010:2,PATTERN011:3,PATTERN100:4,PATTERN101:5,PATTERN110:6,PATTERN111:7};const a={N1:3,N2:3,N3:40,N4:10};e.isValid=function(t){return t!=null&&t!==""&&!isNaN(t)&&t>=0&&t<=7},e.from=function(t){return e.isValid(t)?parseInt(t,10):void 0},e.getPenaltyN1=function(t){const n=t.size;let i=0,s=0,l=0,c=null,h=null;for(let x=0;x<n;x++){s=l=0,c=h=null;for(let g=0;g<n;g++){let b=t.get(x,g);b===c?s++:(s>=5&&(i+=a.N1+(s-5)),c=b,s=1),b=t.get(g,x),b===h?l++:(l>=5&&(i+=a.N1+(l-5)),h=b,l=1)}s>=5&&(i+=a.N1+(s-5)),l>=5&&(i+=a.N1+(l-5))}return i},e.getPenaltyN2=function(t){const n=t.size;let i=0;for(let s=0;s<n-1;s++)for(let l=0;l<n-1;l++){const c=t.get(s,l)+t.get(s,l+1)+t.get(s+1,l)+t.get(s+1,l+1);(c===4||c===0)&&i++}return i*a.N2},e.getPenaltyN3=function(t){const n=t.size;let i=0,s=0,l=0;for(let c=0;c<n;c++){s=l=0;for(let h=0;h<n;h++)s=s<<1&2047|t.get(c,h),h>=10&&(s===1488||s===93)&&i++,l=l<<1&2047|t.get(h,c),h>=10&&(l===1488||l===93)&&i++}return i*a.N3},e.getPenaltyN4=function(t){let n=0;const i=t.data.length;for(let l=0;l<i;l++)n+=t.data[l];return Math.abs(Math.ceil(n*100/i/5)-10)*a.N4};function o(r,t,n){switch(r){case e.Patterns.PATTERN000:return(t+n)%2===0;case e.Patterns.PATTERN001:return t%2===0;case e.Patterns.PATTERN010:return n%3===0;case e.Patterns.PATTERN011:return(t+n)%3===0;case e.Patterns.PATTERN100:return(Math.floor(t/2)+Math.floor(n/3))%2===0;case e.Patterns.PATTERN101:return t*n%2+t*n%3===0;case e.Patterns.PATTERN110:return(t*n%2+t*n%3)%2===0;case e.Patterns.PATTERN111:return(t*n%3+(t+n)%2)%2===0;default:throw new Error("bad maskPattern:"+r)}}e.applyMask=function(t,n){const i=n.size;for(let s=0;s<i;s++)for(let l=0;l<i;l++)n.isReserved(l,s)||n.xor(l,s,o(t,l,s))},e.getBestMask=function(t,n){const i=Object.keys(e.Patterns).length;let s=0,l=1/0;for(let c=0;c<i;c++){n(c),e.applyMask(c,t);const h=e.getPenaltyN1(t)+e.getPenaltyN2(t)+e.getPenaltyN3(t)+e.getPenaltyN4(t);e.applyMask(c,t),h<l&&(l=h,s=c)}return s}})(lt)),lt}var De={},nn;function Nn(){if(nn)return De;nn=1;const e=Nt(),a=[1,1,1,1,1,1,1,1,1,1,2,2,1,2,2,4,1,2,4,4,2,4,4,4,2,4,6,5,2,4,6,6,2,5,8,8,4,5,8,8,4,5,8,11,4,8,10,11,4,9,12,16,4,9,16,16,6,10,12,18,6,10,17,16,6,11,16,19,6,13,18,21,7,14,21,25,8,16,20,25,8,17,23,25,9,17,23,34,9,18,25,30,10,20,27,32,12,21,29,35,12,23,34,37,12,25,34,40,13,26,35,42,14,28,38,45,15,29,40,48,16,31,43,51,17,33,45,54,18,35,48,57,19,37,51,60,19,38,53,63,20,40,56,66,21,43,59,70,22,45,62,74,24,47,65,77,25,49,68,81],o=[7,10,13,17,10,16,22,28,15,26,36,44,20,36,52,64,26,48,72,88,36,64,96,112,40,72,108,130,48,88,132,156,60,110,160,192,72,130,192,224,80,150,224,264,96,176,260,308,104,198,288,352,120,216,320,384,132,240,360,432,144,280,408,480,168,308,448,532,180,338,504,588,196,364,546,650,224,416,600,700,224,442,644,750,252,476,690,816,270,504,750,900,300,560,810,960,312,588,870,1050,336,644,952,1110,360,700,1020,1200,390,728,1050,1260,420,784,1140,1350,450,812,1200,1440,480,868,1290,1530,510,924,1350,1620,540,980,1440,1710,570,1036,1530,1800,570,1064,1590,1890,600,1120,1680,1980,630,1204,1770,2100,660,1260,1860,2220,720,1316,1950,2310,750,1372,2040,2430];return De.getBlocksCount=function(t,n){switch(n){case e.L:return a[(t-1)*4+0];case e.M:return a[(t-1)*4+1];case e.Q:return a[(t-1)*4+2];case e.H:return a[(t-1)*4+3];default:return}},De.getTotalCodewordsCount=function(t,n){switch(n){case e.L:return o[(t-1)*4+0];case e.M:return o[(t-1)*4+1];case e.Q:return o[(t-1)*4+2];case e.H:return o[(t-1)*4+3];default:return}},De}var dt={},$e={},rn;function Na(){if(rn)return $e;rn=1;const e=new Uint8Array(512),a=new Uint8Array(256);return(function(){let r=1;for(let t=0;t<255;t++)e[t]=r,a[r]=t,r<<=1,r&256&&(r^=285);for(let t=255;t<512;t++)e[t]=e[t-255]})(),$e.log=function(r){if(r<1)throw new Error("log("+r+")");return a[r]},$e.exp=function(r){return e[r]},$e.mul=function(r,t){return r===0||t===0?0:e[a[r]+a[t]]},$e}var an;function Ia(){return an||(an=1,(function(e){const a=Na();e.mul=function(r,t){const n=new Uint8Array(r.length+t.length-1);for(let i=0;i<r.length;i++)for(let s=0;s<t.length;s++)n[i+s]^=a.mul(r[i],t[s]);return n},e.mod=function(r,t){let n=new Uint8Array(r);for(;n.length-t.length>=0;){const i=n[0];for(let l=0;l<t.length;l++)n[l]^=a.mul(t[l],i);let s=0;for(;s<n.length&&n[s]===0;)s++;n=n.slice(s)}return n},e.generateECPolynomial=function(r){let t=new Uint8Array([1]);for(let n=0;n<r;n++)t=e.mul(t,new Uint8Array([1,a.exp(n)]));return t}})(dt)),dt}var ut,on;function $a(){if(on)return ut;on=1;const e=Ia();function a(o){this.genPoly=void 0,this.degree=o,this.degree&&this.initialize(this.degree)}return a.prototype.initialize=function(r){this.degree=r,this.genPoly=e.generateECPolynomial(this.degree)},a.prototype.encode=function(r){if(!this.genPoly)throw new Error("Encoder not initialized");const t=new Uint8Array(r.length+this.degree);t.set(r);const n=e.mod(t,this.genPoly),i=this.degree-n.length;if(i>0){const s=new Uint8Array(this.degree);return s.set(n,i),s}return n},ut=a,ut}var ct={},ft={},ht={},sn;function In(){return sn||(sn=1,ht.isValid=function(a){return!isNaN(a)&&a>=1&&a<=40}),ht}var le={},ln;function $n(){if(ln)return le;ln=1;const e="[0-9]+",a="[A-Z $%*+\\-./:]+";let o="(?:[u3000-u303F]|[u3040-u309F]|[u30A0-u30FF]|[uFF00-uFFEF]|[u4E00-u9FAF]|[u2605-u2606]|[u2190-u2195]|u203B|[u2010u2015u2018u2019u2025u2026u201Cu201Du2225u2260]|[u0391-u0451]|[u00A7u00A8u00B1u00B4u00D7u00F7])+";o=o.replace(/u/g,"\\u");const r="(?:(?![A-Z0-9 $%*+\\-./:]|"+o+`)(?:.|[\r
]))+`;le.KANJI=new RegExp(o,"g"),le.BYTE_KANJI=new RegExp("[^A-Z0-9 $%*+\\-./:]+","g"),le.BYTE=new RegExp(r,"g"),le.NUMERIC=new RegExp(e,"g"),le.ALPHANUMERIC=new RegExp(a,"g");const t=new RegExp("^"+o+"$"),n=new RegExp("^"+e+"$"),i=new RegExp("^[A-Z0-9 $%*+\\-./:]+$");return le.testKanji=function(l){return t.test(l)},le.testNumeric=function(l){return n.test(l)},le.testAlphanumeric=function(l){return i.test(l)},le}var dn;function we(){return dn||(dn=1,(function(e){const a=In(),o=$n();e.NUMERIC={id:"Numeric",bit:1,ccBits:[10,12,14]},e.ALPHANUMERIC={id:"Alphanumeric",bit:2,ccBits:[9,11,13]},e.BYTE={id:"Byte",bit:4,ccBits:[8,16,16]},e.KANJI={id:"Kanji",bit:8,ccBits:[8,10,12]},e.MIXED={bit:-1},e.getCharCountIndicator=function(n,i){if(!n.ccBits)throw new Error("Invalid mode: "+n);if(!a.isValid(i))throw new Error("Invalid version: "+i);return i>=1&&i<10?n.ccBits[0]:i<27?n.ccBits[1]:n.ccBits[2]},e.getBestModeForData=function(n){return o.testNumeric(n)?e.NUMERIC:o.testAlphanumeric(n)?e.ALPHANUMERIC:o.testKanji(n)?e.KANJI:e.BYTE},e.toString=function(n){if(n&&n.id)return n.id;throw new Error("Invalid mode")},e.isValid=function(n){return n&&n.bit&&n.ccBits};function r(t){if(typeof t!="string")throw new Error("Param is not a string");switch(t.toLowerCase()){case"numeric":return e.NUMERIC;case"alphanumeric":return e.ALPHANUMERIC;case"kanji":return e.KANJI;case"byte":return e.BYTE;default:throw new Error("Unknown mode: "+t)}}e.from=function(n,i){if(e.isValid(n))return n;try{return r(n)}catch{return i}}})(ft)),ft}var un;function _a(){return un||(un=1,(function(e){const a=ye(),o=Nn(),r=Nt(),t=we(),n=In(),i=7973,s=a.getBCHDigit(i);function l(g,b,S){for(let P=1;P<=40;P++)if(b<=e.getCapacity(P,S,g))return P}function c(g,b){return t.getCharCountIndicator(g,b)+4}function h(g,b){let S=0;return g.forEach(function(P){const F=c(P.mode,b);S+=F+P.getBitsLength()}),S}function x(g,b){for(let S=1;S<=40;S++)if(h(g,S)<=e.getCapacity(S,b,t.MIXED))return S}e.from=function(b,S){return n.isValid(b)?parseInt(b,10):S},e.getCapacity=function(b,S,P){if(!n.isValid(b))throw new Error("Invalid QR Code version");typeof P>"u"&&(P=t.BYTE);const F=a.getSymbolTotalCodewords(b),B=o.getTotalCodewordsCount(b,S),y=(F-B)*8;if(P===t.MIXED)return y;const f=y-c(P,b);switch(P){case t.NUMERIC:return Math.floor(f/10*3);case t.ALPHANUMERIC:return Math.floor(f/11*2);case t.KANJI:return Math.floor(f/13);case t.BYTE:default:return Math.floor(f/8)}},e.getBestVersionForData=function(b,S){let P;const F=r.from(S,r.M);if(Array.isArray(b)){if(b.length>1)return x(b,F);if(b.length===0)return 1;P=b[0]}else P=b;return l(P.mode,P.getLength(),F)},e.getEncodedBits=function(b){if(!n.isValid(b)||b<7)throw new Error("Invalid QR Code version");let S=b<<12;for(;a.getBCHDigit(S)-s>=0;)S^=i<<a.getBCHDigit(S)-s;return b<<12|S}})(ct)),ct}var gt={},cn;function La(){if(cn)return gt;cn=1;const e=ye(),a=1335,o=21522,r=e.getBCHDigit(a);return gt.getEncodedBits=function(n,i){const s=n.bit<<3|i;let l=s<<10;for(;e.getBCHDigit(l)-r>=0;)l^=a<<e.getBCHDigit(l)-r;return(s<<10|l)^o},gt}var bt={},pt,fn;function Ma(){if(fn)return pt;fn=1;const e=we();function a(o){this.mode=e.NUMERIC,this.data=o.toString()}return a.getBitsLength=function(r){return 10*Math.floor(r/3)+(r%3?r%3*3+1:0)},a.prototype.getLength=function(){return this.data.length},a.prototype.getBitsLength=function(){return a.getBitsLength(this.data.length)},a.prototype.write=function(r){let t,n,i;for(t=0;t+3<=this.data.length;t+=3)n=this.data.substr(t,3),i=parseInt(n,10),r.put(i,10);const s=this.data.length-t;s>0&&(n=this.data.substr(t),i=parseInt(n,10),r.put(i,s*3+1))},pt=a,pt}var vt,hn;function Ua(){if(hn)return vt;hn=1;const e=we(),a=["0","1","2","3","4","5","6","7","8","9","A","B","C","D","E","F","G","H","I","J","K","L","M","N","O","P","Q","R","S","T","U","V","W","X","Y","Z"," ","$","%","*","+","-",".","/",":"];function o(r){this.mode=e.ALPHANUMERIC,this.data=r}return o.getBitsLength=function(t){return 11*Math.floor(t/2)+6*(t%2)},o.prototype.getLength=function(){return this.data.length},o.prototype.getBitsLength=function(){return o.getBitsLength(this.data.length)},o.prototype.write=function(t){let n;for(n=0;n+2<=this.data.length;n+=2){let i=a.indexOf(this.data[n])*45;i+=a.indexOf(this.data[n+1]),t.put(i,11)}this.data.length%2&&t.put(a.indexOf(this.data[n]),6)},vt=o,vt}var mt,gn;function Fa(){if(gn)return mt;gn=1;const e=we();function a(o){this.mode=e.BYTE,typeof o=="string"?this.data=new TextEncoder().encode(o):this.data=new Uint8Array(o)}return a.getBitsLength=function(r){return r*8},a.prototype.getLength=function(){return this.data.length},a.prototype.getBitsLength=function(){return a.getBitsLength(this.data.length)},a.prototype.write=function(o){for(let r=0,t=this.data.length;r<t;r++)o.put(this.data[r],8)},mt=a,mt}var yt,bn;function Da(){if(bn)return yt;bn=1;const e=we(),a=ye();function o(r){this.mode=e.KANJI,this.data=r}return o.getBitsLength=function(t){return t*13},o.prototype.getLength=function(){return this.data.length},o.prototype.getBitsLength=function(){return o.getBitsLength(this.data.length)},o.prototype.write=function(r){let t;for(t=0;t<this.data.length;t++){let n=a.toSJIS(this.data[t]);if(n>=33088&&n<=40956)n-=33088;else if(n>=57408&&n<=60351)n-=49472;else throw new Error("Invalid SJIS character: "+this.data[t]+`
Make sure your charset is UTF-8`);n=(n>>>8&255)*192+(n&255),r.put(n,13)}},yt=o,yt}var wt={exports:{}},pn;function ja(){return pn||(pn=1,(function(e){var a={single_source_shortest_paths:function(o,r,t){var n={},i={};i[r]=0;var s=a.PriorityQueue.make();s.push(r,0);for(var l,c,h,x,g,b,S,P,F;!s.empty();){l=s.pop(),c=l.value,x=l.cost,g=o[c]||{};for(h in g)g.hasOwnProperty(h)&&(b=g[h],S=x+b,P=i[h],F=typeof i[h]>"u",(F||P>S)&&(i[h]=S,s.push(h,S),n[h]=c))}if(typeof t<"u"&&typeof i[t]>"u"){var B=["Could not find a path from ",r," to ",t,"."].join("");throw new Error(B)}return n},extract_shortest_path_from_predecessor_list:function(o,r){for(var t=[],n=r;n;)t.push(n),o[n],n=o[n];return t.reverse(),t},find_path:function(o,r,t){var n=a.single_source_shortest_paths(o,r,t);return a.extract_shortest_path_from_predecessor_list(n,t)},PriorityQueue:{make:function(o){var r=a.PriorityQueue,t={},n;o=o||{};for(n in r)r.hasOwnProperty(n)&&(t[n]=r[n]);return t.queue=[],t.sorter=o.sorter||r.default_sorter,t},default_sorter:function(o,r){return o.cost-r.cost},push:function(o,r){var t={value:o,cost:r};this.queue.push(t),this.queue.sort(this.sorter)},pop:function(){return this.queue.shift()},empty:function(){return this.queue.length===0}}};e.exports=a})(wt)),wt.exports}var vn;function Oa(){return vn||(vn=1,(function(e){const a=we(),o=Ma(),r=Ua(),t=Fa(),n=Da(),i=$n(),s=ye(),l=ja();function c(B){return unescape(encodeURIComponent(B)).length}function h(B,y,f){const w=[];let E;for(;(E=B.exec(f))!==null;)w.push({data:E[0],index:E.index,mode:y,length:E[0].length});return w}function x(B){const y=h(i.NUMERIC,a.NUMERIC,B),f=h(i.ALPHANUMERIC,a.ALPHANUMERIC,B);let w,E;return s.isKanjiModeEnabled()?(w=h(i.BYTE,a.BYTE,B),E=h(i.KANJI,a.KANJI,B)):(w=h(i.BYTE_KANJI,a.BYTE,B),E=[]),y.concat(f,w,E).sort(function(u,C){return u.index-C.index}).map(function(u){return{data:u.data,mode:u.mode,length:u.length}})}function g(B,y){switch(y){case a.NUMERIC:return o.getBitsLength(B);case a.ALPHANUMERIC:return r.getBitsLength(B);case a.KANJI:return n.getBitsLength(B);case a.BYTE:return t.getBitsLength(B)}}function b(B){return B.reduce(function(y,f){const w=y.length-1>=0?y[y.length-1]:null;return w&&w.mode===f.mode?(y[y.length-1].data+=f.data,y):(y.push(f),y)},[])}function S(B){const y=[];for(let f=0;f<B.length;f++){const w=B[f];switch(w.mode){case a.NUMERIC:y.push([w,{data:w.data,mode:a.ALPHANUMERIC,length:w.length},{data:w.data,mode:a.BYTE,length:w.length}]);break;case a.ALPHANUMERIC:y.push([w,{data:w.data,mode:a.BYTE,length:w.length}]);break;case a.KANJI:y.push([w,{data:w.data,mode:a.BYTE,length:c(w.data)}]);break;case a.BYTE:y.push([{data:w.data,mode:a.BYTE,length:c(w.data)}])}}return y}function P(B,y){const f={},w={start:{}};let E=["start"];for(let d=0;d<B.length;d++){const u=B[d],C=[];for(let p=0;p<u.length;p++){const N=u[p],k=""+d+p;C.push(k),f[k]={node:N,lastCount:0},w[k]={};for(let z=0;z<E.length;z++){const A=E[z];f[A]&&f[A].node.mode===N.mode?(w[A][k]=g(f[A].lastCount+N.length,N.mode)-g(f[A].lastCount,N.mode),f[A].lastCount+=N.length):(f[A]&&(f[A].lastCount=N.length),w[A][k]=g(N.length,N.mode)+4+a.getCharCountIndicator(N.mode,y))}}E=C}for(let d=0;d<E.length;d++)w[E[d]].end=0;return{map:w,table:f}}function F(B,y){let f;const w=a.getBestModeForData(B);if(f=a.from(y,w),f!==a.BYTE&&f.bit<w.bit)throw new Error('"'+B+'" cannot be encoded with mode '+a.toString(f)+`.
 Suggested mode is: `+a.toString(w));switch(f===a.KANJI&&!s.isKanjiModeEnabled()&&(f=a.BYTE),f){case a.NUMERIC:return new o(B);case a.ALPHANUMERIC:return new r(B);case a.KANJI:return new n(B);case a.BYTE:return new t(B)}}e.fromArray=function(y){return y.reduce(function(f,w){return typeof w=="string"?f.push(F(w,null)):w.data&&f.push(F(w.data,w.mode)),f},[])},e.fromString=function(y,f){const w=x(y,s.isKanjiModeEnabled()),E=S(w),d=P(E,f),u=l.find_path(d.map,"start","end"),C=[];for(let p=1;p<u.length-1;p++)C.push(d.table[u[p]].node);return e.fromArray(b(C))},e.rawSplit=function(y){return e.fromArray(x(y,s.isKanjiModeEnabled()))}})(bt)),bt}var mn;function qa(){if(mn)return nt;mn=1;const e=ye(),a=Nt(),o=Ta(),r=Ea(),t=ka(),n=za(),i=Aa(),s=Nn(),l=$a(),c=_a(),h=La(),x=we(),g=Oa();function b(d,u){const C=d.size,p=n.getPositions(u);for(let N=0;N<p.length;N++){const k=p[N][0],z=p[N][1];for(let A=-1;A<=7;A++)if(!(k+A<=-1||C<=k+A))for(let L=-1;L<=7;L++)z+L<=-1||C<=z+L||(A>=0&&A<=6&&(L===0||L===6)||L>=0&&L<=6&&(A===0||A===6)||A>=2&&A<=4&&L>=2&&L<=4?d.set(k+A,z+L,!0,!0):d.set(k+A,z+L,!1,!0))}}function S(d){const u=d.size;for(let C=8;C<u-8;C++){const p=C%2===0;d.set(C,6,p,!0),d.set(6,C,p,!0)}}function P(d,u){const C=t.getPositions(u);for(let p=0;p<C.length;p++){const N=C[p][0],k=C[p][1];for(let z=-2;z<=2;z++)for(let A=-2;A<=2;A++)z===-2||z===2||A===-2||A===2||z===0&&A===0?d.set(N+z,k+A,!0,!0):d.set(N+z,k+A,!1,!0)}}function F(d,u){const C=d.size,p=c.getEncodedBits(u);let N,k,z;for(let A=0;A<18;A++)N=Math.floor(A/3),k=A%3+C-8-3,z=(p>>A&1)===1,d.set(N,k,z,!0),d.set(k,N,z,!0)}function B(d,u,C){const p=d.size,N=h.getEncodedBits(u,C);let k,z;for(k=0;k<15;k++)z=(N>>k&1)===1,k<6?d.set(k,8,z,!0):k<8?d.set(k+1,8,z,!0):d.set(p-15+k,8,z,!0),k<8?d.set(8,p-k-1,z,!0):k<9?d.set(8,15-k-1+1,z,!0):d.set(8,15-k-1,z,!0);d.set(p-8,8,1,!0)}function y(d,u){const C=d.size;let p=-1,N=C-1,k=7,z=0;for(let A=C-1;A>0;A-=2)for(A===6&&A--;;){for(let L=0;L<2;L++)if(!d.isReserved(N,A-L)){let G=!1;z<u.length&&(G=(u[z]>>>k&1)===1),d.set(N,A-L,G),k--,k===-1&&(z++,k=7)}if(N+=p,N<0||C<=N){N-=p,p=-p;break}}}function f(d,u,C){const p=new o;C.forEach(function(L){p.put(L.mode.bit,4),p.put(L.getLength(),x.getCharCountIndicator(L.mode,d)),L.write(p)});const N=e.getSymbolTotalCodewords(d),k=s.getTotalCodewordsCount(d,u),z=(N-k)*8;for(p.getLengthInBits()+4<=z&&p.put(0,4);p.getLengthInBits()%8!==0;)p.putBit(0);const A=(z-p.getLengthInBits())/8;for(let L=0;L<A;L++)p.put(L%2?17:236,8);return w(p,d,u)}function w(d,u,C){const p=e.getSymbolTotalCodewords(u),N=s.getTotalCodewordsCount(u,C),k=p-N,z=s.getBlocksCount(u,C),A=p%z,L=z-A,G=Math.floor(p/z),re=Math.floor(k/z),ee=re+1,Ce=G-re,te=new l(Ce);let pe=0;const ce=new Array(z),Q=new Array(z);let ae=0;const ze=new Uint8Array(d.buffer);for(let X=0;X<z;X++){const Ae=X<L?re:ee;ce[X]=ze.slice(pe,pe+Ae),Q[X]=te.encode(ce[X]),pe+=Ae,ae=Math.max(ae,Ae)}const ve=new Uint8Array(p);let me=0,V,H;for(V=0;V<ae;V++)for(H=0;H<z;H++)V<ce[H].length&&(ve[me++]=ce[H][V]);for(V=0;V<Ce;V++)for(H=0;H<z;H++)ve[me++]=Q[H][V];return ve}function E(d,u,C,p){let N;if(Array.isArray(d))N=g.fromArray(d);else if(typeof d=="string"){let G=u;if(!G){const re=g.rawSplit(d);G=c.getBestVersionForData(re,C)}N=g.fromString(d,G||40)}else throw new Error("Invalid data");const k=c.getBestVersionForData(N,C);if(!k)throw new Error("The amount of data is too big to be stored in a QR Code");if(!u)u=k;else if(u<k)throw new Error(`
The chosen QR Code version cannot contain this amount of data.
Minimum version required to store current data is: `+k+`.
`);const z=f(u,C,N),A=e.getSymbolSize(u),L=new r(A);return b(L,u),S(L),P(L,u),B(L,C,0),u>=7&&F(L,u),y(L,z),isNaN(p)&&(p=i.getBestMask(L,B.bind(null,L,C))),i.applyMask(p,L),B(L,C,p),{modules:L,version:u,errorCorrectionLevel:C,maskPattern:p,segments:N}}return nt.create=function(u,C){if(typeof u>"u"||u==="")throw new Error("No input text");let p=a.M,N,k;return typeof C<"u"&&(p=a.from(C.errorCorrectionLevel,a.M),N=c.from(C.version),k=i.from(C.maskPattern),C.toSJISFunc&&e.setToSJISFunction(C.toSJISFunc)),E(u,N,p,k)},nt}var Ct={},xt={},yn;function _n(){return yn||(yn=1,(function(e){function a(o){if(typeof o=="number"&&(o=o.toString()),typeof o!="string")throw new Error("Color should be defined as hex string");let r=o.slice().replace("#","").split("");if(r.length<3||r.length===5||r.length>8)throw new Error("Invalid hex color: "+o);(r.length===3||r.length===4)&&(r=Array.prototype.concat.apply([],r.map(function(n){return[n,n]}))),r.length===6&&r.push("F","F");const t=parseInt(r.join(""),16);return{r:t>>24&255,g:t>>16&255,b:t>>8&255,a:t&255,hex:"#"+r.slice(0,6).join("")}}e.getOptions=function(r){r||(r={}),r.color||(r.color={});const t=typeof r.margin>"u"||r.margin===null||r.margin<0?4:r.margin,n=r.width&&r.width>=21?r.width:void 0,i=r.scale||4;return{width:n,scale:n?4:i,margin:t,color:{dark:a(r.color.dark||"#000000ff"),light:a(r.color.light||"#ffffffff")},type:r.type,rendererOpts:r.rendererOpts||{}}},e.getScale=function(r,t){return t.width&&t.width>=r+t.margin*2?t.width/(r+t.margin*2):t.scale},e.getImageWidth=function(r,t){const n=e.getScale(r,t);return Math.floor((r+t.margin*2)*n)},e.qrToImageData=function(r,t,n){const i=t.modules.size,s=t.modules.data,l=e.getScale(i,n),c=Math.floor((i+n.margin*2)*l),h=n.margin*l,x=[n.color.light,n.color.dark];for(let g=0;g<c;g++)for(let b=0;b<c;b++){let S=(g*c+b)*4,P=n.color.light;if(g>=h&&b>=h&&g<c-h&&b<c-h){const F=Math.floor((g-h)/l),B=Math.floor((b-h)/l);P=x[s[F*i+B]?1:0]}r[S++]=P.r,r[S++]=P.g,r[S++]=P.b,r[S]=P.a}}})(xt)),xt}var wn;function Va(){return wn||(wn=1,(function(e){const a=_n();function o(t,n,i){t.clearRect(0,0,n.width,n.height),n.style||(n.style={}),n.height=i,n.width=i,n.style.height=i+"px",n.style.width=i+"px"}function r(){try{return document.createElement("canvas")}catch{throw new Error("You need to specify a canvas element")}}e.render=function(n,i,s){let l=s,c=i;typeof l>"u"&&(!i||!i.getContext)&&(l=i,i=void 0),i||(c=r()),l=a.getOptions(l);const h=a.getImageWidth(n.modules.size,l),x=c.getContext("2d"),g=x.createImageData(h,h);return a.qrToImageData(g.data,n,l),o(x,c,h),x.putImageData(g,0,0),c},e.renderToDataURL=function(n,i,s){let l=s;typeof l>"u"&&(!i||!i.getContext)&&(l=i,i=void 0),l||(l={});const c=e.render(n,i,l),h=l.type||"image/png",x=l.rendererOpts||{};return c.toDataURL(h,x.quality)}})(Ct)),Ct}var Rt={},Cn;function Wa(){if(Cn)return Rt;Cn=1;const e=_n();function a(t,n){const i=t.a/255,s=n+'="'+t.hex+'"';return i<1?s+" "+n+'-opacity="'+i.toFixed(2).slice(1)+'"':s}function o(t,n,i){let s=t+n;return typeof i<"u"&&(s+=" "+i),s}function r(t,n,i){let s="",l=0,c=!1,h=0;for(let x=0;x<t.length;x++){const g=Math.floor(x%n),b=Math.floor(x/n);!g&&!c&&(c=!0),t[x]?(h++,x>0&&g>0&&t[x-1]||(s+=c?o("M",g+i,.5+b+i):o("m",l,0),l=0,c=!1),g+1<n&&t[x+1]||(s+=o("h",h),h=0)):l++}return s}return Rt.render=function(n,i,s){const l=e.getOptions(i),c=n.modules.size,h=n.modules.data,x=c+l.margin*2,g=l.color.light.a?"<path "+a(l.color.light,"fill")+' d="M0 0h'+x+"v"+x+'H0z"/>':"",b="<path "+a(l.color.dark,"stroke")+' d="'+r(h,c,l.margin)+'"/>',S='viewBox="0 0 '+x+" "+x+'"',F='<svg xmlns="http://www.w3.org/2000/svg" '+(l.width?'width="'+l.width+'" height="'+l.width+'" ':"")+S+' shape-rendering="crispEdges">'+g+b+`</svg>
`;return typeof s=="function"&&s(null,F),F},Rt}var xn;function Ha(){if(xn)return Pe;xn=1;const e=Ba(),a=qa(),o=Va(),r=Wa();function t(n,i,s,l,c){const h=[].slice.call(arguments,1),x=h.length,g=typeof h[x-1]=="function";if(!g&&!e())throw new Error("Callback required as last argument");if(g){if(x<2)throw new Error("Too few arguments provided");x===2?(c=s,s=i,i=l=void 0):x===3&&(i.getContext&&typeof c>"u"?(c=l,l=void 0):(c=l,l=s,s=i,i=void 0))}else{if(x<1)throw new Error("Too few arguments provided");return x===1?(s=i,i=l=void 0):x===2&&!i.getContext&&(l=s,s=i,i=void 0),new Promise(function(b,S){try{const P=a.create(s,l);b(n(P,i,l))}catch(P){S(P)}})}try{const b=a.create(s,l);c(null,n(b,i,l))}catch(b){c(b)}}return Pe.create=a.create,Pe.toCanvas=t.bind(null,o.render),Pe.toDataURL=t.bind(null,o.renderToDataURL),Pe.toString=t.bind(null,function(n,i,s){return r.render(n,s)}),Pe}var Ka=Ha();const Ja=Pa(Ka),Ya=ge({__name:"ShareModal",props:{username:{}},emits:["close"],setup(e,{emit:a}){const o=e,r=a,t=En(),n=O(null),i=O(""),s=O(null),l=O(null),c=O(null),h=O("link"),x=O(60),g=[{label:"15 minutes",value:15},{label:"1 hour",value:60},{label:"24 hours",value:1440}],b=O(""),S=O(0),P=O(!1);async function F(d,u){if(!(!d||!u))try{await Ja.toCanvas(d,u,{width:280,margin:1})}catch{}}async function B(){if(await Oe(),h.value==="link"){await F(c.value,b.value);return}n.value&&(n.value.mierusUrls.length&&await F(s.value,n.value.mierusUrls[0]),await F(l.value,n.value.mieruUrl))}zt(async()=>{try{n.value=await Be.get(`/api/users/${encodeURIComponent(o.username)}/share`),await B()}catch(d){i.value=d instanceof _e?d.message:"Failed to build share links"}}),je(h,B);async function y(){P.value=!0;try{const d=await Be.post(`/api/users/${encodeURIComponent(o.username)}/share-token`,{ttlMinutes:x.value});b.value=d.url,S.value=d.expiresAt,await B()}catch(d){t.error(d instanceof _e?d.message:"Failed to create link")}finally{P.value=!1}}function f(){return S.value?`Expires ${new Date(S.value*1e3).toLocaleString()}`:""}function w(d){navigator.clipboard.writeText(d).then(()=>t.success("Copied"),()=>t.error("Copy failed"))}function E(){if(!n.value)return;const d=new Blob([n.value.clientConfigJson],{type:"application/json"}),u=document.createElement("a");u.href=URL.createObjectURL(d),u.download=`mieru-${o.username}.json`,u.click(),URL.revokeObjectURL(u.href)}return(d,u)=>(fe(),Ee(_(_r),{show:!0,preset:"card",title:`Share: ${e.username}`,style:{width:"480px"},onClose:u[6]||(u[6]=C=>r("close")),onMaskClick:u[7]||(u[7]=C=>r("close"))},{default:D(()=>[i.value?(fe(),Ee(_(Wr),{key:0,type:"warning","show-icon":!1},{default:D(()=>[J(Lt(i.value),1)]),_:1})):(fe(),Ee(_(Sa),{key:1,value:h.value,"onUpdate:value":u[5]||(u[5]=C=>h.value=C)},{default:D(()=>[U(_(Fe),{name:"link",tab:"Link"},{default:D(()=>[U(_(se),{vertical:"",align:"center"},{default:D(()=>[U(_(jt),{depth:"3",style:{"font-size":"12px","text-align":"center"}},{default:D(()=>[...u[8]||(u[8]=[J(" Generate a link that expires, so you don't have to send raw credentials around. Anyone with the link can view this user's config until it expires. ",-1)])]),_:1}),U(_(se),{align:"center"},{default:D(()=>[U(_(Vr),{value:x.value,"onUpdate:value":u[0]||(u[0]=C=>x.value=C),options:g,style:{width:"140px"}},null,8,["value"]),U(_(Y),{type:"primary",loading:P.value,onClick:y},{default:D(()=>[...u[9]||(u[9]=[J("Generate link",-1)])]),_:1},8,["loading"])]),_:1}),b.value?(fe(),Pt(We,{key:0},[qe("canvas",{ref_key:"linkCanvas",ref:c},null,512),U(_(Te),{value:b.value,readonly:""},null,8,["value"]),U(_(se),{align:"center"},{default:D(()=>[U(_(Y),{size:"small",onClick:u[1]||(u[1]=C=>w(b.value))},{default:D(()=>[...u[10]||(u[10]=[J("Copy link",-1)])]),_:1}),U(_(jt),{depth:"3",style:{"font-size":"12px"}},{default:D(()=>[J(Lt(f()),1)]),_:1})]),_:1})],64)):Bt("",!0)]),_:1})]),_:1}),U(_(Fe),{name:"mierus",tab:"mierus://",disabled:!n.value},{default:D(()=>[U(_(se),{vertical:"",align:"center"},{default:D(()=>{var C;return[qe("canvas",{ref_key:"mierusCanvas",ref:s},null,512),U(_(Te),{value:(C=n.value)==null?void 0:C.mierusUrls[0],type:"textarea",readonly:"",autosize:{minRows:2,maxRows:4}},null,8,["value"]),U(_(Y),{size:"small",onClick:u[2]||(u[2]=p=>w(n.value.mierusUrls[0]))},{default:D(()=>[...u[11]||(u[11]=[J("Copy link",-1)])]),_:1})]}),_:1})]),_:1},8,["disabled"]),U(_(Fe),{name:"mieru",tab:"mieru://",disabled:!n.value},{default:D(()=>[U(_(se),{vertical:"",align:"center"},{default:D(()=>{var C;return[qe("canvas",{ref_key:"mieruCanvas",ref:l},null,512),U(_(Te),{value:(C=n.value)==null?void 0:C.mieruUrl,type:"textarea",readonly:"",autosize:{minRows:2,maxRows:4}},null,8,["value"]),U(_(Y),{size:"small",onClick:u[3]||(u[3]=p=>w(n.value.mieruUrl))},{default:D(()=>[...u[12]||(u[12]=[J("Copy link",-1)])]),_:1})]}),_:1})]),_:1},8,["disabled"]),U(_(Fe),{name:"json",tab:"Config file",disabled:!n.value},{default:D(()=>[U(_(se),{vertical:""},{default:D(()=>{var C;return[U(_(Te),{value:(C=n.value)==null?void 0:C.clientConfigJson,type:"textarea",readonly:"",autosize:{minRows:8,maxRows:14}},null,8,["value"]),U(_(se),null,{default:D(()=>[U(_(Y),{size:"small",onClick:u[4]||(u[4]=p=>w(n.value.clientConfigJson))},{default:D(()=>[...u[13]||(u[13]=[J("Copy",-1)])]),_:1}),U(_(Y),{size:"small",onClick:E},{default:D(()=>[...u[14]||(u[14]=[J("Download JSON",-1)])]),_:1})]),_:1})]}),_:1})]),_:1},8,["disabled"])]),_:1},8,["value"]))]),_:1},8,["title"]))}}),uo=ge({__name:"UsersPage",setup(e){const a=En(),o=O([]),r=O(!1),t=O(!1),n=O(null),i=O({name:"",password:"",allowPrivateIP:!1}),s=O([]),l=O(null);function c(y){if(y<1024)return`${y} B`;const f=["KiB","MiB","GiB","TiB"];let w=y,E=-1;do w/=1024,E++;while(w>=1024&&E<f.length-1);return`${w.toFixed(1)} ${f[E]}`}function h(y){var d,u;const f=Object.entries(y.metrics);if(!f.length)return"—";const w=((d=f.find(([C])=>C.toLowerCase().includes("download")))==null?void 0:d[1])??0,E=((u=f.find(([C])=>C.toLowerCase().includes("upload")))==null?void 0:u[1])??0;return`↓ ${c(w)} / ↑ ${c(E)}`}const x=[{title:"Name",key:"name"},{title:"Quotas",key:"quotas",render:y=>y.quotas.length?y.quotas.map(f=>T(qr,{size:"small",style:"margin-right:4px"},{default:()=>`${f.megabytes} MB / ${f.days}d`})):"—"},{title:"Traffic",key:"traffic",render:h},{title:"Share",key:"share",render:y=>T(Y,{size:"small",disabled:!y.hasSecret,onClick:()=>l.value=y.name},{default:()=>y.hasSecret?"Share":"No password stored"})},{title:"Actions",key:"actions",render:y=>T(se,null,{default:()=>[T(Y,{size:"small",onClick:()=>S(y)},{default:()=>"Edit"}),T(va,{onPositiveClick:()=>B(y.name)},{trigger:()=>T(Y,{size:"small",type:"error",quaternary:!0},{default:()=>"Delete"}),default:()=>`Delete user ${y.name}?`})]})}];async function g(){r.value=!0;try{o.value=await Be.get("/api/users")}catch(y){a.error(y instanceof _e?y.message:"Failed to load users")}finally{r.value=!1}}function b(){n.value=null,i.value={name:"",password:"",allowPrivateIP:!1},s.value=[],t.value=!0}function S(y){n.value=y.name,i.value={name:y.name,password:"",allowPrivateIP:y.allowPrivateIP},s.value=y.quotas.map(f=>({...f})),t.value=!0}function P(){s.value.push({days:30,megabytes:10240})}async function F(){const y={password:i.value.password,quotas:s.value,allowPrivateIP:i.value.allowPrivateIP};try{n.value===null?(await Be.post("/api/users",{name:i.value.name,...y}),a.success("User created")):(await Be.put(`/api/users/${encodeURIComponent(n.value)}`,y),a.success("User updated")),t.value=!1,await g()}catch(f){a.error(f instanceof _e?f.message:"Save failed")}}async function B(y){try{await Be.del(`/api/users/${encodeURIComponent(y)}`),a.success("User deleted"),await g()}catch(f){a.error(f instanceof _e?f.message:"Delete failed")}}return zt(g),(y,f)=>(fe(),Pt(We,null,[U(_(Lr),{title:"Users"},{"header-extra":D(()=>[U(_(Y),{type:"primary",onClick:b},{default:D(()=>[...f[6]||(f[6]=[J("Add user",-1)])]),_:1})]),default:D(()=>[U(_(Jr),{columns:x,data:o.value,loading:r.value,"row-key":w=>w.name},null,8,["data","loading","row-key"])]),_:1}),U(_(Hr),{show:t.value,"onUpdate:show":f[4]||(f[4]=w=>t.value=w),width:420},{default:D(()=>[U(_(Kr),{title:n.value===null?"Add user":`Edit ${n.value}`},{footer:D(()=>[U(_(se),null,{default:D(()=>[U(_(Y),{onClick:f[3]||(f[3]=w=>t.value=!1)},{default:D(()=>[...f[12]||(f[12]=[J("Cancel",-1)])]),_:1}),U(_(Y),{type:"primary",onClick:F},{default:D(()=>[...f[13]||(f[13]=[J("Save",-1)])]),_:1})]),_:1})]),default:D(()=>[U(_(Yr),null,{default:D(()=>[n.value===null?(fe(),Ee(_(Ue),{key:0,label:"Name"},{default:D(()=>[U(_(Te),{value:i.value.name,"onUpdate:value":f[0]||(f[0]=w=>i.value.name=w),placeholder:"username"},null,8,["value"])]),_:1})):Bt("",!0),U(_(Ue),{label:n.value===null?"Password":"New password (leave empty to keep)"},{default:D(()=>[U(_(Te),{value:i.value.password,"onUpdate:value":f[1]||(f[1]=w=>i.value.password=w),type:"password","show-password-on":"click"},null,8,["value"])]),_:1},8,["label"]),U(_(Ue),{label:"Allow connections to private IPs"},{default:D(()=>[U(_(wa),{value:i.value.allowPrivateIP,"onUpdate:value":f[2]||(f[2]=w=>i.value.allowPrivateIP=w)},null,8,["value"])]),_:1}),U(_(Ue),{label:"Quotas (traffic cap per rolling window)"},{default:D(()=>[U(_(se),{vertical:"",style:{width:"100%"}},{default:D(()=>[(fe(!0),Pt(We,null,Mr(s.value,(w,E)=>(fe(),Ee(_(se),{key:E,align:"center"},{default:D(()=>[U(_(Mt),{value:w.megabytes,"onUpdate:value":d=>w.megabytes=d,min:1,step:1024},{suffix:D(()=>[...f[7]||(f[7]=[J("MB",-1)])]),_:1},8,["value","onUpdate:value"]),f[10]||(f[10]=qe("span",null,"per",-1)),U(_(Mt),{value:w.days,"onUpdate:value":d=>w.days=d,min:1,max:365},{suffix:D(()=>[...f[8]||(f[8]=[J("days",-1)])]),_:1},8,["value","onUpdate:value"]),U(_(Y),{size:"tiny",quaternary:"",type:"error",onClick:d=>s.value.splice(E,1)},{default:D(()=>[...f[9]||(f[9]=[J("✕",-1)])]),_:1},8,["onClick"])]),_:2},1024))),128)),U(_(Y),{size:"small",dashed:"",onClick:P},{default:D(()=>[...f[11]||(f[11]=[J("Add quota",-1)])]),_:1})]),_:1})]),_:1})]),_:1})]),_:1},8,["title"])]),_:1},8,["show"]),l.value?(fe(),Ee(Ya,{key:0,username:l.value,onClose:f[5]||(f[5]=w=>l.value=null)},null,8,["username"])):Bt("",!0)],64))}});export{uo as default};
