<template>
    <div class="wechat_qr_container" v-show="$store.state.flShowQR" @click="hideQRCode">
        <div class="wechat_qr_content">
            <div class="wechat_qr_wrap" v-if="$store.state.setQrImg === 'wechat'">
                <img class="wechat_qr_img" src="../../assets/wechat_qr.png">
                <h2 class="wechat_qr_title">Scan QR Code</h2>
                <p class="wechat_qr_content">to follow our SubScriptions</p>
            </div>
            <div class="wechat_qr_wrap" v-if="$store.state.setQrImg === 'qq'">
                <img v-if="true" class="qq_qr_img" src="../../assets/qq_qr.png">
                <h2 class="wechat_qr_title">Join QQ group</h2>
            </div>
        </div>
    </div>
</template>

<script>
	export default {
		name: "qrComponent",
        data(){
			return{

            }
        },
        methods: {
	        hideQRCode(){
	        	this.$store.commit('flShowQR',false)
            }
        }
	}
</script>

<style scoped lang="scss">
    .wechat_qr_container{
        width: 100%;
        height: 100%;
        background: rgba(0, 0, 0, .6);
        display: flex;
        align-items: center;
        justify-content: center;
        position: fixed;
        z-index: 10000;
        .wechat_qr_content{
            background: #141426;
            border-radius: 0.05rem;
            display: flex;
            align-items: center;
            justify-content: center;
            .wechat_qr_wrap{
                width: 3.7rem;
                height: 3.7rem;
                display: flex;
                flex-direction: column;
                align-items: center;
                justify-content: center;
                .wechat_qr_img{
                    width: 1.8rem;
                    height:1.8rem;
                }
                .qq_qr_img{
                    width: 1.8rem;
                    height:1.8rem;
                }
                .wechat_qr_title{
                    font-size: 0.24rem;
                    color: #5c5c7c;
                    margin-top: 0.1rem;
                }
                .wechat_qr_content{
                    font-size: 0.18rem;
                    color: #5c5c7c;
                    margin-top: 0.1rem;
                }
            }
        }
    }

</style>
