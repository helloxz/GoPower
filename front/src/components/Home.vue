<template>
	<el-row :gutter="0" style="margin-top:4em;">
	    <el-col :xs="{span: 22, offset: 1}" :sm="{span: 18, offset: 3}">
			<div class="spacing">
				<el-button type="primary" class="mybtn" @click="open_setting">添加配置</el-button>
			</div>
			
			<!-- 渲染一个表格 -->
			<el-table :data="s_datas" style="width: 100%">
			    <el-table-column fixed prop="mac" label="Mac地址" width="140" />
				<el-table-column fixed prop="note" label="备注" width="120" />
			    <el-table-column prop="ip" label="IP" width="130" />
				<el-table-column prop="runtime" label="运行时间" :formatter="formatter" width="100" />
			    <el-table-column prop="wol_api" label="唤醒API" width="220" />
				<!-- <el-table-column prop="wol_api_key" label="唤醒API KEY" width="120" /> -->
			    <el-table-column prop="shutdown_api" label="关机API" width="220" />
				<!-- <el-table-column prop="shutdown_api_key" label="关机API KEY" width="120" /> -->
				<el-table-column fixed="right" label="动作">
				  <template #default="scope">
				    <div><el-button type="success" size="small" @click="startup(scope.$index)">唤醒</el-button></div>
				    <div><el-button type="danger" size="small" @click="shutdown(scope.$index)">关机</el-button></div>
				  </template>
				</el-table-column>
			    <el-table-column fixed="right" label="操作">
			      <template #default="scope">
			        <div><el-button type="primary" size="small" @click="edit_sdata(scope.$index)">编辑</el-button></div>
			        <div><el-button type="danger" size="small" @click="del_sdata(scope.row,scope.$index)">删除</el-button></div>
			      </template>
			    </el-table-column>
			</el-table>
			<!-- 渲染表格END -->
			
		</el-col>
	</el-row>
	<!-- 这是抽屉层 -->
	<el-drawer
	    v-model="drawer"
	    title="GoPower参数设置"
		:size="drawer_size"
	  >
	    <!-- 设置表单 -->
		<el-form
		    :label-position="'top'"
		    label-width="100px"
		    style="max-width: 460px"
		  >
		    <el-form-item label="MAC">
		      <el-input v-model="s_data.mac" />
		    </el-form-item>
		    <el-form-item label="IP">
		      <el-input v-model="s_data.ip" />
		    </el-form-item>
			<el-form-item label="唤醒API地址">
			  <el-input v-model="s_data.wol_api" />
			</el-form-item>
			<el-form-item label="唤醒API KEY">
			  <el-input v-model="s_data.wol_api_key" />
			</el-form-item>
			<el-form-item label="关机API地址">
			  <el-input v-model="s_data.shutdown_api" />
			</el-form-item>
			<el-form-item label="关机API KEY">
			  <el-input v-model="s_data.shutdown_api_key" />
			</el-form-item>
			<el-form-item label="备注">
			  <el-input v-model="s_data.note" />
			</el-form-item>
		  </el-form>
		  <el-button type="primary" @click="add_setting" v-if="btn_status.add">添加</el-button>
		  <el-button type="primary" @click="save_setting" v-if="btn_status.save">保存</el-button>
		<!-- 设置表单END -->
		
	</el-drawer>
	<!-- 抽屉层END -->
</template>

<script setup>
import { ref,watch,onMounted,computed } from 'vue'
import axios from 'axios'
import {
  Check,
  Delete,
  Edit,
  Message,
  Search,
  Star,
} from '@element-plus/icons-vue'
import { ElMessage } from 'element-plus'


//抽屉默认关闭
const drawer = ref(false)
const kstatus = ref(false)
const mac = ref("")
const ip = ref("")
//数据全局索引
const global_index = ref()
//抽屉的宽度
const drawer_size = ref("460px")
//按钮状态
const btn_status = ref({
	"add":false,
	"save":false
})

const s_data = ref({
	"mac":"",
	"ip":"",
	"wol_api":"",
	"wol_api_key":"",
	"shutdown_api":"",
	"shutdown_api_key":"",
	"note":""
})
//声明运行时间
const run_time = ref({})
//设置参数合集
const s_datas = ref([])
//判断是否是手机
function isMobilePhone() {
	const ua = navigator.userAgent.toLowerCase();
	const t1 = /android|webos|iphone|ipod|blackberry|iemobile|opera mini/i.test(
	  ua
	);
	// const t2 = !ua.match("iphone") && navigator.maxTouchPoints > 1;
	return t1;
}

