<template>
    <div type="light" class="facet_wrap" :style="showTitle?'':'padding-top:0.38rem;'">
        <div class="faucet text-center" :style="innerWidth<=500?'padding-top:0;':''">
            <div
                class="faucet-title"
                :style="{marginTop: $store.state.isMobile ? '0.6rem' : '0'}"
            >Faucet</div>
            <div class="coin">
                <img src="../../assets/coin.png" />
            </div>
            <p class="hint_content">Get IRIS from this faucet for the latest IRISnet Testnet.</p>
            <p class="hint_content">This faucet will send 10 IRIS to any valid testnet address.</p>
            <p
                class="Balance_number"
                :class="errStyle ? 'err_red' : 'err_black ' "
            >Balance:{{faucetBalance}} {{tokenName}}</p>
            <br />
            <form @submit.prevent="apply">
                <div class="faucet-form">
                    <input type="text" class="form-control" id="token" name="token" hidden />
                    <input
                        type="text"
                        class="form-control"
                        id="session_id"
                        name="session_id"
                        hidden
                    />
                    <input type="text" class="form-control" id="sig" name="sig" hidden />
                    <div class="form-group">
                        <input
                            type="text"
                            class="form-control"
                            id="address"
                            v-model="address"
                            placeholder="Please enter the collection address"
                        />
                        <div
                            class="alert_information"
                            :style="{visibility:alertShowErrMsg}"
                        >{{errMsg}}</div>
                    </div>
                    <button
                        id="submit"
                        type="submit"
                        class="btn btn-primary"
                        :disabled="btnDisabled"
                        :class="showSendingImg ? 'waitingStyle' : ''"
                    >
                        {{btninfo}}
                        <span v-show="showSendingImg" style="padding: 0 0.06rem">
                            <img src="../../assets/wait.gif" />
                        </span>
                    </button>
                </div>
            </form>
        </div>
        <div class="alert_block" v-if="showAlert">
            <div class="img_font_container">
                <div class="img_container">
                    <img v-if="showSuccess" src="../../assets/success.png" />
                    <img v-if="!showSuccess" src="../../assets/x.png" />
                </div>
                <span class="font_style" v-if="showSuccess">Submitted successfully!</span>
                <span class="font_style" v-if="!showSuccess">Failed,try again.</span>
            </div>
        </div>
    </div>
</template>

<script>
import axios from "axios";
import Tools from "../../util/Tools";
import Constant from "../../constant/Constant";

export default {
    name: "FaucetPage",
    $route() {
        this.showTitle = !(
            this.$route.query.flShow &&
            this.$route.query.flShow === "false" &&
            !Tools.currentDeviceIsPersonComputer()
        );
    },
    watch: {
        address(address) {
            if (this.insufficientBalanceStatus === false) {
                if (
                    this.$Crypto
                        .getCrypto(
                            Constant.CHAINNAME,
                            this.$store.state.currentEnv
                        )
                        .isValidAddress(address)
                ) {
                    this.btnDisabled = false;
                    this.alertShowErrMsg = "hidden";
                } else {
                    this.btnDisabled = true;
                    this.errMsg = "Please enter a valid address";
                    this.alertShowErrMsg = "visible";
                }
            }
        }
    },
    data() {
        return {
            address: "",
            errMsg: "",
            alertShowErrMsg: "hidden",
            innerWidth: window.innerWidth,
            showTitle: !(
                this.$route.query.flShow &&
                this.$route.query.flShow === "false" &&
                !Tools.currentDeviceIsPersonComputer()
            ),
            faucetBalance: 0,
            insufficientBalanceStatus: false,
            tokenName: "",
            invalidAddress: false,
            errStyle: false,
            btnDisabled: true,
            btninfo: "Send me IRIS",
            showSendingImg: false,
            showAlert: false,
            showSuccess: false
        };
    },
    beforeCreate() {
        axios
            .get("/api/faucet/account")
            .then(data => {
                if (data.status === 200) {
                    return data.data;
                }
            })
            .then(data => {
                this.errStyle = false;
                this.faucetBalance = `${
                    Tools.formatAccountCoinsAmount(data.coins[0])[0].split(
                        "."
                    )[0]
                }`;
                let faucetQuota = 20;
                if (this.faucetBalance < faucetQuota) {
                    this.errStyle = true;
                    this.btnDisabled = true;
                    this.btninfo = "Insufficient Balance";
                    this.insufficientBalanceStatus = true;
                } else {
                    this.insufficientBalanceStatus = false;
                    this.btninfo = "Send me IRIS";
                }
                this.tokenName = Tools.formatAccountCoinsDenom(
                    data.coins[0]
                )[0].toUpperCase();
            })
            .catch(e => {
                console.log(e);
                this.faucetBalance = "Error";
                this.errStyle = true;
                this.btnDisabled = true;
                this.insufficientBalanceStatus = true;
            });
    },
    created() {
        let addr = this.$route.query.address;
        if (addr && addr !== "") {
            this.address = addr;
        }
    },
    methods: {
        apply() {
            if (!this.errMsg && this.alertShowErrMsg === "hidden") {
                this.alertShowErrMsg = "visible";
                setTimeout(() => {
                    if (this.alertShowErrMsg === "visible") {
                        this.alertShowErrMsg = "hidden";
                        this.errMsg = "";
                    }
                }, 2000);
            }
            this.btninfo = "Sending";
            this.btnDisabled = true;
            this.showSendingImg = true;
            axios
                .post(
                    "/api/faucet/apply/" + this.address,
                    JSON.stringify({
                        address: this.address
                    })
                )
                .then(result => {
                    let data = result.data;
                    let that = this;
                    this.btninfo = "Send me IRIS";
                    this.showSendingImg = false;
                    this.showAlert = true;
                    if (data.code) {
                        this.showSuccess = false;
                    } else {
                        this.showSuccess = true;
                    }
                    setTimeout(function() {
                        that.showAlert = false;
                        that.showSuccess = false;
                        location.reload();
                    }, 2000);
                })
                .catch(e => {
                    console.log(e);
                    this.btninfo = "Send me IRIS";
                    this.showSendingImg = false;
                    this.showAlert = true;
                    this.showSuccess = false;
                    let that = this;
                    setTimeout(function() {
                        that.showAlert = false;
                        that.showSuccess = false;
                        location.reload();
                    }, 2000);
                });
        },
        resize() {
            this.innerWidth = window.innerWidth;
        }
    },
    mounted() {
        window.addEventListener("resize", this.resize);
    },
    beforeDestroy() {
        window.removeEventListener("resize", this.resize);
    }
};
</script>

