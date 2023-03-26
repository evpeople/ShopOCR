import './style.css'
import 'bootstrap';
import 'bootstrap/dist/css/bootstrap.min.css';

import {createApp} from 'vue'
import Vue3BaiduMapGL from 'vue3-baidu-map-gl'
import baiduMap from 'vue3-baidu-map-gl'

import App from './App.vue';
import {FontAwesomeIcon} from './plugins/font-awsome'
import router from './router';
import store from './store';

createApp(App)
    .use(router)
    .use(Vue3BaiduMapGL, {ak: 'ZienfAK0hlDGMlqcjYeepy8PXpPIaFA6'})
    .use(baiduMap, {ak: 'ZienfAK0hlDGMlqcjYeepy8PXpPIaFA6'})
    .use(store)
    .component('font-awesome-icon', FontAwesomeIcon)
    .mount('#app')