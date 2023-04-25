<template>
	<view>
		<view class="btn" @click="btnClick" :style="{'background':bgColor,'color':titleColor,'height':height,'width':width}">{{title}}</view>
	</view>
</template>

<script>
	export default {
		name:"zkml-Button",
		props: {
			// 背景色
			bgColor: {
				type: String,
				default: "#f5f6f9",
				required: false
			},
			// 标题
			title: {
				type: String,
				default: '',
				required: false
			},
			// 标题颜色
			titleColor: {
				type: String,
				default:'#313131',
				required:false
			},
			// 高度
			height:{
				type: String,
				default:'88rpx',
				required:false
			},
			// 宽度
			width:{
				type: String,
				default:'690rpx',
				required:false
			},
			
		},
		data() {
			return {
				
			};
		},
		methods:{
			btnClick(){
				this.$emit('btnClick')
			}
		}
	}
</script>

<style>
    .btn{
		display: flex;
		justify-content: center;
		align-items: center;
		border-radius: 20rpx;
		height: 80rpx;
	}
</style>
