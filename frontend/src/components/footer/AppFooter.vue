<template>
    <div class="footer_container">
<!--        <div class="person_computer_footer">
            <div class="footer_logo_content">
                <a @click="$uMeng.push('Footer_IRISnet','click')" class="irisnet_link_content" href="https://www.irisnet.org/" target="_blank">
                    <img class="irisnet_logo_img" src="../../assets/IRISnet-Rebrand-Capital-Black.png"/>
                </a>
            </div>
            <div class="community_container">
                <div class="community_list_content">
                    <a @click="$uMeng.push('Footer_Github','click')"  target="_blank" href='https://github.com/irisnet'><i class="iconfont icongithub"></i></a>
                    <a @click="$uMeng.push('Footer_Telegram','click')"  target="_blank" href='https://t.me/irisnetwork'><i class="iconfont icontelegram"></i></a>
                    <a @click="$uMeng.push('Footer_Twitter','click')"  target="_blank" href='https://twitter.com/irisnetwork'><i class="iconfont icontuite"></i></a>
                    <a @click="$uMeng.push('Footer_Medium','click')"  target="_blank" href='https://medium.com/irisnet-blog'><i class="iconfont iconmedium"></i></a>
                </div>
            </div>
        </div>-->
        <div class="footer_version_content">
            <div class="footer_version_container_wrap">
              <div class="footer_version_content_warp">
                <div class="irisnet_logo">
                  <a class="irisnet_link_content" href="https://www.irisnet.org/" target="_blank">
                    <img class="irisnet_logo_img" src="../../assets/IRIS_LOGO.png"/>
                  </a>
                </div>
                <div class="footer_copyright_wrap">
                  Copyright © 2022 IOBScan Ltd. All Rights Reserved.
                </div>
                <!-- <div class="footer_chain_id_content">Chain ID {{chainID}}</div>
                 <div class="footer_version_node_tendermint_content"><p>Node Version {{nodeVersion}}</p> <span class="line">|</span>  <p>Tendermint Version {{tendermintVersion}}</p></div>-->
              </div>
              <div class="community_container">
                <div class="community_list_content">
                  <a target="_blank" href='https://github.com/irisnet'><i class="iconfont icon-github"></i></a>
                  <a target="_blank" href='https://t.me/irisnetwork'><i class="iconfont icon-telegram"></i></a>
                  <a target="_blank" href='https://twitter.com/irisnetwork'><i class="iconfont icon-twitter"></i></a>
                  <a target="_blank" href='https://medium.com/irisnet-blog'><i class="iconfont icon-medium"></i></a>
                </div>
              </div>
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
                // this.getVersionInformation()
            }else {
		        let that = this;
		        this.versionTimer = setInterval(() => {
                    // that.getVersionInformation()
                    if(sessionStorage.getItem('skinEnvInformation')){
                        clearInterval(that.versionTimer)
                    }
                },500)
            }
        },
        methods:{
		    /*getVersionInformation(){
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
            },*/
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
        background: rgba(16, 19, 55, 1);
        .person_computer_footer{
            display: flex;
            justify-content: center;
            max-width: 12.8rem;
            margin: 0 auto;
            padding: 0.35rem 0;
            .footer_logo_content{
                display: flex;
                align-items: center;
                padding-left: 0.2rem;
                .irisnet_link_content{
                    display: inline-block;
                    width:1.5rem;
                    .irisnet_logo_img{
                        height: 100%;
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
            background: rgba(0, 0, 0, 1);
            padding: 0.15rem 0;
            @media (max-width: 1280px) {
              padding:0.15rem 0.32rem !important;
            }
            .footer_version_container_wrap{
              display: flex;
              justify-content: space-between;
              margin: 0 auto;
              max-width: 12.8rem;
              @media (max-width: 820px) {
                flex-direction: column-reverse;
                align-items: center;
                justify-content: center;
              }
              .footer_version_content_warp{
                max-width: 12.8rem;
                display: flex;
                color: rgba(255,255,255,0.5);
                align-items: center;
                @media (max-width: 820px) {
                  flex-direction: column;
                  justify-content: center;
                  margin-top: 0.16rem;
                }
                .irisnet_logo{
                  margin-right: 0.32rem;
                  @media (max-width: 820px) {
                    margin-right: 0;
                  }
                  a{
                    display: inline-block;
                    height: 0.4rem;;
                    img{
                      height: 100%;
                    }
                  }
                }
                .footer_copyright_wrap {
                  flex: 1;
                  text-align: center;
                  @media (max-width: 820px) {
                    margin: 0.08rem 0.2rem 0  0.2rem;
                  }
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
              .community_container{
                margin-left: 0.7rem;
                display: flex;
                align-items: flex-start;
                justify-content: center;
                flex-direction: column;
                @media (max-width: 820px) {
                  margin-left: 0;
                }
                .community_title{
                  font-size: 0.16rem;
                  padding-bottom: 0.1rem;
                  color: #fff;
                }
                .community_list_content{
                  display: flex;
                  a{
                    background: url("../../assets/button_bg_regular.png") no-repeat center center;
                    background-size: 100%;
                    margin:0 0.08rem 0 0;
                    box-sizing: border-box;
                    padding: 0.08rem;
                    display: inline-block;
                    &:hover{
                      background: url("../../assets/button_bg_highlight.png") no-repeat center center;
                      background-size: 100%;
                      .iconfont{
                        color: rgba(0, 0, 0, 1);
                      }
                    }
                    .iconfont {
                      font-size: 0.16rem;
                      color: rgba(255, 255, 255, 0.75);
                    }
                  }
                  a:last-child{
                    margin-right: 0;
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
            }


        }
    }
</style>
