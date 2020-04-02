<template>
    <div class="footer_container">
        <div class="person_computer_footer">
            <div class="footer_logo_content">
                <a @click="$uMeng.push('Footer_IRISnet','click')" class="irisnet_link_content" href="https://www.irisnet.org/" target="_blank">
                    <img class="irisnet_logo_img" src="../../assets/IRISnet-Rebrand-Capital-Black.png"/>
                </a>
            </div>
            <div class="community_container">
                <h1 class="community_title">
                    Community
                </h1>
                <div class="community_list_content">
                    <a @click="$uMeng.push('Footer_Github','click')"  target="_blank" href='https://github.com/irisnet'><i class="iconfont icongithub"></i></a>
                    <a @click="$uMeng.push('Footer_Telegram','click')"  target="_blank" href='https://t.me/irisnetwork'><i class="iconfont icontelegram"></i></a>
                    <a @click="$uMeng.push('Footer_Twitter','click')"  target="_blank" href='https://twitter.com/irisnetwork'><i class="iconfont icontuite"></i></a>
                    <a @click="$uMeng.push('Footer_Medium','click')"  target="_blank" href='https://medium.com/irisnet-blog'><i class="iconfont iconmedium"></i></a>
                    <span class="we_chat" @click="showWeChatQRCode"><i class="iconfont iconweixin"></i></span>
                    <span class="qq" @click="showqqQRCode"><i class="iconfont iconQQ"></i></span>
                </div>
            </div>
            <div class="footer_right_content">
                <h1 class="resources_content">Resources</h1>
                <div class="footer_link_wrap">
                    <a @click="$uMeng.push('Footer_Use Testnet','click')"
                       href="https://www.irisnet.org/testnets"
                       target="_blank">
                        <span class="footer_link_contact">Use Testnet</span>
                    </a>
                    <span class="footer_link_join">|</span>
                    <span class="footer_link_privacy">
                        <router-link @click.native="$uMeng.push('Footer_Privacy Policy','click')"
                                     :to="`/privacy_policy`">Privacy Policy</router-link>
                    </span>
                    <span class="footer_link_join">|</span>
                    <span class="footer-faq">
                        <router-link @click.native="$uMeng.push('Footer_FAQ','click')"
                                     :to="`/help`">FAQ</router-link>
                    </span>
                </div>
            </div>
        </div>
        <div class="footer_version_content">
            <div class="footer_version_content_warp">
                <div class="footer_copyright_wrap">
                    ©️ IRISplorer 2019-2020 all rights reserved
                </div>
                <div class="footer_chain_id_content">Chain ID {{chainID}}</div>
                <div class="footer_version_node_tendermint_content"><p>Node Version {{nodeVersion}}</p> <span class="line">|</span>  <p>Tendermint Version {{tendermintVersion}}</p></div>
            </div>
        </div>
    </div>
</template>

<script>
    import Tools from "../../util/Tools"
	export default {
		name: "AppFoooter",
        data () {
			return {
				weChatQRShow:false,
				qqQRShow: false,
                chainID:'',
                tendermintVersion:'',
                nodeVersion:'',
                flShowVersion: false,
                versionTimer: null,
            }
        },

        mounted(){
		    if(sessionStorage.getItem('skinEnvInformation')){
                this.getVersionInformation()
            }else {
		        let that = this;
		        this.versionTimer = setInterval(() => {
                    that.getVersionInformation()
                    if(sessionStorage.getItem('skinEnvInformation')){
                        clearInterval(that.versionTimer)
                    }
                },500)
            }
        },
        methods:{
		    getVersionInformation(){
                if(sessionStorage.getItem('skinEnvInformation')){
                    this.flShowVersion = true;
                    let configs = JSON.parse(sessionStorage.getItem('skinEnvInformation')),
                        currnetEnv = JSON.parse(sessionStorage.getItem('skinEnvInformation')).cur_env,
                        currendChainId = JSON.parse(sessionStorage.getItem('skinEnvInformation')).chain_id;

                    configs.configs.forEach( item => {
                        if(currnetEnv === item.env && currendChainId === item.chain_id){
                            this.chainID = Tools.firstWordUpperCase(item.chain_id);
                            this.tendermintVersion = item.tendermint_version;
                            this.nodeVersion = item.node_version
                        }
                    })
                }else {
                    this.flShowVersion = false;
                }
            },
	        showWeChatQRCode() {
                this.$uMeng.push('Footer_WeChat','click');
                this.$store.commit('flShowQR',true);
	        	this.$store.commit('setQrImg','wechat');
	        },
	        showqqQRCode(){
                this.$uMeng.push('Footer_QQ','click');
		        this.$store.commit('flShowQR',true);
		        this.$store.commit('setQrImg','qq');
	        },
        }
	}
