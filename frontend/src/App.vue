<template>
    <div id="app">
        <loading-component :show-loading="$store.state.flShowLoading"></loading-component>
        <qr-component></qr-component>
        <app-header></app-header>
        <div style="flex: 1;background: #F5F7FD;">
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
	export default {
		name: "App",
		components: {qrComponent, AppFooter, LoadingComponent, AppHeader},
		data () {
            return {

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
        font-size: 16px !important;
        font-family:"Arial" !important;
        /*font-family:"-apple-system","BlinkMacSystemFont","Segoe UI","Helvetica","Arial","sans-serif","Apple Color Emoji","Segoe UI Emoji"!important;*/
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
            height: calc(100% - 0.6rem);
            margin-top: 0.6rem;
            background: #F5F7FD;
        }
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
</style>
