<template>
    <div class="copy_container">
        <img @click="handleCopy(text,$event)" src="../../assets/copy_icon.png" alt="copy" />
        <div class="tips" ref="tip">
            {{tipText}}
            <i></i>
        </div>
    </div>
</template>

<script>
import Clipboard from "clipboard";
export default {
    name: "MClip",
    props: {
        text: {
            type: String,
            default: ""
        }
    },
    data() {
        return {
            timer: null,
            tipText: "Copied",
            resizeTimer: null
        };
    },
    methods: {
        handleCopy(text, event) {
            this.handleClipboard(
                text,
                event,
                this.successHandle,
                this.errorHandle
            );
        },
        successHandle() {
            this.tipText = "Copied";
            this.showTip();
        },
        errorHandle() {
            this.tipText = "failed";
            this.showTip();
        },
        showTip() {
            clearTimeout(this.timer);
            this.$refs.tip.classList.add("show_tips");
            this.timer = setTimeout(() => {
                this.$refs.tip.classList.remove("show_tips");
            }, 600);
        },
        computedTipPosition() {
            clearTimeout(this.resizeTimer);
            this.resizeTimer = setTimeout(() => {
                let w = document.querySelector(".router_view").offsetWidth;
                let tip = this.$refs.tip.getBoundingClientRect();
                let r = tip.right;
                let tip_w = tip.width / 2;
                r = r > w ? r - w : 0;
                let i = this.$refs.tip.querySelector("i");
                let l = tip.left;
                if (r > 0) {
                    this.$refs.tip.style.transform = `translateX(-${tip_w +
                        r}px)`;
                    i.style.transform = `translateX(${r}px)`;
                } else if (l < 0) {
                    this.$refs.tip.style.transform = `translateX(-${tip_w +
                        l}px)`;
                    i.style.transform = `translateX(${l}px)`;
                } else {
                    this.$refs.tip.style.transform = `translateX(-50%)`;
                    i.style.transform = `translateX(0px)`;
                }
            }, 500);
        },
        handleClipboard(text, event, success, error) {
            const clipboard = new Clipboard(event.target, {
                text: () => text
            });
            clipboard.on("success", () => {
                success();
                clipboard.off("error");
                clipboard.off("success");
                clipboard.destroy();
            });
            clipboard.on("error", () => {
                error();
                clipboard.off("error");
                clipboard.off("success");
                clipboard.destroy();
            });
            clipboard.onClick(event);
        }
    },
    mounted() {
        window.addEventListener("resize", this.computedTipPosition);
        this.computedTipPosition();
    },
    destroyed() {
        window.removeEventListener("resize", this.computedTipPosition);
    }
};
</script>

<style lang="scss" scoped>
.copy_container {
    position: relative;
    display: inline;
    img {
        width: 0.11rem;
        height: 0.12rem;
        cursor: pointer;
        vertical-align: baseline;
    }
    .tips {
        opacity: 0;
        position: absolute;
        padding: 0rem 0.15rem;
        top: -0.36rem;
        left: 50%;
        transform: translateX(-50%);
        color: #fff;
        background: rgba(0, 0, 0, 1);
        border-radius: 0.04rem;
        z-index: 10;
        line-height: 32px;
        font-size: 0.14rem;
        white-space: nowrap;
        transition: opacity 0.3s ease-in;
        i {
            width: 0;
            height: 0;
            border: 0.06rem solid transparent;
            content: "";
            display: block;
            position: absolute;
            border-top-color: #000000;
            left: 50%;
            margin-left: -6px;
        }
    }
    .show_tips {
        opacity: 1;
        animation: tip_fading_out 0.3s ease-in 0.3s;
    }
    @keyframes tip_fading_out {
        0% {
            opacity: 1;
        }
        100% {
            opacity: 0;
        }
    }
}
</style>
