import{d as A,h as a,i as to,s as ro,j as no,k as Re,l as oe,m as v,n as I,S as He,p as ie,q,t as Be,v as ae,x as b,r as M,y as J,z as d,A as w,N as Me,C as G,D as F,E as le,F as ee,G as lo,H as Z,I as xe,J as fe,K as io,L as ce,M as ao,V as so,O as Ne,P as co,Q as uo,R as vo,o as K,c as Y,T as ho,a as R,U as mo,W as fo,X as Pe,w as E,u as B,b as L,g as de,Y as ue,Z as po,B as go,_ as bo,$ as X,a0 as xo,f as Co}from"./index-B5SBiCPF.js";import{f as ve,u as pe,a as he}from"./client-DbLYzOci.js";import{C as yo,N as wo,a as je,b as Te}from"./Dropdown-CHh3TWO-.js";import{V as zo,c as me,N as ko}from"./Tag-pmWKzwFY.js";import{u as Io,N as So}from"./Space-C-IUYuEv.js";import{N as Ro}from"./text-BSHOKGhy.js";import{_ as No}from"./_plugin-vue_export-helper-DlAUqK2U.js";const Po=A({name:"ChevronDownFilled",render(){return a("svg",{viewBox:"0 0 16 16",fill:"none",xmlns:"http://www.w3.org/2000/svg"},a("path",{d:"M3.20041 5.73966C3.48226 5.43613 3.95681 5.41856 4.26034 5.70041L8 9.22652L11.7397 5.70041C12.0432 5.41856 12.5177 5.43613 12.7996 5.73966C13.0815 6.0432 13.0639 6.51775 12.7603 6.7996L8.51034 10.7996C8.22258 11.0668 7.77743 11.0668 7.48967 10.7996L3.23966 6.7996C2.93613 6.51775 2.91856 6.0432 3.20041 5.73966Z",fill:"currentColor"}))}});function To(e){const{baseColor:t,textColor2:o,bodyColor:s,cardColor:i,dividerColor:l,actionColor:h,scrollbarColor:u,scrollbarColorHover:c,invertedColor:g}=e;return{textColor:o,textColorInverted:"#FFF",color:s,colorEmbedded:h,headerColor:i,headerColorInverted:g,footerColor:h,footerColorInverted:g,headerBorderColor:l,headerBorderColorInverted:g,footerBorderColor:l,footerBorderColorInverted:g,siderBorderColor:l,siderBorderColorInverted:g,siderColor:i,siderColorInverted:g,siderToggleButtonBorder:`1px solid ${l}`,siderToggleButtonColor:t,siderToggleButtonIconColor:o,siderToggleButtonIconColorInverted:o,siderToggleBarColor:Re(s,u),siderToggleBarColorHover:Re(s,c),__invertScrollbar:"true"}}const Ce=to({name:"Layout",common:no,peers:{Scrollbar:ro},self:To}),$e=oe("n-layout-sider"),ye={type:String,default:"static"},_o=v("layout",`
 color: var(--n-text-color);
 background-color: var(--n-color);
 box-sizing: border-box;
 position: relative;
 z-index: auto;
 flex: auto;
 overflow: hidden;
 transition:
 box-shadow .3s var(--n-bezier),
 background-color .3s var(--n-bezier),
 color .3s var(--n-bezier);
`,[v("layout-scroll-container",`
 overflow-x: hidden;
 box-sizing: border-box;
 height: 100%;
 `),I("absolute-positioned",`
 position: absolute;
 left: 0;
 right: 0;
 top: 0;
 bottom: 0;
 `)]),Ao={embedded:Boolean,position:ye,nativeScrollbar:{type:Boolean,default:!0},scrollbarProps:Object,onScroll:Function,contentClass:String,contentStyle:{type:[String,Object],default:""},hasSider:Boolean,siderPlacement:{type:String,default:"left"}},Ee=oe("n-layout");function Le(e){return A({name:e?"LayoutContent":"Layout",props:Object.assign(Object.assign({},q.props),Ao),setup(t){const o=M(null),s=M(null),{mergedClsPrefixRef:i,inlineThemeDisabled:l}=ie(t),h=q("Layout","-layout",_o,Ce,t,i);function u(k,S){if(t.nativeScrollbar){const{value:O}=o;O&&(S===void 0?O.scrollTo(k):O.scrollTo(k,S))}else{const{value:O}=s;O&&O.scrollTo(k,S)}}J(Ee,t);let c=0,g=0;const H=k=>{var S;const O=k.target;c=O.scrollLeft,g=O.scrollTop,(S=t.onScroll)===null||S===void 0||S.call(t,k)};Be(()=>{if(t.nativeScrollbar){const k=o.value;k&&(k.scrollTop=g,k.scrollLeft=c)}});const T={display:"flex",flexWrap:"nowrap",width:"100%",flexDirection:"row"},f={scrollTo:u},z=b(()=>{const{common:{cubicBezierEaseInOut:k},self:S}=h.value;return{"--n-bezier":k,"--n-color":t.embedded?S.colorEmbedded:S.color,"--n-text-color":S.textColor}}),N=l?ae("layout",b(()=>t.embedded?"e":""),z,t):void 0;return Object.assign({mergedClsPrefix:i,scrollableElRef:o,scrollbarInstRef:s,hasSiderStyle:T,mergedTheme:h,handleNativeElScroll:H,cssVars:l?void 0:z,themeClass:N==null?void 0:N.themeClass,onRender:N==null?void 0:N.onRender},f)},render(){var t;const{mergedClsPrefix:o,hasSider:s}=this;(t=this.onRender)===null||t===void 0||t.call(this);const i=s?this.hasSiderStyle:void 0,l=[this.themeClass,e&&`${o}-layout-content`,`${o}-layout`,`${o}-layout--${this.position}-positioned`];return a("div",{class:l,style:this.cssVars},this.nativeScrollbar?a("div",{ref:"scrollableElRef",class:[`${o}-layout-scroll-container`,this.contentClass],style:[this.contentStyle,i],onScroll:this.handleNativeElScroll},this.$slots):a(He,Object.assign({},this.scrollbarProps,{onScroll:this.onScroll,ref:"scrollbarInstRef",theme:this.mergedTheme.peers.Scrollbar,themeOverrides:this.mergedTheme.peerOverrides.Scrollbar,contentClass:this.contentClass,contentStyle:[this.contentStyle,i]}),this.$slots))}})}const _e=Le(!1),Oo=Le(!0),Ho=v("layout-header",`
 transition:
 color .3s var(--n-bezier),
 background-color .3s var(--n-bezier),
 box-shadow .3s var(--n-bezier),
 border-color .3s var(--n-bezier);
 box-sizing: border-box;
 width: 100%;
 background-color: var(--n-color);
 color: var(--n-text-color);
`,[I("absolute-positioned",`
 position: absolute;
 left: 0;
 right: 0;
 top: 0;
 `),I("bordered",`
 border-bottom: solid 1px var(--n-border-color);
 `)]),Bo={position:ye,inverted:Boolean,bordered:{type:Boolean,default:!1}},Mo=A({name:"LayoutHeader",props:Object.assign(Object.assign({},q.props),Bo),setup(e){const{mergedClsPrefixRef:t,inlineThemeDisabled:o}=ie(e),s=q("Layout","-layout-header",Ho,Ce,e,t),i=b(()=>{const{common:{cubicBezierEaseInOut:h},self:u}=s.value,c={"--n-bezier":h};return e.inverted?(c["--n-color"]=u.headerColorInverted,c["--n-text-color"]=u.textColorInverted,c["--n-border-color"]=u.headerBorderColorInverted):(c["--n-color"]=u.headerColor,c["--n-text-color"]=u.textColor,c["--n-border-color"]=u.headerBorderColor),c}),l=o?ae("layout-header",b(()=>e.inverted?"a":"b"),i,e):void 0;return{mergedClsPrefix:t,cssVars:o?void 0:i,themeClass:l==null?void 0:l.themeClass,onRender:l==null?void 0:l.onRender}},render(){var e;const{mergedClsPrefix:t}=this;return(e=this.onRender)===null||e===void 0||e.call(this),a("div",{class:[`${t}-layout-header`,this.themeClass,this.position&&`${t}-layout-header--${this.position}-positioned`,this.bordered&&`${t}-layout-header--bordered`],style:this.cssVars},this.$slots)}}),jo=v("layout-sider",`
 flex-shrink: 0;
 box-sizing: border-box;
 position: relative;
 z-index: 1;
 color: var(--n-text-color);
 transition:
 color .3s var(--n-bezier),
 border-color .3s var(--n-bezier),
 min-width .3s var(--n-bezier),
 max-width .3s var(--n-bezier),
 transform .3s var(--n-bezier),
 background-color .3s var(--n-bezier);
 background-color: var(--n-color);
 display: flex;
 justify-content: flex-end;
`,[I("bordered",[d("border",`
 content: "";
 position: absolute;
 top: 0;
 bottom: 0;
 width: 1px;
 background-color: var(--n-border-color);
 transition: background-color .3s var(--n-bezier);
 `)]),d("left-placement",[I("bordered",[d("border",`
 right: 0;
 `)])]),I("right-placement",`
 justify-content: flex-start;
 `,[I("bordered",[d("border",`
 left: 0;
 `)]),I("collapsed",[v("layout-toggle-button",[v("base-icon",`
 transform: rotate(180deg);
 `)]),v("layout-toggle-bar",[w("&:hover",[d("top",{transform:"rotate(-12deg) scale(1.15) translateY(-2px)"}),d("bottom",{transform:"rotate(12deg) scale(1.15) translateY(2px)"})])])]),v("layout-toggle-button",`
 left: 0;
 transform: translateX(-50%) translateY(-50%);
 `,[v("base-icon",`
 transform: rotate(0);
 `)]),v("layout-toggle-bar",`
 left: -28px;
 transform: rotate(180deg);
 `,[w("&:hover",[d("top",{transform:"rotate(12deg) scale(1.15) translateY(-2px)"}),d("bottom",{transform:"rotate(-12deg) scale(1.15) translateY(2px)"})])])]),I("collapsed",[v("layout-toggle-bar",[w("&:hover",[d("top",{transform:"rotate(-12deg) scale(1.15) translateY(-2px)"}),d("bottom",{transform:"rotate(12deg) scale(1.15) translateY(2px)"})])]),v("layout-toggle-button",[v("base-icon",`
 transform: rotate(0);
 `)])]),v("layout-toggle-button",`
 transition:
 color .3s var(--n-bezier),
 right .3s var(--n-bezier),
 left .3s var(--n-bezier),
 border-color .3s var(--n-bezier),
 background-color .3s var(--n-bezier);
 cursor: pointer;
 width: 24px;
 height: 24px;
 position: absolute;
 top: 50%;
 right: 0;
 border-radius: 50%;
 display: flex;
 align-items: center;
 justify-content: center;
 font-size: 18px;
 color: var(--n-toggle-button-icon-color);
 border: var(--n-toggle-button-border);
 background-color: var(--n-toggle-button-color);
 box-shadow: 0 2px 4px 0px rgba(0, 0, 0, .06);
 transform: translateX(50%) translateY(-50%);
 z-index: 1;
 `,[v("base-icon",`
 transition: transform .3s var(--n-bezier);
 transform: rotate(180deg);
 `)]),v("layout-toggle-bar",`
 cursor: pointer;
 height: 72px;
 width: 32px;
 position: absolute;
 top: calc(50% - 36px);
 right: -28px;
 `,[d("top, bottom",`
 position: absolute;
 width: 4px;
 border-radius: 2px;
 height: 38px;
 left: 14px;
 transition: 
 background-color .3s var(--n-bezier),
 transform .3s var(--n-bezier);
 `),d("bottom",`
 position: absolute;
 top: 34px;
 `),w("&:hover",[d("top",{transform:"rotate(12deg) scale(1.15) translateY(-2px)"}),d("bottom",{transform:"rotate(-12deg) scale(1.15) translateY(2px)"})]),d("top, bottom",{backgroundColor:"var(--n-toggle-bar-color)"}),w("&:hover",[d("top, bottom",{backgroundColor:"var(--n-toggle-bar-color-hover)"})])]),d("border",`
 position: absolute;
 top: 0;
 right: 0;
 bottom: 0;
 width: 1px;
 transition: background-color .3s var(--n-bezier);
 `),v("layout-sider-scroll-container",`
 flex-grow: 1;
 flex-shrink: 0;
 box-sizing: border-box;
 height: 100%;
 opacity: 0;
 transition: opacity .3s var(--n-bezier);
 max-width: 100%;
 `),I("show-content",[v("layout-sider-scroll-container",{opacity:1})]),I("absolute-positioned",`
 position: absolute;
 left: 0;
 top: 0;
 bottom: 0;
 `)]),$o=A({props:{clsPrefix:{type:String,required:!0},onClick:Function},render(){const{clsPrefix:e}=this;return a("div",{onClick:this.onClick,class:`${e}-layout-toggle-bar`},a("div",{class:`${e}-layout-toggle-bar__top`}),a("div",{class:`${e}-layout-toggle-bar__bottom`}))}}),Eo=A({name:"LayoutToggleButton",props:{clsPrefix:{type:String,required:!0},onClick:Function},render(){const{clsPrefix:e}=this;return a("div",{class:`${e}-layout-toggle-button`,onClick:this.onClick},a(Me,{clsPrefix:e},{default:()=>a(yo,null)}))}}),Lo={position:ye,bordered:Boolean,collapsedWidth:{type:Number,default:48},width:{type:[Number,String],default:272},contentClass:String,contentStyle:{type:[String,Object],default:""},collapseMode:{type:String,default:"transform"},collapsed:{type:Boolean,default:void 0},defaultCollapsed:Boolean,showCollapsedContent:{type:Boolean,default:!0},showTrigger:{type:[Boolean,String],default:!1},nativeScrollbar:{type:Boolean,default:!0},inverted:Boolean,scrollbarProps:Object,triggerClass:String,triggerStyle:[String,Object],collapsedTriggerClass:String,collapsedTriggerStyle:[String,Object],"onUpdate:collapsed":[Function,Array],onUpdateCollapsed:[Function,Array],onAfterEnter:Function,onAfterLeave:Function,onExpand:[Function,Array],onCollapse:[Function,Array],onScroll:Function},Fo=A({name:"LayoutSider",props:Object.assign(Object.assign({},q.props),Lo),setup(e){const t=G(Ee),o=M(null),s=M(null),i=M(e.defaultCollapsed),l=pe(le(e,"collapsed"),i),h=b(()=>ve(l.value?e.collapsedWidth:e.width)),u=b(()=>e.collapseMode!=="transform"?{}:{minWidth:ve(e.width)}),c=b(()=>t?t.siderPlacement:"left");function g(_,C){if(e.nativeScrollbar){const{value:y}=o;y&&(C===void 0?y.scrollTo(_):y.scrollTo(_,C))}else{const{value:y}=s;y&&y.scrollTo(_,C)}}function H(){const{"onUpdate:collapsed":_,onUpdateCollapsed:C,onExpand:y,onCollapse:U}=e,{value:V}=l;C&&F(C,!V),_&&F(_,!V),i.value=!V,V?y&&F(y):U&&F(U)}let T=0,f=0;const z=_=>{var C;const y=_.target;T=y.scrollLeft,f=y.scrollTop,(C=e.onScroll)===null||C===void 0||C.call(e,_)};Be(()=>{if(e.nativeScrollbar){const _=o.value;_&&(_.scrollTop=f,_.scrollLeft=T)}}),J($e,{collapsedRef:l,collapseModeRef:le(e,"collapseMode")});const{mergedClsPrefixRef:N,inlineThemeDisabled:k}=ie(e),S=q("Layout","-layout-sider",jo,Ce,e,N);function O(_){var C,y;_.propertyName==="max-width"&&(l.value?(C=e.onAfterLeave)===null||C===void 0||C.call(e):(y=e.onAfterEnter)===null||y===void 0||y.call(e))}const Q={scrollTo:g},D=b(()=>{const{common:{cubicBezierEaseInOut:_},self:C}=S.value,{siderToggleButtonColor:y,siderToggleButtonBorder:U,siderToggleBarColor:V,siderToggleBarColorHover:se}=C,j={"--n-bezier":_,"--n-toggle-button-color":y,"--n-toggle-button-border":U,"--n-toggle-bar-color":V,"--n-toggle-bar-color-hover":se};return e.inverted?(j["--n-color"]=C.siderColorInverted,j["--n-text-color"]=C.textColorInverted,j["--n-border-color"]=C.siderBorderColorInverted,j["--n-toggle-button-icon-color"]=C.siderToggleButtonIconColorInverted,j.__invertScrollbar=C.__invertScrollbar):(j["--n-color"]=C.siderColor,j["--n-text-color"]=C.textColor,j["--n-border-color"]=C.siderBorderColor,j["--n-toggle-button-icon-color"]=C.siderToggleButtonIconColor),j}),$=k?ae("layout-sider",b(()=>e.inverted?"a":"b"),D,e):void 0;return Object.assign({scrollableElRef:o,scrollbarInstRef:s,mergedClsPrefix:N,mergedTheme:S,styleMaxWidth:h,mergedCollapsed:l,scrollContainerStyle:u,siderPlacement:c,handleNativeElScroll:z,handleTransitionend:O,handleTriggerClick:H,inlineThemeDisabled:k,cssVars:D,themeClass:$==null?void 0:$.themeClass,onRender:$==null?void 0:$.onRender},Q)},render(){var e;const{mergedClsPrefix:t,mergedCollapsed:o,showTrigger:s}=this;return(e=this.onRender)===null||e===void 0||e.call(this),a("aside",{class:[`${t}-layout-sider`,this.themeClass,`${t}-layout-sider--${this.position}-positioned`,`${t}-layout-sider--${this.siderPlacement}-placement`,this.bordered&&`${t}-layout-sider--bordered`,o&&`${t}-layout-sider--collapsed`,(!o||this.showCollapsedContent)&&`${t}-layout-sider--show-content`],onTransitionend:this.handleTransitionend,style:[this.inlineThemeDisabled?void 0:this.cssVars,{maxWidth:this.styleMaxWidth,width:ve(this.width)}]},this.nativeScrollbar?a("div",{class:[`${t}-layout-sider-scroll-container`,this.contentClass],onScroll:this.handleNativeElScroll,style:[this.scrollContainerStyle,{overflow:"auto"},this.contentStyle],ref:"scrollableElRef"},this.$slots):a(He,Object.assign({},this.scrollbarProps,{onScroll:this.onScroll,ref:"scrollbarInstRef",style:this.scrollContainerStyle,contentStyle:this.contentStyle,contentClass:this.contentClass,theme:this.mergedTheme.peers.Scrollbar,themeOverrides:this.mergedTheme.peerOverrides.Scrollbar,builtinThemeOverrides:this.inverted&&this.cssVars.__invertScrollbar==="true"?{colorHover:"rgba(255, 255, 255, .4)",color:"rgba(255, 255, 255, .3)"}:void 0}),this.$slots),s?s==="bar"?a($o,{clsPrefix:t,class:o?this.collapsedTriggerClass:this.triggerClass,style:o?this.collapsedTriggerStyle:this.triggerStyle,onClick:this.handleTriggerClick}):a(Eo,{clsPrefix:t,class:o?this.collapsedTriggerClass:this.triggerClass,style:o?this.collapsedTriggerStyle:this.triggerStyle,onClick:this.handleTriggerClick}):null,this.bordered?a("div",{class:`${t}-layout-sider__border`}):null)}}),te=oe("n-menu"),Fe=oe("n-submenu"),we=oe("n-menu-item-group"),Ae=[w("&::before","background-color: var(--n-item-color-hover);"),d("arrow",`
 color: var(--n-arrow-color-hover);
 `),d("icon",`
 color: var(--n-item-icon-color-hover);
 `),v("menu-item-content-header",`
 color: var(--n-item-text-color-hover);
 `,[w("a",`
 color: var(--n-item-text-color-hover);
 `),d("extra",`
 color: var(--n-item-text-color-hover);
 `)])],Oe=[d("icon",`
 color: var(--n-item-icon-color-hover-horizontal);
 `),v("menu-item-content-header",`
 color: var(--n-item-text-color-hover-horizontal);
 `,[w("a",`
 color: var(--n-item-text-color-hover-horizontal);
 `),d("extra",`
 color: var(--n-item-text-color-hover-horizontal);
 `)])],Ko=w([v("menu",`
 background-color: var(--n-color);
 color: var(--n-item-text-color);
 overflow: hidden;
 transition: background-color .3s var(--n-bezier);
 box-sizing: border-box;
 font-size: var(--n-font-size);
 padding-bottom: 6px;
 `,[I("horizontal",`
 max-width: 100%;
 width: 100%;
 display: flex;
 overflow: hidden;
 padding-bottom: 0;
 `,[v("submenu","margin: 0;"),v("menu-item","margin: 0;"),v("menu-item-content",`
 padding: 0 20px;
 border-bottom: 2px solid #0000;
 `,[w("&::before","display: none;"),I("selected","border-bottom: 2px solid var(--n-border-color-horizontal)")]),v("menu-item-content",[I("selected",[d("icon","color: var(--n-item-icon-color-active-horizontal);"),v("menu-item-content-header",`
 color: var(--n-item-text-color-active-horizontal);
 `,[w("a","color: var(--n-item-text-color-active-horizontal);"),d("extra","color: var(--n-item-text-color-active-horizontal);")])]),I("child-active",`
 border-bottom: 2px solid var(--n-border-color-horizontal);
 `,[v("menu-item-content-header",`
 color: var(--n-item-text-color-child-active-horizontal);
 `,[w("a",`
 color: var(--n-item-text-color-child-active-horizontal);
 `),d("extra",`
 color: var(--n-item-text-color-child-active-horizontal);
 `)]),d("icon",`
 color: var(--n-item-icon-color-child-active-horizontal);
 `)]),ee("disabled",[ee("selected, child-active",[w("&:focus-within",Oe)]),I("selected",[W(null,[d("icon","color: var(--n-item-icon-color-active-hover-horizontal);"),v("menu-item-content-header",`
 color: var(--n-item-text-color-active-hover-horizontal);
 `,[w("a","color: var(--n-item-text-color-active-hover-horizontal);"),d("extra","color: var(--n-item-text-color-active-hover-horizontal);")])])]),I("child-active",[W(null,[d("icon","color: var(--n-item-icon-color-child-active-hover-horizontal);"),v("menu-item-content-header",`
 color: var(--n-item-text-color-child-active-hover-horizontal);
 `,[w("a","color: var(--n-item-text-color-child-active-hover-horizontal);"),d("extra","color: var(--n-item-text-color-child-active-hover-horizontal);")])])]),W("border-bottom: 2px solid var(--n-border-color-horizontal);",Oe)]),v("menu-item-content-header",[w("a","color: var(--n-item-text-color-horizontal);")])])]),ee("responsive",[v("menu-item-content-header",`
 overflow: hidden;
 text-overflow: ellipsis;
 `)]),I("collapsed",[v("menu-item-content",[I("selected",[w("&::before",`
 background-color: var(--n-item-color-active-collapsed) !important;
 `)]),v("menu-item-content-header","opacity: 0;"),d("arrow","opacity: 0;"),d("icon","color: var(--n-item-icon-color-collapsed);")])]),v("menu-item",`
 height: var(--n-item-height);
 margin-top: 6px;
 position: relative;
 `),v("menu-item-content",`
 box-sizing: border-box;
 line-height: 1.75;
 height: 100%;
 display: grid;
 grid-template-areas: "icon content arrow";
 grid-template-columns: auto 1fr auto;
 align-items: center;
 cursor: pointer;
 position: relative;
 padding-right: 18px;
 transition:
 background-color .3s var(--n-bezier),
 padding-left .3s var(--n-bezier),
 border-color .3s var(--n-bezier);
 `,[w("> *","z-index: 1;"),w("&::before",`
 z-index: auto;
 content: "";
 background-color: #0000;
 position: absolute;
 left: 8px;
 right: 8px;
 top: 0;
 bottom: 0;
 pointer-events: none;
 border-radius: var(--n-border-radius);
 transition: background-color .3s var(--n-bezier);
 `),I("disabled",`
 opacity: .45;
 cursor: not-allowed;
 `),I("collapsed",[d("arrow","transform: rotate(0);")]),I("selected",[w("&::before","background-color: var(--n-item-color-active);"),d("arrow","color: var(--n-arrow-color-active);"),d("icon","color: var(--n-item-icon-color-active);"),v("menu-item-content-header",`
 color: var(--n-item-text-color-active);
 `,[w("a","color: var(--n-item-text-color-active);"),d("extra","color: var(--n-item-text-color-active);")])]),I("child-active",[v("menu-item-content-header",`
 color: var(--n-item-text-color-child-active);
 `,[w("a",`
 color: var(--n-item-text-color-child-active);
 `),d("extra",`
 color: var(--n-item-text-color-child-active);
 `)]),d("arrow",`
 color: var(--n-arrow-color-child-active);
 `),d("icon",`
 color: var(--n-item-icon-color-child-active);
 `)]),ee("disabled",[ee("selected, child-active",[w("&:focus-within",Ae)]),I("selected",[W(null,[d("arrow","color: var(--n-arrow-color-active-hover);"),d("icon","color: var(--n-item-icon-color-active-hover);"),v("menu-item-content-header",`
 color: var(--n-item-text-color-active-hover);
 `,[w("a","color: var(--n-item-text-color-active-hover);"),d("extra","color: var(--n-item-text-color-active-hover);")])])]),I("child-active",[W(null,[d("arrow","color: var(--n-arrow-color-child-active-hover);"),d("icon","color: var(--n-item-icon-color-child-active-hover);"),v("menu-item-content-header",`
 color: var(--n-item-text-color-child-active-hover);
 `,[w("a","color: var(--n-item-text-color-child-active-hover);"),d("extra","color: var(--n-item-text-color-child-active-hover);")])])]),I("selected",[W(null,[w("&::before","background-color: var(--n-item-color-active-hover);")])]),W(null,Ae)]),d("icon",`
 grid-area: icon;
 color: var(--n-item-icon-color);
 transition:
 color .3s var(--n-bezier),
 font-size .3s var(--n-bezier),
 margin-right .3s var(--n-bezier);
 box-sizing: content-box;
 display: inline-flex;
 align-items: center;
 justify-content: center;
 `),d("arrow",`
 grid-area: arrow;
 font-size: 16px;
 color: var(--n-arrow-color);
 transform: rotate(180deg);
 opacity: 1;
 transition:
 color .3s var(--n-bezier),
 transform 0.2s var(--n-bezier),
 opacity 0.2s var(--n-bezier);
 `),v("menu-item-content-header",`
 grid-area: content;
 transition:
 color .3s var(--n-bezier),
 opacity .3s var(--n-bezier);
 opacity: 1;
 white-space: nowrap;
 color: var(--n-item-text-color);
 `,[w("a",`
 outline: none;
 text-decoration: none;
 transition: color .3s var(--n-bezier);
 color: var(--n-item-text-color);
 `,[w("&::before",`
 content: "";
 position: absolute;
 left: 0;
 right: 0;
 top: 0;
 bottom: 0;
 `)]),d("extra",`
 font-size: .93em;
 color: var(--n-group-text-color);
 transition: color .3s var(--n-bezier);
 `)])]),v("submenu",`
 cursor: pointer;
 position: relative;
 margin-top: 6px;
 `,[v("menu-item-content",`
 height: var(--n-item-height);
 `),v("submenu-children",`
 overflow: hidden;
 padding: 0;
 `,[lo({duration:".2s"})])]),v("menu-item-group",[v("menu-item-group-title",`
 margin-top: 6px;
 color: var(--n-group-text-color);
 cursor: default;
 font-size: .93em;
 height: 36px;
 display: flex;
 align-items: center;
 transition:
 padding-left .3s var(--n-bezier),
 color .3s var(--n-bezier);
 `)])]),v("menu-tooltip",[w("a",`
 color: inherit;
 text-decoration: none;
 `)]),v("menu-divider",`
 transition: background-color .3s var(--n-bezier);
 background-color: var(--n-divider-color);
 height: 1px;
 margin: 6px 18px;
 `)]);function W(e,t){return[I("hover",e,t),w("&:hover",e,t)]}const Ke=A({name:"MenuOptionContent",props:{collapsed:Boolean,disabled:Boolean,title:[String,Function],icon:Function,extra:[String,Function],showArrow:Boolean,childActive:Boolean,hover:Boolean,paddingLeft:Number,selected:Boolean,maxIconSize:{type:Number,required:!0},activeIconSize:{type:Number,required:!0},iconMarginRight:{type:Number,required:!0},clsPrefix:{type:String,required:!0},onClick:Function,tmNode:{type:Object,required:!0},isEllipsisPlaceholder:Boolean},setup(e){const{props:t}=G(te);return{menuProps:t,style:b(()=>{const{paddingLeft:o}=e;return{paddingLeft:o&&`${o}px`}}),iconStyle:b(()=>{const{maxIconSize:o,activeIconSize:s,iconMarginRight:i}=e;return{width:`${o}px`,height:`${o}px`,fontSize:`${s}px`,marginRight:`${i}px`}})}},render(){const{clsPrefix:e,tmNode:t,menuProps:{renderIcon:o,renderLabel:s,renderExtra:i,expandIcon:l}}=this,h=o?o(t.rawNode):Z(this.icon);return a("div",{onClick:u=>{var c;(c=this.onClick)===null||c===void 0||c.call(this,u)},role:"none",class:[`${e}-menu-item-content`,{[`${e}-menu-item-content--selected`]:this.selected,[`${e}-menu-item-content--collapsed`]:this.collapsed,[`${e}-menu-item-content--child-active`]:this.childActive,[`${e}-menu-item-content--disabled`]:this.disabled,[`${e}-menu-item-content--hover`]:this.hover}],style:this.style},h&&a("div",{class:`${e}-menu-item-content__icon`,style:this.iconStyle,role:"none"},[h]),a("div",{class:`${e}-menu-item-content-header`,role:"none"},this.isEllipsisPlaceholder?this.title:s?s(t.rawNode):Z(this.title),this.extra||i?a("span",{class:`${e}-menu-item-content-header__extra`}," ",i?i(t.rawNode):Z(this.extra)):null),this.showArrow?a(Me,{ariaHidden:!0,class:`${e}-menu-item-content__arrow`,clsPrefix:e},{default:()=>l?l(t.rawNode):a(Po,null)}):null)}}),ne=8;function ze(e){const t=G(te),{props:o,mergedCollapsedRef:s}=t,i=G(Fe,null),l=G(we,null),h=b(()=>o.mode==="horizontal"),u=b(()=>h.value?o.dropdownPlacement:"tmNodes"in e?"right-start":"right"),c=b(()=>{var f;return Math.max((f=o.collapsedIconSize)!==null&&f!==void 0?f:o.iconSize,o.iconSize)}),g=b(()=>{var f;return!h.value&&e.root&&s.value&&(f=o.collapsedIconSize)!==null&&f!==void 0?f:o.iconSize}),H=b(()=>{if(h.value)return;const{collapsedWidth:f,indent:z,rootIndent:N}=o,{root:k,isGroup:S}=e,O=N===void 0?z:N;return k?s.value?f/2-c.value/2:O:l&&typeof l.paddingLeftRef.value=="number"?z/2+l.paddingLeftRef.value:i&&typeof i.paddingLeftRef.value=="number"?(S?z/2:z)+i.paddingLeftRef.value:0}),T=b(()=>{const{collapsedWidth:f,indent:z,rootIndent:N}=o,{value:k}=c,{root:S}=e;return h.value||!S||!s.value?ne:(N===void 0?z:N)+k+ne-(f+k)/2});return{dropdownPlacement:u,activeIconSize:g,maxIconSize:c,paddingLeft:H,iconMarginRight:T,NMenu:t,NSubmenu:i,NMenuOptionGroup:l}}const ke={internalKey:{type:[String,Number],required:!0},root:Boolean,isGroup:Boolean,level:{type:Number,required:!0},title:[String,Function],extra:[String,Function]},Vo=A({name:"MenuDivider",setup(){const e=G(te),{mergedClsPrefixRef:t,isHorizontalRef:o}=e;return()=>o.value?null:a("div",{class:`${t.value}-menu-divider`})}}),Ve=Object.assign(Object.assign({},ke),{tmNode:{type:Object,required:!0},disabled:Boolean,icon:Function,onClick:Function}),Do=xe(Ve),Uo=A({name:"MenuOption",props:Ve,setup(e){const t=ze(e),{NSubmenu:o,NMenu:s,NMenuOptionGroup:i}=t,{props:l,mergedClsPrefixRef:h,mergedCollapsedRef:u}=s,c=o?o.mergedDisabledRef:i?i.mergedDisabledRef:{value:!1},g=b(()=>c.value||e.disabled);function H(f){const{onClick:z}=e;z&&z(f)}function T(f){g.value||(s.doSelect(e.internalKey,e.tmNode.rawNode),H(f))}return{mergedClsPrefix:h,dropdownPlacement:t.dropdownPlacement,paddingLeft:t.paddingLeft,iconMarginRight:t.iconMarginRight,maxIconSize:t.maxIconSize,activeIconSize:t.activeIconSize,mergedTheme:s.mergedThemeRef,menuProps:l,dropdownEnabled:fe(()=>e.root&&u.value&&l.mode!=="horizontal"&&!g.value),selected:fe(()=>s.mergedValueRef.value===e.internalKey),mergedDisabled:g,handleClick:T}},render(){const{mergedClsPrefix:e,mergedTheme:t,tmNode:o,menuProps:{renderLabel:s,nodeProps:i}}=this,l=i==null?void 0:i(o.rawNode);return a("div",Object.assign({},l,{role:"menuitem",class:[`${e}-menu-item`,l==null?void 0:l.class]}),a(wo,{theme:t.peers.Tooltip,themeOverrides:t.peerOverrides.Tooltip,trigger:"hover",placement:this.dropdownPlacement,disabled:!this.dropdownEnabled||this.title===void 0,internalExtraClass:["menu-tooltip"]},{default:()=>s?s(o.rawNode):Z(this.title),trigger:()=>a(Ke,{tmNode:o,clsPrefix:e,paddingLeft:this.paddingLeft,iconMarginRight:this.iconMarginRight,maxIconSize:this.maxIconSize,activeIconSize:this.activeIconSize,selected:this.selected,title:this.title,extra:this.extra,disabled:this.mergedDisabled,icon:this.icon,onClick:this.handleClick})}))}}),De=Object.assign(Object.assign({},ke),{tmNode:{type:Object,required:!0},tmNodes:{type:Array,required:!0}}),Go=xe(De),qo=A({name:"MenuOptionGroup",props:De,setup(e){const t=ze(e),{NSubmenu:o}=t,s=b(()=>o!=null&&o.mergedDisabledRef.value?!0:e.tmNode.disabled);J(we,{paddingLeftRef:t.paddingLeft,mergedDisabledRef:s});const{mergedClsPrefixRef:i,props:l}=G(te);return function(){const{value:h}=i,u=t.paddingLeft.value,{nodeProps:c}=l,g=c==null?void 0:c(e.tmNode.rawNode);return a("div",{class:`${h}-menu-item-group`,role:"group"},a("div",Object.assign({},g,{class:[`${h}-menu-item-group-title`,g==null?void 0:g.class],style:[(g==null?void 0:g.style)||"",u!==void 0?`padding-left: ${u}px;`:""]}),Z(e.title),e.extra?a(io,null," ",Z(e.extra)):null),a("div",null,e.tmNodes.map(H=>Ie(H,l))))}}});function ge(e){return e.type==="divider"||e.type==="render"}function Yo(e){return e.type==="divider"}function Ie(e,t){const{rawNode:o}=e,{show:s}=o;if(s===!1)return null;if(ge(o))return Yo(o)?a(Vo,Object.assign({key:e.key},o.props)):null;const{labelField:i}=t,{key:l,level:h,isGroup:u}=e,c=Object.assign(Object.assign({},o),{title:o.title||o[i],extra:o.titleExtra||o.extra,key:l,internalKey:l,level:h,root:h===0,isGroup:u});return e.children?e.isGroup?a(qo,ce(c,Go,{tmNode:e,tmNodes:e.children,key:l})):a(be,ce(c,Wo,{key:l,rawNodes:o[t.childrenField],tmNodes:e.children,tmNode:e})):a(Uo,ce(c,Do,{key:l,tmNode:e}))}const Ue=Object.assign(Object.assign({},ke),{rawNodes:{type:Array,default:()=>[]},tmNodes:{type:Array,default:()=>[]},tmNode:{type:Object,required:!0},disabled:Boolean,icon:Function,onClick:Function,domId:String,virtualChildActive:{type:Boolean,default:void 0},isEllipsisPlaceholder:Boolean}),Wo=xe(Ue),be=A({name:"Submenu",props:Ue,setup(e){const t=ze(e),{NMenu:o,NSubmenu:s}=t,{props:i,mergedCollapsedRef:l,mergedThemeRef:h}=o,u=b(()=>{const{disabled:f}=e;return s!=null&&s.mergedDisabledRef.value||i.disabled?!0:f}),c=M(!1);J(Fe,{paddingLeftRef:t.paddingLeft,mergedDisabledRef:u}),J(we,null);function g(){const{onClick:f}=e;f&&f()}function H(){u.value||(l.value||o.toggleExpand(e.internalKey),g())}function T(f){c.value=f}return{menuProps:i,mergedTheme:h,doSelect:o.doSelect,inverted:o.invertedRef,isHorizontal:o.isHorizontalRef,mergedClsPrefix:o.mergedClsPrefixRef,maxIconSize:t.maxIconSize,activeIconSize:t.activeIconSize,iconMarginRight:t.iconMarginRight,dropdownPlacement:t.dropdownPlacement,dropdownShow:c,paddingLeft:t.paddingLeft,mergedDisabled:u,mergedValue:o.mergedValueRef,childActive:fe(()=>{var f;return(f=e.virtualChildActive)!==null&&f!==void 0?f:o.activePathRef.value.includes(e.internalKey)}),collapsed:b(()=>i.mode==="horizontal"?!1:l.value?!0:!o.mergedExpandedKeysRef.value.includes(e.internalKey)),dropdownEnabled:b(()=>!u.value&&(i.mode==="horizontal"||l.value)),handlePopoverShowChange:T,handleClick:H}},render(){var e;const{mergedClsPrefix:t,menuProps:{renderIcon:o,renderLabel:s}}=this,i=()=>{const{isHorizontal:h,paddingLeft:u,collapsed:c,mergedDisabled:g,maxIconSize:H,activeIconSize:T,title:f,childActive:z,icon:N,handleClick:k,menuProps:{nodeProps:S},dropdownShow:O,iconMarginRight:Q,tmNode:D,mergedClsPrefix:$,isEllipsisPlaceholder:_,extra:C}=this,y=S==null?void 0:S(D.rawNode);return a("div",Object.assign({},y,{class:[`${$}-menu-item`,y==null?void 0:y.class],role:"menuitem"}),a(Ke,{tmNode:D,paddingLeft:u,collapsed:c,disabled:g,iconMarginRight:Q,maxIconSize:H,activeIconSize:T,title:f,extra:C,showArrow:!h,childActive:z,clsPrefix:$,icon:N,hover:O,onClick:k,isEllipsisPlaceholder:_}))},l=()=>a(ao,null,{default:()=>{const{tmNodes:h,collapsed:u}=this;return u?null:a("div",{class:`${t}-submenu-children`,role:"menu"},h.map(c=>Ie(c,this.menuProps)))}});return this.root?a(je,Object.assign({size:"large",trigger:"hover"},(e=this.menuProps)===null||e===void 0?void 0:e.dropdownProps,{themeOverrides:this.mergedTheme.peerOverrides.Dropdown,theme:this.mergedTheme.peers.Dropdown,builtinThemeOverrides:{fontSizeLarge:"14px",optionIconSizeLarge:"18px"},value:this.mergedValue,disabled:!this.dropdownEnabled,placement:this.dropdownPlacement,keyField:this.menuProps.keyField,labelField:this.menuProps.labelField,childrenField:this.menuProps.childrenField,onUpdateShow:this.handlePopoverShowChange,options:this.rawNodes,onSelect:this.doSelect,inverted:this.inverted,renderIcon:o,renderLabel:s}),{default:()=>a("div",{class:`${t}-submenu`,role:"menu","aria-expanded":!this.collapsed,id:this.domId},i(),this.isHorizontal?null:l())}):a("div",{class:`${t}-submenu`,role:"menu","aria-expanded":!this.collapsed,id:this.domId},i(),l())}}),Xo=Object.assign(Object.assign({},q.props),{options:{type:Array,default:()=>[]},collapsed:{type:Boolean,default:void 0},collapsedWidth:{type:Number,default:48},iconSize:{type:Number,default:20},collapsedIconSize:{type:Number,default:24},rootIndent:Number,indent:{type:Number,default:32},labelField:{type:String,default:"label"},keyField:{type:String,default:"key"},childrenField:{type:String,default:"children"},disabledField:{type:String,default:"disabled"},defaultExpandAll:Boolean,defaultExpandedKeys:Array,expandedKeys:Array,value:[String,Number],defaultValue:{type:[String,Number],default:null},mode:{type:String,default:"vertical"},watchProps:{type:Array,default:void 0},disabled:Boolean,show:{type:Boolean,default:!0},inverted:Boolean,"onUpdate:expandedKeys":[Function,Array],onUpdateExpandedKeys:[Function,Array],onUpdateValue:[Function,Array],"onUpdate:value":[Function,Array],expandIcon:Function,renderIcon:Function,renderLabel:Function,renderExtra:Function,dropdownProps:Object,accordion:Boolean,nodeProps:Function,dropdownPlacement:{type:String,default:"bottom"},responsive:Boolean,items:Array,onOpenNamesChange:[Function,Array],onSelect:[Function,Array],onExpandedNamesChange:[Function,Array],expandedNames:Array,defaultExpandedNames:Array}),Zo=A({name:"Menu",inheritAttrs:!1,props:Xo,setup(e){const{mergedClsPrefixRef:t,inlineThemeDisabled:o}=ie(e),s=q("Menu","-menu",Ko,vo,e,t),i=G($e,null),l=b(()=>{var m;const{collapsed:x}=e;if(x!==void 0)return x;if(i){const{collapseModeRef:r,collapsedRef:p}=i;if(r.value==="width")return(m=p.value)!==null&&m!==void 0?m:!1}return!1}),h=b(()=>{const{keyField:m,childrenField:x,disabledField:r}=e;return me(e.items||e.options,{getIgnored(p){return ge(p)},getChildren(p){return p[x]},getDisabled(p){return p[r]},getKey(p){var P;return(P=p[m])!==null&&P!==void 0?P:p.name}})}),u=b(()=>new Set(h.value.treeNodes.map(m=>m.key))),{watchProps:c}=e,g=M(null);c!=null&&c.includes("defaultValue")?Ne(()=>{g.value=e.defaultValue}):g.value=e.defaultValue;const H=le(e,"value"),T=pe(H,g),f=M([]),z=()=>{f.value=e.defaultExpandAll?h.value.getNonLeafKeys():e.defaultExpandedNames||e.defaultExpandedKeys||h.value.getPath(T.value,{includeSelf:!1}).keyPath};c!=null&&c.includes("defaultExpandedKeys")?Ne(z):z();const N=Io(e,["expandedNames","expandedKeys"]),k=pe(N,f),S=b(()=>h.value.treeNodes),O=b(()=>h.value.getPath(T.value).keyPath);J(te,{props:e,mergedCollapsedRef:l,mergedThemeRef:s,mergedValueRef:T,mergedExpandedKeysRef:k,activePathRef:O,mergedClsPrefixRef:t,isHorizontalRef:b(()=>e.mode==="horizontal"),invertedRef:le(e,"inverted"),doSelect:Q,toggleExpand:$});function Q(m,x){const{"onUpdate:value":r,onUpdateValue:p,onSelect:P}=e;p&&F(p,m,x),r&&F(r,m,x),P&&F(P,m,x),g.value=m}function D(m){const{"onUpdate:expandedKeys":x,onUpdateExpandedKeys:r,onExpandedNamesChange:p,onOpenNamesChange:P}=e;x&&F(x,m),r&&F(r,m),p&&F(p,m),P&&F(P,m),f.value=m}function $(m){const x=Array.from(k.value),r=x.findIndex(p=>p===m);if(~r)x.splice(r,1);else{if(e.accordion&&u.value.has(m)){const p=x.findIndex(P=>u.value.has(P));p>-1&&x.splice(p,1)}x.push(m)}D(x)}const _=m=>{const x=h.value.getPath(m??T.value,{includeSelf:!1}).keyPath;if(!x.length)return;const r=Array.from(k.value),p=new Set([...r,...x]);e.accordion&&u.value.forEach(P=>{p.has(P)&&!x.includes(P)&&p.delete(P)}),D(Array.from(p))},C=b(()=>{const{inverted:m}=e,{common:{cubicBezierEaseInOut:x},self:r}=s.value,{borderRadius:p,borderColorHorizontal:P,fontSize:Qe,itemHeight:eo,dividerColor:oo}=r,n={"--n-divider-color":oo,"--n-bezier":x,"--n-font-size":Qe,"--n-border-color-horizontal":P,"--n-border-radius":p,"--n-item-height":eo};return m?(n["--n-group-text-color"]=r.groupTextColorInverted,n["--n-color"]=r.colorInverted,n["--n-item-text-color"]=r.itemTextColorInverted,n["--n-item-text-color-hover"]=r.itemTextColorHoverInverted,n["--n-item-text-color-active"]=r.itemTextColorActiveInverted,n["--n-item-text-color-child-active"]=r.itemTextColorChildActiveInverted,n["--n-item-text-color-child-active-hover"]=r.itemTextColorChildActiveInverted,n["--n-item-text-color-active-hover"]=r.itemTextColorActiveHoverInverted,n["--n-item-icon-color"]=r.itemIconColorInverted,n["--n-item-icon-color-hover"]=r.itemIconColorHoverInverted,n["--n-item-icon-color-active"]=r.itemIconColorActiveInverted,n["--n-item-icon-color-active-hover"]=r.itemIconColorActiveHoverInverted,n["--n-item-icon-color-child-active"]=r.itemIconColorChildActiveInverted,n["--n-item-icon-color-child-active-hover"]=r.itemIconColorChildActiveHoverInverted,n["--n-item-icon-color-collapsed"]=r.itemIconColorCollapsedInverted,n["--n-item-text-color-horizontal"]=r.itemTextColorHorizontalInverted,n["--n-item-text-color-hover-horizontal"]=r.itemTextColorHoverHorizontalInverted,n["--n-item-text-color-active-horizontal"]=r.itemTextColorActiveHorizontalInverted,n["--n-item-text-color-child-active-horizontal"]=r.itemTextColorChildActiveHorizontalInverted,n["--n-item-text-color-child-active-hover-horizontal"]=r.itemTextColorChildActiveHoverHorizontalInverted,n["--n-item-text-color-active-hover-horizontal"]=r.itemTextColorActiveHoverHorizontalInverted,n["--n-item-icon-color-horizontal"]=r.itemIconColorHorizontalInverted,n["--n-item-icon-color-hover-horizontal"]=r.itemIconColorHoverHorizontalInverted,n["--n-item-icon-color-active-horizontal"]=r.itemIconColorActiveHorizontalInverted,n["--n-item-icon-color-active-hover-horizontal"]=r.itemIconColorActiveHoverHorizontalInverted,n["--n-item-icon-color-child-active-horizontal"]=r.itemIconColorChildActiveHorizontalInverted,n["--n-item-icon-color-child-active-hover-horizontal"]=r.itemIconColorChildActiveHoverHorizontalInverted,n["--n-arrow-color"]=r.arrowColorInverted,n["--n-arrow-color-hover"]=r.arrowColorHoverInverted,n["--n-arrow-color-active"]=r.arrowColorActiveInverted,n["--n-arrow-color-active-hover"]=r.arrowColorActiveHoverInverted,n["--n-arrow-color-child-active"]=r.arrowColorChildActiveInverted,n["--n-arrow-color-child-active-hover"]=r.arrowColorChildActiveHoverInverted,n["--n-item-color-hover"]=r.itemColorHoverInverted,n["--n-item-color-active"]=r.itemColorActiveInverted,n["--n-item-color-active-hover"]=r.itemColorActiveHoverInverted,n["--n-item-color-active-collapsed"]=r.itemColorActiveCollapsedInverted):(n["--n-group-text-color"]=r.groupTextColor,n["--n-color"]=r.color,n["--n-item-text-color"]=r.itemTextColor,n["--n-item-text-color-hover"]=r.itemTextColorHover,n["--n-item-text-color-active"]=r.itemTextColorActive,n["--n-item-text-color-child-active"]=r.itemTextColorChildActive,n["--n-item-text-color-child-active-hover"]=r.itemTextColorChildActiveHover,n["--n-item-text-color-active-hover"]=r.itemTextColorActiveHover,n["--n-item-icon-color"]=r.itemIconColor,n["--n-item-icon-color-hover"]=r.itemIconColorHover,n["--n-item-icon-color-active"]=r.itemIconColorActive,n["--n-item-icon-color-active-hover"]=r.itemIconColorActiveHover,n["--n-item-icon-color-child-active"]=r.itemIconColorChildActive,n["--n-item-icon-color-child-active-hover"]=r.itemIconColorChildActiveHover,n["--n-item-icon-color-collapsed"]=r.itemIconColorCollapsed,n["--n-item-text-color-horizontal"]=r.itemTextColorHorizontal,n["--n-item-text-color-hover-horizontal"]=r.itemTextColorHoverHorizontal,n["--n-item-text-color-active-horizontal"]=r.itemTextColorActiveHorizontal,n["--n-item-text-color-child-active-horizontal"]=r.itemTextColorChildActiveHorizontal,n["--n-item-text-color-child-active-hover-horizontal"]=r.itemTextColorChildActiveHoverHorizontal,n["--n-item-text-color-active-hover-horizontal"]=r.itemTextColorActiveHoverHorizontal,n["--n-item-icon-color-horizontal"]=r.itemIconColorHorizontal,n["--n-item-icon-color-hover-horizontal"]=r.itemIconColorHoverHorizontal,n["--n-item-icon-color-active-horizontal"]=r.itemIconColorActiveHorizontal,n["--n-item-icon-color-active-hover-horizontal"]=r.itemIconColorActiveHoverHorizontal,n["--n-item-icon-color-child-active-horizontal"]=r.itemIconColorChildActiveHorizontal,n["--n-item-icon-color-child-active-hover-horizontal"]=r.itemIconColorChildActiveHoverHorizontal,n["--n-arrow-color"]=r.arrowColor,n["--n-arrow-color-hover"]=r.arrowColorHover,n["--n-arrow-color-active"]=r.arrowColorActive,n["--n-arrow-color-active-hover"]=r.arrowColorActiveHover,n["--n-arrow-color-child-active"]=r.arrowColorChildActive,n["--n-arrow-color-child-active-hover"]=r.arrowColorChildActiveHover,n["--n-item-color-hover"]=r.itemColorHover,n["--n-item-color-active"]=r.itemColorActive,n["--n-item-color-active-hover"]=r.itemColorActiveHover,n["--n-item-color-active-collapsed"]=r.itemColorActiveCollapsed),n}),y=o?ae("menu",b(()=>e.inverted?"a":"b"),C,e):void 0,U=co(),V=M(null),se=M(null);let j=!0;const Se=()=>{var m;j?j=!1:(m=V.value)===null||m===void 0||m.sync({showAllItemsBeforeCalculate:!0})};function Ge(){return document.getElementById(U)}const re=M(-1);function qe(m){re.value=e.options.length-m}function Ye(m){m||(re.value=-1)}const We=b(()=>{const m=re.value;return{children:m===-1?[]:e.options.slice(m)}}),Xe=b(()=>{const{childrenField:m,disabledField:x,keyField:r}=e;return me([We.value],{getIgnored(p){return ge(p)},getChildren(p){return p[m]},getDisabled(p){return p[x]},getKey(p){var P;return(P=p[r])!==null&&P!==void 0?P:p.name}})}),Ze=b(()=>me([{}]).treeNodes[0]);function Je(){var m;if(re.value===-1)return a(be,{root:!0,level:0,key:"__ellpisisGroupPlaceholder__",internalKey:"__ellpisisGroupPlaceholder__",title:"···",tmNode:Ze.value,domId:U,isEllipsisPlaceholder:!0});const x=Xe.value.treeNodes[0],r=O.value,p=!!(!((m=x.children)===null||m===void 0)&&m.some(P=>r.includes(P.key)));return a(be,{level:0,root:!0,key:"__ellpisisGroup__",internalKey:"__ellpisisGroup__",title:"···",virtualChildActive:p,tmNode:x,domId:U,rawNodes:x.rawNode.children||[],tmNodes:x.children||[],isEllipsisPlaceholder:!0})}return{mergedClsPrefix:t,controlledExpandedKeys:N,uncontrolledExpanededKeys:f,mergedExpandedKeys:k,uncontrolledValue:g,mergedValue:T,activePath:O,tmNodes:S,mergedTheme:s,mergedCollapsed:l,cssVars:o?void 0:C,themeClass:y==null?void 0:y.themeClass,overflowRef:V,counterRef:se,updateCounter:()=>{},onResize:Se,onUpdateOverflow:Ye,onUpdateCount:qe,renderCounter:Je,getCounter:Ge,onRender:y==null?void 0:y.onRender,showOption:_,deriveResponsiveState:Se}},render(){const{mergedClsPrefix:e,mode:t,themeClass:o,onRender:s}=this;s==null||s();const i=()=>this.tmNodes.map(c=>Ie(c,this.$props)),h=t==="horizontal"&&this.responsive,u=()=>a("div",uo(this.$attrs,{role:t==="horizontal"?"menubar":"menu",class:[`${e}-menu`,o,`${e}-menu--${t}`,h&&`${e}-menu--responsive`,this.mergedCollapsed&&`${e}-menu--collapsed`],style:this.cssVars}),h?a(zo,{ref:"overflowRef",onUpdateOverflow:this.onUpdateOverflow,getCounter:this.getCounter,onUpdateCount:this.onUpdateCount,updateCounter:this.updateCounter,style:{width:"100%",display:"flex",overflow:"hidden"}},{default:i,counter:this.renderCounter}):i());return h?a(so,{onResize:this.onResize},{default:u}):u()}}),Jo={xmlns:"http://www.w3.org/2000/svg","xmlns:xlink":"http://www.w3.org/1999/xlink",viewBox:"0 0 512 512"},Qo=A({name:"GitNetworkOutline",render:function(t,o){return K(),Y("svg",Jo,o[0]||(o[0]=[ho('<circle cx="128" cy="96" r="48" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="32"></circle><circle cx="256" cy="416" r="48" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="32"></circle><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="32" d="M256 256v112"></path><circle cx="384" cy="96" r="48" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="32"></circle><path d="M128 144c0 74.67 68.92 112 128 112" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="32"></path><path d="M384 144c0 74.67-68.92 112-128 112" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="32"></path>',6)]))}}),et={xmlns:"http://www.w3.org/2000/svg","xmlns:xlink":"http://www.w3.org/1999/xlink",viewBox:"0 0 512 512"},ot=A({name:"GridOutline",render:function(t,o){return K(),Y("svg",et,o[0]||(o[0]=[R("rect",{x:"48",y:"48",width:"176",height:"176",rx:"20",ry:"20",fill:"none",stroke:"currentColor","stroke-linecap":"round","stroke-linejoin":"round","stroke-width":"32"},null,-1),R("rect",{x:"288",y:"48",width:"176",height:"176",rx:"20",ry:"20",fill:"none",stroke:"currentColor","stroke-linecap":"round","stroke-linejoin":"round","stroke-width":"32"},null,-1),R("rect",{x:"48",y:"288",width:"176",height:"176",rx:"20",ry:"20",fill:"none",stroke:"currentColor","stroke-linecap":"round","stroke-linejoin":"round","stroke-width":"32"},null,-1),R("rect",{x:"288",y:"288",width:"176",height:"176",rx:"20",ry:"20",fill:"none",stroke:"currentColor","stroke-linecap":"round","stroke-linejoin":"round","stroke-width":"32"},null,-1)]))}}),tt={xmlns:"http://www.w3.org/2000/svg","xmlns:xlink":"http://www.w3.org/1999/xlink",viewBox:"0 0 512 512"},rt=A({name:"LinkOutline",render:function(t,o){return K(),Y("svg",tt,o[0]||(o[0]=[R("path",{d:"M208 352h-64a96 96 0 0 1 0-192h64",fill:"none",stroke:"currentColor","stroke-linecap":"round","stroke-linejoin":"round","stroke-width":"36"},null,-1),R("path",{d:"M304 160h64a96 96 0 0 1 0 192h-64",fill:"none",stroke:"currentColor","stroke-linecap":"round","stroke-linejoin":"round","stroke-width":"36"},null,-1),R("path",{fill:"none",stroke:"currentColor","stroke-linecap":"round","stroke-linejoin":"round","stroke-width":"36",d:"M163.29 256h187.42"},null,-1)]))}}),nt={xmlns:"http://www.w3.org/2000/svg","xmlns:xlink":"http://www.w3.org/1999/xlink",viewBox:"0 0 512 512"},lt=A({name:"LogOutOutline",render:function(t,o){return K(),Y("svg",nt,o[0]||(o[0]=[R("path",{d:"M304 336v40a40 40 0 0 1-40 40H104a40 40 0 0 1-40-40V136a40 40 0 0 1 40-40h152c22.09 0 48 17.91 48 40v40",fill:"none",stroke:"currentColor","stroke-linecap":"round","stroke-linejoin":"round","stroke-width":"32"},null,-1),R("path",{fill:"none",stroke:"currentColor","stroke-linecap":"round","stroke-linejoin":"round","stroke-width":"32",d:"M368 336l80-80l-80-80"},null,-1),R("path",{fill:"none",stroke:"currentColor","stroke-linecap":"round","stroke-linejoin":"round","stroke-width":"32",d:"M176 256h256"},null,-1)]))}}),it={xmlns:"http://www.w3.org/2000/svg","xmlns:xlink":"http://www.w3.org/1999/xlink",viewBox:"0 0 512 512"},at=A({name:"PeopleOutline",render:function(t,o){return K(),Y("svg",it,o[0]||(o[0]=[R("path",{d:"M402 168c-2.93 40.67-33.1 72-66 72s-63.12-31.32-66-72c-3-42.31 26.37-72 66-72s69 30.46 66 72z",fill:"none",stroke:"currentColor","stroke-linecap":"round","stroke-linejoin":"round","stroke-width":"32"},null,-1),R("path",{d:"M336 304c-65.17 0-127.84 32.37-143.54 95.41c-2.08 8.34 3.15 16.59 11.72 16.59h263.65c8.57 0 13.77-8.25 11.72-16.59C463.85 335.36 401.18 304 336 304z",fill:"none",stroke:"currentColor","stroke-miterlimit":"10","stroke-width":"32"},null,-1),R("path",{d:"M200 185.94c-2.34 32.48-26.72 58.06-53 58.06s-50.7-25.57-53-58.06C91.61 152.15 115.34 128 147 128s55.39 24.77 53 57.94z",fill:"none",stroke:"currentColor","stroke-linecap":"round","stroke-linejoin":"round","stroke-width":"32"},null,-1),R("path",{d:"M206 306c-18.05-8.27-37.93-11.45-59-11.45c-52 0-102.1 25.85-114.65 76.2c-1.65 6.66 2.53 13.25 9.37 13.25H154",fill:"none",stroke:"currentColor","stroke-linecap":"round","stroke-miterlimit":"10","stroke-width":"32"},null,-1)]))}}),st={xmlns:"http://www.w3.org/2000/svg","xmlns:xlink":"http://www.w3.org/1999/xlink",viewBox:"0 0 512 512"},ct=A({name:"PersonCircleOutline",render:function(t,o){return K(),Y("svg",st,o[0]||(o[0]=[R("path",{d:"M258.9 48C141.92 46.42 46.42 141.92 48 258.9c1.56 112.19 92.91 203.54 205.1 205.1c117 1.6 212.48-93.9 210.88-210.88C462.44 140.91 371.09 49.56 258.9 48zm126.42 327.25a4 4 0 0 1-6.14-.32a124.27 124.27 0 0 0-32.35-29.59C321.37 329 289.11 320 256 320s-65.37 9-90.83 25.34a124.24 124.24 0 0 0-32.35 29.58a4 4 0 0 1-6.14.32A175.32 175.32 0 0 1 80 259c-1.63-97.31 78.22-178.76 175.57-179S432 158.81 432 256a175.32 175.32 0 0 1-46.68 119.25z",fill:"currentColor"},null,-1),R("path",{d:"M256 144c-19.72 0-37.55 7.39-50.22 20.82s-19 32-17.57 51.93C191.11 256 221.52 288 256 288s64.83-32 67.79-71.24c1.48-19.74-4.8-38.14-17.68-51.82C293.39 151.44 275.59 144 256 144z",fill:"currentColor"},null,-1)]))}}),dt={xmlns:"http://www.w3.org/2000/svg","xmlns:xlink":"http://www.w3.org/1999/xlink",viewBox:"0 0 512 512"},ut=A({name:"SettingsOutline",render:function(t,o){return K(),Y("svg",dt,o[0]||(o[0]=[R("path",{d:"M262.29 192.31a64 64 0 1 0 57.4 57.4a64.13 64.13 0 0 0-57.4-57.4zM416.39 256a154.34 154.34 0 0 1-1.53 20.79l45.21 35.46a10.81 10.81 0 0 1 2.45 13.75l-42.77 74a10.81 10.81 0 0 1-13.14 4.59l-44.9-18.08a16.11 16.11 0 0 0-15.17 1.75A164.48 164.48 0 0 1 325 400.8a15.94 15.94 0 0 0-8.82 12.14l-6.73 47.89a11.08 11.08 0 0 1-10.68 9.17h-85.54a11.11 11.11 0 0 1-10.69-8.87l-6.72-47.82a16.07 16.07 0 0 0-9-12.22a155.3 155.3 0 0 1-21.46-12.57a16 16 0 0 0-15.11-1.71l-44.89 18.07a10.81 10.81 0 0 1-13.14-4.58l-42.77-74a10.8 10.8 0 0 1 2.45-13.75l38.21-30a16.05 16.05 0 0 0 6-14.08c-.36-4.17-.58-8.33-.58-12.5s.21-8.27.58-12.35a16 16 0 0 0-6.07-13.94l-38.19-30A10.81 10.81 0 0 1 49.48 186l42.77-74a10.81 10.81 0 0 1 13.14-4.59l44.9 18.08a16.11 16.11 0 0 0 15.17-1.75A164.48 164.48 0 0 1 187 111.2a15.94 15.94 0 0 0 8.82-12.14l6.73-47.89A11.08 11.08 0 0 1 213.23 42h85.54a11.11 11.11 0 0 1 10.69 8.87l6.72 47.82a16.07 16.07 0 0 0 9 12.22a155.3 155.3 0 0 1 21.46 12.57a16 16 0 0 0 15.11 1.71l44.89-18.07a10.81 10.81 0 0 1 13.14 4.58l42.77 74a10.8 10.8 0 0 1-2.45 13.75l-38.21 30a16.05 16.05 0 0 0-6.05 14.08c.33 4.14.55 8.3.55 12.47z",fill:"none",stroke:"currentColor","stroke-linecap":"round","stroke-linejoin":"round","stroke-width":"32"},null,-1)]))}}),vt={xmlns:"http://www.w3.org/2000/svg","xmlns:xlink":"http://www.w3.org/1999/xlink",viewBox:"0 0 512 512"},ht=A({name:"SwapHorizontalOutline",render:function(t,o){return K(),Y("svg",vt,o[0]||(o[0]=[R("path",{fill:"none",stroke:"currentColor","stroke-linecap":"round","stroke-linejoin":"round","stroke-width":"32",d:"M304 48l112 112l-112 112"},null,-1),R("path",{fill:"none",stroke:"currentColor","stroke-linecap":"round","stroke-linejoin":"round","stroke-width":"32",d:"M398.87 160H96"},null,-1),R("path",{fill:"none",stroke:"currentColor","stroke-linecap":"round","stroke-linejoin":"round","stroke-width":"32",d:"M208 464L96 352l112-112"},null,-1),R("path",{fill:"none",stroke:"currentColor","stroke-linecap":"round","stroke-linejoin":"round","stroke-width":"32",d:"M114 352h302"},null,-1)]))}}),mt=A({__name:"MainLayout",setup(e){const t=bo(),o=Co(),s=M(""),i=M(""),l=M("");let h;function u(z){return()=>a(Te,null,{default:()=>a(z)})}const c=[{label:()=>a(X,{to:"/"},{default:()=>"Dashboard"}),key:"/",icon:u(ot)},{label:()=>a(X,{to:"/users"},{default:()=>"Users"}),key:"/users",icon:u(at)},{label:()=>a(X,{to:"/network"},{default:()=>"Network"}),key:"/network",icon:u(Qo)},{label:()=>a(X,{to:"/outbounds"},{default:()=>"Outbounds"}),key:"/outbounds",icon:u(ht)},{label:()=>a(X,{to:"/chain"},{default:()=>"Chain"}),key:"/chain",icon:u(rt)},{label:()=>a(X,{to:"/settings"},{default:()=>"Settings"}),key:"/settings",icon:u(ut)}],g=b(()=>i.value==="RUNNING"?"success":i.value==="IDLE"?"warning":i.value?"error":"default"),H=[{label:"Log out",key:"logout",icon:u(lt)}];async function T(){try{const z=await he.get("/api/dashboard");i.value=z.mitaStatus,l.value=z.mitaVersion}catch{}}mo(async()=>{try{const z=await he.get("/api/me");s.value=z.username}catch{}T(),h=window.setInterval(()=>{document.hidden||T()},8e3)}),fo(()=>window.clearInterval(h));async function f(z){if(z==="logout")try{await he.post("/api/logout")}finally{o.push("/login")}}return(z,N)=>{const k=xo("router-view");return K(),Pe(B(_e),{position:"absolute","has-sider":""},{default:E(()=>[L(B(Fo),{bordered:"","collapsed-width":64,width:220,"collapse-mode":"width","show-trigger":"bar"},{default:E(()=>[N[0]||(N[0]=R("div",{class:"brand"},[R("span",{class:"dot"}),R("span",{class:"name"},"mieru")],-1)),L(B(Zo),{value:B(t).path,options:c,indent:20},null,8,["value"])]),_:1}),L(B(_e),null,{default:E(()=>[L(B(Mo),{bordered:"",class:"header"},{default:E(()=>[L(B(So),{align:"center",size:10},{default:E(()=>[L(B(ko),{type:g.value,size:"small",round:""},{default:E(()=>[de(ue(i.value||"…"),1)]),_:1},8,["type"]),l.value?(K(),Pe(B(Ro),{key:0,depth:"3",style:{"font-size":"12px"}},{default:E(()=>[de("mita v"+ue(l.value),1)]),_:1})):po("",!0)]),_:1}),L(B(je),{options:H,onSelect:f,trigger:"click"},{default:E(()=>[L(B(go),{text:""},{icon:E(()=>[L(B(Te),{component:B(ct)},null,8,["component"])]),default:E(()=>[de(" "+ue(s.value||"admin"),1)]),_:1})]),_:1})]),_:1}),L(B(Oo),{class:"content","content-style":"padding: 24px; max-width: 1100px; margin: 0 auto;"},{default:E(()=>[L(k)]),_:1})]),_:1})]),_:1})}}}),wt=No(mt,[["__scopeId","data-v-bca71a67"]]);export{wt as default};
