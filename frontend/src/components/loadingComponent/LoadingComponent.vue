<template>
    <div class="spin-component" v-show="showLoading">
        <div class="spin-content">
            <img style="with:48px;height: 48px;" src="../../assets/loading.gif">
            <p>Loading,please wait...</p>
        </div>
    </div>
</template>

<script>
	export default {
		name: "LoadingComponent",
        data() {
			return {

            }
        },
		props: ['showLoading'],
	}
</script>

<style scoped lang="scss">
    @import '../../style/mixin';
    .spin-component {
        position: fixed;
        top: 50%;
        left: 50%;
        transform: translate(-50%,-50%);
        z-index: 1000;
        @include flex;
        align-items: center;
        justify-content: center;
        .spin-content {
            background: rgba(255, 255, 255, 1);
            box-shadow:0px 0px 6px 0px rgba(10, 97, 142, 0.15);
            width: 2.54rem;
            height: 1.32rem;
            @include borderRadius(0.1rem);
            @include flex;
            flex-direction: column;
            align-items: center;
            justify-content: center;
            img {

            }
            p {
                color: #000;
                font-size: 0.14rem;
                font-family: ArialMT;
                margin-top: 0.19rem;
            }
        }
    }
</style>
