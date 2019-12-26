<template>
    <div class="footer_container">
        <div class="person_computer_footer">
            <div class="footer_logo_content">
                <a class="irisnet_link_content" href="https://www.irisnet.org/" target="_blank">
                    <img class="irisnet_logo_img" src="../../assets/IRISnet-Rebrand-Capital-Black.png"/>
                </a>
            </div>
            <div class="community_container">
                <h1 class="community_title">
                    Community
                </h1>
                <div class="community_list_content">
                    <a target="_blank" href='https://github.com/irisnet'><i class="iconfont icongithub"></i></a>
                    <a target="_blank" href='https://t.me/irisnetwork'><i class="iconfont icontelegram"></i></a>
                    <a target="_blank" href='https://twitter.com/irisnetwork'><i class="iconfont icontuite"></i></a>
                    <a target="_blank" href='https://medium.com/irisnet-blog'><i class="iconfont iconmedium"></i></a>
                    <span class="we_chat" @click="showWeChatQRCode"><i class="iconfont iconweixin"></i></span>
                    <span class="qq" @click="showqqQRCode"><i class="iconfont iconQQ"></i></span>
                </div>
            </div>
            <div class="footer_right_content">
                <h1 class="resources_content">Resources</h1>
                <div class="footer_link_wrap">
                    <a href="https://www.irisnet.org/testnets" target="_blank">
                        <span class="footer_link_contact">Use Testnet</span>
                    </a>
                    <span class="footer_link_join">|</span>
                    <span class="footer_link_privacy"><router-link :to="`/privacy_policy`">Privacy Policy</router-link></span>
                    <span class="footer_link_join">|</span>
                    <span class="footer-faq"><router-link :to="`/help`">FAQ</router-link></span>
                </div>
            </div>
        </div>
        <div class="foot_version" v-if="flShowVersion">
            <div class="foot_version_chain_id">Chain ID: {{chainID}}</div>
            <span class="footer_link_join">|</span>
            <div class="foot_version_tendermint">Tendermint Version: {{tendermintVersion}}</div>
            <span class="footer_link_join">|</span>
            <div class="foot_version_node">Node Version:{{nodeVersion}}</div>
        </div>
        <p class="footer_copyright_wrap">
            ©️ IRISplorer 2019 all rights reserved
        </p>
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
            }
        },
        mounted(){
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
        methods:{
	        showWeChatQRCode() {
	        	this.$store.commit('flShowQR',true);
	        	this.$store.commit('setQrImg','wechat');
	        },
	        showqqQRCode(){
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
        .footer_copyright_wrap {
            background: #363a3d;
            border-top: 0.01rem  solid rgba(255,255,255,0.2);
            padding: 0.15rem 0;
            text-align: center;
            color: rgba(255,255,255,0.5);
        }
        .foot_version{
            font-size: 0.14rem;
            border-top: 0.01rem  solid rgba(255,255,255,0.2);
            padding: 0.16rem 0;
            display: flex;
            align-items: center;
            justify-content: center;
            color: rgba(255,255,255,0.5);
            .foot_version_chain_id{
                padding-right: 0.2rem;
            }
            .foot_version_tendermint{
                padding:0 0.2rem;
            }
            .foot_version_node{
                padding-left: 0.2rem;
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
        .foot_version{
            padding: 0.06rem;
            display: flex;
            flex-direction: column;
            align-items: center;
            justify-content: center;
            font-size: 0.14rem;
            .footer_link_join{
                display: none;
            }
            .foot_version_chain_id{
                padding: 0;
            }
            .foot_version_chain_id{
                padding: 0;
            }
            .foot_version_node{
                padding: 0;
            }
        }
    }

}
</style>
