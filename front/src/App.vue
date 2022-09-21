<script setup>
import Menu from "../src/components/Menu.vue"
import Footer from "../src/components/Footer.vue"
</script>

<template>
	<!-- 公共菜单 -->
	<Menu></Menu>
	<!-- 公共菜单END -->
	<!-- 这里是要显示的内容 -->
	<RouterView />
	<!-- 这里是要显示的内容END -->
	<Footer></Footer>
</template>

<style scoped>
	
</style>
