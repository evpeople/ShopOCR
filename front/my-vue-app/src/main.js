import './style.css'
import 'bootstrap';
import 'bootstrap/dist/css/bootstrap.min.css';

import {createApp} from 'vue'

import App from './App.vue';
import {FontAwesomeIcon} from './plugins/font-awsome'
import router from './router';
import store from './store';

createApp(App)
    .use(router)
    .use(store)
    .component('font-awesome-icon', FontAwesomeIcon)
    .mount('#app')
