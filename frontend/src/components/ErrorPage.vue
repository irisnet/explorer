<template>
    <div>
        <div class="error_container">
            <span class="error_content_text">404</span>
            <p>Page not found</p>
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
            font-size: 2rem;
        }
        .error_content_btn{
            max-width: 1.5rem;
            padding: 0.1rem 0.2rem;
            background: var(--bgColor);
            border: 0.01rem solid var(--bgColor);
            border-radius: 0.06rem;
            color: #fff;
        }
    }
</style>
