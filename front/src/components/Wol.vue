<template>
	<el-row :gutter="0" style="margin-top:2em;">
		<el-col :xs="{ span: 22, offset: 1 }" :sm="{ span: 18, offset: 3 }">
			<div class="spacing">
				<el-input v-model="s_data.mac" placeholder="请输入要唤醒的MAC地址,如00:01:6C:06:A6:29">
					<template #prepend>MAC地址</template>
				</el-input>
			</div>
			<div class="spacing">
				<el-input v-model="s_data.ip" placeholder="请输入需要唤醒的IP地址">
					<template #prepend>IP地址</template>
				</el-input>
			</div>
			<div class="spacing">
				<el-input v-model="s_data.mask" placeholder="请输入需要唤醒的IP地址的子网掩码或掩码位元数">
					<template #prepend>子网掩码</template>
				</el-input>
			</div>
			<div class="spacing">
				<el-input v-model="s_data.wol_api" placeholder="请输入API地址,格式为http://192.168.20.3:886">
					<template #prepend>API地址</template>
				</el-input>
			</div>
			<div class="spacing">
				<el-input v-model="s_data.wol_api_key" placeholder="请输入API KEY">
					<template #prepend>KEY</template>
				</el-input>
			</div>

			<div class="spacing">
				<el-button type="success" class="mybtn" @click="manual_startup">手动唤醒</el-button>
			</div>
		</el-col>
	</el-row>
</template>

<script setup>
import { ref, watch, onMounted, computed } from 'vue'
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
	"mac": "",
	"ip": "",
	"mask": "",
	"wol_api": "",
	"wol_api_key": "",
	"broadcast": "",
})


//计算广播地址
function calculateBroadcastAddress(ip, mask) {
	let ipParts = ip.split('.').map(part => parseInt(part, 10));
	let maskParts = (mask.indexOf('.') !== -1) ?
		mask.split('.').map(part => parseInt(part, 10)) :
		cidrToMask(mask);

	let broadcastAddressParts = ipParts.map((part, index) => part | (255 - maskParts[index]));

	return broadcastAddressParts.join('.');
}
function cidrToMask(cidr) {
	let mask = [];
	for (let i = 0; i < 4; i++) {
		if (cidr >= 8) {
			mask.push(255);
			cidr -= 8;
		} else {
			mask.push(256 - Math.pow(2, 8 - cidr));
			cidr = 0;
		}
	}
	return mask;
}
//唤醒事件
function manual_startup() {
	//验证参数是否合法
	if (!v_params()) {
		return false;
	}

	//数据保存到浏览器方便下次使用
	localStorage.setItem("s_data", JSON.stringify(s_data.value));

	let wol_api = s_data.value.wol_api + "/api/startup"
	let wol_api_key = s_data.value.wol_api_key;


	const params = new URLSearchParams();
	params.append('key', wol_api_key);
	params.append('mac', s_data.value.mac);
	params.append('broadcast', s_data.value.broadcast);

	axios.post(wol_api, params)
		.then(function (response) {
			let data = response.data
			if (data.code == 200) {
				ElMessage({
					message: '唤醒成功！',
					type: 'success',
				})
			}
			else {
				ElMessage.error(data.msg)
			}
		})
		.catch(function (error) {
			ElMessage.error("网络请求失败！")
		});
}

onMounted(() => {
	//加载缓存数据
	let value = localStorage.getItem("s_data");

	//赋值数据
	if (value !== null) {
		value = JSON.parse(value)
		s_data.value = value
	}

	console.log(s_data.value)
})

//验证参数是否合法
function v_params() {
	//如果MAC地址为空
	if (s_data.value.mac == "") {
		ElMessage.error("MAC地址不能为空！");
		return false;
	}
	//判断MAC地址是否合法
	let mac_pattern = /^[0-9a-zA-Z]{2}:[0-9a-zA-Z]{2}:[0-9a-zA-Z]{2}:[0-9a-zA-Z]{2}:[0-9a-zA-Z]{2}:[0-9a-zA-Z]{2}$/
	if (!mac_pattern.test(s_data.value.mac)) {
		ElMessage.error("MAC地址格式不正确！");
		return false;
	}
	//判断IP是否正确
	let ip_pattern = /^(25[0-5]|2[0-4]\d|[0-1]?\d?\d)(\.(25[0-5]|2[0-4]\d|[0-1]?\d?\d)){3}$/
	//如果IP为空
	if (s_data.value.ip == "") {
		ElMessage.error("IP地址不能为空！");
		return false;
	}
	if (!ip_pattern.test(s_data.value.ip)) {
		ElMessage.error("IP地址格式不正确！");
		return false;
	}
	// 判断子网掩码是否正确
	if (s_data.value.mask == "" || s_data.value.mask == undefined) {
		ElMessage.error("子网掩码不能为空！");
		return false;
	}
	let mask_pattern1 = /^(\d{1,2}|1\d{2}|2[0-4]\d|25[0-5])(\.(\d{1,2}|1\d{2}|2[0-4]\d|25[0-5])){3}$/;
	let mask_pattern2 = /^([1-9]|[1-2][0-9]|3[0-2])$/;
	if ((s_data.value.mask != "") && !mask_pattern1.test(s_data.value.mask) && !mask_pattern2.test(s_data.value.mask)) {
		ElMessage.error("子网掩码格式不正确！");
		return false;
	}
	s_data.value.broadcast = calculateBroadcastAddress(s_data.value.ip, s_data.value.mask);
	//ElMessage.info(s_data.value.broadcast)
	//判断唤醒api地址
	let url_pattern = /^(http:\/\/|https:\/\/).+[0-9a-zA-Z]$/
	if (s_data.value.wol_api == "") {
		ElMessage.error("唤醒API地址不能为空！");
		return false;
	}
	if (!url_pattern.test(s_data.value.wol_api)) {
		ElMessage.error("唤醒API地址格式不正确！");
		return false;
	}
	//判断key
	if (s_data.value.wol_api_key == "") {
		ElMessage.error("唤醒API KEY不能为空！");
		return false;
	}
	else {
		return true;
	}
}
</script>

<style></style>