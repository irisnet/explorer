<template>
    <loading-error-dialog>
        <img class="loading-img" :src="loadingImg" alt="" />
        <span class="loading-text">{{ loadingText }}</span>
    </loading-error-dialog>
</template>

<script>
import LoadingErrorDialog from "../loadingErrorDialog/LoadingErrorDialog";
export default {
    name: "LoadingComponent",
    data() {
        return {};
    },
    props: ["loadingImg", "loadingText"],
    components: {
        LoadingErrorDialog,
    },
};
</script>

<style scoped lang="scss">
@import "../../style/mixin";
.loading-img {
    width: 42px;
    height: 40px;
}
.loading-text {
    margin-top: 0.2rem;
    font-size: 0.14rem;
    font-family: ArialMT;
    color: #fff;
    line-height: 0.18rem;
}
</style>
