<template>
	<form action="/api/info" method="post" @submit.prevent="submit">
		<input type="url" name="url">
		<button type="submit">Info</button>
	</form>
</template>

<script>
const {fetch} = require('fetch-ponyfill')()

module.exports = {
	props: ['value'],
	methods: {
		submit(event) {
			fetch('/api/info', {method: 'post', body: new FormData(event.target)})
			.then(res => res.json())
			.then(data => this.$emit('info', data))
			.catch(err => this.$emit('error', err))
		}
	}
}
</script>
