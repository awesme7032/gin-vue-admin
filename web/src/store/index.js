import Vue from 'vue'
import Vuex from 'vuex'
import VuexPersistence from 'vuex-persist'

import { user } from "@/store/module/user"
import { router } from "@/store/module/router"
import { dictionary } from "@/store/module/dictionary"
Vue.use(Vuex)



const vuexLocal = new VuexPersistence({
    storage: window.localStorage,
    modules: ['user']
})
export const store = new Vuex.Store({
    modules: {
        user,
        router,
        dictionary
    },
    state: {
        imgPrefix: 'http://127.0.0.1:8888/',
        uploadUrl: 'http://127.0.0.1:8888/uploadImg',
        uploadFullUrl: "http://127.0.0.1:8888/uploadImgFull", // 不压缩的上传图片接口

        // uploadFullUrl: "http://api.asfenfm.com/uploadImgFull", // 不压缩的上传图片接口
        // imgPrefix: 'http://api.asfenfm.com/'
        // uploadUrl: 'http://localhost:19000/uploadImg'
    },
    plugins: [vuexLocal.plugin]
})