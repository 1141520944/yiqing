<template>
	<view class="content">
		<view class="btn">
			<button type="default" @click="test">扫描</button>
		</view>
		<view class="admin">
			<view>{{msg}}</view>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				Code: '',
				openid: '',
				session_key: ''
			}
		},
		methods: {
			test:function(){
			
			},
		}
	}
</script>

<style>

</style>
