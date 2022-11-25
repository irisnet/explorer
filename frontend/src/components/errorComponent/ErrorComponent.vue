<template>
    <loading-error-dialog>
        <span class="error-text">{{ errorText }}</span>
        <span class="error-btn" @click="updateRefresh">{{ errorBtnText }}</span>
    </loading-error-dialog>
</template>

<script>
import LoadingErrorDialog from "../loadingErrorDialog/LoadingErrorDialog";
export default {
    name: "ErrorComponent",
    data() {
        return {};
    },
    props: ["errorText", "errorBtnText"],
    methods: {
        updateRefresh: function () {
            this.$emit("refresh");
        },
    },
    components: {
        LoadingErrorDialog,
    },
};
</script>

<style scoped lang="scss">
@import "../../style/mixin";
.error-text {
    font-size: 0.14rem;
    font-family: ArialMT;
    color: #fff;
    line-height: 0.18rem;
    white-space: pre-line;
}
.error-btn {
    @include flex;
    justify-content: center;
    align-items: center;
    margin-top: 0.22rem;
    padding: 3px 0;
    min-width: 100px;
    min-height: 24px;
    font-size: 0.14rem;
    font-family: ArialMT;
    color: #fff;
    line-height: 0.18rem;
    background: linear-gradient(315deg, rgba(33, 38, 121, 0.7) 0%, rgba(15, 13, 59, 0.7) 100%);
    border: 1px solid #4153cd;
    border-radius: 20%/50%;
    cursor: pointer;
}
</style>
