<template>
	<view>
		<view class="top-swiper">
			 <view class="bg">
				 <view class="placeholder"></view>
				 <view class="image">
					 <image :src="swiper.list[swiper.index]" mode="aspectFill"></image>
				 </view>
			 </view>
			 <view class="box">
				<view style="height: 44px;"></view>
			 	<swiper class="swiper" :previous-margin="swiper.margin" :next-margin='swiper.margin' :circular="true" @change="swiperChange">
			 		<swiper-item v-for="(item,index) in swiper.list" :key="index">
			 			<image class='le-img' :src='item' :class="{'le-active':swiper.index == index}"></image>
			 		</swiper-item>
			 	</swiper> 
			 </view>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				swiper: {
					margin: "150rpx",
					index: 0,
					list: [
						"/static/images/douyin/1.jpg",
						"/static/images/douyin/3.jpg",
						"/static/images/douyin/4.jpg",
					]
				}
			}
		},
		components: {

		},
		mounted() {

		},
		methods: {
			//swiper滑动事件
			swiperChange: function(e) {
				this.swiper.index = e.detail.current;
			},
		}
	}
</script>

<style lang="scss">
	.top-swiper{
		margin-bottom: 30rpx; 
		
		.bg{
			padding-top: var(--status-bar-height);
			box-sizing: content-box;
			width: 100%;
			position: relative;
			
			.placeholder{
				box-sizing: content-box;
				padding-top: 600rpx;
				height: 44px;
			}
			
			.image{
				box-sizing: content-box;
				position: absolute;
				z-index: 1;
				top: 0;
				left: 0;
				width: 100%;
				height: 100%;
				overflow: hidden;
				
				&::after{
					content: '';
					position: absolute;
					width: 100%;
					height: 100%;
					z-index: 1;
					bottom: 0;
					left: 0;
					height: 65%;
					background-image: linear-gradient(to bottom ,transparent, #FFF);
				}
				
				> image{
					position: absolute;
					box-sizing: content-box;
					padding: 60px;
					top: 0;
					left: 0;
					width: 100%;
					height: 80%;
					top: -60px;
					left: -60px;
					filter: blur(50px);
				}
			}
		}
		
		.box{
			padding-top: var(--status-bar-height);
			box-sizing: content-box;
			position: absolute;
			z-index: 5;
			top: 0;
			left: 0;
			width: 100%;
			height: auto;
		}
		
		.swiper {
			height: 600rpx;
			margin: 0 20rpx;
		
			.le-img {
				width: 100%;
				height: 100%;
				display: block;
				transform: scale(0.9);
				transition: transform 0.3s ease-in-out 0s;
				border-radius: 4px;
		
				&.le-active {
					transform: scale(1);
				}
			}
		
		}
	}
	
</style>