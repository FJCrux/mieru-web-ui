import{C as we,x as _,a4 as hn,d as I,h as a,cf as fn,m as x,A as z,z as s,av as vn,ay as mn,aq as J,aI as Re,N as ae,E as ge,az as pn,l as gn,n as T,F as q,a5 as be,r as C,aa as bn,ar as re,S as yn,K as wn,V as xn,p as Cn,q as We,cg as Pn,ax as Sn,J as ze,U as Mn,ae as zn,O as Ae,a6 as An,v as Fn,aL as kn,aF as Fe,b5 as ke,D as P,bg as Te,aA as me,aG as Tn,y as _n,ac as Rn,ch as Wn}from"./index-B5SBiCPF.js";import{u as Dn}from"./client-DbLYzOci.js";const En={name:"en-US",global:{undo:"Undo",redo:"Redo",confirm:"Confirm",clear:"Clear"},Popconfirm:{positiveText:"Confirm",negativeText:"Cancel"},Cascader:{placeholder:"Please Select",loading:"Loading",loadingRequiredMessage:t=>`Please load all ${t}'s descendants before checking it.`},Time:{dateFormat:"yyyy-MM-dd",dateTimeFormat:"yyyy-MM-dd HH:mm:ss"},DatePicker:{yearFormat:"yyyy",monthFormat:"MMM",dayFormat:"eeeeee",yearTypeFormat:"yyyy",monthTypeFormat:"yyyy-MM",dateFormat:"yyyy-MM-dd",dateTimeFormat:"yyyy-MM-dd HH:mm:ss",quarterFormat:"yyyy-qqq",weekFormat:"YYYY-w",clear:"Clear",now:"Now",confirm:"Confirm",selectTime:"Select Time",selectDate:"Select Date",datePlaceholder:"Select Date",datetimePlaceholder:"Select Date and Time",monthPlaceholder:"Select Month",yearPlaceholder:"Select Year",quarterPlaceholder:"Select Quarter",weekPlaceholder:"Select Week",startDatePlaceholder:"Start Date",endDatePlaceholder:"End Date",startDatetimePlaceholder:"Start Date and Time",endDatetimePlaceholder:"End Date and Time",startMonthPlaceholder:"Start Month",endMonthPlaceholder:"End Month",monthBeforeYear:!0,firstDayOfWeek:6,today:"Today"},DataTable:{checkTableAll:"Select all in the table",uncheckTableAll:"Unselect all in the table",confirm:"Confirm",clear:"Clear"},LegacyTransfer:{sourceTitle:"Source",targetTitle:"Target"},Transfer:{selectAll:"Select all",unselectAll:"Unselect all",clearAll:"Clear",total:t=>`Total ${t} items`,selected:t=>`${t} items selected`},Empty:{description:"No Data"},Select:{placeholder:"Please Select"},TimePicker:{placeholder:"Select Time",positiveText:"OK",negativeText:"Cancel",now:"Now",clear:"Clear"},Pagination:{goto:"Goto",selectionSuffix:"page"},DynamicTags:{add:"Add"},Log:{loading:"Loading"},Input:{placeholder:"Please Input"},InputNumber:{placeholder:"Please Input"},DynamicInput:{create:"Create"},ThemeEditor:{title:"Theme Editor",clearAllVars:"Clear All Variables",clearSearch:"Clear Search",filterCompName:"Filter Component Name",filterVarName:"Filter Variable Name",import:"Import",export:"Export",restore:"Reset to Default"},Image:{tipPrevious:"Previous picture (←)",tipNext:"Next picture (→)",tipCounterclockwise:"Counterclockwise",tipClockwise:"Clockwise",tipZoomOut:"Zoom out",tipZoomIn:"Zoom in",tipDownload:"Download",tipClose:"Close (Esc)",tipOriginalSize:"Zoom to original size"},Heatmap:{less:"less",more:"more",monthFormat:"MMM",weekdayFormat:"eee"}};function pe(t){return(l={})=>{const o=l.width?String(l.width):t.defaultWidth;return t.formats[o]||t.formats[t.defaultWidth]}}function Y(t){return(l,o)=>{const c=o!=null&&o.context?String(o.context):"standalone";let b;if(c==="formatting"&&t.formattingValues){const d=t.defaultFormattingWidth||t.defaultWidth,r=o!=null&&o.width?String(o.width):d;b=t.formattingValues[r]||t.formattingValues[d]}else{const d=t.defaultWidth,r=o!=null&&o.width?String(o.width):t.defaultWidth;b=t.values[r]||t.values[d]}const u=t.argumentCallback?t.argumentCallback(l):l;return b[u]}}function X(t){return(l,o={})=>{const c=o.width,b=c&&t.matchPatterns[c]||t.matchPatterns[t.defaultMatchWidth],u=l.match(b);if(!u)return null;const d=u[0],r=c&&t.parsePatterns[c]||t.parsePatterns[t.defaultParseWidth],f=Array.isArray(r)?In(r,y=>y.test(d)):Bn(r,y=>y.test(d));let A;A=t.valueCallback?t.valueCallback(f):f,A=o.valueCallback?o.valueCallback(A):A;const S=l.slice(d.length);return{value:A,rest:S}}}function Bn(t,l){for(const o in t)if(Object.prototype.hasOwnProperty.call(t,o)&&l(t[o]))return o}function In(t,l){for(let o=0;o<t.length;o++)if(l(t[o]))return o}function $n(t){return(l,o={})=>{const c=l.match(t.matchPattern);if(!c)return null;const b=c[0],u=l.match(t.parsePattern);if(!u)return null;let d=t.valueCallback?t.valueCallback(u[0]):u[0];d=o.valueCallback?o.valueCallback(d):d;const r=l.slice(b.length);return{value:d,rest:r}}}const Vn={lessThanXSeconds:{one:"less than a second",other:"less than {{count}} seconds"},xSeconds:{one:"1 second",other:"{{count}} seconds"},halfAMinute:"half a minute",lessThanXMinutes:{one:"less than a minute",other:"less than {{count}} minutes"},xMinutes:{one:"1 minute",other:"{{count}} minutes"},aboutXHours:{one:"about 1 hour",other:"about {{count}} hours"},xHours:{one:"1 hour",other:"{{count}} hours"},xDays:{one:"1 day",other:"{{count}} days"},aboutXWeeks:{one:"about 1 week",other:"about {{count}} weeks"},xWeeks:{one:"1 week",other:"{{count}} weeks"},aboutXMonths:{one:"about 1 month",other:"about {{count}} months"},xMonths:{one:"1 month",other:"{{count}} months"},aboutXYears:{one:"about 1 year",other:"about {{count}} years"},xYears:{one:"1 year",other:"{{count}} years"},overXYears:{one:"over 1 year",other:"over {{count}} years"},almostXYears:{one:"almost 1 year",other:"almost {{count}} years"}},Ln=(t,l,o)=>{let c;const b=Vn[t];return typeof b=="string"?c=b:l===1?c=b.one:c=b.other.replace("{{count}}",l.toString()),o!=null&&o.addSuffix?o.comparison&&o.comparison>0?"in "+c:c+" ago":c},Nn={lastWeek:"'last' eeee 'at' p",yesterday:"'yesterday at' p",today:"'today at' p",tomorrow:"'tomorrow at' p",nextWeek:"eeee 'at' p",other:"P"},On=(t,l,o,c)=>Nn[t],jn={narrow:["B","A"],abbreviated:["BC","AD"],wide:["Before Christ","Anno Domini"]},Un={narrow:["1","2","3","4"],abbreviated:["Q1","Q2","Q3","Q4"],wide:["1st quarter","2nd quarter","3rd quarter","4th quarter"]},Hn={narrow:["J","F","M","A","M","J","J","A","S","O","N","D"],abbreviated:["Jan","Feb","Mar","Apr","May","Jun","Jul","Aug","Sep","Oct","Nov","Dec"],wide:["January","February","March","April","May","June","July","August","September","October","November","December"]},Kn={narrow:["S","M","T","W","T","F","S"],short:["Su","Mo","Tu","We","Th","Fr","Sa"],abbreviated:["Sun","Mon","Tue","Wed","Thu","Fri","Sat"],wide:["Sunday","Monday","Tuesday","Wednesday","Thursday","Friday","Saturday"]},qn={narrow:{am:"a",pm:"p",midnight:"mi",noon:"n",morning:"morning",afternoon:"afternoon",evening:"evening",night:"night"},abbreviated:{am:"AM",pm:"PM",midnight:"midnight",noon:"noon",morning:"morning",afternoon:"afternoon",evening:"evening",night:"night"},wide:{am:"a.m.",pm:"p.m.",midnight:"midnight",noon:"noon",morning:"morning",afternoon:"afternoon",evening:"evening",night:"night"}},Yn={narrow:{am:"a",pm:"p",midnight:"mi",noon:"n",morning:"in the morning",afternoon:"in the afternoon",evening:"in the evening",night:"at night"},abbreviated:{am:"AM",pm:"PM",midnight:"midnight",noon:"noon",morning:"in the morning",afternoon:"in the afternoon",evening:"in the evening",night:"at night"},wide:{am:"a.m.",pm:"p.m.",midnight:"midnight",noon:"noon",morning:"in the morning",afternoon:"in the afternoon",evening:"in the evening",night:"at night"}},Xn=(t,l)=>{const o=Number(t),c=o%100;if(c>20||c<10)switch(c%10){case 1:return o+"st";case 2:return o+"nd";case 3:return o+"rd"}return o+"th"},Jn={ordinalNumber:Xn,era:Y({values:jn,defaultWidth:"wide"}),quarter:Y({values:Un,defaultWidth:"wide",argumentCallback:t=>t-1}),month:Y({values:Hn,defaultWidth:"wide"}),day:Y({values:Kn,defaultWidth:"wide"}),dayPeriod:Y({values:qn,defaultWidth:"wide",formattingValues:Yn,defaultFormattingWidth:"wide"})},Gn=/^(\d+)(th|st|nd|rd)?/i,Qn=/\d+/i,Zn={narrow:/^(b|a)/i,abbreviated:/^(b\.?\s?c\.?|b\.?\s?c\.?\s?e\.?|a\.?\s?d\.?|c\.?\s?e\.?)/i,wide:/^(before christ|before common era|anno domini|common era)/i},er={any:[/^b/i,/^(a|c)/i]},tr={narrow:/^[1234]/i,abbreviated:/^q[1234]/i,wide:/^[1234](th|st|nd|rd)? quarter/i},nr={any:[/1/i,/2/i,/3/i,/4/i]},rr={narrow:/^[jfmasond]/i,abbreviated:/^(jan|feb|mar|apr|may|jun|jul|aug|sep|oct|nov|dec)/i,wide:/^(january|february|march|april|may|june|july|august|september|october|november|december)/i},or={narrow:[/^j/i,/^f/i,/^m/i,/^a/i,/^m/i,/^j/i,/^j/i,/^a/i,/^s/i,/^o/i,/^n/i,/^d/i],any:[/^ja/i,/^f/i,/^mar/i,/^ap/i,/^may/i,/^jun/i,/^jul/i,/^au/i,/^s/i,/^o/i,/^n/i,/^d/i]},ar={narrow:/^[smtwf]/i,short:/^(su|mo|tu|we|th|fr|sa)/i,abbreviated:/^(sun|mon|tue|wed|thu|fri|sat)/i,wide:/^(sunday|monday|tuesday|wednesday|thursday|friday|saturday)/i},ir={narrow:[/^s/i,/^m/i,/^t/i,/^w/i,/^t/i,/^f/i,/^s/i],any:[/^su/i,/^m/i,/^tu/i,/^w/i,/^th/i,/^f/i,/^sa/i]},lr={narrow:/^(a|p|mi|n|(in the|at) (morning|afternoon|evening|night))/i,any:/^([ap]\.?\s?m\.?|midnight|noon|(in the|at) (morning|afternoon|evening|night))/i},sr={any:{am:/^a/i,pm:/^p/i,midnight:/^mi/i,noon:/^no/i,morning:/morning/i,afternoon:/afternoon/i,evening:/evening/i,night:/night/i}},cr={ordinalNumber:$n({matchPattern:Gn,parsePattern:Qn,valueCallback:t=>parseInt(t,10)}),era:X({matchPatterns:Zn,defaultMatchWidth:"wide",parsePatterns:er,defaultParseWidth:"any"}),quarter:X({matchPatterns:tr,defaultMatchWidth:"wide",parsePatterns:nr,defaultParseWidth:"any",valueCallback:t=>t+1}),month:X({matchPatterns:rr,defaultMatchWidth:"wide",parsePatterns:or,defaultParseWidth:"any"}),day:X({matchPatterns:ar,defaultMatchWidth:"wide",parsePatterns:ir,defaultParseWidth:"any"}),dayPeriod:X({matchPatterns:lr,defaultMatchWidth:"any",parsePatterns:sr,defaultParseWidth:"any"})},ur={full:"EEEE, MMMM do, y",long:"MMMM do, y",medium:"MMM d, y",short:"MM/dd/yyyy"},dr={full:"h:mm:ss a zzzz",long:"h:mm:ss a z",medium:"h:mm:ss a",short:"h:mm a"},hr={full:"{{date}} 'at' {{time}}",long:"{{date}} 'at' {{time}}",medium:"{{date}}, {{time}}",short:"{{date}}, {{time}}"},fr={date:pe({formats:ur,defaultWidth:"full"}),time:pe({formats:dr,defaultWidth:"full"}),dateTime:pe({formats:hr,defaultWidth:"full"})},vr={code:"en-US",formatDistance:Ln,formatLong:fr,formatRelative:On,localize:Jn,match:cr,options:{weekStartsOn:0,firstWeekContainsDate:1}},mr={name:"en-US",locale:vr};function pr(t){const{mergedLocaleRef:l,mergedDateLocaleRef:o}=we(hn,null)||{},c=_(()=>{var u,d;return(d=(u=l==null?void 0:l.value)===null||u===void 0?void 0:u[t])!==null&&d!==void 0?d:En[t]});return{dateLocaleRef:_(()=>{var u;return(u=o==null?void 0:o.value)!==null&&u!==void 0?u:mr}),localeRef:c}}const gr=I({name:"ChevronDown",render(){return a("svg",{viewBox:"0 0 16 16",fill:"none",xmlns:"http://www.w3.org/2000/svg"},a("path",{d:"M3.14645 5.64645C3.34171 5.45118 3.65829 5.45118 3.85355 5.64645L8 9.79289L12.1464 5.64645C12.3417 5.45118 12.6583 5.45118 12.8536 5.64645C13.0488 5.84171 13.0488 6.15829 12.8536 6.35355L8.35355 10.8536C8.15829 11.0488 7.84171 11.0488 7.64645 10.8536L3.14645 6.35355C2.95118 6.15829 2.95118 5.84171 3.14645 5.64645Z",fill:"currentColor"}))}}),br=fn("clear",()=>a("svg",{viewBox:"0 0 16 16",version:"1.1",xmlns:"http://www.w3.org/2000/svg"},a("g",{stroke:"none","stroke-width":"1",fill:"none","fill-rule":"evenodd"},a("g",{fill:"currentColor","fill-rule":"nonzero"},a("path",{d:"M8,2 C11.3137085,2 14,4.6862915 14,8 C14,11.3137085 11.3137085,14 8,14 C4.6862915,14 2,11.3137085 2,8 C2,4.6862915 4.6862915,2 8,2 Z M6.5343055,5.83859116 C6.33943736,5.70359511 6.07001296,5.72288026 5.89644661,5.89644661 L5.89644661,5.89644661 L5.83859116,5.9656945 C5.70359511,6.16056264 5.72288026,6.42998704 5.89644661,6.60355339 L5.89644661,6.60355339 L7.293,8 L5.89644661,9.39644661 L5.83859116,9.4656945 C5.70359511,9.66056264 5.72288026,9.92998704 5.89644661,10.1035534 L5.89644661,10.1035534 L5.9656945,10.1614088 C6.16056264,10.2964049 6.42998704,10.2771197 6.60355339,10.1035534 L6.60355339,10.1035534 L8,8.707 L9.39644661,10.1035534 L9.4656945,10.1614088 C9.66056264,10.2964049 9.92998704,10.2771197 10.1035534,10.1035534 L10.1035534,10.1035534 L10.1614088,10.0343055 C10.2964049,9.83943736 10.2771197,9.57001296 10.1035534,9.39644661 L10.1035534,9.39644661 L8.707,8 L10.1035534,6.60355339 L10.1614088,6.5343055 C10.2964049,6.33943736 10.2771197,6.07001296 10.1035534,5.89644661 L10.1035534,5.89644661 L10.0343055,5.83859116 C9.83943736,5.70359511 9.57001296,5.72288026 9.39644661,5.89644661 L9.39644661,5.89644661 L8,7.293 L6.60355339,5.89644661 Z"}))))),yr=I({name:"Eye",render(){return a("svg",{xmlns:"http://www.w3.org/2000/svg",viewBox:"0 0 512 512"},a("path",{d:"M255.66 112c-77.94 0-157.89 45.11-220.83 135.33a16 16 0 0 0-.27 17.77C82.92 340.8 161.8 400 255.66 400c92.84 0 173.34-59.38 221.79-135.25a16.14 16.14 0 0 0 0-17.47C428.89 172.28 347.8 112 255.66 112z",fill:"none",stroke:"currentColor","stroke-linecap":"round","stroke-linejoin":"round","stroke-width":"32"}),a("circle",{cx:"256",cy:"256",r:"80",fill:"none",stroke:"currentColor","stroke-miterlimit":"10","stroke-width":"32"}))}}),wr=I({name:"EyeOff",render(){return a("svg",{xmlns:"http://www.w3.org/2000/svg",viewBox:"0 0 512 512"},a("path",{d:"M432 448a15.92 15.92 0 0 1-11.31-4.69l-352-352a16 16 0 0 1 22.62-22.62l352 352A16 16 0 0 1 432 448z",fill:"currentColor"}),a("path",{d:"M255.66 384c-41.49 0-81.5-12.28-118.92-36.5c-34.07-22-64.74-53.51-88.7-91v-.08c19.94-28.57 41.78-52.73 65.24-72.21a2 2 0 0 0 .14-2.94L93.5 161.38a2 2 0 0 0-2.71-.12c-24.92 21-48.05 46.76-69.08 76.92a31.92 31.92 0 0 0-.64 35.54c26.41 41.33 60.4 76.14 98.28 100.65C162 402 207.9 416 255.66 416a239.13 239.13 0 0 0 75.8-12.58a2 2 0 0 0 .77-3.31l-21.58-21.58a4 4 0 0 0-3.83-1a204.8 204.8 0 0 1-51.16 6.47z",fill:"currentColor"}),a("path",{d:"M490.84 238.6c-26.46-40.92-60.79-75.68-99.27-100.53C349 110.55 302 96 255.66 96a227.34 227.34 0 0 0-74.89 12.83a2 2 0 0 0-.75 3.31l21.55 21.55a4 4 0 0 0 3.88 1a192.82 192.82 0 0 1 50.21-6.69c40.69 0 80.58 12.43 118.55 37c34.71 22.4 65.74 53.88 89.76 91a.13.13 0 0 1 0 .16a310.72 310.72 0 0 1-64.12 72.73a2 2 0 0 0-.15 2.95l19.9 19.89a2 2 0 0 0 2.7.13a343.49 343.49 0 0 0 68.64-78.48a32.2 32.2 0 0 0-.1-34.78z",fill:"currentColor"}),a("path",{d:"M256 160a95.88 95.88 0 0 0-21.37 2.4a2 2 0 0 0-1 3.38l112.59 112.56a2 2 0 0 0 3.38-1A96 96 0 0 0 256 160z",fill:"currentColor"}),a("path",{d:"M165.78 233.66a2 2 0 0 0-3.38 1a96 96 0 0 0 115 115a2 2 0 0 0 1-3.38z",fill:"currentColor"}))}}),xr=x("base-clear",`
 flex-shrink: 0;
 height: 1em;
 width: 1em;
 position: relative;
`,[z(">",[s("clear",`
 font-size: var(--n-clear-size);
 height: 1em;
 width: 1em;
 cursor: pointer;
 color: var(--n-clear-color);
 transition: color .3s var(--n-bezier);
 display: flex;
 `,[z("&:hover",`
 color: var(--n-clear-color-hover)!important;
 `),z("&:active",`
 color: var(--n-clear-color-pressed)!important;
 `)]),s("placeholder",`
 display: flex;
 `),s("clear, placeholder",`
 position: absolute;
 left: 50%;
 top: 50%;
 transform: translateX(-50%) translateY(-50%);
 `,[vn({originalTransform:"translateX(-50%) translateY(-50%)",left:"50%",top:"50%"})])])]),ye=I({name:"BaseClear",props:{clsPrefix:{type:String,required:!0},show:Boolean,onClear:Function},setup(t){return Re("-base-clear",xr,ge(t,"clsPrefix")),{handleMouseDown(l){l.preventDefault()}}},render(){const{clsPrefix:t}=this;return a("div",{class:`${t}-base-clear`},a(mn,null,{default:()=>{var l,o;return this.show?a("div",{key:"dismiss",class:`${t}-base-clear__clear`,onClick:this.onClear,onMousedown:this.handleMouseDown,"data-clear":!0},J(this.$slots.icon,()=>[a(ae,{clsPrefix:t},{default:()=>a(br,null)})])):a("div",{key:"icon",class:`${t}-base-clear__placeholder`},(o=(l=this.$slots).placeholder)===null||o===void 0?void 0:o.call(l))}}))}}),Cr=I({name:"InternalSelectionSuffix",props:{clsPrefix:{type:String,required:!0},showArrow:{type:Boolean,default:void 0},showClear:{type:Boolean,default:void 0},loading:{type:Boolean,default:!1},onClear:Function},setup(t,{slots:l}){return()=>{const{clsPrefix:o}=t;return a(pn,{clsPrefix:o,class:`${o}-base-suffix`,strokeWidth:24,scale:.85,show:t.loading},{default:()=>t.showArrow?a(ye,{clsPrefix:o,show:t.showClear,onClear:t.onClear},{placeholder:()=>a(ae,{clsPrefix:o,class:`${o}-base-suffix__arrow`},{default:()=>J(l.default,()=>[a(gr,null)])})}):null})}}}),De=gn("n-input"),Pr=x("input",`
 max-width: 100%;
 cursor: text;
 line-height: 1.5;
 z-index: auto;
 outline: none;
 box-sizing: border-box;
 position: relative;
 display: inline-flex;
 border-radius: var(--n-border-radius);
 background-color: var(--n-color);
 transition: background-color .3s var(--n-bezier);
 font-size: var(--n-font-size);
 font-weight: var(--n-font-weight);
 --n-padding-vertical: calc((var(--n-height) - 1.5 * var(--n-font-size)) / 2);
`,[s("input, textarea",`
 overflow: hidden;
 flex-grow: 1;
 position: relative;
 `),s("input-el, textarea-el, input-mirror, textarea-mirror, separator, placeholder",`
 box-sizing: border-box;
 font-size: inherit;
 line-height: 1.5;
 font-family: inherit;
 border: none;
 outline: none;
 background-color: #0000;
 text-align: inherit;
 transition:
 -webkit-text-fill-color .3s var(--n-bezier),
 caret-color .3s var(--n-bezier),
 color .3s var(--n-bezier),
 text-decoration-color .3s var(--n-bezier);
 `),s("input-el, textarea-el",`
 -webkit-appearance: none;
 scrollbar-width: none;
 width: 100%;
 min-width: 0;
 text-decoration-color: var(--n-text-decoration-color);
 color: var(--n-text-color);
 caret-color: var(--n-caret-color);
 background-color: transparent;
 `,[z("&::-webkit-scrollbar, &::-webkit-scrollbar-track-piece, &::-webkit-scrollbar-thumb",`
 width: 0;
 height: 0;
 display: none;
 `),z("&::placeholder",`
 color: #0000;
 -webkit-text-fill-color: transparent !important;
 `),z("&:-webkit-autofill ~",[s("placeholder","display: none;")])]),T("round",[q("textarea","border-radius: calc(var(--n-height) / 2);")]),s("placeholder",`
 pointer-events: none;
 position: absolute;
 left: 0;
 right: 0;
 top: 0;
 bottom: 0;
 overflow: hidden;
 color: var(--n-placeholder-color);
 `,[z("span",`
 width: 100%;
 display: inline-block;
 `)]),T("textarea",[s("placeholder","overflow: visible;")]),q("autosize","width: 100%;"),T("autosize",[s("textarea-el, input-el",`
 position: absolute;
 top: 0;
 left: 0;
 height: 100%;
 `)]),x("input-wrapper",`
 overflow: hidden;
 display: inline-flex;
 flex-grow: 1;
 position: relative;
 padding-left: var(--n-padding-left);
 padding-right: var(--n-padding-right);
 `),s("input-mirror",`
 padding: 0;
 height: var(--n-height);
 line-height: var(--n-height);
 overflow: hidden;
 visibility: hidden;
 position: static;
 white-space: pre;
 pointer-events: none;
 `),s("input-el",`
 padding: 0;
 height: var(--n-height);
 line-height: var(--n-height);
 `,[z("&[type=password]::-ms-reveal","display: none;"),z("+",[s("placeholder",`
 display: flex;
 align-items: center; 
 `)])]),q("textarea",[s("placeholder","white-space: nowrap;")]),s("eye",`
 display: flex;
 align-items: center;
 justify-content: center;
 transition: color .3s var(--n-bezier);
 `),T("textarea","width: 100%;",[x("input-word-count",`
 position: absolute;
 right: var(--n-padding-right);
 bottom: var(--n-padding-vertical);
 `),T("resizable",[x("input-wrapper",`
 resize: vertical;
 min-height: var(--n-height);
 `)]),s("textarea-el, textarea-mirror, placeholder",`
 height: 100%;
 padding-left: 0;
 padding-right: 0;
 padding-top: var(--n-padding-vertical);
 padding-bottom: var(--n-padding-vertical);
 word-break: break-word;
 display: inline-block;
 vertical-align: bottom;
 box-sizing: border-box;
 line-height: var(--n-line-height-textarea);
 margin: 0;
 resize: none;
 white-space: pre-wrap;
 scroll-padding-block-end: var(--n-padding-vertical);
 `),s("textarea-mirror",`
 width: 100%;
 pointer-events: none;
 overflow: hidden;
 visibility: hidden;
 position: static;
 white-space: pre-wrap;
 overflow-wrap: break-word;
 `)]),T("pair",[s("input-el, placeholder","text-align: center;"),s("separator",`
 display: flex;
 align-items: center;
 transition: color .3s var(--n-bezier);
 color: var(--n-text-color);
 white-space: nowrap;
 `,[x("icon",`
 color: var(--n-icon-color);
 `),x("base-icon",`
 color: var(--n-icon-color);
 `)])]),T("disabled",`
 cursor: not-allowed;
 background-color: var(--n-color-disabled);
 `,[s("border","border: var(--n-border-disabled);"),s("input-el, textarea-el",`
 cursor: not-allowed;
 color: var(--n-text-color-disabled);
 text-decoration-color: var(--n-text-color-disabled);
 `),s("placeholder","color: var(--n-placeholder-color-disabled);"),s("separator","color: var(--n-text-color-disabled);",[x("icon",`
 color: var(--n-icon-color-disabled);
 `),x("base-icon",`
 color: var(--n-icon-color-disabled);
 `)]),x("input-word-count",`
 color: var(--n-count-text-color-disabled);
 `),s("suffix, prefix","color: var(--n-text-color-disabled);",[x("icon",`
 color: var(--n-icon-color-disabled);
 `),x("internal-icon",`
 color: var(--n-icon-color-disabled);
 `)])]),q("disabled",[s("eye",`
 color: var(--n-icon-color);
 cursor: pointer;
 `,[z("&:hover",`
 color: var(--n-icon-color-hover);
 `),z("&:active",`
 color: var(--n-icon-color-pressed);
 `)]),z("&:hover",[s("state-border","border: var(--n-border-hover);")]),T("focus","background-color: var(--n-color-focus);",[s("state-border",`
 border: var(--n-border-focus);
 box-shadow: var(--n-box-shadow-focus);
 `)])]),s("border, state-border",`
 box-sizing: border-box;
 position: absolute;
 left: 0;
 right: 0;
 top: 0;
 bottom: 0;
 pointer-events: none;
 border-radius: inherit;
 border: var(--n-border);
 transition:
 box-shadow .3s var(--n-bezier),
 border-color .3s var(--n-bezier);
 `),s("state-border",`
 border-color: #0000;
 z-index: 1;
 `),s("prefix","margin-right: 4px;"),s("suffix",`
 margin-left: 4px;
 `),s("suffix, prefix",`
 transition: color .3s var(--n-bezier);
 flex-wrap: nowrap;
 flex-shrink: 0;
 line-height: var(--n-height);
 white-space: nowrap;
 display: inline-flex;
 align-items: center;
 justify-content: center;
 color: var(--n-suffix-text-color);
 `,[x("base-loading",`
 font-size: var(--n-icon-size);
 margin: 0 2px;
 color: var(--n-loading-color);
 `),x("base-clear",`
 font-size: var(--n-icon-size);
 `,[s("placeholder",[x("base-icon",`
 transition: color .3s var(--n-bezier);
 color: var(--n-icon-color);
 font-size: var(--n-icon-size);
 `)])]),z(">",[x("icon",`
 transition: color .3s var(--n-bezier);
 color: var(--n-icon-color);
 font-size: var(--n-icon-size);
 `)]),x("base-icon",`
 font-size: var(--n-icon-size);
 `)]),x("input-word-count",`
 pointer-events: none;
 line-height: 1.5;
 font-size: .85em;
 color: var(--n-count-text-color);
 transition: color .3s var(--n-bezier);
 margin-left: 4px;
 font-variant: tabular-nums;
 `),["warning","error"].map(t=>T(`${t}-status`,[q("disabled",[x("base-loading",`
 color: var(--n-loading-color-${t})
 `),s("input-el, textarea-el",`
 caret-color: var(--n-caret-color-${t});
 `),s("state-border",`
 border: var(--n-border-${t});
 `),z("&:hover",[s("state-border",`
 border: var(--n-border-hover-${t});
 `)]),z("&:focus",`
 background-color: var(--n-color-focus-${t});
 `,[s("state-border",`
 box-shadow: var(--n-box-shadow-focus-${t});
 border: var(--n-border-focus-${t});
 `)]),T("focus",`
 background-color: var(--n-color-focus-${t});
 `,[s("state-border",`
 box-shadow: var(--n-box-shadow-focus-${t});
 border: var(--n-border-focus-${t});
 `)])])]))]),Sr=x("input",[T("disabled",[s("input-el, textarea-el",`
 -webkit-text-fill-color: var(--n-text-color-disabled);
 `)])]);function Mr(t){let l=0;for(const o of t)l++;return l}function oe(t){return t===""||t==null}function zr(t){const l=C(null);function o(){const{value:u}=t;if(!(u!=null&&u.focus)){b();return}const{selectionStart:d,selectionEnd:r,value:f}=u;if(d==null||r==null){b();return}l.value={start:d,end:r,beforeText:f.slice(0,d),afterText:f.slice(r)}}function c(){var u;const{value:d}=l,{value:r}=t;if(!d||!r)return;const{value:f}=r,{start:A,beforeText:S,afterText:y}=d;let M=f.length;if(f.endsWith(y))M=f.length-y.length;else if(f.startsWith(S))M=S.length;else{const w=S[A-1],h=f.indexOf(w,A-1);h!==-1&&(M=h+1)}(u=r.setSelectionRange)===null||u===void 0||u.call(r,M,M)}function b(){l.value=null}return be(t,b),{recordCursor:o,restoreCursor:c}}const _e=I({name:"InputWordCount",setup(t,{slots:l}){const{mergedValueRef:o,maxlengthRef:c,mergedClsPrefixRef:b,countGraphemesRef:u}=we(De),d=_(()=>{const{value:r}=o;return r===null||Array.isArray(r)?0:(u.value||Mr)(r)});return()=>{const{value:r}=c,{value:f}=o;return a("span",{class:`${b.value}-input-word-count`},bn(l.default,{value:f===null||Array.isArray(f)?"":f},()=>[r===void 0?d.value:`${d.value} / ${r}`]))}}}),Ar=Object.assign(Object.assign({},We.props),{bordered:{type:Boolean,default:void 0},type:{type:String,default:"text"},placeholder:[Array,String],defaultValue:{type:[String,Array],default:null},value:[String,Array],disabled:{type:Boolean,default:void 0},size:String,rows:{type:[Number,String],default:3},round:Boolean,minlength:[String,Number],maxlength:[String,Number],clearable:Boolean,autosize:{type:[Boolean,Object],default:!1},pair:Boolean,separator:String,readonly:{type:[String,Boolean],default:!1},passivelyActivated:Boolean,showPasswordOn:String,stateful:{type:Boolean,default:!0},autofocus:Boolean,inputProps:Object,resizable:{type:Boolean,default:!0},showCount:Boolean,loading:{type:Boolean,default:void 0},allowInput:Function,renderCount:Function,onMousedown:Function,onKeydown:Function,onKeyup:[Function,Array],onInput:[Function,Array],onFocus:[Function,Array],onBlur:[Function,Array],onClick:[Function,Array],onChange:[Function,Array],onClear:[Function,Array],countGraphemes:Function,status:String,"onUpdate:value":[Function,Array],onUpdateValue:[Function,Array],textDecoration:[String,Array],attrSize:{type:Number,default:20},onInputBlur:[Function,Array],onInputFocus:[Function,Array],onDeactivate:[Function,Array],onActivate:[Function,Array],onWrapperFocus:[Function,Array],onWrapperBlur:[Function,Array],internalDeactivateOnEnter:Boolean,internalForceFocus:Boolean,internalLoadingBeforeSuffix:{type:Boolean,default:!0},showPasswordToggle:Boolean}),Tr=I({name:"Input",props:Ar,slots:Object,setup(t){const{mergedClsPrefixRef:l,mergedBorderedRef:o,inlineThemeDisabled:c,mergedRtlRef:b,mergedComponentPropsRef:u}=Cn(t),d=We("Input","-input",Pr,kn,t,l);Pn&&Re("-input-safari",Sr,l);const r=C(null),f=C(null),A=C(null),S=C(null),y=C(null),M=C(null),w=C(null),h=zr(w),m=C(null),{localeRef:F}=pr("Input"),k=C(t.defaultValue),ie=ge(t,"value"),R=Dn(ie,k),N=Sn(t,{mergedSize:e=>{var n,i;const{size:p}=t;if(p)return p;const{mergedSize:g}=e||{};if(g!=null&&g.value)return g.value;const v=(i=(n=u==null?void 0:u.value)===null||n===void 0?void 0:n.Input)===null||i===void 0?void 0:i.size;return v||"medium"}}),{mergedSizeRef:le,mergedDisabledRef:$,mergedStatusRef:Ee}=N,V=C(!1),O=C(!1),W=C(!1),j=C(!1);let se=null;const ce=_(()=>{const{placeholder:e,pair:n}=t;return n?Array.isArray(e)?e:e===void 0?["",""]:[e,e]:e===void 0?[F.value.placeholder]:[e]}),Be=_(()=>{const{value:e}=W,{value:n}=R,{value:i}=ce;return!e&&(oe(n)||Array.isArray(n)&&oe(n[0]))&&i[0]}),Ie=_(()=>{const{value:e}=W,{value:n}=R,{value:i}=ce;return!e&&i[1]&&(oe(n)||Array.isArray(n)&&oe(n[1]))}),ue=ze(()=>t.internalForceFocus||V.value),$e=ze(()=>{if($.value||t.readonly||!t.clearable||!ue.value&&!O.value)return!1;const{value:e}=R,{value:n}=ue;return t.pair?!!(Array.isArray(e)&&(e[0]||e[1]))&&(O.value||n):!!e&&(O.value||n)}),de=_(()=>{const{showPasswordOn:e}=t;if(e)return e;if(t.showPasswordToggle)return"click"}),U=C(!1),Ve=_(()=>{const{textDecoration:e}=t;return e?Array.isArray(e)?e.map(n=>({textDecoration:n})):[{textDecoration:e}]:["",""]}),xe=C(void 0),Le=()=>{var e,n;if(t.type==="textarea"){const{autosize:i}=t;if(i&&(xe.value=(n=(e=m.value)===null||e===void 0?void 0:e.$el)===null||n===void 0?void 0:n.offsetWidth),!f.value||typeof i=="boolean")return;const{paddingTop:p,paddingBottom:g,lineHeight:v}=window.getComputedStyle(f.value),D=Number(p.slice(0,-2)),E=Number(g.slice(0,-2)),B=Number(v.slice(0,-2)),{value:H}=A;if(!H)return;if(i.minRows){const K=Math.max(i.minRows,1),ve=`${D+E+B*K}px`;H.style.minHeight=ve}if(i.maxRows){const K=`${D+E+B*i.maxRows}px`;H.style.maxHeight=K}}},Ne=_(()=>{const{maxlength:e}=t;return e===void 0?void 0:Number(e)});Mn(()=>{const{value:e}=R;Array.isArray(e)||fe(e)});const Oe=zn().proxy;function G(e,n){const{onUpdateValue:i,"onUpdate:value":p,onInput:g}=t,{nTriggerFormInput:v}=N;i&&P(i,e,n),p&&P(p,e,n),g&&P(g,e,n),k.value=e,v()}function Q(e,n){const{onChange:i}=t,{nTriggerFormChange:p}=N;i&&P(i,e,n),k.value=e,p()}function je(e){const{onBlur:n}=t,{nTriggerFormBlur:i}=N;n&&P(n,e),i()}function Ue(e){const{onFocus:n}=t,{nTriggerFormFocus:i}=N;n&&P(n,e),i()}function He(e){const{onClear:n}=t;n&&P(n,e)}function Ke(e){const{onInputBlur:n}=t;n&&P(n,e)}function qe(e){const{onInputFocus:n}=t;n&&P(n,e)}function Ye(){const{onDeactivate:e}=t;e&&P(e)}function Xe(){const{onActivate:e}=t;e&&P(e)}function Je(e){const{onClick:n}=t;n&&P(n,e)}function Ge(e){const{onWrapperFocus:n}=t;n&&P(n,e)}function Qe(e){const{onWrapperBlur:n}=t;n&&P(n,e)}function Ze(){W.value=!0}function et(e){W.value=!1,e.target===M.value?Z(e,1):Z(e,0)}function Z(e,n=0,i="input"){const p=e.target.value;if(fe(p),e instanceof InputEvent&&!e.isComposing&&(W.value=!1),t.type==="textarea"){const{value:v}=m;v&&v.syncUnifiedContainer()}if(se=p,W.value)return;h.recordCursor();const g=tt(p);if(g)if(!t.pair)i==="input"?G(p,{source:n}):Q(p,{source:n});else{let{value:v}=R;Array.isArray(v)?v=[v[0],v[1]]:v=["",""],v[n]=p,i==="input"?G(v,{source:n}):Q(v,{source:n})}Oe.$forceUpdate(),g||Fe(h.restoreCursor)}function tt(e){const{countGraphemes:n,maxlength:i,minlength:p}=t;if(n){let v;if(i!==void 0&&(v===void 0&&(v=n(e)),v>Number(i))||p!==void 0&&(v===void 0&&(v=n(e)),v<Number(i)))return!1}const{allowInput:g}=t;return typeof g=="function"?g(e):!0}function nt(e){Ke(e),e.relatedTarget===r.value&&Ye(),e.relatedTarget!==null&&(e.relatedTarget===y.value||e.relatedTarget===M.value||e.relatedTarget===f.value)||(j.value=!1),ee(e,"blur"),w.value=null}function rt(e,n){qe(e),V.value=!0,j.value=!0,Xe(),ee(e,"focus"),n===0?w.value=y.value:n===1?w.value=M.value:n===2&&(w.value=f.value)}function ot(e){t.passivelyActivated&&(Qe(e),ee(e,"blur"))}function at(e){t.passivelyActivated&&(V.value=!0,Ge(e),ee(e,"focus"))}function ee(e,n){e.relatedTarget!==null&&(e.relatedTarget===y.value||e.relatedTarget===M.value||e.relatedTarget===f.value||e.relatedTarget===r.value)||(n==="focus"?(Ue(e),V.value=!0):n==="blur"&&(je(e),V.value=!1))}function it(e,n){Z(e,n,"change")}function lt(e){Je(e)}function st(e){He(e),Ce()}function Ce(){t.pair?(G(["",""],{source:"clear"}),Q(["",""],{source:"clear"})):(G("",{source:"clear"}),Q("",{source:"clear"}))}function ct(e){const{onMousedown:n}=t;n&&n(e);const{tagName:i}=e.target;if(i!=="INPUT"&&i!=="TEXTAREA"){if(t.resizable){const{value:p}=r;if(p){const{left:g,top:v,width:D,height:E}=p.getBoundingClientRect(),B=14;if(g+D-B<e.clientX&&e.clientX<g+D&&v+E-B<e.clientY&&e.clientY<v+E)return}}e.preventDefault(),V.value||Pe()}}function ut(){var e;O.value=!0,t.type==="textarea"&&((e=m.value)===null||e===void 0||e.handleMouseEnterWrapper())}function dt(){var e;O.value=!1,t.type==="textarea"&&((e=m.value)===null||e===void 0||e.handleMouseLeaveWrapper())}function ht(){$.value||de.value==="click"&&(U.value=!U.value)}function ft(e){if($.value)return;e.preventDefault();const n=p=>{p.preventDefault(),Te("mouseup",document,n)};if(ke("mouseup",document,n),de.value!=="mousedown")return;U.value=!0;const i=()=>{U.value=!1,Te("mouseup",document,i)};ke("mouseup",document,i)}function vt(e){t.onKeyup&&P(t.onKeyup,e)}function mt(e){switch(t.onKeydown&&P(t.onKeydown,e),e.key){case"Escape":he();break;case"Enter":pt(e);break}}function pt(e){var n,i;if(t.passivelyActivated){const{value:p}=j;if(p){t.internalDeactivateOnEnter&&he();return}e.preventDefault(),t.type==="textarea"?(n=f.value)===null||n===void 0||n.focus():(i=y.value)===null||i===void 0||i.focus()}}function he(){t.passivelyActivated&&(j.value=!1,Fe(()=>{var e;(e=r.value)===null||e===void 0||e.focus()}))}function Pe(){var e,n,i;$.value||(t.passivelyActivated?(e=r.value)===null||e===void 0||e.focus():((n=f.value)===null||n===void 0||n.focus(),(i=y.value)===null||i===void 0||i.focus()))}function gt(){var e;!((e=r.value)===null||e===void 0)&&e.contains(document.activeElement)&&document.activeElement.blur()}function bt(){var e,n;(e=f.value)===null||e===void 0||e.select(),(n=y.value)===null||n===void 0||n.select()}function yt(){$.value||(f.value?f.value.focus():y.value&&y.value.focus())}function wt(){const{value:e}=r;e!=null&&e.contains(document.activeElement)&&e!==document.activeElement&&he()}function xt(e){if(t.type==="textarea"){const{value:n}=f;n==null||n.scrollTo(e)}else{const{value:n}=y;n==null||n.scrollTo(e)}}function fe(e){const{type:n,pair:i,autosize:p}=t;if(!i&&p)if(n==="textarea"){const{value:g}=A;g&&(g.textContent=`${e??""}\r
`)}else{const{value:g}=S;g&&(e?g.textContent=e:g.innerHTML="&nbsp;")}}function Ct(){Le()}const Se=C({top:"0"});function Pt(e){var n;const{scrollTop:i}=e.target;Se.value.top=`${-i}px`,(n=m.value)===null||n===void 0||n.syncUnifiedContainer()}let te=null;Ae(()=>{const{autosize:e,type:n}=t;e&&n==="textarea"?te=be(R,i=>{!Array.isArray(i)&&i!==se&&fe(i)}):te==null||te()});let ne=null;Ae(()=>{t.type==="textarea"?ne=be(R,e=>{var n;!Array.isArray(e)&&e!==se&&((n=m.value)===null||n===void 0||n.syncUnifiedContainer())}):ne==null||ne()}),_n(De,{mergedValueRef:R,maxlengthRef:Ne,mergedClsPrefixRef:l,countGraphemesRef:ge(t,"countGraphemes")});const St={wrapperElRef:r,inputElRef:y,textareaElRef:f,isCompositing:W,clear:Ce,focus:Pe,blur:gt,select:bt,deactivate:wt,activate:yt,scrollTo:xt},Mt=An("Input",b,l),Me=_(()=>{const{value:e}=le,{common:{cubicBezierEaseInOut:n},self:{color:i,borderRadius:p,textColor:g,caretColor:v,caretColorError:D,caretColorWarning:E,textDecorationColor:B,border:H,borderDisabled:K,borderHover:ve,borderFocus:zt,placeholderColor:At,placeholderColorDisabled:Ft,lineHeightTextarea:kt,colorDisabled:Tt,colorFocus:_t,textColorDisabled:Rt,boxShadowFocus:Wt,iconSize:Dt,colorFocusWarning:Et,boxShadowFocusWarning:Bt,borderWarning:It,borderFocusWarning:$t,borderHoverWarning:Vt,colorFocusError:Lt,boxShadowFocusError:Nt,borderError:Ot,borderFocusError:jt,borderHoverError:Ut,clearSize:Ht,clearColor:Kt,clearColorHover:qt,clearColorPressed:Yt,iconColor:Xt,iconColorDisabled:Jt,suffixTextColor:Gt,countTextColor:Qt,countTextColorDisabled:Zt,iconColorHover:en,iconColorPressed:tn,loadingColor:nn,loadingColorError:rn,loadingColorWarning:on,fontWeight:an,[me("padding",e)]:ln,[me("fontSize",e)]:sn,[me("height",e)]:cn}}=d.value,{left:un,right:dn}=Tn(ln);return{"--n-bezier":n,"--n-count-text-color":Qt,"--n-count-text-color-disabled":Zt,"--n-color":i,"--n-font-size":sn,"--n-font-weight":an,"--n-border-radius":p,"--n-height":cn,"--n-padding-left":un,"--n-padding-right":dn,"--n-text-color":g,"--n-caret-color":v,"--n-text-decoration-color":B,"--n-border":H,"--n-border-disabled":K,"--n-border-hover":ve,"--n-border-focus":zt,"--n-placeholder-color":At,"--n-placeholder-color-disabled":Ft,"--n-icon-size":Dt,"--n-line-height-textarea":kt,"--n-color-disabled":Tt,"--n-color-focus":_t,"--n-text-color-disabled":Rt,"--n-box-shadow-focus":Wt,"--n-loading-color":nn,"--n-caret-color-warning":E,"--n-color-focus-warning":Et,"--n-box-shadow-focus-warning":Bt,"--n-border-warning":It,"--n-border-focus-warning":$t,"--n-border-hover-warning":Vt,"--n-loading-color-warning":on,"--n-caret-color-error":D,"--n-color-focus-error":Lt,"--n-box-shadow-focus-error":Nt,"--n-border-error":Ot,"--n-border-focus-error":jt,"--n-border-hover-error":Ut,"--n-loading-color-error":rn,"--n-clear-color":Kt,"--n-clear-size":Ht,"--n-clear-color-hover":qt,"--n-clear-color-pressed":Yt,"--n-icon-color":Xt,"--n-icon-color-hover":en,"--n-icon-color-pressed":tn,"--n-icon-color-disabled":Jt,"--n-suffix-text-color":Gt}}),L=c?Fn("input",_(()=>{const{value:e}=le;return e[0]}),Me,t):void 0;return Object.assign(Object.assign({},St),{wrapperElRef:r,inputElRef:y,inputMirrorElRef:S,inputEl2Ref:M,textareaElRef:f,textareaMirrorElRef:A,textareaScrollbarInstRef:m,rtlEnabled:Mt,uncontrolledValue:k,mergedValue:R,passwordVisible:U,mergedPlaceholder:ce,showPlaceholder1:Be,showPlaceholder2:Ie,mergedFocus:ue,isComposing:W,activated:j,showClearButton:$e,mergedSize:le,mergedDisabled:$,textDecorationStyle:Ve,mergedClsPrefix:l,mergedBordered:o,mergedShowPasswordOn:de,placeholderStyle:Se,mergedStatus:Ee,textAreaScrollContainerWidth:xe,handleTextAreaScroll:Pt,handleCompositionStart:Ze,handleCompositionEnd:et,handleInput:Z,handleInputBlur:nt,handleInputFocus:rt,handleWrapperBlur:ot,handleWrapperFocus:at,handleMouseEnter:ut,handleMouseLeave:dt,handleMouseDown:ct,handleChange:it,handleClick:lt,handleClear:st,handlePasswordToggleClick:ht,handlePasswordToggleMousedown:ft,handleWrapperKeydown:mt,handleWrapperKeyup:vt,handleTextAreaMirrorResize:Ct,getTextareaScrollContainer:()=>f.value,mergedTheme:d,cssVars:c?void 0:Me,themeClass:L==null?void 0:L.themeClass,onRender:L==null?void 0:L.onRender})},render(){var t,l,o,c,b,u,d;const{mergedClsPrefix:r,mergedStatus:f,themeClass:A,type:S,countGraphemes:y,onRender:M}=this,w=this.$slots;return M==null||M(),a("div",{ref:"wrapperElRef",class:[`${r}-input`,`${r}-input--${this.mergedSize}-size`,A,f&&`${r}-input--${f}-status`,{[`${r}-input--rtl`]:this.rtlEnabled,[`${r}-input--disabled`]:this.mergedDisabled,[`${r}-input--textarea`]:S==="textarea",[`${r}-input--resizable`]:this.resizable&&!this.autosize,[`${r}-input--autosize`]:this.autosize,[`${r}-input--round`]:this.round&&S!=="textarea",[`${r}-input--pair`]:this.pair,[`${r}-input--focus`]:this.mergedFocus,[`${r}-input--stateful`]:this.stateful}],style:this.cssVars,tabindex:!this.mergedDisabled&&this.passivelyActivated&&!this.activated?0:void 0,onFocus:this.handleWrapperFocus,onBlur:this.handleWrapperBlur,onClick:this.handleClick,onMousedown:this.handleMouseDown,onMouseenter:this.handleMouseEnter,onMouseleave:this.handleMouseLeave,onCompositionstart:this.handleCompositionStart,onCompositionend:this.handleCompositionEnd,onKeyup:this.handleWrapperKeyup,onKeydown:this.handleWrapperKeydown},a("div",{class:`${r}-input-wrapper`},re(w.prefix,h=>h&&a("div",{class:`${r}-input__prefix`},h)),S==="textarea"?a(yn,{ref:"textareaScrollbarInstRef",class:`${r}-input__textarea`,container:this.getTextareaScrollContainer,theme:(l=(t=this.theme)===null||t===void 0?void 0:t.peers)===null||l===void 0?void 0:l.Scrollbar,themeOverrides:(c=(o=this.themeOverrides)===null||o===void 0?void 0:o.peers)===null||c===void 0?void 0:c.Scrollbar,triggerDisplayManually:!0,useUnifiedContainer:!0,internalHoistYRail:!0},{default:()=>{var h,m;const{textAreaScrollContainerWidth:F}=this,k={width:this.autosize&&F&&`${F}px`};return a(wn,null,a("textarea",Object.assign({},this.inputProps,{ref:"textareaElRef",class:[`${r}-input__textarea-el`,(h=this.inputProps)===null||h===void 0?void 0:h.class],autofocus:this.autofocus,rows:Number(this.rows),placeholder:this.placeholder,value:this.mergedValue,disabled:this.mergedDisabled,maxlength:y?void 0:this.maxlength,minlength:y?void 0:this.minlength,readonly:this.readonly,tabindex:this.passivelyActivated&&!this.activated?-1:void 0,style:[this.textDecorationStyle[0],(m=this.inputProps)===null||m===void 0?void 0:m.style,k],onBlur:this.handleInputBlur,onFocus:ie=>{this.handleInputFocus(ie,2)},onInput:this.handleInput,onChange:this.handleChange,onScroll:this.handleTextAreaScroll})),this.showPlaceholder1?a("div",{class:`${r}-input__placeholder`,style:[this.placeholderStyle,k],key:"placeholder"},this.mergedPlaceholder[0]):null,this.autosize?a(xn,{onResize:this.handleTextAreaMirrorResize},{default:()=>a("div",{ref:"textareaMirrorElRef",class:`${r}-input__textarea-mirror`,key:"mirror"})}):null)}}):a("div",{class:`${r}-input__input`},a("input",Object.assign({type:S==="password"&&this.mergedShowPasswordOn&&this.passwordVisible?"text":S},this.inputProps,{ref:"inputElRef",class:[`${r}-input__input-el`,(b=this.inputProps)===null||b===void 0?void 0:b.class],style:[this.textDecorationStyle[0],(u=this.inputProps)===null||u===void 0?void 0:u.style],tabindex:this.passivelyActivated&&!this.activated?-1:(d=this.inputProps)===null||d===void 0?void 0:d.tabindex,placeholder:this.mergedPlaceholder[0],disabled:this.mergedDisabled,maxlength:y?void 0:this.maxlength,minlength:y?void 0:this.minlength,value:Array.isArray(this.mergedValue)?this.mergedValue[0]:this.mergedValue,readonly:this.readonly,autofocus:this.autofocus,size:this.attrSize,onBlur:this.handleInputBlur,onFocus:h=>{this.handleInputFocus(h,0)},onInput:h=>{this.handleInput(h,0)},onChange:h=>{this.handleChange(h,0)}})),this.showPlaceholder1?a("div",{class:`${r}-input__placeholder`},a("span",null,this.mergedPlaceholder[0])):null,this.autosize?a("div",{class:`${r}-input__input-mirror`,key:"mirror",ref:"inputMirrorElRef"}," "):null),!this.pair&&re(w.suffix,h=>h||this.clearable||this.showCount||this.mergedShowPasswordOn||this.loading!==void 0?a("div",{class:`${r}-input__suffix`},[re(w["clear-icon-placeholder"],m=>(this.clearable||m)&&a(ye,{clsPrefix:r,show:this.showClearButton,onClear:this.handleClear},{placeholder:()=>m,icon:()=>{var F,k;return(k=(F=this.$slots)["clear-icon"])===null||k===void 0?void 0:k.call(F)}})),this.internalLoadingBeforeSuffix?null:h,this.loading!==void 0?a(Cr,{clsPrefix:r,loading:this.loading,showArrow:!1,showClear:!1,style:this.cssVars}):null,this.internalLoadingBeforeSuffix?h:null,this.showCount&&this.type!=="textarea"?a(_e,null,{default:m=>{var F;const{renderCount:k}=this;return k?k(m):(F=w.count)===null||F===void 0?void 0:F.call(w,m)}}):null,this.mergedShowPasswordOn&&this.type==="password"?a("div",{class:`${r}-input__eye`,onMousedown:this.handlePasswordToggleMousedown,onClick:this.handlePasswordToggleClick},this.passwordVisible?J(w["password-visible-icon"],()=>[a(ae,{clsPrefix:r},{default:()=>a(yr,null)})]):J(w["password-invisible-icon"],()=>[a(ae,{clsPrefix:r},{default:()=>a(wr,null)})])):null]):null)),this.pair?a("span",{class:`${r}-input__separator`},J(w.separator,()=>[this.separator])):null,this.pair?a("div",{class:`${r}-input-wrapper`},a("div",{class:`${r}-input__input`},a("input",{ref:"inputEl2Ref",type:this.type,class:`${r}-input__input-el`,tabindex:this.passivelyActivated&&!this.activated?-1:void 0,placeholder:this.mergedPlaceholder[1],disabled:this.mergedDisabled,maxlength:y?void 0:this.maxlength,minlength:y?void 0:this.minlength,value:Array.isArray(this.mergedValue)?this.mergedValue[1]:void 0,readonly:this.readonly,style:this.textDecorationStyle[1],onBlur:this.handleInputBlur,onFocus:h=>{this.handleInputFocus(h,1)},onInput:h=>{this.handleInput(h,1)},onChange:h=>{this.handleChange(h,1)}}),this.showPlaceholder2?a("div",{class:`${r}-input__placeholder`},a("span",null,this.mergedPlaceholder[1])):null),re(w.suffix,h=>(this.clearable||h)&&a("div",{class:`${r}-input__suffix`},[this.clearable&&a(ye,{clsPrefix:r,show:this.showClearButton,onClear:this.handleClear},{icon:()=>{var m;return(m=w["clear-icon"])===null||m===void 0?void 0:m.call(w)},placeholder:()=>{var m;return(m=w["clear-icon-placeholder"])===null||m===void 0?void 0:m.call(w)}}),h]))):null,this.mergedBordered?a("div",{class:`${r}-input__border`}):null,this.mergedBordered?a("div",{class:`${r}-input__state-border`}):null,this.showCount&&S==="textarea"?a(_e,null,{default:h=>{var m;const{renderCount:F}=this;return F?F(h):(m=w.count)===null||m===void 0?void 0:m.call(w,h)}}):null)}});function _r(){const t=we(Wn,null);return t===null&&Rn("use-message","No outer <n-message-provider /> founded. See prerequisite in https://www.naiveui.com/en-US/os-theme/components/message for more details. If you want to use `useMessage` outside setup, please check https://www.naiveui.com/zh-CN/os-theme/components/message#Q-&-A."),t}export{gr as C,Tr as N,pr as a,Cr as b,_r as u};
