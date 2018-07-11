<!--头部页面-->
<template>
    <div class="app_header">
        <header class="app_header_person_computer" v-if="devicesShow === 1">
            <div class="imageWrap">
                <img src="../assets/logo.png" alt="失去网络了..."/>
            </div>
            <div class="useFeature">
                <div class="navSearch">
                    <input type="text" class="search_input"
                           placeholder="Search by Address / Txhash / Block"
                           v-model="searchInputValue"
                           @keyup.enter="onInputChange"
                    >
                    <i class="search_icon" @click="getData(searchInputValue)"></i>
                </div>
                <div class="navButton">
                    <span class="nav_item">Home</span>
                    <span class="flag">|</span>
                    <span class="nav_item">Block</span>
                    <b-dropdown id="ddown-left" text="Transaction" variant="primary" class="m-2">
                        <b-dropdown-item >Action</b-dropdown-item>
                        <b-dropdown-item >Another action</b-dropdown-item>
                        <b-dropdown-item >Something else here</b-dropdown-item>
                    </b-dropdown>
                    <span class="flag">|</span>
                    <b-dropdown text="Left align" variant="Validators" class="m-2">
                        <b-dropdown-item >Action</b-dropdown-item>
                        <b-dropdown-item >Another action</b-dropdown-item>
                        <b-dropdown-item >Something else here</b-dropdown-item>
                    </b-dropdown>
                    <span class="flag">|</span>
                    <span class="nav_item">Nodes</span>
                    <span class="flag">|</span>
                    <span class="nav_item" style='padding-right: 0'>Resources</span>
                </div>
            </div>

        </header>
        <div class="app_header_mobile" v-if="devicesShow === 0">
            <div class="feature_btn" @click="showFeature"></div>
            <div class="image_wrap_mobile">
                <!--<img src="../assets/logo.png" alt="失去网络了..."/>-->
                logo
            </div>

            <div class="use_feature_mobile" v-show="featureShow">
                <span class="feature_btn_mobile feature_nav" @click="featureButtonClick('/home')">Home</span>
                <span class="feature_btn_mobile feature_nav" @click="featureButtonClick('/blocks')">Blocks</span>
                <span class="feature_btn_mobile feature_nav" @click="transactionShow =! transactionShow">Transactions</span>
                <span class="feature_btn_mobile feature_subNav" v-show="transactionShow" @click="featureButtonClick('/recent_transactions')">Recent Transactions</span>
                <span class="feature_btn_mobile feature_subNav" v-show="transactionShow" @click="featureButtonClick('/transfer_transactions')">Transfer Transactions</span>
                <span class="feature_btn_mobile feature_subNav" v-show="transactionShow" @click="featureButtonClick('/delegate_transactions')">Delegate Transactions</span>
                <span class="feature_btn_mobile feature_nav"  @click="validatorsShow =! validatorsShow">Validators</span>
                <span class="feature_btn_mobile feature_subNav"  v-show="validatorsShow" @click="featureButtonClick('/validators')">Validators</span>
                <span class="feature_btn_mobile feature_subNav"  v-show="validatorsShow" @click="featureButtonClick('/candidates')">Candidates</span>
                <span class="feature_btn_mobile feature_nav" @click="featureButtonClick('/faucet')">Faucet</span>
            </div>
            <div class="search_input_mobile">
                <input type="text" class="search_input"
                       placeholder="Search by Address / Txhash / Block"
                >
            </div>
        </div>
    </div>
