import { createApp } from 'vue'
import ElementPlus from 'element-plus'
import router from './router'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import 'element-plus/dist/index.css'


import App from './App.vue'

const app = createApp(App)

app.use(router)
app.use(ElementPlus)
app.mount('#app')