<template>
	<div class="info">
		<h1>{{ info.title }}</h1>
		<img :src="info.thumbnails[0].url" v-once alt="Thumbnail"/>
		<p>{{ info.description }}</p>

		<select v-model="format">
			<option v-for="f in info.formats" :value="f">
				{{ f.ext }} {{ f.format }}
			</option>
		</select>

		<select v-model="subtitles">
			<option :value="null">None</option>
			<option v-for="(s, lang) in info.subtitles" :value="s">
				{{ lang }}
			</option>
		</select>

		<p>
			<code>{{ args() }}</code>
		</p>
	</div>
</template>

<script>
module.exports = {
	props: ['info'],
	data() {
		return {
			format: null,
			subtitles: null
		}
	},
	methods: {
		args() {
			const args = ['mpv']

			if (this.format) {
				args.push('"'+this.format.url+'"')
			}
			if (this.subtitles) {
				const sub = this.subtitles.filter(s => ['srt', 'ass', 'vtt'].indexOf(s.ext) > 0)[0]

				if (sub) {
					args.push('--sub-file')
					args.push('"'+sub.url+'"')
				}
			}

			return args.join(' ')
		}
	}
}
</script>

<style scoped>
img {
	max-width: 300px;
}
p {
	white-space: pre-wrap;
}
</style>
