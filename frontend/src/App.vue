<template>
    <div id="app">
        <loading-component :show-loading="$store.state.flShowLoading"></loading-component>
        <qr-component></qr-component>
        <!--<app-header></app-header>-->
        <goz-header></goz-header>
        <div style="flex: 1;background: #202342;">
            <router-view id="router_wrap" class="router-view" :key="key"/>
        </div>
        <app-footer></app-footer>
    </div>
</template>

<script>
	import AppHeader from "./components/header/AppHeader";
	import LoadingComponent from "./components/loadingComponent/LoadingComponent";
	import AppFooter from "./components/footer/AppFooter";
	import qrComponent from "./components/qrComponent/qrComponent";
    import GozHeader from "./components/gozHeader/GozHeader";
	export default {
		name: "App",
		components: {GozHeader, qrComponent, AppFooter, LoadingComponent, AppHeader},
		data () {
            return {
                timer:null,
            }
        },
		beforeMount(){
			window.addEventListener('resize', this.onresize);
			if (window.innerWidth > 910) {
				this.$store.commit('isMobile',false);
			} else {
				this.$store.commit('isMobile',true);
			}
		},
        mounted () {
            let uMengID;
		    this.timer = setInterval(() =>{
                if(sessionStorage.getItem('UMengID')){
                    uMengID = sessionStorage.getItem('UMengID');
                    const script = document.createElement('script');
                    script.src = `https://s95.cnzz.com/z_stat.php?id=${uMengID}&web_id=${uMengID}`;
                    script.language = 'JavaScript';
                    document.body.appendChild(script)
                    clearInterval(this.timer)
                }
            },500)
        },
        watch: {
            '$route' () {
                if (window._czc) {
                    let location = window.location;
                    let locationHash;
                    if(location.hash.includes('=') && !location.hash.split('?')){
                        locationHash = location.hash.split('=')[0];
                    }else if(location.hash.includes('/address/') || location.hash.includes('/#/address/')){
                        locationHash = 'address/'
                    }else if(location.hash.includes('/block/') || location.hash.includes('/#/block/')){
                        locationHash = 'block/'
                    }else if(location.hash.includes('/asset/') || location.hash.includes('/#/asset/')){
                        locationHash = 'asset/'
                    }else if(location.hash.includes('/ProposalsDetail/') || location.hash.includes('/#/ProposalsDetail/')){
                        locationHash = 'ProposalsDetail/'
                    }else if(location.hash.includes('/htlc/') || location.hash.includes('/#/htlc/')){
                        locationHash = 'htlc/'
                    }else if(location.hash.includes('/validators/') || location.hash.includes('/#/validators/')){
                        locationHash = 'validators/'
                    }else if(location.hash.includes('?')){
                        locationHash = location.hash.split('?')[0];
                    }else {
                        locationHash = location.hash
                    }
                    let contentUrl = location.pathname + locationHash;
                    let refererUrl = '/';
                    window._czc.push(['_trackPageview', contentUrl, refererUrl])
                }
            }
        },
		computed: {
			onresize(){
				if (window.innerWidth > 910) {
					this.$store.commit('isMobile',false);
				} else {
					this.$store.commit('isMobile',true);
				}
			},
			key() {
				return this.$route.name !== undefined ? this.$route.name + new Date() : this.$route + new Date()
			}
		},
	}
</script>

