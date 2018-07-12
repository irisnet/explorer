<template>
    <div class="home_wrap">
        <div :class="pageClassName">
            <div :class="module_item_wrap">
                <div class="home_module_item home_module_item_status">
                    <article class="home_module_item_status_title">
                        <span>Blockchain Status</span>
                        <span>Fuxi 2000</span>
                    </article>
                    <section>
                        <div class="home_status_module_content_left">
                            <span>Market Cap :</span>
                            <span>¥ {{marketCapValue}}</span>
                            <span>Validators :</span>
                            <span>{{validatorsValue}}</span>
                            <span>Transactions :</span>
                            <span>{{transactionValue}}</span>
                        </div>
                        <div class="home_status_module_content_right">
                            <span>Price :</span>
                            <span>¥ {{priceValue}}</span>
                            <span>Voting Power :</span>
                            <span>{{votingPowerValue}}</span>
                        </div>
                    </section>
                    <p class="current_block">
                        Current Block
                    </p>
                    <div class="home_status_bottom">
                        <div class="home_status_bottom_content">
                            <span>Block Height :</span>
                            <span>{{blockHeightValue}}</span>
                        </div>
                        <div class="home_status_bottom_content">
                            <span>Timestamp</span>
                            <span>{{timestampValue}}</span>
                        </div>
                    </div>

                </div>
                <div class="home_module_item">transaction History</div>
            </div>
            <div :class="module_item_wrap">
                <div class="home_module_item">Validators</div>
                <div class="home_module_item">Blocks</div>
            </div>
          <div>hello</div>
          <div>hello</div>
          <div>hello</div>
          <div>hello</div>
          <div>hello</div>
          <div>hello</div>
          <div>hello</div>
          <div>hello</div>
          <div>hello</div>
          <div>hello</div>
          <div>hello</div>
          <div>hello</div>
          <div>hello</div>
          <div>hello</div>
          <div>hello</div>
          <div>hello</div>
          <div>hello</div>
          <div>hello</div>
          <div>hello</div>
          <div>hello</div>
          <div>hello</div>
          <div>hello</div>
          <div>hello</div>
          <div>hello</div>
          <div>hello</div>
          <div>hello</div>



        </div>
    </div>

</template>

<script>

    export default {
        name: 'app-header',
        data() {
            return {
                devicesWidth: window.innerWidth,
                pageClassName: 'personal_computer_home_wrap',//1是显示pc端，0是移动端
                module_item_wrap: 'module_item_wrap_computer',//module_item_wrap_computer是pc端，module_item_wrap_mobile是移动端
                marketCapValue: '3,222,333.02',
                validatorsValue: 'hello',
                transactionValue: 'world',
                priceValue: '23,222,444.21',
                votingPowerValue: 'hello world',
                blockHeightValue: 'aaa',
                timestampValue: 'bbb',

            }
        },
        beforeMount() {
            if (this.devicesWidth > 500) {
                this.pageClassName = 'personal_computer_home_wrap';
                this.module_item_wrap = 'module_item_wrap_computer';
            } else {
                this.pageClassName = 'mobile_home_wrap';
                this.module_item_wrap = 'module_item_wrap_mobile';
            }
        },
        mounted() {
            document.getElementById('router_wrap').addEventListener('click', this.hideFeature)
        },

        methods: {}
    }
</script>
<style lang="scss">
    @import '../style/mixin.scss';

    .home_wrap {
        width: 100%;
        @include flex();
        flex-direction: column;
        align-items: center;
        .personal_computer_home_wrap {
            width: 80%;
            min-width: 600px;
            max-width: 1000px;
            border: 1px solid #3190e8;
            margin-top: 20px;
            .module_item_wrap_computer {
                width: 100%;
                @include flex();
                justify-content: space-between;
                .home_module_item {
                    width: 45%;
                    margin-bottom: 10px;
                    border: 1px solid #eee;
                }
            }

        }
        .mobile_home_wrap {
            flex-direction: column;
            justify-content: space-between;
            width: 100%;
            padding: 10px;
            .module_item_wrap_mobile {
                @include flex();
                flex-direction: column;
                align-items: center;
                .home_module_item {
                    width: 98%;
                    margin-bottom: 10px;
                    border: 1px solid #eee;
                }
            }
        }
        .home_module_item_status {
            padding: 10px;
            background: #3190e8;
            .home_module_item_status_title {
                @include flex();
                border-bottom: 1px solid #ffffff;
                justify-content: space-between;
                span {
                    font-weight: 500;
                    color: #fff;
                    display: inline-block;
                    height: 28px;
                    line-height: 28px;
                }
            }
            section {
                @include flex();
                .home_status_module_content_left {
                    flex: 1;
                    @include flex();
                    flex-direction: column;
                    span {
                        &:nth-child(2n+1) {
                            margin-top: 10px;
                            font-size: 12px;
                            color: #ffffff;
                            font-weight: normal;
                        }
                        &:nth-child(2n) {
                            font-size: 12px;
                            color: #cccccc;
                        }
                    }
                }
                .home_status_module_content_right {
                    flex: 1;
                    @include flex();
                    flex-direction: column;
                    span {
                        &:nth-child(2n+1) {
                            margin-top: 10px;
                            font-size: 12px;
                            color: #ffffff;
                            font-weight: normal;
                        }
                        &:nth-child(2n) {
                            font-size: 12px;
                            color: #cccccc;
                        }
                    }
                }
            }
            .current_block {
                font-weight: 500;
                color: #fff;
                display: inline-block;
                height: 28px;
                line-height: 28px;
                border-bottom: 1px solid #ffffff;
                width: 100%;

            }
            .home_status_bottom {
                @include flex();

                .home_status_bottom_content {
                    flex: 1;
                    @include flex();
                    flex-direction:column;
                    span {
                        font-size: 12px;
                        color: #ffffff;
                        font-weight: normal;
                    }
                }

            }
        }
    }
</style>
