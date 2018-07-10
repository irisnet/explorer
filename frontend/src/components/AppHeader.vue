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
                    >
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
                    <span class="nav_item">Resources</span>
                </div>
            </div>

        </header>
        <div class="app_header_mobile" v-if="devicesShow === 0">
            这是移动端显示端头部
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
            }
        },
        beforeMount() {
            if (this.devicesWidth > 500) {
                this.devicesShow = 1;
            } else {
                this.devicesShow = 0;
            }
        },
        methods: {
            searchClick() {
                console.log(this.searchValue)
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
    }



</style>
