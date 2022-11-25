<template>
    <div class="loading-error-dialog">
        <div class="loading-error-outline">
            <div class="loading-error-inline"></div>
            <div class="loading-error-content">
                <slot></slot>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    name: "LoadingErrorDialog",
    data() {
        return {};
    },
};
</script>

<style scoped lang="scss">
@import "../../style/mixin";
.loading-error-dialog {
    @include flex;
    justify-content: center;
    align-items: center;
    width: 100%;
    height: 100%;
    background: url("../../assets/GOZ_bg.png") no-repeat center 35% / cover;
    z-index: 1000;
    .loading-error-outline {
        @include flex;
        position: relative;
        padding: 4px;
        min-width: 260px;
        min-height: 132px;
        border: 1px solid;
        border-image: linear-gradient(180deg, rgba(74, 97, 231, 0.6), rgba(74, 97, 231, 0.35)) 1 1;
        background-color: transparent;
        transform: translateY(-40px);
        &::before {
            position: absolute;
            top: -2px;
            left: -2px;
            width: 10px;
            height: 10px;
            border-top: 2px solid #4a61e7;
            border-left: 2px solid #4a61e7;
            content: "";
        }
        &::after {
            position: absolute;
            position: absolute;
            top: -2px;
            right: -2px;
            width: 10px;
            height: 10px;
            border-top: 2px solid #4a61e7;
            border-right: 2px solid #4a61e7;
            content: "";
        }
        .loading-error-inline {
            position: absolute;
            left: 0;
            bottom: 0;
            width: 100%;
            &::before {
                position: absolute;
                left: -2px;
                bottom: -2px;
                width: 10px;
                height: 10px;
                border-left: 2px solid #4a61e7;
                border-bottom: 2px solid #4a61e7;
                content: "";
            }
            &::after {
                position: absolute;
                position: absolute;
                right: -2px;
                bottom: -2px;
                width: 10px;
                height: 10px;
                border-right: 2px solid #4a61e7;
                border-bottom: 2px solid #4a61e7;
                content: "";
            }
        }
        .loading-error-content {
            flex: 1;
            @include flex;
            flex-direction: column;
            justify-content: center;
            align-items: center;
            background: linear-gradient(315deg, rgba(33, 38, 121, 0.7) 0%, rgba(15, 13, 59, 0.7) 100%);
            border: 1px solid;
            border-image: linear-gradient(134deg, rgba(32, 37, 137, 1), rgba(16, 14, 51, 0), rgba(32, 37, 137, 1)) 1 1;
        }
    }
}
</style>