//清空临时数据
function clear_s_data() {
	s_data.value = {
		"mac":"",
		"ip":"",
		"wol_api":"",
		"wol_api_key":"",
		"shutdown_api":"",
		"shutdown_api_key":"",
		"note":""
	}
}
//验证参数是否合法
function v_params() {
	//如果MAC地址为空
	if(s_data.value.mac == "") {
		ElMessage.error("MAC地址不能为空！");
		return false;
	}
	//判断MAC地址是否合法
	let mac_pattern = /^[0-9a-zA-Z]{2}:[0-9a-zA-Z]{2}:[0-9a-zA-Z]{2}:[0-9a-zA-Z]{2}:[0-9a-zA-Z]{2}:[0-9a-zA-Z]{2}$/
	if( !mac_pattern.test(s_data.value.mac) ) {
		ElMessage.error("MAC地址格式不正确！");
		return false;
	}
	//判断IP是否正确
	if( s_data.value.ip == "" || s_data.value.ip == undefined ) {
		ElMessage.error("IP不能为空！");
		return false;
	}
	let ip_pattern = /^(25[0-5]|2[0-4]\d|[0-1]?\d?\d)(\.(25[0-5]|2[0-4]\d|[0-1]?\d?\d)){3}$/
	if( (s_data.value.ip != "") && !ip_pattern.test(s_data.value.ip) ) {
		ElMessage.error("IP地址格式不正确！");
		return false;
	}
	//判断唤醒api地址
	let url_pattern = /^(http:\/\/|https:\/\/).+[0-9a-zA-Z]$/
	if( s_data.value.wol_api == "" ) {
		ElMessage.error("唤醒API地址不能为空！");
		return false;
	}
	if( !url_pattern.test(s_data.value.wol_api) ) {
		ElMessage.error("唤醒API地址格式不正确！");
		return false;
	}
	//判断关机API地址
	if( s_data.value.shutdown_api == "" ) {
		ElMessage.error("关机API地址不能为空！");
		return false;
	}
	if( !url_pattern.test(s_data.value.shutdown_api) ) {
		ElMessage.error("关机API地址格式不正确！");
		return false;
	}
	//判断key
	if( s_data.value.wol_api_key == "") {
		ElMessage.error("唤醒API KEY不能为空！");
		return false;
	}
	if( s_data.value.shutdown_api_key == "" ) {
		ElMessage.error("关机API KEY不能为空！");
		return false;
	}
	else{
		return true;
	}
}
//手动唤醒
function manual_startup() {
	wol_dialog_visible.value = true;
}
//开机函数
function startup(index){
	let info = s_datas.value[index];
	//开机API
	let wol_api = info.wol_api + "/api/startup"
	
	const params = new URLSearchParams();
	params.append('key', info.wol_api_key);
	params.append('mac', info.mac);
	params.append('ip', info.ip);
	
	axios.post(wol_api, params)
	  .then(function (response) {
		let data = response.data
		if( data.code == 200 ) {
			ElMessage({
				message: '唤醒成功！',
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
//关机函数
function shutdown(index){
	let info = s_datas.value[index];
	//关机API
	let shutdown_api = info.shutdown_api + "/api/shutdown"
	const params = new URLSearchParams();
	params.append('key', info.shutdown_api_key);
	
	axios.post(shutdown_api, params)
		  .then(function (response) {
			let data = response.data
			if( data.code == 200 ) {
				ElMessage({
					message: '关机成功！',
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
//保存设置参数，接收数据的索引
function save_setting() {
	//改变数据
	s_datas.value[global_index] = {
		"mac":s_data.value.mac,
		"ip":s_data.value.ip,
		"wol_api":s_data.value.wol_api,
		"wol_api_key":s_data.value.wol_api_key,
		"shutdown_api":s_data.shutdown_api,
		"shutdown_api_key":s_data.shutdown_api_key,
		"note":s_data.note
	}
	//验证参数
	if( !v_params() ) {
		return false;
	}
	//重新保存数据
	localStorage.setItem("s_datas",JSON.stringify(s_datas.value));
	ElMessage({
		message: '设置已保存！',
		type: 'success',
	})
}
//打开空设置选项
function open_setting() {
	//清空数据
	clear_s_data()
	//显示添加按钮
	btn_status.value.add = true
	btn_status.value.save = false
	//打开抽屉
	drawer.value = true
}
//添加一个设置
function add_setting() {
	//先获取浏览器缓存数据
	let value = localStorage.getItem("s_datas");
	//转为对象
	value = JSON.parse(value)
	//如果数据不为空，则赋值给s_datas
	if( value !== null ) {
		s_datas.value = value;
	}
	//验证参数
	if( !v_params() ) {
		return false;
	}
	//追加数据
	s_datas.value.push(s_data.value);
	//重新保存到浏览器存储
	localStorage.setItem("s_datas",JSON.stringify(s_datas.value));
	//清空数据
	clear_s_data()
	ElMessage({
		message: '添加成功！',
		type: 'success',
	})
}
//计算运行时间
const c_runtime = computed((shutdown_api,shutdown_api_key)=>{
	
})
const formatter = (row, column) => {
	let shutdown_api = row.shutdown_api + "/api/health";
	let shutdown_api_key = row.shutdown_api_key;
	
	//索引值
	let key = row.mac.replace(/:/g,"");
	//请求健康检查接口
	const params = new URLSearchParams();
	params.append('key', shutdown_api_key);
	axios.post(shutdown_api, params)
		  .then(function (response) {
			let data = response.data
			if( data.code == 200 ) {
				run_time.value[key] = data.data + "小时";
			}
			else{
				run_time.value[key] = data.msg
			}
		  })
		  .catch(function (error) {
				run_time.value[key] = "获取失败"
		});
	return run_time.value[key]
}

onMounted(()=>{
	//判断设备类型
	let device = isMobilePhone();
	//如果是手机，则设置抽屉宽度为90%
	if( device ) {
		drawer_size.value = "80%";
	}
	
	let value = localStorage.getItem("s_datas");
	
	//转为对象
	value = JSON.parse(value)
	//console.log(value)
	//console.log(value.key);
	if( value !== null ) {
		s_datas.value = value;
		//console.log(s_datas.value);
	}
})

//删除数据
function del_sdata(row,index) {
	s_datas.value.splice(index,1)
	//重新保存到浏览器
	localStorage.setItem("s_datas",JSON.stringify(s_datas.value));
}
//编辑数据
function edit_sdata(index){
	//显示保存按钮
	btn_status.value.save = true
	btn_status.value.add = false
	//console.log(s_datas.value[index]);
	s_data.value = s_datas.value[index]
	//显示抽屉
	drawer.value = true;
	//全局索引赋值
	global_index.value = index;
}
</script>

<style>
body{
	margin: 0;
	padding:0;
}
.spacing{
	margin-bottom: 14px;
}
.mybtn{
	width: 100%;
}
</style>