<style lang="scss" scoped>
@import "../../style/mixin";

.faucet {
    max-width: 12.8rem;
    margin: 0 auto;
    text-align: center;
    padding: 0.7rem 0;
    .coin {
        display: flex;
        justify-content: center;
        margin-bottom: 10px;
        margin-top: 0.35rem;
    }
    .hint_content{
        font-size:0.14rem;
        color:var(--contentColor);
        padding:0 0.1rem;
    }
    .faucet-title {
        font-size: 22px;
    }
}

.faucet-form {
    margin: 0 auto;
    width: 3rem;

    .form-group {
        text-align: left;
        margin-bottom: 0 !important;
        height: 63px;
        .alert_information {
            display: inline-block;
            height: 0.14rem;
            line-height: 0.14rem;
            color: red;
            font-size: 0.14rem;
            margin-top: 0.05rem;
        }
        .form-control {
            height: 0.36rem;
            line-height: 0.36rem;
            font-size: 0.14rem;
            @include borderRadius(0.04rem);
        }
    }
    .btn-primary {
        border:none;
        margin-top: 0.2rem;
        /*box-shadow: 0 0 0 transparent;*/
        @include borderRadius(0.04rem);
        padding: 0.09rem 0.19rem !important;
        background: var(--bgColor);
        @include fontSize;
        color: #fff;
        &:disabled {
            background: rgba(214, 217, 224, 1);
            border-color: rgba(214, 217, 224, 1);
        }
    }
}
.err_red {
    color: red !important;
}
.err_black {
    color: #000;
}
.err-msg {
    color: red;
}
.btn {
    width: 1.58rem;
    height: 0.36rem;
    font-size: 0.14rem !important;
}
.form-control {
    padding: 0.0375rem 0.075rem !important;
    width: 3rem !important;
    @include borderRadius(0.025rem);
    font-size: 0.14rem !important;
}
a,
input {
    -webkit-box-shadow: 0 0 0 transparent !important;
    -moz-box-shadow: 0 0 0 transparent !important;
    box-shadow: 0 0 0 transparent !important;
}

#sc {
    height: 0.52rem;

    #SM_BTN_WRAPPER_1 {
        height: 0.52rem;
        #SM_BTN_1 {
            @include flex;
            align-items: center;
            height: 0.52rem !important;
        }
        #sm-btn-bg {
            height: 0.52rem !important;
        }
    }
}
.Balance_number {
    font-size: 0.14rem;
    color: #000;
    padding-top: 0.04rem;
}
.faucet_address {
    font-size: 0.14rem;
    color: #000;
}
.btn-primary:focus {
    -webkit-box-shadow: 0 0 0 0.2rem rgba(255, 255, 255, 0.5);
    box-shadow: 0 0 0 0.2rem rgba(255, 255, 255, 0.5);
}
.alert_block {
    width: 2.54rem;
    height: 1.32rem;
    background: rgba(0, 0, 0, 0.6);
    @include fixedCenter;
    border-radius: 0.1rem;
    z-index: 2000;
    .img_container {
        width: 36px;
        height: 36px;
        margin: auto;
        margin-bottom: 0.1rem;
        img {
            width: 100%;
        }
    }
    .font_style {
        color: #fff;
        @include fontSize;
    }
    .img_font_container {
        width: 100%;
        text-align: center;
        @include center;
        padding-top: 0.02rem;
    }
}
.waitingStyle {
    color: #000 !important;
}
#address {
    border: 0.01rem solid #d7d9e0;
}
#address:-webkit-autofill {
    background-color: rgb(255, 255, 255) !important;
    background-image: none !important;
    color: rgb(0, 0, 0) !important;
}
</style>
