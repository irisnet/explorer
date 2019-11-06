<template>
    <div>
        <div class="error_container">
            <span class="error_content_text">
                <img src="../../assets/error_page.png">
            </span>
            <p class="error_page_hint">page not found</p>
            <button class="error_content_btn" @click="goHome()">HOME</button>
        </div>
    </div>
</template>

<script>
	export default {
		name: "ErrorPage",
        data () {
            return {

            }
        },
        methods:{
	        goHome(){
		        this.$router.push("/home")
            }
        }
	}
</script>

<style scoped lang="scss">
    .error_container{
        position: absolute;
        left: 50%;
        top: 50%;
        transform: translate(-50%,-50%);
        display: flex;
        flex-direction: column;
        align-items: center;
        .error_content_text{
            display: inline-block;
            width: 70%;
            img{
                width: 100%;
                font-size: 4rem;
                color: #E5ECFF;
            }
        }
        .error_page_hint{
            margin-top: 0.2rem;
            color: var(--contentColor);
        }

        .error_content_btn{
            margin-top: 0.5rem;
            max-width: 1.5rem;
            padding: 0.1rem 0.2rem;
            background: var(--bgColor);
            border: 0.01rem solid var(--bgColor);
            border-radius: 0.06rem;
            color: #fff;
        }
    }
</style>
