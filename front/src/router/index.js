import {
	createRouter,
	createWebHistory,
	createWebHashHistory
} from 'vue-router'
import App from '../App.vue'
import Home from '../components/Home.vue'
import Wol from '../components/Wol.vue'
import Shutdown from '../components/Shutdown.vue'

//在很多场景中会有一部分公用的展示内容，例如导航，菜单，这个时候可以用子路由解决，我们可以自己建立一个公共文件夹layout
const router = createRouter({
	history: createWebHashHistory(),
	routes: [{
		path: "/",
		component: Home,
		//redirect:'/result/',//这个就是默认重定向,这个地方重定向不过去
		meta:{
			title: 'GoPower WEB控制台'
		}
	},
	{
		path:"/wol",
		component:Wol,
		meta:{
			title: '手动唤醒 - GoPower WEB控制台'
		}
	},
	{
		path:"/shutdown",
		component:Shutdown,
		meta:{
			title: '手动关机 - GoPower WEB控制台'
		}
	}
	]
})
router.beforeEach((from,to,next)=>{
	//注意一下，在这个地方如果你没有定义title的话会报错，所以要先判断是否现存在title，在做修改动作，
	const {title='默认名称'} = from?.meta;
	document.title = title;
	//这个和webman的洋葱模型有点像，要调用next;
	next();//这里还可以做权限拦截，如归没权限，可以吧跳转的地址传递到next里面
})
export default router
