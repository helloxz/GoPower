<template>
	<el-row :gutter="0" style="margin-top:2em;">
	    <el-col :xs="{span: 22, offset: 1}" :sm="{span: 18, offset: 3}">
			<div class="spacing">
			    <el-input v-model="s_data.shutdown_api" placeholder="请输入API地址,格式为http://192.168.20.3:886">
			      <template #prepend>API地址</template>
			    </el-input>
			</div>
			<div class="spacing">
			    <el-input v-model="s_data.shutdown_api_key" placeholder="请输入API KEY">
			      <template #prepend>KEY</template>
			    </el-input>
			</div>
			
			<div class="spacing">
				<el-button type="danger" class="mybtn" @click="manual_shutdown">手动关机</el-button>
			</div>
		</el-col>
	</el-row>		
</template>

<script setup>
	import { ref,watch,onMounted,computed } from 'vue'
	import {
	  Check,
	  Delete,
	  Edit,
	  Message,
	  Search,
	  Star,
	} from '@element-plus/icons-vue'
	import { ElMessage } from 'element-plus'
	import axios from 'axios'
	
	const s_data = ref({
		"shutdown_api":"",
		"shutdown_api_key":"",
	})
	
	//唤醒事件
	function manual_shutdown() {
		//验证参数是否合法
		if( !v_params() ) {
			return false;
		}
		
		//数据保存到浏览器方便下次使用
		localStorage.setItem("s_data",JSON.stringify(s_data.value));
		
		let api = s_data.value.shutdown_api + "/api/shutdown"
		let api_key = s_data.value.shutdown_api_key;
		
		
		const params = new URLSearchParams();
		params.append('key', api_key);
		
		axios.post(api, params)
			  .then(function (response) {
				let data = response.data
				if( data.code == 200 ) {
					ElMessage({
						message: '执行成功，将在1分钟后关机！',
						type: 'success',
					})
				}
				else{
					ElMessage.error(data.msg)
				}
			  })
			  .catch(function (error) {
				ElMessage.error("网络请求失败！")
			});
	}
	
	onMounted(()=>{
		//加载缓存数据
		let value = localStorage.getItem("s_data");
		
		//赋值数据
		if( value !== null ) {
			value = JSON.parse(value)
			s_data.value = value
		}
		
	})
	
	//验证参数是否合法
	function v_params() {
		//判断唤醒api地址
		let url_pattern = /^(http:\/\/|https:\/\/).+[0-9a-zA-Z]$/
		if( s_data.value.shutdown_api == "" ) {
			ElMessage.error("关机API地址不能为空！");
			return false;
		}
		if( !url_pattern.test(s_data.value.shutdown_api) ) {
			ElMessage.error("关机API地址格式不正确！");
			return false;
		}
		//判断key
		if( s_data.value.shutdown_api_key == "" || s_data.value.shutdown_api_key === undefined ) {
			ElMessage.error("API KEY不能为空！");
			return false;
		}
		else{
			return true;
		}
	}
</script>

<style>
</style>