const Vue = require('vue')
const infoForm = require('./components/info-form.vue')
const infoFrame = require('./components/info-frame.vue')

const app = new Vue({
	el: '#app',
	data() {
		return {info: null}
	},
	components: {
		'info-form': infoForm,
		'info-frame': infoFrame
	},
	methods: {
		onInfo: function (info) {
			this.info = info
		}
	}
})
