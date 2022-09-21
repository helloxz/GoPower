<template>
  <el-menu
    :default-active="activeIndex"
    class="h-menu"
    mode="horizontal"
    :ellipsis="false"
    @select="handleSelect"
  >
    <el-menu-item index="0"><router-link to="/"><span style="font-weight: bold;">GoPower</span></router-link></el-menu-item>
    <div class="flex-grow" />
    <el-menu-item index="1"><router-link to="/wol">手动唤醒</router-link></el-menu-item>
	<el-menu-item index="2"><router-link to="/shutdown">手动关机</router-link></el-menu-item>
	<el-menu-item>
		<a href="https://github.com/helloxz/GoPower" title="" rel="nofollow" target="_blank">
			Github
		</a>
	</el-menu-item>
  </el-menu>
</template>

<script lang="ts" setup>
import { ref } from 'vue'

const activeIndex = ref('0')
const handleSelect = (key: string, keyPath: string[]) => {
	activeIndex.value = key;
}
</script>

<style>
.flex-grow {
  flex-grow: 1;
}
.h-menu h1{
	font-size: 24px;
}
.h-menu a{
	text-decoration:none;
}
</style>