<style lang="scss">
    @import './style/mixin.scss';
    html {
        font-size: 625%;
        -webkit-text-size-adjust: none;
        overflow-y: auto;
        position: relative;
    }
    body, html {
        width: 100%;
        height: 100%;
        color: var(--moduleColor);
    }
    body {
        text-shadow:none !important;
        font-size: 16px !important;
        font-family:"Arial" !important;
        /*-webkit-tap-highlight-color:rgba(0,0,0,0);*/
        /*font-family:"-apple-system","BlinkMacSystemFont","Segoe UI","Helvetica","Arial","sans-serif","Apple Color Emoji","Segoe UI Emoji"!important;*/
    }
    //解决ios点击出现蓝色边框
    * {
        -webkit-tap-highlight-color: rgba(0,0,0,0);
    }
    #app{
        width: 100%;
        height: 100%;
        display: flex;
        flex-direction: column;
        justify-content: space-between;
        .router-view{
            flex: 1;
            width: 100%;
            background: #202342;
        }
    }
    .el-scrollbar {
        > .el-scrollbar__bar {
            opacity: 1 !important;
        }
    }
    .el-select-dropdown{
        .el-scrollbar{
            .el-select-dropdown__wrap{
                .el-select-dropdown__list{
                    .el-select-dropdown__item{
                        color: var(--bgColor);
                    }
                }
            }
        }
    }
    .el-date-picker{
        width: 2.46rem !important;
    }
    .el-cascader-menu__wrap{
        height: 2.2rem !important;
    }
    .el-scrollbar__bar.is-horizontal>div {
        display: none;
    }
    .el-picker-panel{
        .el-picker-panel__body-wrapper{
            .el-picker-panel__body{
                .el-date-picker__header{
                    width: 2.16rem;
                    margin: 0.06rem 0.12rem;
                    .el-picker-panel__icon-btn{
                        color: #A3A6B9;
                        padding: 0;
                        &:hover{
                            color: var(--bgColor) !important;
                        }
                    }
                    .el-date-picker__header-label{
                        color: #171D44 !important;
                        &:hover{
                            color: var(--bgColor) !important;
                        }
                        font-size: 0.14rem !important;
                    }
                }
                .el-picker-panel__content{
                    width: 2.16rem;
                    margin: 0 0.15rem 0.15rem 0.15rem;
                    .el-date-table{
                        tbody{
                            tr{
                                th{
                                    color:#171D44 !important;
                                }
                                td{
                                    opacity: 1 !important;
                                }
                            }
                            .el-date-table__row{
                                .prev-month{
                                    color: #A3A6B9 !important;
                                    div{
                                        height: 0.2rem;
                                    }
                                }
                                .next-month{
                                    div{
                                        height: 0.2rem;
                                    }
                                }
                                .available {
                                    width: 0.24rem;
                                    height: 0.2rem;
                                    div{
                                        height: 0.2rem;
                                        color: #171D44 !important;
                                        &:hover{
                                            color: var(--bgColor) !important;
                                        }
                                    }
                                }
                                .today{
                                    div{
                                        span{
                                            color:var(--bgColor) !important;
                                        }
                                    }
                                }
                                .current{
                                    div{
                                        span{
                                            color:#fff !important;
                                            background:var(--bgColor) !important;
                                        }
                                    }
                                }
                            }
                        }
                    }
                }
            }
        }
    }

     .el-date-table td, .el-date-table td div{
         height: 0.24rem!important;
         padding: 0 !important;
     }
    @media screen and (max-width: 910px){
        #app{
            .router-view{
                margin-top: 0;
                height: auto;
            }
        }
        //解决在移动端，日期选择器会移出到窗口以外
        .ivu-date-picker:nth-of-type(2){
            .ivu-select-dropdown{
                left: 37% !important;
            }
        }
    }
    //全局列表跳转样式
    .skip_route{
        a{
            font-size: 0.14rem;
            color: var(--bgColor) !important;
        }
    }
    .no_skip{
        color: var(--titleColore) !important;
        a{
            color: var(--titleColore) !important;
        }
    }
    .ivu-date-picker-focused input{
        border-color: var(--bgColor) !important;
    }
    .ivu-select-selection-focused{
        border-color: var(--bgColor) !important;
    }
    .is-active{
        color: var(--bgColor)!important;
    }
    .in-active-path{
        color: var(--bgColor)!important;
    }
    .el-cascader__suggestion-item.is-checked{
        color:var(--bgColor) !important;
    }
    @media screen and (max-width: 910px){
        .el-cascader__dropdown{
            .el-cascader-panel{
                div:nth-of-type(1){
                    min-width: 1.2rem;
                    div{
                        ul{
                            li{
                                padding-right: 0;
                                padding-left: 0.15rem;
                            }
                        }
                    }
                }
                div:nth-of-type(2){
                    min-width: 1.2rem;
                    div{
                        ul{
                            li{
                                padding-right: 0;
                                padding-left: 0.15rem;
                            }
                        }
                    }
                }

            }
        }
    }
</style>