</template>
<script>

    export default {
        name: 'app-header',
        data() {
            return {
                devicesWidth: window.innerWidth,
                devicesShow: 1,//1是显示pc端，0是移动端
                searchValue: '',
                featureShow:false,//是否显示功能菜单栏
                transactionShow:false,//点击显示Transactions菜单
                validatorsShow:false,//点击显示validators菜单
                searchInputValue:'',
            }
        },
        beforeMount() {
            if (this.devicesWidth > 500) {
                this.devicesShow = 1;
            } else {
                this.devicesShow = 0;
            }
        },
        mounted(){
            document.getElementById('router_wrap').addEventListener('click',this.hideFeature)
        },
        beforeDestroy(){
            document.getElementById('router_wrap').removeEventListener('click',this.hideFeature)
        },
        methods: {
            hideFeature(){
                if(this.featureShow){
                    this.featureShow = false;
                }
            },
            showFeature(){
                this.featureShow = !this.featureShow;
            },
            featureButtonClick(path){
                this.featureShow = !this.featureShow;
                console.log(path)

            },
            getData(data){
                console.log(data)
            },
            onInputChange(){
                console.log(this.searchInputValue)
            }


        }
    }
</script>
<style lang="scss">
    @import '../style/mixin.scss';
    .app_header{
        height: 100px;
        width: 100%;
        display:flex;
        flex-direction:column;
        align-items: center;
        background: yellow;
        .app_header_person_computer {
            width:80%;
            min-width:600px;
            max-width:1000px;
            height:100%;
            border:1px solid red;
            @include flex();
            justify-content: space-between;
            padding-top:10px;
            .imageWrap {
                width: 100px;
                height: 50px;
                img {
                    width: 100px;
                    height: 50px;
                }
            }
            .useFeature{
                @include flex();
                flex-direction:column;
                align-items: flex-end;
                .navSearch {
                    margin-bottom:10px;
                    position:relative;
                    input::-webkit-input-placeholder{
                        text-align:center;
                        font-size:13px;
                        color:#cccccc;
                    }

                    .search_input{
                        @include borderRadius(4px);
                        width:300px;
                        height:28px;
                        line-height:28px;
                        text-indent:10px;
                        outline:none;
                        border:1px solid #dddddd;
                    }
                    .search_icon{
                        position:absolute;
                        top:5px;
                        right:10px;
                        width:15px;
                        height:15px;
                        background: lawngreen;
                    }
                }
                .navButton {
                    @include flex();
                    .nav_item{
                        display:inline-block;
                        height:30px;
                        line-height:30px;
                        padding:0 15px;
                        text-align: center;
                        font-size:14px;
                        cursor:pointer;
                    }
                    .m-2{
                        margin:0 !important;
                        button{
                            padding:0 15px;
                        }
                        .btn{
                            background-color: transparent;
                            color:#dddddd;
                            border:none;
                            font-weight:normal;
                        }
                    }
                    .flag{
                        display:inline-block;
                        height:30px;
                        line-height:30px;
                        font-size:12px;
                        color:#dddddd;
                    }
                }

            }

        }
        .app_header_mobile{
            width:100%;
            padding:10px 0;
            @include flex();
            flex-direction:column;
            position:relative;
            height:100px;
            .feature_btn{
                position:absolute;
                width:40px;
                height:40px;
                top:0;
                right:0;
                background: lawngreen;
            }
            .image_wrap_mobile{
                width: 70vw;
                height:40px;
                img{
                    width:100%;
                    height:100%;
                }
            }
            .search_input_mobile{
                width:100%;
                margin-top:10px;
                @include flex();
                flex-direction:column;
                align-items:center;
                input::-webkit-input-placeholder{
                    text-align:center;
                    font-size:14px;
                    color:#cccccc;
                    line-height:28px;
                }
                input{
                    width:95%;
                    @include borderRadius(5px);
                    border:1px solid #eee;
                    font-size:14px;
                    &:focus{
                        border:1px solid #3190e8;
                        outline:none;
                    }
                }
            }
            .use_feature_mobile{
                position:absolute;
                width:100%;
                top:100px;
                left:0;
                background:#3190e8;
                @include flex();
                flex-direction:column;
                .feature_btn_mobile{
                    border-bottom:1px solid #eee;
                    height:39px;
                    line-height:39px;
                    padding-left:15px;
                    color:#555;
                }
                .feature_subNav{
                    padding-left:30px;
                }
            }
        }
    }



</style>