</script>

<style scoped lang="scss">
    .footer_container{
        width: 100%;
        background: #363a3d;
        .person_computer_footer{
            display: flex;
            max-width: 12.8rem;
            margin: 0 auto;
            padding: 0.35rem 0;
            .footer_logo_content{
                flex: 1;
                display: flex;
                align-items: center;
                padding-left: 0.2rem;
                .irisnet_link_content{
                    display: inline-block;
                    width:33%;
                    .irisnet_logo_img{
                        width:100%;
                    }
                }
            }
            .community_container{
                flex: 1;
                display: flex;
                align-items: flex-start;
                justify-content: center;
                flex-direction: column;
                .community_title{
                    font-size: 0.16rem;
                    padding-bottom: 0.1rem;
                    color: #fff;
                }
                .community_list_content{
                    display: flex;
                    a{
                        margin-right: 0.25rem;
                        i{
                            font-size: 0.25rem;
                            color: rgba(255,255,255,0.5);
                            &:hover{
                                color: #fff;
                            }
                        }
                    }
                }
                span{
                    margin-right: 0.25rem;
                    cursor: pointer;
                    i{
                        font-size: 0.25rem;
                        color: rgba(255,255,255,0.5);
                        &:hover{
                            color: #fff;
                        }
                    }
                }
            }
            .footer_right_content{
                flex: 1;
                display: flex;
                flex-direction: column;
                justify-content: center;
                .resources_content{
                    font-size: 0.16rem;
                    color: #fff;
                    margin-bottom: 0.1rem;
                }
                .footer_link_wrap{
                    color: rgba(255,255,255,0.5);
                    a{
                        padding: 0 0.2rem;
                        color: rgba(255,255,255,0.5) !important;
                        &:hover{
                            color: #fff !important;
                        }
                    }
                    a:first-child{
                        padding-left: 0;
                    }
                    .footer_link_privacy{
                        padding-left: 0.2rem;
                    }
                    .footer-faq{
                        padding-left: 0.2rem;
                    }

                }
            }
        }

        .footer_version_content{
            border-top: 0.01rem  solid rgba(255,255,255,0.2);
            background: #2B2E31;
            padding: 0.15rem 0;
            .footer_version_content_warp{
                max-width: 12.8rem;
                margin: 0 auto;
                display: flex;
                color: rgba(255,255,255,0.5);
                .footer_copyright_wrap {
                    flex: 1;
                }
                .footer_chain_id_content{
                    flex: 1;
                    padding-left: 0.2rem;
                }
                .footer_version_node_tendermint_content{
                    flex: 1;
                    display: flex;
                    white-space: nowrap;
                    .line{
                        padding: 0 0.2rem;
                    }
                }
            }

        }
    }
@media screen  and (max-width: 910px){
    .footer_container{
        .person_computer_footer{
            flex-direction: column;
            align-items: center;
            padding: 0.1rem 0 0 0;
            .footer_logo_content{
                padding-left: 0;
                display: flex;
                justify-content: center;
                .irisnet_link_content{
                    width: 1.5rem;
                }
            }
            .community_container{
                margin-top: 0.1rem;
                justify-content: center;
                align-items: center;
                .community_list_content{
                    justify-content: center;
                    .qq{
                        margin-right: 0;
                    }
                }
            }
            .footer_right_content{
                justify-content: center;
                align-items: center;
                margin: 0.2rem 0;
                .resources_content{
                    text-align: center;
                }
                .footer_link_wrap{
                    text-align: center;
                }
            }
        }
        .footer_version_content{
            .footer_version_content_warp{
                flex-direction: column;
                align-items: center;
                justify-content: center;
                .footer_chain_id_content{
                    margin: 0.1rem 0;
                    padding-left: 0;
                }
                .footer_version_node_tendermint_content{
                    display: flex;
                    flex-direction: column;
                    align-items: center;
                    .line{
                        display: none;
                    }
                    p:last-child{
                        margin-top: 0.1rem;
                    }
                }

            }
        }
    }

}
</style>
